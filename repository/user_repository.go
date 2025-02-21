// repository/user_repository.go
package repository

import (
	"log"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		log.Println("Failed to query users:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			log.Println("Failed to scan user:", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByID retrieves a single user by ID from the database
func GetUserByID(id int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser inserts a new user into the database
func CreateUser(user *User) error {
	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(id int, user *User) error {
	_, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the database
func DeleteUser(id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
