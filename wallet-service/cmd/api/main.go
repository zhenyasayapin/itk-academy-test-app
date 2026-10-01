package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port string
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bonjour tout le monde !")
}

func main() {
	app := createConfig()

	http.HandleFunc("/", hello)

	log.Printf("Server is strating on port :%s", app.Port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", app.Port), nil)

	if err != nil {
		log.Fatal(err)
	}
}

func createConfig() Config {

	return Config{
		Port: os.Getenv("WEB_SERVER_PORT"),
	}
}
