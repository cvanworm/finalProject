package send

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"crypto/hmac"
    "crypto/sha256"
    "encoding/hex"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/session"
)

// SendMessage takes a destination username and will
// prompt the user for a message to send to that user
func SendMessage(to string) {
	if !db.UserExists(to) {
		log.Println("Destination user does not exist: " + to + "\n")
		log.Fatalf("Destination user does not exist")
	}

	username := loginUserName()
	if !db.UserExists(username) {
		log.Println("User tried to send a message from a user that does not exist: " + username + "\n")
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(username)
	if err != nil {
		log.Println(username + " failed to login to send a message\n")
		log.Fatalf("Unable to authenticate user")
	}

	message := getUserMessage()
	
	h := hmac.New(sha256.New, []byte(os.Getenv("KEY")))
	h.Write([]byte(message))

	hash := hex.EncodeToString(h.Sum(nil))

	
	log.Println(username + " sent a message to " + to + "\n")
	db.SaveMessage(message, to, username, hash)
}

// getUserMessage prompts the user for the message to send
// and returns it
func getUserMessage() string {
	fmt.Println("Enter your message: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return text
}

func loginUserName() string{
	fmt.Println("Please log in to send a message.\n Username: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return strings.Trim(text, "\n\t ")
}