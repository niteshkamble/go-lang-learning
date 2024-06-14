package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type WeatherStructure struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		Temp      float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Hour []struct {
				Temp       float64 `json:"temp_c"`
				Time_epoch int64   `json:"time_epoch"`
				Condition  struct {
					Text string `json:"text"`
				} `json:"condition"`
				FeelsLikeC   float64 `json:"feelslike_c"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		fmt.Println("Set API_KEY in .env file (create new api here www.weatherapi.com)")
	}

	//Run program with --city <city_name> flag for particular city, default city is mumbai
	city := flag.String("city", "mumbai", "city provided by user")
	flag.Parse()

	const baseUrl string = "https://api.weatherapi.com/v1/forecast.json"
	var parms string = "&q=" + *city + "&days=1&aqi=yes&alerts=no"
	var url string = fmt.Sprintf("%s?key=%s%s", baseUrl, API_KEY, parms)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather Api Failed.")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weater WeatherStructure
	err = json.Unmarshal(body, &weater)

	if err != nil {
		panic(err)
	}

	_location, _current, _hours := weater.Location, weater.Current, weater.Forecast.ForecastDay[0].Hour

	fmt.Printf("\n%s, %s: %0.fC, %s\n",
		_location.Name,
		_location.Country,
		_current.Temp,
		_current.Condition.Text)
	fmt.Println("")

	for _, hour := range _hours {
		dateFormatter := time.Unix(hour.Time_epoch, 0)
		if dateFormatter.Before(time.Now()) {
			continue
		}

		data := fmt.Sprintf(
			"%s - %0.fC, %0.f%%, %s\n",
			dateFormatter.Format("03:04 PM"),
			hour.Temp,
			hour.ChanceOfRain,
			hour.Condition.Text)

		if hour.ChanceOfRain <= 50 {
			if hour.FeelsLikeC <= 36 {
				fmt.Print(data)
			} else {
				//if the tepm feels like >=36 (note:not temp. its feelsliketemp)
				color.Yellow(data)
			}
		} else {
			//if there are chances of rain
			color.Red(data)
		}

	}
}
