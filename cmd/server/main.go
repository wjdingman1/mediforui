package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("hey")

	r := mux.NewRouter()
	r.HandleFunc("/", Handler)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	log.Println("Starting API Server")
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}

}

// Handler does a thing
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("hit endpoint")
	w.Write([]byte("Gorilla!\n"))
}
