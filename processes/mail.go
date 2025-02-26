package processes // import "odk_mailer/processes"

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"odk_mailer/models"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(message models.Message, job models.Job) bool {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	pass := os.Getenv("SMTP_PASS")

	from := mail.Address{Name: "Mailer", Address: job.Sender}
	to := mail.Address{Name: "", Address: message.Recipient}
	subject := message.Subject

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	messageBody := ""

	for k, v := range headers {
		messageBody += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	messageBody += "\r\n" + message.Content
	serverName := smtpHost + ":" + smtpPort

	host, _, _ := net.SplitHostPort(serverName)

	auth := smtp.PlainAuth("", from.Address, pass, smtpHost)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	con, err := tls.Dial("tcp", serverName, tlsconfig)
	if err != nil {
		panic(err)
	}

	c, err := smtp.NewClient(con, host)
	if err != nil {
		panic(err)
	}

	if err = c.Auth(auth); err != nil {
		panic(err)
	}

	if err = c.Mail(from.Address); err != nil {
		panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		panic(err)
	}

	w, err := c.Data()
	if err != nil {
		panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		panic(err)
	}

	err = w.Close()
	if err != nil {
		panic(err)
	}

	c.Close()

	return true
}
