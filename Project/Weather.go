package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// func city();
type ResponseCity struct {
	Results []Coordinates `json:"results"`
}

type Coordinates struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type CurrentWeather struct {
	CurrentWeather Weather `json:"current_weather"`
}

type Weather struct {
	Temperature   float32 `json:"temperature"`
	Windspeed     float32 `json:"windspeed"`
	Winddirection float32 `json:"winddirection"`
	Weathercode   float32 `json:"weathercode"`
}

func weather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

		// Bagian mengubah kota menjadi koordinat
		var city = r.FormValue("city")
		fmt.Fprintln(w, city)

		responseCity, err := http.Get("https://geocoding-api.open-meteo.com/v1/search?count=1&name=" + city)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		// Mengambil hasil koordinat
		var resultCoordinatesObject ResponseCity
		json.NewDecoder(responseCity.Body).Decode(&resultCoordinatesObject)

		// Bagian mengubah koordinat menjadi cuaca
		responseWeather, err := http.Get(fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", resultCoordinatesObject.Results[0].Latitude, resultCoordinatesObject.Results[0].Longitude))
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		// Mengambil kondisi cuaca
		var resultWeatherObject CurrentWeather
		json.NewDecoder(responseWeather.Body).Decode(&resultWeatherObject)

		var WeatherCode int = int(resultWeatherObject.CurrentWeather.Weathercode)
		var WeatherType string
		switch WeatherCode {
		case 0:
			WeatherType = "Clear sky"
		case 1:
			WeatherType = "Mainly clear"
		case 2:
			WeatherType = "Partly cloudy"
		case 3:
			WeatherType = "Overcast"
		case 45:
			WeatherType = "Fog"
		case 48:
			WeatherType = "Depositing rime fog"
		case 51:
			WeatherType = "Drizzle, light"
		case 53:
			WeatherType = "Drizzle, moderate"
		case 55:
			WeatherType = "Drizzle, dense"
		case 56:
			WeatherType = "Freezing Dreezle, light"
		case 57:
			WeatherType = "Freezing Drizzle, dense"
		case 61:
			WeatherType = "Rain, slight"
		case 63:
			WeatherType = "Rain, moderate"
		case 65:
			WeatherType = "Rain, heavy"
		case 66:
			WeatherType = "Freezing Rain, light"
		case 67:
			WeatherType = "Freezing Rain, heavy"
		case 71:
			WeatherType = "Snow fall, slight"
		case 73:
			WeatherType = "Snow fall, moderate"
		case 75:
			WeatherType = "Snow fall, heavy"
		case 77:
			WeatherType = "Snow grains"
		case 80:
			WeatherType = "Rain showers, slight"
		case 81:
			WeatherType = "Rain showers, moderate"
		case 82:
			WeatherType = "Rain showers, violent"
		case 85:
			WeatherType = "Snow showers, slight"
		case 86:
			WeatherType = "Snow showers, heavy"
		case 95:
			WeatherType = "Thunderstorm, slight or Moderate"
		case 96:
			WeatherType = "Thunderstorm with slight hail"
		case 99:
			WeatherType = "Thunderstorm with heavy hail"
		default:
			WeatherType = "Unknown"
		}

		fmt.Fprintln(w, "Weather: "+WeatherType)
		fmt.Fprintln(w, "Temperature: "+fmt.Sprintf("%f", resultWeatherObject.CurrentWeather.Temperature))
		fmt.Fprintln(w, "Wind Speed: "+fmt.Sprintf("%f", resultWeatherObject.CurrentWeather.Windspeed))
		fmt.Fprintln(w, "Wind Direction: "+fmt.Sprintf("%f", resultWeatherObject.CurrentWeather.Winddirection))
	}
}

func main() {
	http.HandleFunc("/weather", weather)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
