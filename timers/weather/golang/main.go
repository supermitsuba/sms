package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	message := "{ \"duration\":30, \"text\":\"" + GetWeather() + "\" }"
	log.Printf("Here is the message: " + message)
	UpdateMessage(message)
}

type WeatherData struct {
	Weather Conditions `json:"weather"`
	Main    Stats      `json:"main"`
}

type Conditions struct {
	Main string `json:"main"`
}

type Stats struct {
	Temp float64 `json:"temp"`
}

func GetWeather() string {
	weatherURL := os.Args[2]
	resp, err := http.Get(weatherURL)
	failOnError(err, "Could not send weather get http request.")
	io.Copy(os.Stdout, resp.Body)

	m := &WeatherData{}
	err1 := json.NewDecoder(resp.Body).Decode(&m)
	failOnError(err1, "Could not parse json.")

	log.Printf("Got the parsed json: %s", m)
	temperature := (m.Main.Temp * 9 / 5) - 459.67
	conditions := m.Weather.Main

	return "Now:     " + strconv.FormatFloat(temperature, 'f', 1, 64) + " F " + conditions
}

func UpdateMessage(message string) {
	b := new(bytes.Buffer)
	b.WriteString(message)

	log.Printf("The message is: %s", message)

	URL := os.Args[1]
	resp, err := http.Post(URL, "application/json", b)
	failOnError(err, "Could not send post http request.")
	io.Copy(os.Stdout, resp.Body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
