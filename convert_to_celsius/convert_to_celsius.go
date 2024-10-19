package main

import "fmt"

func convertToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func getFahrenheit() (float64, error) {
	var fahrenheit float64
	_, err := fmt.Scan(&fahrenheit)
	if err != nil {
		return fahrenheit, err
	}
	return fahrenheit, nil
}

func main() {
	fahrenheit, err := getFahrenheit()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(convertToCelsius(fahrenheit))
}
