package main

import (
	"database/sql"
	"fmt"
	"log"

	// Postgres driver for Go, the underscore is used to import the package for its side effects (registering the driver)
	// Side effects are changes that occur outside of a function's scope, such as modifying global variables or performing I/O operations.
	// In this case, the side effect is registering the PostgreSQL driver with the database/sql package, allowing it to be used for database operations.
	_ "github.com/lib/pq"
)

type User struct {
	name     string
	password string
	email    string
	role     string
}

// Go to go docs to find best functions needed
func main() {
	// WARNING: Do not use this in production code, password is hard coded in this url
	connStr := "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"

	// Open a connection to the database
	// first argument is the driver name, second is the connection string
	// open returns two values, a pointer to the database and an error
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		// handle error
		log.Fatal(err)
	}

	// close the database connection when the function returns
	defer db.Close()

	//db.ping is a method that checks if the database is reachable, its a health function
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	createUserTable(db)

	// testUser := User{
	// 	name:     "John Doe",
	// 	password: "password",
	// 	email:    "test21@gmail.com",
	// 	role:     "admin",
	// }
	getAllUsers(db)
}

// the param is a sql database pointer, this is the connection to the database
func createUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT NOT NULL DEFAULT 'client'
)`

	// Execute the query, this will create the table if it does not exist
	// The Exec method is used to execute a query that doesn't return rows, such as an INSERT, UPDATE, or DELETE statement.
	// It returns a Result and an error. The Result contains information about the execution of the query, such as the number of rows affected.
	// Since we are creating a table, we don't need to check the result, we just need to check for errors.
	// The underscore is used to ignore the result, since we don't need it in this case.
	_, err := db.Exec(query)

	if err != nil {
		// handle error
		log.Fatal(err)
	}
}

func insertUser(db *sql.DB, user User) int {
	// The $1, $2, $3, $4 are placeholders for the values that will be passed to the query.
	query := `INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`

	var pk int

	// The QueryRow method is used to execute a query that returns a single row.
	// First argument is the query string, the rest are the values to be passed to the query, in order.
	// The Scan method is used to copy the columns from the row into the variables passed as arguments.
	// &pk is a pointer to the variable above where the value will be stored.
	err := db.QueryRow(query, user.name, user.email, user.password, user.role).Scan(&pk)

	if err != nil {
		// handle error
		log.Fatal(err)
	}
	return pk
}

func getUser(db *sql.DB, id int) {

	var name string
	var email string
	var password string
	var role string

	query := "SELECT name, email, password, role FROM users WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&name, &email, &password, &role)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("no rows were returned for id %d", id)
		}
		log.Fatal(err)
	}
	fmt.Printf(" User: %s, %s, %s, %s\n", name, email, password, role)
}

func getAllUsers(db *sql.DB) {

	data := []User{}                                                       // sets up an empy slice of users
	rows, err := db.Query("SELECT name, email, password, role FROM users") // query gets all users from the database
	if err != nil {
		log.Fatal(err)
	}
	// Close the row when the function returns
	defer rows.Close()

	var name string
	var email string
	var password string
	var role string

	// iterates over all the returned data
	for rows.Next() {
		err := rows.Scan(&name, &email, &password, &role)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, User{name, email, password, role})

	}

	fmt.Println("Users:", data)
}
