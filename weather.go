package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetchWeatherData(city string, apiKey string) (*WeatherData, error) {
	openWeatherBaseURL := "https://api.openweathermap.org/data/2.5/weather"
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", openWeatherBaseURL, city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OpenWeather API returned an error: Status %d, Body: %s", resp.StatusCode, string(body))
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return &weatherData, nil
}
