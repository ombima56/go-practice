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
	for {
		fmt.Println("Currency Converter")
		fmt.Println("1. Convert Currency")
		fmt.Println("2. Update Exchange Rate")
		fmt.Println("3. Exit")
		fmt.Println("Choose an option:")

		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			// Clear input buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch choice {
		case 1:
			performConversion()
		case 2:
			updateCurrency()
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
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

func performConversion() {
	fmt.Printf("\nAvailable currencies for conversion: %v\n", getCurrency())

	var source, target string
	var amount float64

	fmt.Println("Enter Source Currency: ")
	fmt.Scan(&source)
	source = strings.ToUpper(source)

	fmt.Println("Enter Target Currency: ")
	fmt.Scan(&target)
	target = strings.ToUpper(target)

	fmt.Println("Enter Amount: ")
	if _, err := fmt.Scan(&amount); err != nil {
		fmt.Println("Error: Invalid amount entered.")
		return
	}

	if !isValidCurrency(source) {
		fmt.Printf("Error: Unsupported currency: %s\n", source)
		return
	}

	if !isValidCurrency(target) {
		fmt.Printf("Error: Unsupported currency: %s\n", target)
		return
	}

	if amount <= 0 {
		fmt.Println("Error: Amount must be more than zero.")
		return
	}

	result := currencyConverter(source, target, amount)
	fmt.Printf("Conversion from %s to %s: %.2f %s\n", source, target, result, target)
}

func isValidCurrency(currency string) bool {
	_, ok := exchangeRate[currency]
	return ok
}

func updateCurrency() {
	fmt.Printf("\nAvailable currencies for update: %v\n", getCurrency())

	var currency string
	var rate float64

	fmt.Println("Enter Currency to Update (e.g., USD): ")
	fmt.Scan(&currency)
	currency = strings.ToUpper(currency)

	if !isValidCurrency(currency) {
		fmt.Printf("Error: Unsupported currency: %s\n", currency)
		return
	}

	fmt.Println("Enter New Exchange Rate: ")
	if _, err := fmt.Scan(&rate); err != nil {
		fmt.Println("Error: Invalid rate entered.")
		return
	}

	if rate <= 0 {
		fmt.Println("Error: Rate must be more than zero.")
		return
	}

	exchangeRate[currency] = rate
	fmt.Printf("Exchange rate updated: %s = %.2f\n", currency, rate)
}
