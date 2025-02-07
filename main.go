package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type USER struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

var db *sql.DB

func main() {
	var err error
	// Update the connection string with your MySQL credentials
	db, err = sql.Open("mysql", "root:910095879@tcp(127.0.0.1:3306)/go_crud_api")
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping MySQL:", err)
	}
	log.Println("Connected to MySQL")

	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []USER
	rows, err := db.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed to query users:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user USER
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Failed to scan user:", err)
			return
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user USER
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = int(id)
	user.CreatedAt = "now"
	json.NewEncoder(w).Encode(user)
}
