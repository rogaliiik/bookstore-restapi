package controllers

import (
	"log"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(data)
	if err != nil {
		log.Print(err)
	}
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	handler := &GetAllBooksHandler{}
	data := handler.Handle(r)
	WriteJSONResponse(w, data)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	handler := &GetBookByIDHandler{}
	data := handler.Handle(r)
	WriteJSONResponse(w, data)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	handler := &CreateBookHandler{}
	data := handler.Handle(r)
	WriteJSONResponse(w, data)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	handler := &DeleteBookHandler{}
	data := handler.Handle(r)
	WriteJSONResponse(w, data)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	handler := &UpdateBookHandler{}
	data := handler.Handle(r)
	WriteJSONResponse(w, data)
}
