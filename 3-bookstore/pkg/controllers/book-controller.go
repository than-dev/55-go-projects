package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/than-dev/55-go-projects/3-api+mysql/pkg/models"
	"github.com/than-dev/55-go-projects/3-api+mysql/pkg/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks();
	res, _ := json.Marshal(newBooks)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
} 

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	
	book, _ := models.GetBookById(bookId);
	
	res, _ := json.Marshal(book)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
} 

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	
	book := models.DeleteBook(bookId);

	res, _ := json.Marshal(book)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Book := &models.Book{}
	utils.ParseBody(r, Book)
	book := Book.CreateBook()
	
	res, _ := json.Marshal(book)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)

	bookInfo, db := models.GetBookById(bookId)

	if updatedBook.Name != "" {
		bookInfo.Name = updatedBook.Name
	}

	if updatedBook.Author != "" {
		bookInfo.Name = updatedBook.Author
	}

	if updatedBook.Publication != "" {
		bookInfo.Name = updatedBook.Publication
	}

	db.Save(&bookInfo)
	res, _ := json.Marshal(bookInfo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}




var NewBook models.Book