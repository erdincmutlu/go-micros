package main

import (
	"fmt"
	"log"
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
	log.Println("Connected to RabbitMQ")

	// start listening for messages

	// create consumer

	// watch the queue and consume events
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second

	// don't continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			return c, nil
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		log.Println("Backing off...")
		time.Sleep(backOff)
		backOff *= 2
	}
}
