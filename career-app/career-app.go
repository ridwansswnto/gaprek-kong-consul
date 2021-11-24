package main

import (
	"log"
	"net/http"
)

func career(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Haloo Ini Halaman Karir, Kita lagi open rekrutmen lhoo!!!"))
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Haloo Ini BLOG kita, ini berisi catatan-catatan engineer kita lhoo!!"))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/career", career)
	mux.HandleFunc("/blog", blog)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}