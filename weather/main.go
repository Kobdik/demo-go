package main

import (
	"demo-go/weather/geo"
	"demo-go/weather/wttr"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hi from weather app.")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода")
	flag.Parse()
	fmt.Println(*city)
	fmt.Println(*format)

	wttr.GetFullQuery()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(geoData)

	weatherData := wttr.GetWeather(geoData, *format)
	fmt.Println(weatherData)
}
