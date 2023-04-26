package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"github.com/andey-robins/deaddrop-go/new"
	"github.com/andey-robins/deaddrop-go/read"
	"github.com/andey-robins/deaddrop-go/send"
	"github.com/joho/godotenv"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

	err = godotenv.Load(".env")
	if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)
	defer file.Close()

	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var to, user string
	var new, send, read, help bool
	flag.BoolVar(&help, "help", false, "Get help")
	flag.StringVar(&to, "to", "void", "the username to send data to")
	flag.StringVar(&user, "user", "void", "the username to retrieve data for")
	flag.BoolVar(&new, "new", false, "run the utility in add user mode")
	flag.BoolVar(&send, "send", false, "run the utility in send mode")
	flag.BoolVar(&read, "read", false, "run the utility in read mode")
	flag.Parse()

	if help {
		pad := func() {
			fmt.Printf("\n\n")
		}

		pad()
		fmt.Println(" Welcome to deaddrop, your solution to all deaddrop data needs")
		fmt.Println(" This code is licensed under GPLv3")
		pad()
		fmt.Println("Args:")
		fmt.Println("  -to    The username of the user to send this data to")
		fmt.Println("  -user  The username of the user currently using the system")
		fmt.Println("  -new   The verb flag to specify you want to create a new user")
		fmt.Println("  -send  The verb flag to specify you want to send data")
		fmt.Println("  -read  The verb flag to specify you want to retreive your data")
		pad()
		fmt.Println(" Option -send must include the flags: -to")
		fmt.Println(" Option -new  must include the flags: -user")
		fmt.Println(" Option -read must include the flags: -user")
		pad()
		return
	}

	if !read && !send && !new {
		fmt.Println("Please specify a verb for the utility.")
		fmt.Println("Valid verbs: send, read")
		return
	}

	if read && send || new && read || new && send || new && send && read {
		fmt.Println("Please specify only one verb")
		return
	}

	if read {
		readMode(user)
	} else if send {
		sendMode(to)
	} else if new {
		newMode(user)
	}
}

func readMode(user string) {
	read.ReadMessages(user)
}

func sendMode(to string) {
	send.SendMessage(to)
}

func newMode(user string) {
	new.NewUser(user)
}
