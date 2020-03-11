package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Take dump
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	dyno := os.Getenv("DYNO")
	fmt.Println("DYNO:", dyno)

	fmt.Println(string(dump))
	fmt.Fprintf(w, "go!\n")
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	d, err := ioutil.ReadFile("assets/main.css")

	if err != nil {
		fmt.Fprintf(w, "No such file\n")
	} else {
		fmt.Fprintf(w, string(d))
	}
}

func main() {
	http.HandleFunc("/css", cssHandler)
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Running on:", port, "...")

	http.ListenAndServe(":"+port, nil)
}
