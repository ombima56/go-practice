package main

import (
	"fmt"
	"strings"
)

var exchangeRate = map[string]float64{
	"USD": 1,
	"EUR": 1.5,
	"JPY": 130,
	"KES": 140,
}

func main() {
	fmt.Printf("The following are the only available conversion: %s\n", getCurrency())

	var source, target string
	var amount, rate float64
	fmt.Println("Enter Source: ")
	fmt.Scan(&source)
	source = strings.ToUpper(source)

	fmt.Println("Enter Target: ")
	fmt.Scan(&target)
	target = strings.ToUpper(target)

	fmt.Println("Enter Amount: ")
	fmt.Scan(&amount)

	if _, ok := exchangeRate[source]; !ok {
		fmt.Printf("Error: unsupported currencies: %s\n", source)
		return
	}

	if _, ok := exchangeRate[target]; !ok {
		fmt.Printf("Error: unsupported currencies: %s\n", target)
		return
	}

	if amount <= 0 {
		fmt.Println("Error: amount must be more than zero")
		return
	}

	rate = currencyConverter(source, target, amount)
	fmt.Printf("Conversion from %s to %s is %.2f %s\n", source, target, rate, source)
}

func getCurrency() (rates []string) {
	for rate := range exchangeRate {
		rates = append(rates, rate)
	}
	return
}

func currencyConverter(source, target string, amount float64) (rate float64) {
	sourceRate := exchangeRate[source]
	targetRate := exchangeRate[target]

	rate = (amount / sourceRate) * targetRate
	return
}
