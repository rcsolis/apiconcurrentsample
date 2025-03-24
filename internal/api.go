package internal

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

const mode = "xml"
const base_url = "http://api.openweathermap.org/data/2.5/weather"

func GetWeather(city string) (Weather, error) {
	appid := os.Getenv("OPENWEATHERMAP_API_KEY")
	if appid == "" {
		return Weather{}, fmt.Errorf("OPENWEATHERMAP_API_KEY is not set")
	}

	url := fmt.Sprintf("%s?q=%s&appid=%s&mode=%s", base_url, city, appid, mode)
	resp, err := http.Get(url)
	if err != nil {
		return Weather{}, err
	}
	defer resp.Body.Close()

	var weather Weather
	if err := xml.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return Weather{}, err
	}

	return weather, nil
}
