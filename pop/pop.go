package pop

import (
	"fmt"
	"io"
	"log"
	"mime/quotedprintable"

	"github.com/joho/godotenv"
	"github.com/knadh/go-pop3"

	"mailer/environment"
)

var lastReadMessage int = 0

func GetEmails(config environment.POP3Config) {
	fmt.Println(config)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the client.
	p := pop3.New(pop3.Opt{
		Host:       config.Host,
		Port:       config.Port,
		TLSEnabled: config.TLSEnabled,
	})

	// Create a new connection. POP3 connections are stateful and should end
	// with a Quit() once the opreations are done.
	c, err := p.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Quit()

	// Authenticate.
	if err := c.Auth(config.Username, config.Password); err != nil {
		log.Fatal(err)
	}

	// Print the total number of messages and their size.
	count, size, _ := c.Stat()
	fmt.Println("total messages=", count, "size=", size)

	// Pull the list of all message IDs and their sizes.
	msgs, _ := c.List(0)
	for _, m := range msgs {
		fmt.Println("id=", m.ID, "size=", m.Size)
	}

	// Pull all messages on the server. Message IDs go from 1 to N.
	for id := 1; id <= count; id++ {
		m, _ := c.Retr(id)

		fmt.Println(id, "=", m.Header.Get("subject"))

		message, _ := io.ReadAll(quotedprintable.NewReader(m.Body))

		fmt.Printf("%s\n", string(message))

		lastReadMessage = id
	}

	fmt.Println("Last read message", lastReadMessage)
}