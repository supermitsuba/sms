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
	"strings"
	"time"
)

func main() {
	weather := GetWeather(os.Args[2])

	for i := 0; i < len(weather.WeatherForcast); i++ {
		item := weather.WeatherForcast[i]

		splitDateTime := strings.Split(item.DateOfTemperature, " ")
		splitHour := strings.Split(splitDateTime[1], ":")[0]
		if splitHour == "15" {
			t1, err := time.Parse("2006-01-02", splitDateTime[0])
			failOnError(err, "Could not parse datetime properly.")
			dayOfTheWeek := t1.Format("Mon")
			temp := (item.Main.Temperature * 9 / 5) - 459.17
			temp1 := strconv.FormatFloat(temp, 'f', 0, 64)
			message := Message{
				Text:     dayOfTheWeek + ":       " + temp1 + " F " + item.WeatherSections[0].Main,
				Duration: 30,
			}

			fmt.Println("Getting Weather: ", message)
			SendMessage(message, os.Args[1])
		}
	}
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

type Message struct {
	Text     string  `json:"text"`
	Duration float64 `json:"duration"`
}

type WeatherModel struct {
	WeatherForcast []ListWeather `json:"list"`
}

type ListWeather struct {
	WeatherSections   []WeatherSection `json:"weather"`
	Main              MainSection      `json:"main"`
	DateOfTemperature string           `json:"dt_txt"`
}

type WeatherSection struct {
	Main string `json:"main"`
}

type MainSection struct {
	Temperature float64 `json:"temp"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
