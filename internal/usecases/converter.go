package usecases

import (
	"currency-converter/internal/infrastructure"
	"fmt"
)

type CurrencyConverter struct {
	rateProvider *infrastructure.ExchangeRateAPI
}

func NewCurrencyConverter(rateProvider *infrastructure.ExchangeRateAPI) *CurrencyConverter {
	return &CurrencyConverter{rateProvider: rateProvider}
}

func (c *CurrencyConverter) GetExchangeRate(base, target string) (float64, error) {
	return c.rateProvider.GetExchangeRate(base, target)
}

func (c *CurrencyConverter) Convert(amount float64, base, target string) (float64, error) {
	rate, err := c.GetExchangeRate(base, target)
	if err != nil {
		return 0, fmt.Errorf("ошибка получения курса валют: %v", err)
	}

	return amount * rate, nil
}
