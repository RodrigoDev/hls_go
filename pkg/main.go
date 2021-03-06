package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	server := &http.Server{
		Handler:      handlers(),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Starting server at 0.0.0.0:8080")
	log.Fatal(server.ListenAndServe())
}

func handlers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", videoEmbedHandler).Methods("GET")
	router.HandleFunc("/media/{mId:[0-9]+}/stream/", streamHandler).Methods("GET")
	router.HandleFunc("/media/{mId:[0-9]+}/stream/{segName:index[0-9]+.ts}", streamHandler).Methods("GET")

	router.HandleFunc("/upload", uploadFormHandler).Methods("GET")
	router.HandleFunc("/video/upload/", uploadHandler).Methods("POST")

	return router
}
