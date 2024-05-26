package port

import "github.com/phelipperibeiro/lab-01-temperatureSystemByCEP/internal/domain/entity"

type WeatherGateway interface {
	GetWeather(city string) (*entity.Weather, error)
}
