package main

import "fmt"

func exchangeMoney(money int) [4]int {
	var tenCoin, fiveCoin, twoCoin, oneCoin int
	if money/10 > 0 {
		tenCoin += money / 10
		money = money % 10
	}
	if money/5 > 0 {
		fiveCoin += money / 5
		money = money % 5
	}
	if money/2 > 0 {
		twoCoin += money / 2
		money = money % 2
	}
	oneCoin += money
	return [4]int{tenCoin, fiveCoin, twoCoin, oneCoin}
}

func getMoney() (int, error) {
	var money int
	_, err := fmt.Scan(&money)
	if err != nil {
		return money, err
	}
	return money, nil
}

func main() {
	money, err := getMoney()
	if err != nil {
		fmt.Println(err)
	}
	exchangeCoin := exchangeMoney(money)
	fmt.Println("10 baht =", exchangeCoin[0])
	fmt.Println("5 baht =", exchangeCoin[1])
	fmt.Println("2 baht =", exchangeCoin[2])
	fmt.Println("1 baht =", exchangeCoin[3])
}
