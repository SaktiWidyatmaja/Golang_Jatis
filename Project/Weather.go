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

type Weather struct {
	Time          int     `json:"Time"`
	Temperature   float32 `json:"Temperature"`
	Weathercode   int     `json:"Weathercode"`
	Windspeed     float32 `json:"Windspeed"`
	Winddirection int     `json:"Winddirection"`
}

func weather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

		// Bagian mengubah kota menjadi koordinat
		var city = r.FormValue("city")
		fmt.Println(city)

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

		// resultWeatherData, err := ioutil.ReadAll(responseWeather.Body)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// Mengambil kondisi cuaca
		var resultWeatherObject Weather
		json.NewDecoder(responseWeather.Body).Decode(&resultWeatherObject)

		fmt.Fprintln(w, resultWeatherObject.Time)
		fmt.Fprintln(w, "halo!")
		fmt.Fprintln(w, "halo!")
		fmt.Fprintln(w, "halo!")
		fmt.Fprintln(w, "halo!")

		// resultWeatherData, err := json.Marshal(resultWeatherObject)
		// var resultWeatherData http.ResponseWriter
		// json.NewEncoder(w).Encode(resultWeatherObject)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// w.Write()
	}

}

func main() {
	http.HandleFunc("/weather", weather)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
