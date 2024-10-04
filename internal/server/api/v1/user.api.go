package userApi

// TODO: custom error handling

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"example.com/m/ent"
	"example.com/m/ent/user"
	"example.com/m/internal/database"
	"github.com/gorilla/mux"
)

func New(r *mux.Router) {
	user := r.NewRoute().PathPrefix("/user").Subrouter()
	user.HandleFunc("", getUsers).Methods(http.MethodGet)
	user.HandleFunc("/{id:[0-9]}", getUser).Methods(http.MethodGet)
	user.HandleFunc("", createUser).Methods(http.MethodPost)
	user.HandleFunc("/{id:(?:[0-9]{1,})}", deleteUser).Methods(http.MethodDelete)
}

func Get(id int) (*ent.User, error) {
	return database.DB.Client.User.Get(context.Background(), id)
}

func GetAll() ([]*ent.User, error) {
	return database.DB.Client.User.Query().All(context.Background())
}

func Create(u ent.User) (*ent.User, error) {
	_, err := database.DB.Client.User.
		Query().
		Where(func(s *sql.Selector) { sql.EQ(s.C(user.FieldUsername), u.Username) }).
		Only(context.Background())
	if err == nil {
		return nil, fmt.Errorf("user with this username already exists")
	}

	return database.DB.Client.User.Create().
		SetUsername(u.Username).
		SetEmail(u.Email).
		SetPassword(u.Password).
		Save(context.Background())
}

func DeleteOne(id int) error {
	deleteOne := database.DB.Client.User.DeleteOneID(id)
	return deleteOne.Exec(context.Background())
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	u, err := GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 32)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	u, err := Get(int(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user ent.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	u, err := Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 32)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if err := DeleteOne(int(id)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
