package main

import (
	"testing"
)

func TestConvertToCelsius(t *testing.T) {
	fahrenheit := 68.5
	expected := 20.27777777777778

	celsius := convertToCelsius(fahrenheit)
	if celsius != expected {
		t.Fatalf(`Expected %f, got %f`, expected, celsius)
	}

	fahrenheit = 80
	expected = 26.666666666666668

	celsius = convertToCelsius(fahrenheit)
	if celsius != expected {
		t.Fatalf(`Expected %f, got %f`, expected, celsius)
	}
}
