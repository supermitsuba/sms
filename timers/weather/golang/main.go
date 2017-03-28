package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func main() {
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?zip=48307,us&appid=480b42d2dc3eb6aae6664c3921733971")
	if err != nil {
    		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
    		panic(err.Error())
	}

	var s = new (WeatherModel)
	err1 := json.Unmarshal(body, &s)
	if(err1 != nil){
        	fmt.Println("whoops:", err1)
    	}
	fmt.Println("Hello, playground", s.WeatherSections[0].Main, s.Main.Temperature )
}

type WeatherModel struct {
	WeatherSections []WeatherSection `json:"weather"`
	Main MainSection `json:"main"`
}

type WeatherSection struct {
	Main string `json:"main"`
}

type MainSection struct {
	Temperature float64 `json:"temp"`
}
