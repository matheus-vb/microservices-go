package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/matheus-vb/microservices-go/logger-service/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	PORT     = "80"
	rpcPORT  = "5001"
	mongoURL = "mongodb://mongo:27017"
	grpcPORT = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	log.Println("Starting logger service")

	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	app.setupServer()
}

func (app *Config) setupServer() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.setRoutes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panicln(err)
	}
}

func connectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
