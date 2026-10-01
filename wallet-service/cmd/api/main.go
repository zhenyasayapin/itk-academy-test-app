package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port string
	DSN  string
	DB   DatabaseRepo
}

func main() {
	app, err := createConfig()
	if err != nil {
		log.Fatal("Failed to create config:", err)
	}

	log.Printf("Server is strating on port :%s", app.Port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", app.Port), app.routes())

	if err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

func createConfig() (*Config, error) {
	app := Config{
		Port: os.Getenv("WEB_SERVER_PORT"),
		DSN:  os.Getenv("POSTGRES_DSN"),
	}

	conn, err := app.connectToDB()
	if err != nil {
		return nil, err
	}

	app.DB = &PostgresDBRepo{DB: conn}

	return &app, nil
}
