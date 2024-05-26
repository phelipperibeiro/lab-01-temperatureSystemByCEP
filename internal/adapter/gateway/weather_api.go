package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/phelipperibeiro/lab-01-temperatureSystemByCEP/internal/domain/entity"
)

func buildQuery(params map[string]string) string {
	var parts []string
	for key, value := range params {
		parts = append(parts, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
	}
	return strings.Join(parts, "&")
}

type WeatherAPIGateway struct{}

func (w *WeatherAPIGateway) GetWeather(city string) (*entity.Weather, error) {

	apiKey := "776617dd5d694eaa94d33907242605" // token de acesso da API WeatherAPI

	params := map[string]string{
		"q":    city,
		"lang": "en",
		"key":  apiKey,
	}

	resp, err := http.Get(fmt.Sprintf("%s?%s", "http://api.weatherapi.com/v1/current.json", buildQuery(params)))

	if err != nil {
		fmt.Println("Erro ao buscar informações do clima:", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erro ao buscar informações do clima:", resp.Status)
		return nil, fmt.Errorf("cannot get weather information")
	}

	var weatherResponse struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		fmt.Println("Erro ao decodificar resposta do clima:", err)
		return nil, err
	}

	weather := &entity.Weather{
		TempC: weatherResponse.Current.TempC,
		TempF: weatherResponse.Current.TempC*1.8 + 32,
		TempK: weatherResponse.Current.TempC + 273,
	}

	return weather, nil
}

func NewWeatherAPIGateway() *WeatherAPIGateway {
	return &WeatherAPIGateway{}
}
