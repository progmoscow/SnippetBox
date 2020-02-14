package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Домашняя страница"))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on : 4000")
	err := http.ListenAndServe(":4000", mux)

	if(err != nil){
		log.Fatal(err)
	}
}
