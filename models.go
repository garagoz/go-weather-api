package main

type WeatherData struct {
	Weather []WeatherDescription `json:"weather"`
	Main    Main                 `json:"main"`
	Wind    Wind                 `json:"wind"`
	Name    string               `json:"name"`
}

type WeatherDescription struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}
