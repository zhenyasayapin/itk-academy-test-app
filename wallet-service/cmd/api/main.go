package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bonjour tout le monde !")
}

func main() {
	http.HandleFunc("/", hello)

	log.Println("Server is strating on port", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
