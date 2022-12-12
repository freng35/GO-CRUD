package user

import (
	database "app/pkg/db"
	"app/pkg/model"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.Id = uuid.New().String()

	database.GetUserByPhone(user)
	database.CreateUser(user)
}
