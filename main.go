package main

import (
	"encoding/json"
	"go-crud-api/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	err := repository.InitDB("root:910095879@tcp(127.0.0.1:3306)/go_crud_api")
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}
	defer repository.GetDB().Close()

	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user repository.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repository.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user repository.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repository.UpdateUser(id, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = id
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
