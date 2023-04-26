package new

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/session"
)

// Create a NewUser as authorized by the user 'user'
func NewUser(user string) {

	log.Println(user + " is creating a new user\n")

	if !db.NoUsers() && !db.UserExists(user) {
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		log.Println(user + " couldn't log in while creating a new user\n")
		log.Fatalf("Unable to authenticate user")
	}

	newEmail := getNewEmail()
	newUser := getNewUsername()
	newPassHash, err := session.GetPassword()
	if err != nil {
		log.Fatalf("Unable to get password hash")
	}

	err = db.SetUserPassHash(newUser, newEmail, newPassHash)
	if err != nil {
		log.Fatalf("Unable to create new user")
	}

	log.Println(user + " successfully created a new user: " + newUser + "\n")
}

// getUserMessage prompts the user for the message to send
// and returns it
func getNewUsername() string {

	fmt.Println("Enter the username for the new user: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	return strings.Trim(text, "\n\t ")
}

func getNewEmail() string {

	fmt.Println("Enter an email for the new user: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	return strings.Trim(text, "\n\t ")
}
