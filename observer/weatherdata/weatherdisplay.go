package weatherdata

import "fmt"

type Observer interface {
	Name() string
	Update(data map[string]string)
}

type ConsoleWeatherDisplay struct {
	name string
}

func NewConsoleWeatherDisplay() *ConsoleWeatherDisplay {
	return &ConsoleWeatherDisplay{name: "console"}
}

func (c *ConsoleWeatherDisplay) Name() string {
	return c.name
}

func (c *ConsoleWeatherDisplay) Update(data map[string]string) {
	c.Display(data)
}

func (c *ConsoleWeatherDisplay) Display(data map[string]string) {
	for k, v := range data {
		fmt.Printf("%s:%s\n", k, v)
	}
}
