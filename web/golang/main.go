package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

type Message struct {
	Duration int    `json:"duration"` // 1 to 60 validation
	Text     string `json:"text"`     //32 max validation
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/api/test", Test).Methods("GET")
	router.HandleFunc("/api/message", MessageFunc).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+os.Args[1], router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	filename := "index.html"
	body, _ := ioutil.ReadFile(filename)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, string(body))
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Working...")
}

func MessageFunc(w http.ResponseWriter, r *http.Request) {
	var priority = r.URL.Query().Get("priority")
	var item = new(Message)
	json.NewDecoder(r.Body).Decode(item)

	if priority == "" {
		priority = "false"
	}

	var isBool, err = strconv.ParseBool(priority)
	if err != nil {
		http.Error(w, "Invalid priority", 400)
		return
	}

	if item.Duration < 1 || item.Duration > 60 {
		http.Error(w, "Invalid duration", 400)
		return
	}

	if len(item.Text) > 32 {
		http.Error(w, "Invalid text", 400)
		return
	}

	if isBool {
		str, err := json.Marshal(item)
		if err != nil {
			http.Error(w, "Invalid item", 400)
			return
		}
		SendMessage(os.Args[2], "priority", str)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK Priority"))
	} else {
		str, err := json.Marshal(item)
		if err != nil {
			http.Error(w, "Invalid item", 400)
			return
		}
		SendMessage(os.Args[2], "messages", str)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}

	return
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func SendMessage(amqpUrl string, queueName string, body []byte) {
	conn, err := amqp.Dial(amqpUrl)
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

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
