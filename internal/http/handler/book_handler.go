package handler

import (
	"RestApi/internal/models"
	"RestApi/internal/storage"
	"RestApi/pkg/logger"
	"encoding/json"
	"net/http"
	"strconv"
)

type BookHandler struct {
	log     *logger.Logger
	storage *storage.MemoryStorage
}

func NewBookHandler(log *logger.Logger, storage *storage.MemoryStorage) *BookHandler {
	return &BookHandler{log: log, storage: storage}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Fetching all books")
	books := h.storage.GetBooks()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.storage.AddBooks(book)
	h.log.Info("Book added successfully")

	w.WriteHeader(http.StatusCreated)
}

func (h *BookHandler) GetBookId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) //
		return
	}

	book, err := h.storage.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) //404
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.storage.UpdateBook(id, updatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.log.Info("Book updated successfully")
	w.WriteHeader(http.StatusOK)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if err := h.storage.DeleteBook(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.log.Info("Book deleted successfully")
	w.WriteHeader(http.StatusOK)
}
