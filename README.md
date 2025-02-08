# Go CRUD API with MySQL and Gorilla Mux

This project is a simple CRUD (Create, Read, Update, Delete) API built in Go. It uses MySQL as the database and the Gorilla Mux package for HTTP routing. The API allows you to create, read, update, and delete users.

## Features

- **Get All Users**: Retrieve a list of all users.
- **Get User by ID**: Retrieve details for a specific user by their ID.
- **Create New User**: Add a new user to the database.
- **Update Existing User**: Update the details of an existing user.
- **Delete User**: Remove a user from the database.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later recommended)
- [MySQL](https://dev.mysql.com/downloads/)
- Go packages:
    - [Gorilla Mux](https://github.com/gorilla/mux)
    - [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-crud-api.git
cd go-crud-api
