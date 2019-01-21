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

	. "github.com/supermitsuba/go-linq"
)

func main() {
	weather := GetWeather(os.Args[2])
	//weather := GetWeatherTest()
	defaultTakeNumber := 5

	myList := []Group{}

	fmt.Println("--------", weather)
	From(weather.WeatherForcast).
		GroupByT(func(item ListWeather) string { // key
			splitDateTime := strings.Split(item.DateOfTemperature, " ")
			return splitDateTime[0]
		}, func(listWeather ListWeather) ListWeather { // value
			return listWeather
		}).
		OrderByDescendingT(func(dateOfTemp Group) int64 {
			stringTime := (dateOfTemp.Key).(string)
			t1, _ := time.Parse("2006-01-02", stringTime)
			return -(t1.Unix())
		}).
		Take(defaultTakeNumber).
		ToSlice(&myList)

	fmt.Println("--------", len(myList))
	for i := 0; i < len(myList); i++ {

		fmt.Println("--------")
		t1, _ := time.Parse("2006-01-02", myList[i].Key.(string))
		dayOfTheWeek := t1.Format("Mon")

		MaxTemp := From(myList[i].Group).
			SelectT(func(x ListWeather) float64 { return x.Main.Temperature }).
			Max()

		MinTemp := From(myList[i].Group).
			SelectT(func(x ListWeather) float64 { return x.Main.Temperature }).
			Min()

		Conditions := ""
		hasCondition := From(myList[i].Group).
			AnyWithT(func(x ListWeather) bool {
				splitDateTime := strings.Split(x.DateOfTemperature, " ")
				splitHour := strings.Split(splitDateTime[1], ":")[0]
				return splitHour == "15"
			})

		if hasCondition {
			Conditions = From(myList[i].Group).
				FirstWithT(func(x ListWeather) bool {
					splitDateTime := strings.Split(x.DateOfTemperature, " ")
					splitHour := strings.Split(splitDateTime[1], ":")[0]
					return splitHour == "15"
				}).(ListWeather).WeatherSections[0].Main
		}

		temp := dayOfTheWeek + " H " + ConvertTempToFahrenheit(MaxTemp.(float64)) + " L " + ConvertTempToFahrenheit(MinTemp.(float64)) + "   " + Conditions

		myMessage := Message{Text: temp, Duration: 15}
		fmt.Println(temp)

		SendMessage(myMessage, os.Args[1])
	}
}

func ConvertTempToFahrenheit(tempCelsius float64) string {
	temp := (tempCelsius * 9 / 5) - 459.17
	return strconv.FormatFloat(tempCelsius, 'f', 0, 64)
	// use this if you are doing Kelvin
	// strconv.FormatFloat(temp, 'f', 0, 64)
}

