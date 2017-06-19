package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

type Message struct {
	Duration int    `json:"duration"` // 1 to 60 validation
	Text     string `json:"text"`     //32 max validation
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/api/test", Test).Methods("GET")
	router.HandleFunc("/api/message", MessageFunc).Methods("POST")
	router.HandleFunc("/api/weather", WeatherFunc).Methods("POST")
	router.HandleFunc("/api/forecast", ForecastFunc).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))
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
		SendMessage(os.Args[1], "priority", str)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK Priority"))
	} else {
		str, err := json.Marshal(item)
		if err != nil {
			http.Error(w, "Invalid item", 400)
			return
		}
		SendMessage(os.Args[1], "messages", str)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}

	return
}

func WeatherFunc(w http.ResponseWriter, r *http.Request) {
	weather := GetWeather(os.Args[2])
	temp1 := (weather.Main.Temperature * 9 / 5) - 459.17
	temp := strconv.FormatFloat(temp1, 'f', 0, 64)
	var item = new(Message)
	item.Text = "Now:     " + temp + " F   " + weather.WeatherSections[0].Main
	item.Duration = 15

	str, err := json.Marshal(item)
	if err != nil {
		http.Error(w, "Invalid item", 400)
		return
	}
	SendMessage(os.Args[1], "messages", str)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

	return
}

func ForecastFunc(w http.ResponseWriter, r *http.Request) {
	weather := GetForecast(os.Args[3])

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
			var myMessage = new(Message)
			myMessage.Text = dayOfTheWeek + ":       " + temp1 + " F " + item.WeatherSections[0].Main
			myMessage.Duration = 15

			str, err := json.Marshal(myMessage)
			if err != nil {
				http.Error(w, "Invalid item", 400)
				return
			}
			SendMessage(os.Args[1], "messages", str)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
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

func GetForecast(URL string) ForecastModel {
	res, err := http.Get(URL)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var s = new(ForecastModel)
	err1 := json.Unmarshal(body, &s)
	if err1 != nil {
		fmt.Println("whoops:", err1)
	}

	return *s
}

type WeatherModel struct {
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

type ForecastModel struct {
	WeatherForcast []WeatherModel `json:"list"`
}
