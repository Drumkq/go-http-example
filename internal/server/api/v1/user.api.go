package userApi

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"example.com/m/ent"
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

func GetUser(id int) (*ent.User, error) {
	return database.DB.Client.User.Get(context.Background(), id)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("users got")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Printf("user with id %v got", mux.Vars(r)["id"])
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("user created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		// TODO: error throwing and error handling
		log.Fatalf("Internal server error: %v", err.Error())
		return
	}

	log.Printf("user with id %v deleted", id)
	w.WriteHeader(http.StatusOK)
}
