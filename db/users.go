package db

import (
	"database/sql"
	"fmt"
	"log"
)

// ErrNoUser is a generic error for no user existing
type ErrNoUser struct{}

func (e *ErrNoUser) Error() string {
	return "user does not exist"
}

// UserExists takes a username 'user' and returns true if that
// user exists and false otherwise
func UserExists(user string) bool {
	db := Connect().Db

	var id int
	if err := db.QueryRow("SELECT id FROM Users WHERE user=?;", user).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatalf("An unexpected error occured checking the database")
		}
	}
	return true
}

// GetUserId will return the UserId number for the user 'user'
// if that user exists and an error otherwise
func GetUserId(user string) (int, error) {
	db := Connect().Db

	var id int
	if err := db.QueryRow("SELECT id FROM Users WHERE user = ?;", user).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return 0, &ErrNoUser{}
		} else {
			log.Fatalf("An unexpected error occured checking the database")
		}
	}
	return id, nil
}

// GetUserPassHash will return the string representation of a password
// hash if the user exists and an error otherwise
func GetUserPassHash(user string) (string, error) {
	db := Connect().Db

	var hash string
	if err := db.QueryRow("SELECT hash FROM Users WHERE user = ?;", user).Scan(&hash); err != nil {
		if err == sql.ErrNoRows {
			return "", &ErrNoUser{}
		} else {
			log.Fatalf("An unexpected error occured checking the database")
		}
	}
	return hash, nil
}

// SetUsePassHash will allow for adding a new user to the Users table
func SetUserPassHash(user, hash string) error {
	db := Connect().Db

	_, err := db.Exec(`
		INSERT INTO Users (user, hash)
		VALUES (
			?,
			?
		);
	`, user, hash)
	return err
}

// Returns true if no users have been registered and false otherwise
func NoUsers() bool {
	db := Connect().Db

	var usersExist bool
	err := db.QueryRow("SELECT IIF(COUNT(*),'true', 'false') FROM Users;").Scan(&usersExist)
	if err != nil {
		fmt.Println(err)
		return true
	}

	return !usersExist
}
