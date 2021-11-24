package main

import (
	"log"
	"net/http"
)

func product(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halaman PRODUCT!!! AYOO KITA SERBUUUU"))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/product", product)
	log.Println("Starting server on :800")
	err := http.ListenAndServe(":8070", mux)
	log.Fatal(err)
}