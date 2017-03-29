package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	URL := os.Args[1]
	t := time.Now()
	message := "{ \"duration\":30, \"text\":\"Time: " + t.Format("03:04PM") + "   Date: " + t.Format("01/02/06") + "\" }"
	b := new(bytes.Buffer)
	b.WriteString(message)

	log.Printf("The message is: %s", message)
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