func GetWeatherTest() WeatherModel {
	b := []byte(`{"cod":"200","message":0.0049,"cnt":35,"list":[{"dt":1499353200,"main":{"temp":299.72,"temp_min":299.72,"temp_max":301.451,"pressure":998.36,"sea_level":1029.2,"grnd_level":998.36,"humidity":45,"temp_kf":-1.73},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"clouds":{"all":24},"wind":{"speed":2.36,"deg":222.501},"sys":{"pod":"d"},"dt_txt":"2017-07-06 15:00:00"},{"dt":1499364000,"main":{"temp":302.71,"temp_min":302.71,"temp_max":304.003,"pressure":996.47,"sea_level":1027.13,"grnd_level":996.47,"humidity":40,"temp_kf":-1.3},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03d"}],"clouds":{"all":32},"wind":{"speed":3.01,"deg":220.501},"sys":{"pod":"d"},"dt_txt":"2017-07-06 18:00:00"},{"dt":1499374800,"main":{"temp":303.79,"temp_min":303.79,"temp_max":304.654,"pressure":995.18,"sea_level":1025.75,"grnd_level":995.18,"humidity":34,"temp_kf":-0.86},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"clouds":{"all":12},"wind":{"speed":4.01,"deg":226.506},"sys":{"pod":"d"},"dt_txt":"2017-07-06 21:00:00"},{"dt":1499385600,"main":{"temp":302.77,"temp_min":302.77,"temp_max":303.205,"pressure":993.52,"sea_level":1024.13,"grnd_level":993.52,"humidity":33,"temp_kf":-0.43},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"02n"}],"clouds":{"all":8},"wind":{"speed":3.91,"deg":224.502},"sys":{"pod":"n"},"dt_txt":"2017-07-07 00:00:00"},{"dt":1499396400,"main":{"temp":299.12,"temp_min":299.12,"temp_max":299.12,"pressure":992.63,"sea_level":1023.35,"grnd_level":992.63,"humidity":42,"temp_kf":0},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02n"}],"clouds":{"all":12},"wind":{"speed":4.56,"deg":211.502},"sys":{"pod":"n"},"dt_txt":"2017-07-07 03:00:00"},{"dt":1499407200,"main":{"temp":297.696,"temp_min":297.696,"temp_max":297.696,"pressure":991.38,"sea_level":1022.23,"grnd_level":991.38,"humidity":54,"temp_kf":0},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02n"}],"clouds":{"all":24},"wind":{"speed":5.52,"deg":230.005},"sys":{"pod":"n"},"dt_txt":"2017-07-07 06:00:00"},{"dt":1499418000,"main":{"temp":292.75,"temp_min":292.75,"temp_max":292.75,"pressure":991.07,"sea_level":1021.88,"grnd_level":991.07,"humidity":90,"temp_kf":0},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10n"}],"clouds":{"all":56},"wind":{"speed":3.12,"deg":274.5},"rain":{"3h":2.11},"sys":{"pod":"n"},"dt_txt":"2017-07-07 09:00:00"},{"dt":1499428800,"main":{"temp":292.486,"temp_min":292.486,"temp_max":292.486,"pressure":991.54,"sea_level":1022.56,"grnd_level":991.54,"humidity":97,"temp_kf":0},"weather":[{"id":501,"main":"Rain","description":"moderate rain","icon":"10d"}],"clouds":{"all":80},"wind":{"speed":3.8,"deg":302.005},"rain":{"3h":7.645},"sys":{"pod":"d"},"dt_txt":"2017-07-07 12:00:00"},{"dt":1499439600,"main":{"temp":294.276,"temp_min":294.276,"temp_max":294.276,"pressure":992.7,"sea_level":1023.55,"grnd_level":992.7,"humidity":90,"temp_kf":0},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10d"}],"clouds":{"all":36},"wind":{"speed":4.46,"deg":347.506},"rain":{"3h":0.02},"sys":{"pod":"d"},"dt_txt":"2017-07-07 15:00:00"},{"dt":1499450400,"main":{"temp":296.84,"temp_min":296.84,"temp_max":296.84,"pressure":993.1,"sea_level":1023.75,"grnd_level":993.1,"humidity":83,"temp_kf":0},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"clouds":{"all":20},"wind":{"speed":3.91,"deg":354.502},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-07 18:00:00"},{"dt":1499461200,"main":{"temp":296.821,"temp_min":296.821,"temp_max":296.821,"pressure":993.11,"sea_level":1023.87,"grnd_level":993.11,"humidity":74,"temp_kf":0},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"clouds":{"all":92},"wind":{"speed":4.22,"deg":344.501},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-07 21:00:00"},{"dt":1499472000,"main":{"temp":296.045,"temp_min":296.045,"temp_max":296.045,"pressure":993.92,"sea_level":1024.76,"grnd_level":993.92,"humidity":66,"temp_kf":0},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10n"}],"clouds":{"all":92},"wind":{"speed":3.32,"deg":357.001},"rain":{"3h":0.145},"sys":{"pod":"n"},"dt_txt":"2017-07-08 00:00:00"},{"dt":1499482800,"main":{"temp":292.761,"temp_min":292.761,"temp_max":292.761,"pressure":995.41,"sea_level":1026.56,"grnd_level":995.41,"humidity":76,"temp_kf":0},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10n"}],"clouds":{"all":44},"wind":{"speed":3.77,"deg":19.5115},"rain":{"3h":0.07},"sys":{"pod":"n"},"dt_txt":"2017-07-08 03:00:00"},{"dt":1499493600,"main":{"temp":289.306,"temp_min":289.306,"temp_max":289.306,"pressure":995.53,"sea_level":1026.86,"grnd_level":995.53,"humidity":92,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.22,"deg":353.5},"rain":{},"sys":{"pod":"n"},"dt_txt":"2017-07-08 06:00:00"},{"dt":1499504400,"main":{"temp":287.725,"temp_min":287.725,"temp_max":287.725,"pressure":996.04,"sea_level":1027.41,"grnd_level":996.04,"humidity":94,"temp_kf":0},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03n"}],"clouds":{"all":36},"wind":{"speed":1.66,"deg":283.501},"rain":{},"sys":{"pod":"n"},"dt_txt":"2017-07-08 09:00:00"},{"dt":1499515200,"main":{"temp":289.598,"temp_min":289.598,"temp_max":289.598,"pressure":996.77,"sea_level":1028.02,"grnd_level":996.77,"humidity":83,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":1.86,"deg":321.002},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-08 12:00:00"},{"dt":1499526000,"main":{"temp":294.323,"temp_min":294.323,"temp_max":294.323,"pressure":997.07,"sea_level":1028.21,"grnd_level":997.07,"humidity":70,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":2.01,"deg":301.502},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-08 15:00:00"},{"dt":1499536800,"main":{"temp":296.871,"temp_min":296.871,"temp_max":296.871,"pressure":996.59,"sea_level":1027.6,"grnd_level":996.59,"humidity":56,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":2.76,"deg":300.5},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-08 18:00:00"},{"dt":1499547600,"main":{"temp":298.104,"temp_min":298.104,"temp_max":298.104,"pressure":995.81,"sea_level":1026.78,"grnd_level":995.81,"humidity":42,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":3.65,"deg":294.5},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-08 21:00:00"},{"dt":1499558400,"main":{"temp":297.193,"temp_min":297.193,"temp_max":297.193,"pressure":995.38,"sea_level":1026.25,"grnd_level":995.38,"humidity":37,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":3.16,"deg":288.002},"rain":{},"sys":{"pod":"n"},"dt_txt":"2017-07-09 00:00:00"},{"dt":1499569200,"main":{"temp":291.174,"temp_min":291.174,"temp_max":291.174,"pressure":995.28,"sea_level":1026.29,"grnd_level":995.28,"humidity":58,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.62,"deg":235.504},"rain":{},"sys":{"pod":"n"},"dt_txt":"2017-07-09 03:00:00"},{"dt":1499580000,"main":{"temp":289.408,"temp_min":289.408,"temp_max":289.408,"pressure":994.72,"sea_level":1025.78,"grnd_level":994.72,"humidity":67,"temp_kf":0},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10n"}],"clouds":{"all":80},"wind":{"speed":1.96,"deg":205.001},"rain":{"3h":0.135},"sys":{"pod":"n"},"dt_txt":"2017-07-09 06:00:00"},{"dt":1499590800,"main":{"temp":290.23,"temp_min":290.23,"temp_max":290.23,"pressure":994.38,"sea_level":1025.58,"grnd_level":994.38,"humidity":90,"temp_kf":0},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10n"}],"clouds":{"all":36},"wind":{"speed":1.57,"deg":233.5},"rain":{"3h":0.745},"sys":{"pod":"n"},"dt_txt":"2017-07-09 09:00:00"},{"dt":1499601600,"main":{"temp":291.533,"temp_min":291.533,"temp_max":291.533,"pressure":995.19,"sea_level":1026.32,"grnd_level":995.19,"humidity":89,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":3.31,"deg":287.004},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-09 12:00:00"},{"dt":1499612400,"main":{"temp":297.552,"temp_min":297.552,"temp_max":297.552,"pressure":995.92,"sea_level":1026.91,"grnd_level":995.92,"humidity":63,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":3.31,"deg":311},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-09 15:00:00"},{"dt":1499623200,"main":{"temp":300.062,"temp_min":300.062,"temp_max":300.062,"pressure":996.13,"sea_level":1027.04,"grnd_level":996.13,"humidity":51,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":5.06,"deg":303.002},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-09 18:00:00"},{"dt":1499634000,"main":{"temp":300.329,"temp_min":300.329,"temp_max":300.329,"pressure":995.72,"sea_level":1026.58,"grnd_level":995.72,"humidity":39,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":5.21,"deg":305.002},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-09 21:00:00"},{"dt":1499644800,"main":{"temp":298.369,"temp_min":298.369,"temp_max":298.369,"pressure":995.92,"sea_level":1026.74,"grnd_level":995.92,"humidity":36,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":4.36,"deg":307.505},"rain":{},"sys":{"pod":"n"},"dt_txt":"2017-07-10 00:00:00"},{"dt":1499655600,"main":{"temp":292.389,"temp_min":292.389,"temp_max":292.389,"pressure":996.9,"sea_level":1027.95,"grnd_level":996.9,"humidity":50,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.85,"deg":324},"rain":{},"sys":{"pod":"n"},"dt_txt":"2017-07-10 03:00:00"},{"dt":1499666400,"main":{"temp":288.756,"temp_min":288.756,"temp_max":288.756,"pressure":996.85,"sea_level":1028.14,"grnd_level":996.85,"humidity":66,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"02n"}],"clouds":{"all":8},"wind":{"speed":1.62,"deg":250.502},"rain":{},"sys":{"pod":"n"},"dt_txt":"2017-07-10 06:00:00"},{"dt":1499677200,"main":{"temp":286.26,"temp_min":286.26,"temp_max":286.26,"pressure":997.13,"sea_level":1028.42,"grnd_level":997.13,"humidity":76,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.2,"deg":260.501},"rain":{},"sys":{"pod":"n"},"dt_txt":"2017-07-10 09:00:00"},{"dt":1499688000,"main":{"temp":290.885,"temp_min":290.885,"temp_max":290.885,"pressure":997.83,"sea_level":1028.97,"grnd_level":997.83,"humidity":56,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":1.67,"deg":267.001},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-10 12:00:00"},{"dt":1499698800,"main":{"temp":298.64,"temp_min":298.64,"temp_max":298.64,"pressure":997.84,"sea_level":1028.8,"grnd_level":997.84,"humidity":46,"temp_kf":0},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"clouds":{"all":20},"wind":{"speed":2.62,"deg":213.002},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-10 15:00:00"},{"dt":1499709600,"main":{"temp":302.971,"temp_min":302.971,"temp_max":302.971,"pressure":996.76,"sea_level":1027.51,"grnd_level":996.76,"humidity":40,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":5.12,"deg":250.001},"rain":{},"sys":{"pod":"d"},"dt_txt":"2017-07-10 18:00:00"},{"dt":1499720400,"main":{"temp":301.228,"temp_min":301.228,"temp_max":301.228,"pressure":996.07,"sea_level":1026.82,"grnd_level":996.07,"humidity":44,"temp_kf":0},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10d"}],"clouds":{"all":56},"wind":{"speed":5.57,"deg":268.005},"rain":{"3h":0.54},"sys":{"pod":"d"},"dt_txt":"2017-07-10 21:00:00"}],"city":{"name":"Rochester","coord":{"lat":42.6593,"lon":-83.1225},"country":"US"}}`)
	var s = new(WeatherModel)
	err1 := json.Unmarshal(b, &s)
	if err1 != nil {
		fmt.Println("whoops:", err1)
	}

	return *s
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
