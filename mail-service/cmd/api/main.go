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
		Mailer: setupMail(),
	}

	log.Println("Mail service running")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.setRoutes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func setupMail() Mail {
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		log.Panic(err)
	}

	mail := Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		FromName:    os.Getenv("FROM_NAME"),
		FromAddress: os.Getenv("FROM_ADDRESS"),
	}

	return mail
}
