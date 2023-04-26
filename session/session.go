package session

import (
	"fmt"
	"log"
	"os"

	"github.com/andey-robins/deaddrop-go/db"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

// GetPassword will read in a password from stdin using the terminal
// no-echo utility ReadPassword. it will then salt and hash it with
// bcrypt
func GetPassword() (string, error) {
	pass, err := readPass()
	if err != nil {
		return "", err
	}

	return saltAndHash(pass)
}

// a nice wrapper to encapsulate the bcrypt generateFromPassword func
// to make salting and hashing easier
func saltAndHash(pass []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Authenticate takes in the username of a user and returns nil
// if the given password matches the user and an error otherwise
func Authenticate(user string) error {

	// bypass authentication if no users exist as this means
	// there is no data being stored and we can let the user create
	// a new user without auth
	if db.NoUsers() {
		return nil
	}

	pass, err := readPass()
	if err != nil {
		log.Fatalf("Error reading in password")
	}

	hash, err := db.GetUserPassHash(user)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(hash), pass)
}

// using the built in password read utility, get a password from stdin
func readPass() ([]byte, error) {
	fmt.Println("Password: ")
	pass, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}
	return pass, nil
}
