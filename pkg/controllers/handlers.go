package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rogaliiik/bookstore/pkg/repository"
	"github.com/rogaliiik/bookstore/pkg/utils"
	"log"
	"net/http"
	"strconv"
)

type Handler interface {
	Handle(*http.Request) []byte
}

type GetAllBooksHandler struct{}

func (h *GetAllBooksHandler) Handle(r *http.Request) []byte {
	NewBooks := repository.Repo.GetAllBooks()

	data, err := json.Marshal(NewBooks)
	if err != nil {
		log.Print(err)
	}
	return data
}

type GetBookByIDHandler struct{}

func (h *GetBookByIDHandler) Handle(r *http.Request) []byte {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		log.Print(err)
	}
	bookDetails := repository.Repo.GetBookById(ID)
	data, err := json.Marshal(bookDetails)
	if err != nil {
		log.Print(err)
	}
	return data
}

type CreateBookHandler struct{}

func (h *CreateBookHandler) Handle(r *http.Request) []byte {
	newBook := &repository.Book{}
	utils.ParseBody(r, newBook)
	newBook = repository.Repo.CreateBook(newBook)
	data, err := json.Marshal(newBook)
	if err != nil {
		log.Print(err)
	}
	return data
}

type DeleteBookHandler struct{}

func (h *DeleteBookHandler) Handle(r *http.Request) []byte {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		log.Print(err)
	}

	deletedBook := repository.Repo.GetBookById(ID)
	_ = repository.Repo.DeleteBook(ID)
	data, err := json.Marshal(deletedBook)
	if err != nil {
		log.Print(err)
	}
	return data
}

type UpdateBookHandler struct{}

func (h *UpdateBookHandler) Handle(r *http.Request) []byte {
	var updateBook = &repository.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		log.Print(err)
	}
	bookDetails := repository.Repo.GetBookById(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	repository.Repo.DB.Save(&bookDetails)
	data, err := json.Marshal(bookDetails)
	if err != nil {
		log.Print(err)
	}
	return data
}
