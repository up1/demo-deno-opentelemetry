package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/user/:id", func(c echo.Context) error {
		id := c.Param("id")
		// Create the database if it doesn't exist
		db, err := sql.Open("sqlite3", "./foo.db")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Database error with "+err.Error())
		}
		// Create the table if it doesn't exist
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY, name TEXT)")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Database error with "+err.Error())
		}
		// Insert a user if it doesn't exist
		_, err = db.Exec("INSERT OR IGNORE INTO users (id, name) VALUES (?, ?)", id, "Default Name")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Database error with "+err.Error())
		}
		defer db.Close()
		// Query the database
		row := db.QueryRow("SELECT name FROM users WHERE id = ?", id)
		var name string
		err = row.Scan(&name)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.String(http.StatusNotFound, "User not found")
			}
			return c.String(http.StatusInternalServerError, "Database error")
		}
		// Return the user name
		return c.String(http.StatusOK, "User ID: "+id+", Name: "+name)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
