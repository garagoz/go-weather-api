package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3" // For YAML config
)

type Config struct {
	OpenWeatherAPIKey string `yaml:"open_weather_api_key"`
	Port              string `yaml:"port"`
}

func loadConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var config Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	config, err := loadConfig("config.yaml") // Load config from YAML
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/weather/{city}", func(w http.ResponseWriter, r *http.Request) {
		getWeatherHandler(w, r, config.OpenWeatherAPIKey) // Pass API key to handler
	}).Methods("GET")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	fmt.Println("Server listening on port " + config.Port) // Use port from config

	srv := &http.Server{
		Handler:      r,
		Addr:         config.Port, // Use port from config
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
