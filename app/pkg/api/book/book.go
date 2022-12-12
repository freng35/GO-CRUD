package book

import (
	database "app/pkg/db"
	"app/pkg/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func addBook(book model.Book) {
	data := database.GetBookByName(book)
	if data == (model.Book{}) {
		database.CreateBook(book)
	} else {
		database.AddBook(book)
	}
}

func AddBooks(w http.ResponseWriter, r *http.Request) {
	var books []model.Book
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&books)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO AMOUNT < 0
	for _, book := range books {
		addBook(book)
	}

	w.WriteHeader(http.StatusOK)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := database.GetBooks()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]model.Book{"items": books})
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	bookName := mux.Vars(r)["name"]
	data := database.GetBookByName(model.Book{Name: bookName})

	w.WriteHeader(http.StatusOK)

	if data == (model.Book{}) {
		json.NewEncoder(w).Encode(make(map[string]string))
	} else {
		json.NewEncoder(w).Encode(data)
	}
}
