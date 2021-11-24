package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Haloo Ini Halaman Home, ###SELAMAT DATANG IR01"))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Starting server on :8070")
	err := http.ListenAndServe(":8070", mux)
	log.Fatal(err)
}