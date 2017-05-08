package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

func main() { // ./main 5000, "containers_weather_timer_1", "containers_forecast_timer_1"

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/test", Test).Methods("GET")
	router.HandleFunc("/api/weather/current", weatherCurrentFunc).Methods("POST")
	router.HandleFunc("/api/weather/forecast", weatherForecastFunc).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+os.Args[1], router))
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Working...")
}

func weatherCurrentFunc(w http.ResponseWriter, r *http.Request) {
	currentWeatherContainerName := os.Args[2]
	cmd := exec.Command("/bin/bash", "-c", "docker start "+currentWeatherContainerName)
	if output, err := cmd.Output(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: " + err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK: " + string(output)))
	}

	return
}

func weatherForecastFunc(w http.ResponseWriter, r *http.Request) {
	forecastWeatherContainerName := os.Args[3]
	cmd := exec.Command("/bin/bash", "-c", "docker start "+forecastWeatherContainerName)
	if output, err := cmd.Output(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: " + err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK: " + string(output)))
	}

	return
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
