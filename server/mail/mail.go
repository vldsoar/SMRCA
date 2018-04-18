package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"os"
	"encoding/json"
)

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func New(to string, body string) Mail {
	mail := Mail{}
	mail.senderId = ""
	mail.toIds = []string{to}
	mail.subject = "[AVISO] - SMCRA"
	mail.body = body

	return mail
}

func ConnectAndSend(mail Mail)  {
	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}

	f, err := os.Open("./mail/config.json")

	if err != nil {
		panic(err)
		return
	}

	var conf = struct {
		Email string
		Pass string
	}{}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&conf)

	mail.senderId = conf.Email
	messageBody := mail.BuildMessage()

	log.Println(smtpServer.host)
	//build an auth
	auth := smtp.PlainAuth("", mail.senderId, conf.Pass, smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		log.Panic(err)
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Email enviado com sucesso")
}