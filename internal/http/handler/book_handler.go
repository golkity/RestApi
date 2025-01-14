package handler

import (
	"RestApi/internal/models"
	"RestApi/internal/storage"
	"RestApi/pkg/logger"
	"encoding/json"
	"net/http"
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
