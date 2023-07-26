package main

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// try to connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	// start listening for messages, as we get them dirrectly from the queue
	log.Println("listening for and consummitng RabbitMq messasges")

	// create consummer, consumes messages from the queue
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// watch the queue
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERRROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// wait till rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("not yet ready - RabbitMq")
			counts++
		} else {
			connection = c
			log.Println("Connected!")
			break
		}

		// smth is wrong
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Printf("backing off ...[%d]\n", backOff)
		time.Sleep(backOff)
	}

	return connection, nil
}
