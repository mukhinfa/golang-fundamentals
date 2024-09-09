package main

import "fmt"

const fromUSDtoEUR = 0.9
const fromUSDtoRUB = 88.25

type currency map[string]map[string]float64

var originalCurrency string

func main() {
	currency := currency{
		"USD": {
			"EUR": fromUSDtoEUR,
			"RUB": fromUSDtoRUB,
		},
		"EUR": {
			"USD": 1 / fromUSDtoEUR,
			"RUB": fromUSDtoEUR / fromUSDtoRUB,
		},
		"RUB": {
			"EUR": fromUSDtoRUB / fromUSDtoEUR,
			"USD": 1 / fromUSDtoRUB,
		},
	}
	originalCurrency = getOriginalCurrency()
	originalCurrencyValue := getOriginalCurrencyValue()
	targetCurrency := getTargetCurrency()
	calculateAmmount(originalCurrencyValue, originalCurrency, targetCurrency, currency)
}

func getOriginalCurrency() string {
	var v string
	fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
	for {
		fmt.Scan(&v)
		if v == "USD" || v == "EUR" || v == "RUB" {
			return v
		}
		fmt.Printf("Валюты %s у нас нет\n", v)
		fmt.Print("Введите исходную валюту еще раз (USD, EUR, RUB): ")
	}
}
func getOriginalCurrencyValue() float64 {
	var v float64
	fmt.Print("Введите исходную сумму: ")
	for {
		fmt.Scan(&v)
		if v > 0 {
			return v
		}
		fmt.Print("Некорректная сумма.\nПожалуйста, введите исходную суммму еще раз: ")
	}
}
func getTargetCurrency() string {
	var v string
	var acceptableCurrencies string
	switch {
	case originalCurrency == "USD":
		acceptableCurrencies = "(EUR, RUB): "
	case originalCurrency == "EUR":
		acceptableCurrencies = "(USD, RUB): "
	case originalCurrency == "RUB":
		acceptableCurrencies = "(EUR, USD): "
	}
	fmt.Printf("Введите целевую валюту %s", acceptableCurrencies)

	for {
		fmt.Scan(&v)
		if (v == "USD" || v == "EUR" || v == "RUB") && v != originalCurrency {
			return v
		}
		if v == originalCurrency {
			fmt.Printf("Валюта должна отличаться от исходной (%s)\n", v)
		} else {
			fmt.Printf("Валюты %s у нас нет\n", v)
		}
		fmt.Printf("Введите целевую валюту еще раз %s", acceptableCurrencies)
	}
}

func calculateAmmount(originalCurrencyValue float64, originalCurrency, targetCurrency string, currency currency) {
	var result float64
	result = originalCurrencyValue * currency[originalCurrency][targetCurrency]
	fmt.Printf("%.2f %s это %.2f %s", originalCurrencyValue, originalCurrency, result, targetCurrency)
}
