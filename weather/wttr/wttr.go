package wttr

import (
	"demo-go/weather/geo"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetFullQuery() {
	baseURL := "https://api.weather.mock/forecast"
	params := map[string]string{
		"city":  "Moscow",
		"units": "metric",
		"lang":  "ru",
	}

	// Ваш код здесь
	bU, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	bU.RawQuery = values.Encode()
	fullURL := bU.String()
	fmt.Println(fullURL)
}

func GetWeather(geoData *geo.GeoData, format int) string {

	baseUrl, err := url.Parse("https://wttr.in/" + geoData.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(body)
}
