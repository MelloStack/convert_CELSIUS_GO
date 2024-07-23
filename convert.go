package main

import "fmt"

func main() {
	result := fmt.Sprintf("Kelvin to celsius: %.2f °C", convert(373.2))
	fmt.Println(result)
}

func convert(kel float64) float64 {
	var cel float64 = kel - 273.15
	return cel
}
