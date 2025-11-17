package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lonewolf-nb/Golang_Journey/tree/master/Go_Projects/2_Bookstore_Management/pkg/models"
	"github.com/lonewolf-nb/Golang_Journey/tree/master/Go_Projects/2_Bookstore_Management/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBookById handles GET requests to fetch a single book by its ID.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	// Extract all path variables from the request URL using Gorilla Mux.
	vars := mux.Vars(r)

	// Get the "bookId" variable value from the URL (e.g., /book/{bookId}).
	bookId := vars["bookId"]

	// Convert the bookId (string) into an integer.
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// Log error if conversion fails.
		fmt.Println("error while parsing book ID")
	}

	// Fetch book details from the database using the models package.
	bookDetails, _ := models.GetBookById(ID)

	// Convert the book details struct into JSON format.
	res, _ := json.Marshal(bookDetails)

	// Set response header to indicate JSON content (⚠ Typo: should be "application/json").
	w.Header().Set("Content-Type", "pkglication/json")

	// Set HTTP status to 200 OK.
	w.WriteHeader(http.StatusOK)

	// Write the JSON response body to the client.
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r * http.Request){
	vars :=  mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r * http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}