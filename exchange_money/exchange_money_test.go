package main

import "testing"

func TestExchangeMoney(t *testing.T) {
	money := 28
	expected := [4]int{2, 1, 1, 1}

	exchangedCoin := exchangeMoney(money)
	if exchangedCoin != expected {
		t.Fatalf("Exchange money not correct. Expected %d, got %d", expected, exchangedCoin)
	}

	money = 6
	expected = [4]int{0, 1, 0, 1}

	exchangedCoin = exchangeMoney(money)
	if exchangedCoin != expected {
		t.Fatalf("Exchange money not correct. Expected %d, got %d", expected, exchangedCoin)
	}
}
