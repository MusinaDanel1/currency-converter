package infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const apiURL = "https://api.exchangeratesapi.io/v1/latest"

type ExchangeRateAPI struct {
	apiKey string
	client *http.Client
}

func NewExchangeRateAPI(apiKey string) *ExchangeRateAPI {
	return &ExchangeRateAPI{
		apiKey: apiKey,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (api *ExchangeRateAPI) GetExchangeRate(base, target string) (float64, error) {
	url := fmt.Sprintf("%s?access_key=%s&base=%s&symbols=%s", apiURL, api.apiKey, base, target)

	resp, err := api.client.Get(url)
	if err != nil {
		return 0, fmt.Errorf("не удалось отправить запрос: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("неправильный статус ответа: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("не удалось прочитать тело ответа: %v", err)
	}

	var response struct {
		Success bool               `json:"success"`
		Rates   map[string]float64 `json:"rates"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, fmt.Errorf("не удалось разобрать ответ: %v", err)
	}

	if !response.Success {
		return 0, fmt.Errorf("не удалось получить курсы валют")
	}

	rate, found := response.Rates[target]
	if !found {
		return 0, fmt.Errorf("не найдена целевая валюта: %s", target)
	}

	return rate, nil
}
