package main

import "github.com/s-buhar0v/golang-design-patters/observer/weatherdata"

func main() {
	w := weatherdata.NewWeatherData()
	c := weatherdata.NewConsoleWeatherDisplay()

	w.RegisterObserver(c)

	w.SetWeatherData(map[string]string{ // ConsoleWeatherDisplay will print
		"temperature": "+12", // temperature:+12
		"rainy":       "no",  // rainy:no
		"windy":       "yes", // windy:yes
	})
	w.SetWeatherData(map[string]string{ // ConsoleWeatherDisplay will print
		"temperature": "+8",  // temperature:+8
		"windy":       "yes", // windy:yes
		"rainy":       "yes", // rainy:yes
	})

	w.RemoveObserver(c)
	w.SetWeatherData(map[string]string{ // ConsoleWeatherDisplay will print nothing
		"temperature": "-12",
		"windy":       "no",
		"rainy":       "no",
	})
}
