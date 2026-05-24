package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	if dyno := os.Getenv("DYNO"); dyno != "" {
		log.Println("DYNO:", dyno)
	}
	log.Println(string(dump))

	fmt.Fprintf(w, "go!\n")
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	d, err := os.ReadFile("assets/main.css")
	if err != nil {
		http.Error(w, "asset not found", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/css")
	w.Write(d)
}

func main() {
	http.HandleFunc("/css", cssHandler)
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Running on:", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
