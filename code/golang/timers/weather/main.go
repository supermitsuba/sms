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
	weather := GetWeather(os.Args[2])
	// temp1 := (weather.Main.Temperature * 9 / 5) - 459.17
	temp := strconv.FormatFloat(weather.Main.Temperature, 'f', 0, 64)
	message := Message{
		Text:     "Now:     " + temp + " F   " + weather.WeatherSections[0].Main,
		Duration: 30,
	}
	SendMessage(message, os.Args[1])
	fmt.Println("Getting Weather: ", weather.WeatherSections[0].Main, weather.Main.Temperature)

}

func GetWeather(URL string) WeatherModel {
	res, err := http.Get(URL)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var s = new(WeatherModel)
	err1 := json.Unmarshal(body, &s)
	if err1 != nil {
		fmt.Println("whoops:", err1)
	}

	return *s
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

type WeatherModel struct {
	WeatherSections []WeatherSection `json:"weather"`
	Main            MainSection      `json:"main"`
}

type WeatherSection struct {
	Main string `json:"main"`
}

type MainSection struct {
	Temperature float64 `json:"temp"`
}
