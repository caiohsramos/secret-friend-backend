package main

import (
	"os"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var db *mgo.Database

// InitDB creates and test a MongoDB connection
func InitDB() *mgo.Database {
	connString := os.Getenv("MONGO_CONNECTION")
	session, _ := mgo.Dial(connString)
	db = session.DB("main")
	return db
}

// FindAll return all the SecretFriends
func FindAll() ([]SecretFriend, error) {
	var sf []SecretFriend
	err := db.C("secret-friend").Find(bson.M{}).All(&sf)

	return sf, err
}

// FindByID returns the SecretFriend based on id
func FindByID(id string) (SecretFriend, error) {
	var sf SecretFriend
	err := db.C("secret-friend").FindId(bson.ObjectIdHex(id)).One(&sf)
	return sf, err
}

// DeleteByID deletes a record based on a ID
func DeleteByID(id string) error {
	err := db.C("secret-friend").RemoveId(bson.ObjectIdHex(id))

	return err
}

// Insert puts a new secret friend in the collection
func Insert(sf SecretFriend) error {
	err := db.C("secret-friend").Insert(&sf)
	return err
}
