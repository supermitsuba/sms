package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	message := Message{
		Text:     os.Args[2],
		Duration: 30,
	}
	SendMessage(message, os.Args[1])

}

func SendMessage(message Message, URL string) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(message)
	resp, err := http.Post(URL, "application/json", b)
	failOnError(err, "Could not send http request.")
	io.Copy(os.Stdout, resp.Body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

type Message struct {
	Text     string  `json:"text"`
	Duration float64 `json:"duration"`
}
