package main

import "fmt"

// exercise 6: Convert Celsius to Fahrenheit.

func convertTemperature(celsius float64) float64 {
	fahrenheit := (celsius * 1.8) + 32

	return fahrenheit
}

func main() {
	var celsiusTemperature float64

	fmt.Print("Enter the current celsius temperature of your city? ")
	fmt.Scan(&celsiusTemperature)

	fahrenheitTemperature := convertTemperature(celsiusTemperature)

	fmt.Printf("The current fahrenheit temperature of your city is: %.1f\n", fahrenheitTemperature)
}
