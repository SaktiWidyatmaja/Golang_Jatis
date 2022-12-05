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
	Timezone  string  `json:"timezone"`
}

type Time struct {
	Time     string `json:"time"`
	TimeZone string `json:"timeZone"`
}

type CurrentWeather struct {
	CurrentWeather Weather `json:"current_weather"`
}

type Weather struct {
	Temperature   float32 `json:"temperature"`
	Windspeed     float32 `json:"windspeed"`
	Winddirection float32 `json:"winddirection"`
	Weathercode   float32 `json:"weathercode"`
	Time          string  `json:"time"`
}

func time(w http.ResponseWriter, r *http.Request) {
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

		fmt.Println(resultCoordinatesObject.Results)

		// Bagian mengubah koordinat menjadi waktu
		responseTime, err := http.Get(fmt.Sprintf("https://timeapi.io/api/Time/current/coordinate?latitude=%f&longitude=%f", resultCoordinatesObject.Results[0].Latitude, resultCoordinatesObject.Results[0].Longitude))
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		// Bagian mengubah koordinat menjadi cuaca
		responseWeather, err := http.Get(fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", resultCoordinatesObject.Results[0].Latitude, resultCoordinatesObject.Results[0].Longitude))
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		// Mengambil waktu dari http response
		var resultTimeObject Time
		json.NewDecoder(responseTime.Body).Decode(&resultTimeObject)

		// Mengambil kondisi cuaca
		var resultWeatherObject CurrentWeather
		json.NewDecoder(responseWeather.Body).Decode(&resultWeatherObject)

		// var WeatherCode int = int(resultWeatherObject.CurrentWeather.Weathercode)
		// var WeatherType string

		// switch WeatherCode {
		// case 0:
		// 	WeatherType = "Clear sky"
		// case 1:
		// 	WeatherType = "Mainly clear"
		// default:
		// 	WeatherType = "Unknown"
		// }

		fmt.Fprintln(w, "Time: "+resultTimeObject.Time)
		fmt.Fprintln(w, "Time Zone: "+resultTimeObject.TimeZone)
	}
}

func main() {
	http.HandleFunc("/time", time)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
