package weather

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
)

const OpenWeatherApiEndpoint = "https://api.openweathermap.org/data/2.5/weather"

type Conditions struct {
	Summary               Main
	TemperatureCelsius    float64
	TemperatureFahrenheit float64
}

type CurrentWeatherData struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  float64 `json:"pressure"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
}

func GetCurrent(location, token string) (Conditions, error) {
	url := FormatURL(location, token)
	resp, err := http.Get(url)
	if err != nil {
		return Conditions{}, nil
	}
	data := ParseJSON(resp.Body)
	return data, nil
}
func FormatURL(location string, token string) string {
	return OpenWeatherApiEndpoint + "?q=" + location + "&appid=" + token
}
func ParseJSON(r io.Reader) Conditions {
	current := CurrentWeatherData{
		Weather: []Weather{},
		Main:    Main{},
	}
	json.NewDecoder(r).Decode(&current)

	celsius := math.Round((current.Main.Temp-273.15)*100) / 100
	fahrenheit := math.Round((1.8*(current.Main.Temp-273.15)+32)*100) / 100
	data := Conditions{
		Summary:               current.Main,
		TemperatureCelsius:    celsius,
		TemperatureFahrenheit: fahrenheit,
	}
	return data
}
