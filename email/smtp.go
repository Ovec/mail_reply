package email

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"

	"mailer/environment"

	"github.com/emersion/go-message"
)

func SendMail(config environment.ConnectionConfig, email *message.Entity) {
	// Message to send
	// message := []byte("To: " + email.To + "\r\n" +
	// 	"Subject: " + email.Subject + "\r\n" +
	// 	"\r\n" + email.Body)

	// // Authenticate with the SMTP server
	// fmt.Println("", email.From, config.Password, config.Host, config.Port)
	// auth := smtp.PlainAuth("", email.From, config.Password, config.Host)

	// // Send the email
	// err := smtp.SendMail(config.Host+":"+string(rune(config.Port)), auth, email.From, []string{email.To}, message)
	// if err != nil {
	// 	log.Fatal("Error sending email:", err)
	// } else {
	// 	log.Println("Email sent successfully!")
	// }

	// from := email.From
	// to := email.To
	// from := email.Headers.Get("from")
	// to := email.Headers.Get("to")
	// subj := email.Headers.Get("subject")
	// body := email.Body

	// // Setup headers
	// headers := make(map[string]string)
	// headers["From"] = from
	// headers["To"] = to
	// headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range email.Headers.Map() {
		fmt.Println(fmt.Sprintf("%s: %s\r\n", k, v))
		message += fmt.Sprintf("%s: %s\r\n", k, v)

	}
	message += "\r\n" + email.Body

	// Connect to the SMTP Server
	// servername := "smtp.example.tld:465"
	servername := config.Host + ":" + fmt.Sprint(config.Port)
	// servername := config.Host + ":" + string(config.Port)

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	fmt.Println("address", servername)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	// if err = c.Mail(from.Address); err != nil {
	if err = c.Mail("m@mireksirina.cz"); err != nil {
		log.Panic(err)
	}

	// if err = c.Rcpt(to.Address); err != nil {
	if err = c.Rcpt("mirekovec@gmail.com"); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

}
