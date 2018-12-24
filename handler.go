package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

// GetAllSecertFriend retrieves all SecretFriends
func GetAllSecertFriend(w http.ResponseWriter, r *http.Request) {
	sf, err := FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, sf)
}

// GetSecretFriend retrieves ones SecretFriend
func GetSecretFriend(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sf, err := FindByID(params["ID"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Secret Friend ID")
		return
	}
	respondWithJSON(w, http.StatusOK, sf)
}

// CreateSecretFriend adds one secret friend
func CreateSecretFriend(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var sf SecretFriend
	if err := json.NewDecoder(r.Body).Decode(&sf); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	sf.ID = bson.NewObjectId()
	if err := Insert(sf); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, sf)
}

// DeleteSecretFriend is a handlers tha deletes a Secret-Friend
func DeleteSecretFriend(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := DeleteByID(params["ID"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Secret Friend ID")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"msg": "Deleted!"})
}

// RandSecretFriend combines and sends emails to the guests
func RandSecretFriend(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sf, err := FindByID(params["ID"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Secret Friend ID")
		return
	}
	sf.RandEmailSend()
	respondWithJSON(w, http.StatusOK, map[string]string{"msg": "Emails sent!"})
}
