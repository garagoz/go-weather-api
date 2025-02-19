package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getWeatherHandler(w http.ResponseWriter, r *http.Request, apiKey string) {
	vars := mux.Vars(r)
	city := vars["city"]

	weatherData, err := fetchWeatherData(city, apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}
