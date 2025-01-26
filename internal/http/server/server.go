package server

import (
	"RestApi/internal/http/handler"
	storage "RestApi/internal/storage"
	"RestApi/pkg/logger"
	"net/http"
)

func RegRoutes(mux *http.ServeMux, log *logger.Logger) {
	storage := storage.NewMemoryStorage()
	bookHandler := handler.NewBookHandler(log, storage)

	mux.HandleFunc("/api/v1/books", bookHandler.GetBooks)
	mux.HandleFunc("/api/v1/books/add", bookHandler.AddBook)
	mux.HandleFunc("/api/v1/books/get", bookHandler.GetBookId)
	mux.HandleFunc("/api/v1/books/update", bookHandler.UpdateBook)
	mux.HandleFunc("/api/v1/books/delete", bookHandler.DeleteBook)

	mux.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Health server - OK"))
	})
}
