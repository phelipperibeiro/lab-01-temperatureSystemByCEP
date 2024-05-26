package service

import (
	"errors"

	"github.com/phelipperibeiro/lab-01-temperatureSystemByCEP/internal/domain/entity"
	"github.com/phelipperibeiro/lab-01-temperatureSystemByCEP/internal/port"
)

// func dd(data interface{}) {
// 	jsonData, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		fmt.Println("Erro ao serializar dados:", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(string(jsonData))
// }

type WeatherService struct {
	LocationGateway port.LocationGateway
	WeatherGateway  port.WeatherGateway
}

func (weatherService *WeatherService) GetWeather(zipcode string) (*entity.Weather, error) {
	if len(zipcode) != 8 {
		return nil, errors.New("invalid zipcode")
	}

	location, err := weatherService.LocationGateway.GetLocation(zipcode)
	if err != nil {
		return nil, err
	}

	if location == nil {
		return nil, errors.New("cannot find zipcode")
	}

	weather, err := weatherService.WeatherGateway.GetWeather(location.Localidade)
	if err != nil {
		return nil, errors.New("cannot get weather information")
	}

	return weather, nil
}

func NewWeatherService(locationGateway port.LocationGateway, weatherGateway port.WeatherGateway) *WeatherService {
	return &WeatherService{
		LocationGateway: locationGateway,
		WeatherGateway:  weatherGateway,
	}
}
