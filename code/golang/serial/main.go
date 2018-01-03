package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"encoding/json"

	"github.com/supermitsuba/amqp"
	"github.com/supermitsuba/go-serial/serial"
)

// used to capture errors and print it to console
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// func for receiving a message from message queue
type receiveMessageFunc func(<-chan amqp.Delivery)

func main() {
	fmt.Println(os.Args)
	amqpURL := os.Args[1]   // url to the amqp
	statusUrl := os.Args[3] // url to the status

	listenForMesages(amqpURL, "messages", func(msgs <-chan amqp.Delivery) {
		for d := range msgs { // foreach message
			status := GetStatus(statusUrl)
			// Get the status for the led
			// if it is false, do not display anything on the led
			if status.IsLEDActive != false {
				receivedMessage(d.Body) // display message
			}
			d.Ack(false) // acknowledge the message on the queue
			log.Printf("Done")
		}
	})

	forever := make(chan bool) // listen forever
	<-forever
}

// listen for messages on the queue
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

// unmarshal message and display message on led
func receivedMessage(message []byte) {
	var m Message
	json.Unmarshal(message, &m)
	log.Printf("Received a message: %s", m.Text)
	displayMessage(" " + m.Text)
	timeout := time.Duration(m.Duration) * time.Second
	time.Sleep(timeout)
	displayMessage(" ")
}

// if console, display in console for debugging
// otherwise display message to led sign on the serial port
func displayMessage(message string) {
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

// Call url and deserialize json into LEDStatus
func GetStatus(URL string) LEDStatus {
	var s LEDStatus
	resp, err := http.Get(URL)
	failOnError(err, "Could not send http request.")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	failOnError(err, "Could not send http request.")

	err2 := json.Unmarshal(body, &s)
	failOnError(err2, "Could not send http request.")

	log.Printf("Get Body status '%v' = s '%v'", string(body), s.IsLEDActive)
	return s
}

type LEDStatus struct {
	IsLEDActive bool `json:"isLEDActive"`
}

type Message struct {
	Duration int
	Text     string
}
