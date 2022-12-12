package user

import (
	database "app/pkg/db"
	"app/pkg/model"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func checkUser(user model.User) (string, error) {
	data := database.GetUserByPhone(user)
	if data == (model.User{}) {
		return "", nil
	}

	return data.Id, &ExistUserError{}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := checkUser(user)
	if err == nil {
		user.Id = uuid.New().String()
		database.CreateUser(user)
	} else {
		user.Id = id
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"id": user.Id})
}
