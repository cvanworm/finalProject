package db

import (
	"log"
	//"fmt"
)

type MessageInfo struct{
	Message string
	Sender string
	Hash string
}
// GetMessagesForUser assumes that a user has already been
// authenticated through a call to session.Authenticate(user)
// and then returns all the messages stored for that user
func GetMessagesForUser(user string) []MessageInfo {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT Users.user, Messages.data, Messages.Hash
		FROM Messages
		INNER JOIN Users ON Messages.sender=Users.id
		AND Messages.recipient = (
			SELECT id FROM Users WHERE user = ?
		)
	`, user)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer rows.Close()

	// marshall rows into an array
	messages := []MessageInfo{}
	for rows.Next(){
		message := MessageInfo{}
		if err := rows.Scan(&message.Sender, &message.Message, &message.Hash); err != nil{
			log.Fatal("Could not scan row")
		}
		messages = append(messages, message)
	}

	return messages
}


// saveMessage will process the transaction to place a message
// into the database
func SaveMessage(message, recipient, sender, hash string) {
	database := Connect().Db
	database.Exec(`
		INSERT INTO Messages (recipient, sender, data, hash)
		VALUES (
			(SELECT id FROM Users WHERE user = ?),
			(SELECT id FROM Users WHERE user = ?),
			?,
			?
		);
	`, recipient, sender, message, hash)
}
