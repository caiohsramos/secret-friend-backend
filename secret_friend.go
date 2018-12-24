package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/globalsign/mgo/bson"
)

// SecretFriend defines a MongoDB document
type SecretFriend struct {
	ID     bson.ObjectId `bson:"_id" json:"_id"`
	Name   string        `bson:"name" json:"name"`
	Guests []Guest       `bson:"guests" json:"guests"`
}

// Guest define a single Guest of a secret friend
type Guest struct {
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}

// RandEmailSend is a random choice of the Guests
func (sf SecretFriend) RandEmailSend() {
	guestNumber := len(sf.Guests)
	shuffle := shuffleNoRep(guestNumber)

	sampleMessage := `Olá %s, sorteamos seu amigo secreto %s de ID %s! Aqui estão as informações dele:
    Nome: %s
    Email: %s`
	// Use shuffle to send the emails
	for idx, value := range shuffle {
		from := sf.Guests[idx]
		to := sf.Guests[value]
		message := fmt.Sprintf(sampleMessage, from.Name, sf.Name, sf.ID.Hex(), to.Name, to.Email)

		Send(message, from.Email)
	}
}

func shuffleNoRep(n int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	var shuffle []int
	good := false
	for !good {
		good = true
		shuffle = r.Perm(n)
		// Assert index != value
		for idx, value := range shuffle {
			if idx == value {
				good = false
			}
		}
	}

	return shuffle
}
