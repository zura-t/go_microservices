package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	Mailer Mail
}

const webPort = "80"

func main() {
	app := Config{
		Mailer: createMail(),
	}

	log.Printf("Starting service on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func createMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAILER_PORT"))
	m := Mail{
		Domain:      os.Getenv("MAILER_DOMAIN"),
		Host:        os.Getenv("MAILER_HOST"),
		Port:        port,
		Username:    os.Getenv("MAILER_USERNAME"),
		Password:    os.Getenv("MAILER_PASSWORD"),
		Encryption:  os.Getenv("MAILER_ENCRYPTION"),
		FromName:    os.Getenv("FROM_NAME"),
		FromAddress: os.Getenv("FROM_ADDRESS"),
	}

	return m
}
