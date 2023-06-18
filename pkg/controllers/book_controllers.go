package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/rogaliiik/bookstore/pkg/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rogaliiik/bookstore/pkg/utils"
)

func WriteJSONResponse(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	NewBooks := repository.Repo.GetAllBooks()

	res, err := json.Marshal(NewBooks)
	if err != nil {
		log.Print(err)
	}

	WriteJSONResponse(w, res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails := repository.Repo.GetBookById(ID)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		log.Print(err)
	}

	WriteJSONResponse(w, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &repository.Book{}
	utils.ParseBody(r, newBook)
	newBook = repository.Repo.CreateBook(newBook)
	res, _ := json.Marshal(newBook)

	WriteJSONResponse(w, res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("error while parsing")
	}

	deletedBook := repository.Repo.GetBookById(ID)
	_ = repository.Repo.DeleteBook(ID)
	res, err := json.Marshal(deletedBook)
	if err != nil {
		log.Print(err)
	}

	WriteJSONResponse(w, res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &repository.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("error while parsing")
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
	res, err := json.Marshal(bookDetails)
	if err != nil {
		log.Print(err)
	}

	WriteJSONResponse(w, res)
}
