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

func main() {
	app := createConfig()

	log.Printf("Server is strating on port :%s", app.Port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", app.Port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}

func createConfig() Config {

	return Config{
		Port: os.Getenv("WEB_SERVER_PORT"),
	}
}
