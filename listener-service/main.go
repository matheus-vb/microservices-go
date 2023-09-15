package main

import (
	"log"
	"math"
	"os"
	"time"

	"github.com/matheus-vb/microservices-go/listener-service/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitConn, err := connectToMQ()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	log.Println("Connected.")
	//listen for messaages

	//create consumer
	consumer, err := event.SetupNewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	//watch queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connectToMQ() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		conn, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			log.Println("RabbitMQ not ready...")
			counts++
		} else {
			connection = conn
			break
		}

		if counts > 5 {
			log.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off...")

		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
