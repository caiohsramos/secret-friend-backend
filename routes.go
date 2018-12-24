package main

import "net/http"

// Route defines a http route
type Route struct {
	Methods []string
	Path    string
	Handler http.HandlerFunc
}

// Routes is an slice of Route's
type Routes []Route

var routes = Routes{
	Route{
		Methods: []string{"GET", "OPTIONS"},
		Path:    "/",
		Handler: Logger(Cors(GetAllSecertFriend)),
	},
	Route{
		Methods: []string{"GET", "OPTIONS"},
		Path:    "/{ID}",
		Handler: Logger(Cors(GetSecretFriend)),
	},
	Route{
		Methods: []string{"POST", "OPTIONS"},
		Path:    "/",
		Handler: Logger(Cors(CreateSecretFriend)),
	},
	Route{
		Methods: []string{"GET", "OPTIONS"},
		Path:    "/{ID}/rand",
		Handler: Logger(Cors(RandSecretFriend)),
	},
	Route{
		Methods: []string{"DELETE", "OPTIONS"},
		Path:    "/{ID}",
		Handler: Logger(Cors(DeleteSecretFriend)),
	},
}
