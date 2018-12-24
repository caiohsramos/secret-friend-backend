package main

import (
	"log"
	"net/smtp"
	"os"
)

// Send is a email sender
func Send(body, to string) {
	from := "noreply.amigo.secreto@gmail.com"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Seu sorteio do amigo secreto!\n\n" +
		body
	password := os.Getenv("GMAIL_PASSWORD")
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", "caioramos97@gmail.com", password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Printf("sent to %s", to)
}
