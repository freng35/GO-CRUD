package order

import (
	database "app/pkg/db"
	"app/pkg/model"
	"encoding/json"
	"net/http"
)

func checkOrder(order model.Order) error {
	userData := database.GetUserById(model.User{Id: order.UserId})

	if userData == (model.User{}) {
		return &NoUserError{}
	}

	bookData := database.GetBookByName(model.Book{Name: order.BookName})

	if bookData == (model.Book{}) {
		return &NoBookError{}
	}

	if bookData.Amount == 0 {
		return &ZeroBooksError{}
	}

	return nil
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = checkOrder(order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	database.DelBook(database.GetBookByName(model.Book{Name: order.BookName}))
	database.CreateOrder(order)
}
