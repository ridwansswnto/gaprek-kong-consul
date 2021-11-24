package main

import (
	"log"
	"net/http"
)

func product(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Haloo Ini halaman PRODUCT, KITA JUALAN BANYAK JANJI LHOO"))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/product", product)
	log.Println("Starting server on :8090")
	err := http.ListenAndServe(":8090", mux)
	log.Fatal(err)
}