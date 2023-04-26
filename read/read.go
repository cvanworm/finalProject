package read

import (
	"fmt"
	"log"
	"os"
	"crypto/hmac"
    "crypto/sha256"
    "encoding/hex"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/session"
)

func ReadMessages(user string) {

	if !db.UserExists(user) {
		log.Println("Failed to read messages for a user that doesn't exist: " + user + "\n")
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		log.Println("Failed to read messages with the wrong password for: " + user + "\n")
		log.Fatalf("Unable to authenticate user")
	}

	messages := db.GetMessagesForUser(user)
	for _, message := range messages {
		MAC := verify([]byte(message.Message), []byte(os.Getenv("KEY")), message.Hash)
		if(MAC == true){
			fmt.Println(message.Sender + ": " + message.Message)
		}else{
			log.Println("Message sent to " + user + " could not be verified.")
			fmt.Println("Integrity of message could not be verified: " + "\n" + message.Sender + ": " + message.Message)
		}
		

	}

	log.Println(user + " read their messages successfully\n")
}

func verify(msg, key []byte, hash string) bool {
	sig, err := hex.DecodeString(hash)
	if err != nil {
		return false
	}
 
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
 
	return hmac.Equal(sig, mac.Sum(nil))
}