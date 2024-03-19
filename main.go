package main

import (
	"fmt"
)

func celsiusToKelvin(degrees float64) float64 {
	return degrees + 273.15
}

func main() {

	celsiusBoilingPoint := 100.0
	fmt.Printf("O ponto de ebulicão da água em Celsius é: %.2f °C\n", celsiusBoilingPoint)
	kelvinBoilingPoint := celsiusToKelvin(celsiusBoilingPoint)
	fmt.Printf("O ponto de ebulicão da água em Kelvin é: %.2f °K\n", kelvinBoilingPoint)
}
