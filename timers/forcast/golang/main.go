package main

import (
	"fmt"
	"encoding/json"
	"strings"
)

func main() {
	weather :=  Weather()
	var s = new(WeatherModel)
	err1 := json.Unmarshal([]byte(weather), &s)
	if err1 != nil {
		fmt.Println("whoops:", err1)
	}
	
	for i := 0; i < len(s.WeatherForcast); i++ {
		item := s.WeatherForcast[i]
		
		d := strings.Split(item.DateOfTemperature, " ")[1]
		e := strings.Split(d, ":")[0]
		if e == "15" {
			fmt.Println("Hello, playground ", item)
		}
	}
}

type Message struct {
	Text     string  `json:"text"`
	Duration float64 `json:"duration"`
}

type WeatherModel struct {
	WeatherForcast    []ListWeather    `json:"list"`
}

type ListWeather struct {
	WeatherSections []WeatherSection `json:"weather"`
	Main            MainSection      `json:"main"`
	DateOfTemperature  string           `json:"dt_txt"`
}

type WeatherSection struct {
	Main string `json:"main"`
}

type MainSection struct {
	Temperature float64 `json:"temp"`
}


func Weather() string {
	return `{
  "cod": "200",
  "message": 0.1214,
  "cnt": 35,
  "list": [
    {
      "dt": 1490886000,
      "main": {
        "temp": 275.79,
        "temp_min": 275.79,
        "temp_max": 276.033,
        "pressure": 999.08,
        "sea_level": 1031.75,
        "grnd_level": 999.08,
        "humidity": 98,
        "temp_kf": -0.24
      },
      "weather": [
        {
          "id": 501,
          "main": "Rain",
          "description": "moderate rain",
          "icon": "10d"
        }
      ],
      "clouds": {
        "all": 100
      },
      "wind": {
        "speed": 6.66,
        "deg": 97.5053
      },
      "rain": {
        "3h": 5.92
      },
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-03-30 15:00:00"
    },
    {
      "dt": 1490896800,
      "main": {
        "temp": 276.89,
        "temp_min": 276.89,
        "temp_max": 277.072,
        "pressure": 996.16,
        "sea_level": 1028.62,
        "grnd_level": 996.16,
        "humidity": 100,
        "temp_kf": -0.18
      },
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10d"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 7.88,
        "deg": 98.5029
      },
      "rain": {
        "3h": 1.505
      },
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-03-30 18:00:00"
    },
    {
      "dt": 1490907600,
      "main": {
        "temp": 277.14,
        "temp_min": 277.14,
        "temp_max": 277.257,
        "pressure": 992.67,
        "sea_level": 1024.85,
        "grnd_level": 992.67,
        "humidity": 99,
        "temp_kf": -0.12
      },
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10d"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 7.75,
        "deg": 97.0012
      },
      "rain": {
        "3h": 2.795
      },
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-03-30 21:00:00"
    },
    {
      "dt": 1490918400,
      "main": {
        "temp": 277.69,
        "temp_min": 277.69,
        "temp_max": 277.746,
        "pressure": 990.23,
        "sea_level": 1022.4,
        "grnd_level": 990.23,
        "humidity": 99,
        "temp_kf": -0.06
      },
      "weather": [
        {
          "id": 501,
          "main": "Rain",
          "description": "moderate rain",
          "icon": "10n"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 6.86,
        "deg": 93.5033
      },
      "rain": {
        "3h": 4.955
      },
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-03-31 00:00:00"
    },
    {
      "dt": 1490929200,
      "main": {
        "temp": 279.057,
        "temp_min": 279.057,
        "temp_max": 279.057,
        "pressure": 988.67,
        "sea_level": 1020.9,
        "grnd_level": 988.67,
        "humidity": 97,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 501,
          "main": "Rain",
          "description": "moderate rain",
          "icon": "10n"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 6.01,
        "deg": 98.5006
      },
      "rain": {
        "3h": 3.01
      },
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-03-31 03:00:00"
    },
    {
      "dt": 1490940000,
      "main": {
        "temp": 281.177,
        "temp_min": 281.177,
        "temp_max": 281.177,
        "pressure": 986.68,
        "sea_level": 1018.74,
        "grnd_level": 986.68,
        "humidity": 96,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10n"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 3.97,
        "deg": 99.5048
      },
      "rain": {
        "3h": 0.75
      },
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-03-31 06:00:00"
    },
    {
      "dt": 1490950800,
      "main": {
        "temp": 283.348,
        "temp_min": 283.348,
        "temp_max": 283.348,
        "pressure": 984.94,
        "sea_level": 1016.8,
        "grnd_level": 984.94,
        "humidity": 98,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 501,
          "main": "Rain",
          "description": "moderate rain",
          "icon": "10n"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 2.96,
        "deg": 110.502
      },
      "rain": {
        "3h": 4.845
      },
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-03-31 09:00:00"
    },
    {
      "dt": 1490961600,
      "main": {
        "temp": 285.386,
        "temp_min": 285.386,
        "temp_max": 285.386,
        "pressure": 984.56,
        "sea_level": 1016.44,
        "grnd_level": 984.56,
        "humidity": 96,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 501,
          "main": "Rain",
          "description": "moderate rain",
          "icon": "10d"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 3.41,
        "deg": 189.504
      },
      "rain": {
        "3h": 4.815
      },
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-03-31 12:00:00"
    },
    {
      "dt": 1490972400,
      "main": {
        "temp": 288.012,
        "temp_min": 288.012,
        "temp_max": 288.012,
        "pressure": 984.87,
        "sea_level": 1016.63,
        "grnd_level": 984.87,
        "humidity": 100,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10d"
        }
      ],
      "clouds": {
        "all": 64
      },
      "wind": {
        "speed": 2.31,
        "deg": 197.003
      },
      "rain": {
        "3h": 0.505
      },
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-03-31 15:00:00"
    },
    {
      "dt": 1490983200,
      "main": {
        "temp": 288.883,
        "temp_min": 288.883,
        "temp_max": 288.883,
        "pressure": 984.2,
        "sea_level": 1015.8,
        "grnd_level": 984.2,
        "humidity": 95,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10d"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 1.94,
        "deg": 168.003
      },
      "rain": {
        "3h": 0.02
      },
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-03-31 18:00:00"
    },
    {
      "dt": 1490994000,
      "main": {
        "temp": 288.594,
        "temp_min": 288.594,
        "temp_max": 288.594,
        "pressure": 983.78,
        "sea_level": 1015.51,
        "grnd_level": 983.78,
        "humidity": 94,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10d"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 1.26,
        "deg": 203.502
      },
      "rain": {
        "3h": 0.59
      },
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-03-31 21:00:00"
    },
    {
      "dt": 1491004800,
      "main": {
        "temp": 285.706,
        "temp_min": 285.706,
        "temp_max": 285.706,
        "pressure": 984.54,
        "sea_level": 1016.65,
        "grnd_level": 984.54,
        "humidity": 97,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 501,
          "main": "Rain",
          "description": "moderate rain",
          "icon": "10n"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 1.6,
        "deg": 353.503
      },
      "rain": {
        "3h": 3.055
      },
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-01 00:00:00"
    },
    {
      "dt": 1491015600,
      "main": {
        "temp": 283.482,
        "temp_min": 283.482,
        "temp_max": 283.482,
        "pressure": 986.58,
        "sea_level": 1018.65,
        "grnd_level": 986.58,
        "humidity": 94,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 501,
          "main": "Rain",
          "description": "moderate rain",
          "icon": "10n"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 2.87,
        "deg": 69.002
      },
      "rain": {
        "3h": 4.38
      },
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-01 03:00:00"
    },
    {
      "dt": 1491026400,
      "main": {
        "temp": 282.152,
        "temp_min": 282.152,
        "temp_max": 282.152,
        "pressure": 988.31,
        "sea_level": 1020.52,
        "grnd_level": 988.31,
        "humidity": 99,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10n"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 2.96,
        "deg": 62.0016
      },
      "rain": {
        "3h": 1.58
      },
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-01 06:00:00"
    },
    {
      "dt": 1491037200,
      "main": {
        "temp": 280.873,
        "temp_min": 280.873,
        "temp_max": 280.873,
        "pressure": 990.52,
        "sea_level": 1022.96,
        "grnd_level": 990.52,
        "humidity": 100,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10n"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 3.88,
        "deg": 29.0089
      },
      "rain": {
        "3h": 0.245
      },
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-01 09:00:00"
    },
    {
      "dt": 1491048000,
      "main": {
        "temp": 278.925,
        "temp_min": 278.925,
        "temp_max": 278.925,
        "pressure": 993.64,
        "sea_level": 1026.22,
        "grnd_level": 993.64,
        "humidity": 98,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10d"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 4.57,
        "deg": 23.5049
      },
      "rain": {
        "3h": 0.12
      },
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-01 12:00:00"
    },
    {
      "dt": 1491058800,
      "main": {
        "temp": 279.28,
        "temp_min": 279.28,
        "temp_max": 279.28,
        "pressure": 996.44,
        "sea_level": 1028.92,
        "grnd_level": 996.44,
        "humidity": 96,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10d"
        }
      ],
      "clouds": {
        "all": 80
      },
      "wind": {
        "speed": 4.61,
        "deg": 20.5017
      },
      "rain": {
        "3h": 0.030000000000001
      },
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-01 15:00:00"
    },
    {
      "dt": 1491069600,
      "main": {
        "temp": 281.108,
        "temp_min": 281.108,
        "temp_max": 281.108,
        "pressure": 997.77,
        "sea_level": 1030.03,
        "grnd_level": 997.77,
        "humidity": 91,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 802,
          "main": "Clouds",
          "description": "scattered clouds",
          "icon": "03d"
        }
      ],
      "clouds": {
        "all": 48
      },
      "wind": {
        "speed": 4.86,
        "deg": 20.5017
      },
      "rain": {},
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-01 18:00:00"
    },
    {
      "dt": 1491080400,
      "main": {
        "temp": 282.524,
        "temp_min": 282.524,
        "temp_max": 282.524,
        "pressure": 998.26,
        "sea_level": 1030.54,
        "grnd_level": 998.26,
        "humidity": 88,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 802,
          "main": "Clouds",
          "description": "scattered clouds",
          "icon": "03d"
        }
      ],
      "clouds": {
        "all": 36
      },
      "wind": {
        "speed": 3.85,
        "deg": 17
      },
      "rain": {},
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-01 21:00:00"
    },
    {
      "dt": 1491091200,
      "main": {
        "temp": 281.043,
        "temp_min": 281.043,
        "temp_max": 281.043,
        "pressure": 1000.1,
        "sea_level": 1032.55,
        "grnd_level": 1000.1,
        "humidity": 84,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 804,
          "main": "Clouds",
          "description": "overcast clouds",
          "icon": "04n"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 3.41,
        "deg": 1.00635
      },
      "rain": {},
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-02 00:00:00"
    },
    {
      "dt": 1491102000,
      "main": {
        "temp": 277.446,
        "temp_min": 277.446,
        "temp_max": 277.446,
        "pressure": 1001.12,
        "sea_level": 1034.03,
        "grnd_level": 1001.12,
        "humidity": 88,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 800,
          "main": "Clear",
          "description": "clear sky",
          "icon": "01n"
        }
      ],
      "clouds": {
        "all": 0
      },
      "wind": {
        "speed": 2.78,
        "deg": 13.0003
      },
      "rain": {},
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-02 03:00:00"
    },
    {
      "dt": 1491112800,
      "main": {
        "temp": 273.971,
        "temp_min": 273.971,
        "temp_max": 273.971,
        "pressure": 1001.22,
        "sea_level": 1034.23,
        "grnd_level": 1001.22,
        "humidity": 94,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 801,
          "main": "Clouds",
          "description": "few clouds",
          "icon": "02n"
        }
      ],
      "clouds": {
        "all": 20
      },
      "wind": {
        "speed": 1.26,
        "deg": 353.003
      },
      "rain": {},
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-02 06:00:00"
    },
    {
      "dt": 1491123600,
      "main": {
        "temp": 272.367,
        "temp_min": 272.367,
        "temp_max": 272.367,
        "pressure": 1001.94,
        "sea_level": 1035.12,
        "grnd_level": 1001.94,
        "humidity": 92,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 800,
          "main": "Clear",
          "description": "clear sky",
          "icon": "02n"
        }
      ],
      "clouds": {
        "all": 8
      },
      "wind": {
        "speed": 1.24,
        "deg": 308
      },
      "rain": {},
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-02 09:00:00"
    },
    {
      "dt": 1491134400,
      "main": {
        "temp": 272.835,
        "temp_min": 272.835,
        "temp_max": 272.835,
        "pressure": 1002.64,
        "sea_level": 1035.74,
        "grnd_level": 1002.64,
        "humidity": 87,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 804,
          "main": "Clouds",
          "description": "overcast clouds",
          "icon": "04d"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 1.26,
        "deg": 310.007
      },
      "rain": {},
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-02 12:00:00"
    },
    {
      "dt": 1491145200,
      "main": {
        "temp": 279.289,
        "temp_min": 279.289,
        "temp_max": 279.289,
        "pressure": 1002.56,
        "sea_level": 1035.25,
        "grnd_level": 1002.56,
        "humidity": 88,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 804,
          "main": "Clouds",
          "description": "overcast clouds",
          "icon": "04d"
        }
      ],
      "clouds": {
        "all": 92
      },
      "wind": {
        "speed": 2.07,
        "deg": 21.5027
      },
      "rain": {},
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-02 15:00:00"
    },
    {
      "dt": 1491156000,
      "main": {
        "temp": 283.771,
        "temp_min": 283.771,
        "temp_max": 283.771,
        "pressure": 1002.16,
        "sea_level": 1034.41,
        "grnd_level": 1002.16,
        "humidity": 79,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 802,
          "main": "Clouds",
          "description": "scattered clouds",
          "icon": "03d"
        }
      ],
      "clouds": {
        "all": 48
      },
      "wind": {
        "speed": 1.91,
        "deg": 96.0019
      },
      "rain": {},
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-02 18:00:00"
    },
    {
      "dt": 1491166800,
      "main": {
        "temp": 285.398,
        "temp_min": 285.398,
        "temp_max": 285.398,
        "pressure": 1000.52,
        "sea_level": 1032.7,
        "grnd_level": 1000.52,
        "humidity": 67,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 800,
          "main": "Clear",
          "description": "clear sky",
          "icon": "02d"
        }
      ],
      "clouds": {
        "all": 8
      },
      "wind": {
        "speed": 1.86,
        "deg": 116.004
      },
      "rain": {},
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-02 21:00:00"
    },
    {
      "dt": 1491177600,
      "main": {
        "temp": 282.348,
        "temp_min": 282.348,
        "temp_max": 282.348,
        "pressure": 1000.3,
        "sea_level": 1032.57,
        "grnd_level": 1000.3,
        "humidity": 66,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 800,
          "main": "Clear",
          "description": "clear sky",
          "icon": "02n"
        }
      ],
      "clouds": {
        "all": 8
      },
      "wind": {
        "speed": 1.32,
        "deg": 106
      },
      "rain": {},
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-03 00:00:00"
    },
    {
      "dt": 1491188400,
      "main": {
        "temp": 277.706,
        "temp_min": 277.706,
        "temp_max": 277.706,
        "pressure": 1000.37,
        "sea_level": 1033.1,
        "grnd_level": 1000.37,
        "humidity": 81,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 802,
          "main": "Clouds",
          "description": "scattered clouds",
          "icon": "03n"
        }
      ],
      "clouds": {
        "all": 32
      },
      "wind": {
        "speed": 2.53,
        "deg": 144.502
      },
      "rain": {},
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-03 03:00:00"
    },
    {
      "dt": 1491199200,
      "main": {
        "temp": 275.317,
        "temp_min": 275.317,
        "temp_max": 275.317,
        "pressure": 1000.09,
        "sea_level": 1032.88,
        "grnd_level": 1000.09,
        "humidity": 92,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 802,
          "main": "Clouds",
          "description": "scattered clouds",
          "icon": "03n"
        }
      ],
      "clouds": {
        "all": 44
      },
      "wind": {
        "speed": 1.21,
        "deg": 104.003
      },
      "rain": {},
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-03 06:00:00"
    },
    {
      "dt": 1491210000,
      "main": {
        "temp": 274.746,
        "temp_min": 274.746,
        "temp_max": 274.746,
        "pressure": 999.89,
        "sea_level": 1032.77,
        "grnd_level": 999.89,
        "humidity": 89,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 803,
          "main": "Clouds",
          "description": "broken clouds",
          "icon": "04n"
        }
      ],
      "clouds": {
        "all": 64
      },
      "wind": {
        "speed": 1.26,
        "deg": 52.5041
      },
      "rain": {},
      "sys": {
        "pod": "n"
      },
      "dt_txt": "2017-04-03 09:00:00"
    },
    {
      "dt": 1491220800,
      "main": {
        "temp": 275.16,
        "temp_min": 275.16,
        "temp_max": 275.16,
        "pressure": 999.64,
        "sea_level": 1032.51,
        "grnd_level": 999.64,
        "humidity": 90,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 803,
          "main": "Clouds",
          "description": "broken clouds",
          "icon": "04d"
        }
      ],
      "clouds": {
        "all": 64
      },
      "wind": {
        "speed": 1.66,
        "deg": 68.5001
      },
      "rain": {},
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-03 12:00:00"
    },
    {
      "dt": 1491231600,
      "main": {
        "temp": 280.806,
        "temp_min": 280.806,
        "temp_max": 280.806,
        "pressure": 999.1,
        "sea_level": 1031.43,
        "grnd_level": 999.1,
        "humidity": 82,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 803,
          "main": "Clouds",
          "description": "broken clouds",
          "icon": "04d"
        }
      ],
      "clouds": {
        "all": 64
      },
      "wind": {
        "speed": 3.02,
        "deg": 94.0001
      },
      "rain": {},
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-03 15:00:00"
    },
    {
      "dt": 1491242400,
      "main": {
        "temp": 284.479,
        "temp_min": 284.479,
        "temp_max": 284.479,
        "pressure": 997.87,
        "sea_level": 1029.76,
        "grnd_level": 997.87,
        "humidity": 71,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 802,
          "main": "Clouds",
          "description": "scattered clouds",
          "icon": "03d"
        }
      ],
      "clouds": {
        "all": 44
      },
      "wind": {
        "speed": 3.92,
        "deg": 107.5
      },
      "rain": {},
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-03 18:00:00"
    },
    {
      "dt": 1491253200,
      "main": {
        "temp": 285.668,
        "temp_min": 285.668,
        "temp_max": 285.668,
        "pressure": 995.75,
        "sea_level": 1027.54,
        "grnd_level": 995.75,
        "humidity": 61,
        "temp_kf": 0
      },
      "weather": [
        {
          "id": 802,
          "main": "Clouds",
          "description": "scattered clouds",
          "icon": "03d"
        }
      ],
      "clouds": {
        "all": 48
      },
      "wind": {
        "speed": 4.32,
        "deg": 93.0003
      },
      "rain": {},
      "sys": {
        "pod": "d"
      },
      "dt_txt": "2017-04-03 21:00:00"
    }
  ],
  "city": {
    "id": 5007402,
    "name": "Rochester Hills",
    "coord": {
      "lat": 42.6584,
      "lon": -83.15
    },
    "country": "US",
    "population": 70995
  }
}`
}
