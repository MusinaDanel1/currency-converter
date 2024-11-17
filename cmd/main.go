package main

import (
	"currency-converter/internal/infrastructure"
	"currency-converter/internal/interfaces"
	"currency-converter/internal/usecases"
	"flag"
	"log"
)

func main() {

	apiKey := flag.String("apikey", "", "API ключ для доступа к курсам валют")
	flag.Parse()

	if *apiKey == "" {
		log.Fatal("Необходимо указать API ключ с помощью флага -apikey")
	}

	rateProvider := infrastructure.NewExchangeRateAPI(*apiKey)

	converter := usecases.NewCurrencyConverter(rateProvider)

	app := interfaces.NewConsoleApp(converter)

	app.Run()
}
