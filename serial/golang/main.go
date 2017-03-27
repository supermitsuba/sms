package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"encoding/json"

	"github.com/jacobsa/go-serial/serial"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

type receiveMessageFunc func(<-chan amqp.Delivery)

func main() {
	fmt.Println(os.Args)
	amqpURL := os.Args[1]

	listenForMesages(amqpURL, "messages", func(msgs <-chan amqp.Delivery) {
		for d := range msgs {
			receivedMessage(d.Body)
			d.Ack(false)
			log.Printf("Done")
		}
	})

	forever := make(chan bool)
	<-forever
}

func listenForMesages(amqpURL string, queueName string, myFunc receiveMessageFunc) {
	conn, err := amqp.Dial(amqpURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	myFunc(msgs)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}

func receivedMessage(message []byte) {
	var m Message
	json.Unmarshal(message, &m)
	log.Printf("Received a message: %s", m.Text)
	derp(" " + m.Text)
	timeout := time.Duration(m.Duration) * time.Second
	time.Sleep(timeout)
	derp(" ")
}

func derp(message string) {
	serialPort := os.Args[2]
	if serialPort == "console" {
		log.Printf("This is printing just to console. Message: %s", message)
	} else {
		log.Printf("The USB port is: %s", serialPort)
		options := serial.OpenOptions{
			PortName:        serialPort,
			BaudRate:        9600,
			DataBits:        8,
			StopBits:        1,
			MinimumReadSize: 4,
		}

		s, err := serial.Open(options)
		if err != nil {
			log.Fatal(err)
		}

		n, err := s.Write([]byte(message))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", n)

		s.Close()
	}
}

type Message struct {
	Duration int
	Text     string
}
