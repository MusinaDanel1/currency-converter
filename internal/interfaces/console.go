package interfaces

import (
	"currency-converter/internal/usecases"
	"fmt"
	"log"
)

type ConsoleApp struct {
	converter *usecases.CurrencyConverter
}

func NewConsoleApp(converter *usecases.CurrencyConverter) *ConsoleApp {
	return &ConsoleApp{converter: converter}
}

func (app *ConsoleApp) Run() {
	var amount float64
	var base, target string

	fmt.Print("Введите сумму для конвертации: ")
	_, err := fmt.Scanf("%f", &amount)
	if err != nil {
		log.Fatalf("Ошибка ввода суммы. Пожалуйста, введите корректное число: %v", err)
	}

	fmt.Print("Введите исходную валюту (например, USD): ")
	_, err = fmt.Scan(&base)
	if err != nil {
		log.Fatalf("Ошибка ввода исходной валюты: %v", err)
	}

	fmt.Print("Введите целевую валюту (например, EUR): ")
	_, err = fmt.Scan(&target)
	if err != nil {
		log.Fatalf("Ошибка ввода целевой валюты: %v", err)
	}

	rate, err := app.converter.GetExchangeRate(base, target)
	if err != nil {
		log.Fatalf("Ошибка получения курса валют: %v", err)
	}

	result := amount * rate

	fmt.Printf("%.2f %s = %.2f %s\n", amount, base, result, target)
}
