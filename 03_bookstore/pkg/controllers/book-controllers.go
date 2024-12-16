package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitishb2m/golang_projects/03_bookstore/pkg/models"
	"github.com/nitishb2m/golang_projects/03_bookstore/pkg/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	utils.JsonResponse(books, w)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		return
	}
	bookDetail, _ := models.GetBookById(ID)
	utils.JsonResponse(bookDetail, w)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	b, _ := newBook.CreateBook()
	utils.JsonResponse(b, w)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	if err := models.DeleteBook(ID); err != nil {
		http.Error(w, fmt.Sprintf("Error deleting book: %v", err), http.StatusInternalServerError)
		return
	}
	utils.JsonResponse(map[string]string{"message": "Book deleted successfully"}, w)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}
	bookDetail, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetail.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetail.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetail.Publication = updateBook.Publication
	}

	db.Save(&bookDetail)
	utils.JsonResponse(bookDetail, w)
}
