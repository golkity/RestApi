package storage

import (
	"RestApi/internal/Error"
	"RestApi/internal/models"
)

type MemoryStorage struct {
	books []models.Book
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{books: make([]models.Book, 0)}
}

func (s *MemoryStorage) AddBooks(book models.Book) {
	book.ID = len(s.books) + 1
	s.books = append(s.books, book)
}

func (s *MemoryStorage) GetBooks() []models.Book {
	return s.books
}

func (s *MemoryStorage) GetBook(id int) (*models.Book, error) {
	for _, book := range s.books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, Error.BookNotFoundError
}

func (s *MemoryStorage) UpdateBook(id int, updateBook models.Book) error {
	for i, book := range s.books {
		if book.ID == id {
			s.books[i] = updateBook
			s.books[i].ID = id
			return nil
		}
	}
	return Error.BookNotFoundError
}

func (s *MemoryStorage) DeleteBook(id int) error {
	for i, book := range s.books {
		if book.ID == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return nil
		}
	}
	return Error.BookNotFoundError
}
