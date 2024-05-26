package service

import (
	"errors"
	"testing"

	"github.com/phelipperibeiro/lab-01-temperatureSystemByCEP/internal/domain/entity"

	"github.com/stretchr/testify/assert"
)

// ///////////////////////////////////////////
// ////////// mockLocationGateway ////////////
// ///////////////////////////////////////////
type mockLocationGateway struct {
	mockGetLocation func(zipcode string) (*entity.Location, error)
}

func (mock *mockLocationGateway) GetLocation(zipcode string) (*entity.Location, error) {
	return mock.mockGetLocation(zipcode)
}

// ////////////////////////////////////////////
// //////////// mockWeatherGateway ////////////
// ////////////////////////////////////////////

type mockWeatherGateway struct {
	mockGetWeather func(city string) (*entity.Weather, error)
}

func (mock *mockWeatherGateway) GetWeather(city string) (*entity.Weather, error) {
	return mock.mockGetWeather(city)
}

// ////////////////////////////////////////////
// ///////////////// TESTS ////////////////////
// ////////////////////////////////////////////

func TestWeatherService_GetWeather_Success(t *testing.T) {
	locationGateway := &mockLocationGateway{
		mockGetLocation: func(zipcode string) (*entity.Location, error) {
			return &entity.Location{Localidade: "Test City"}, nil
		},
	}
	weatherGateway := &mockWeatherGateway{
		mockGetWeather: func(city string) (*entity.Weather, error) {
			return &entity.Weather{TempC: 25, TempF: 77, TempK: 298}, nil
		},
	}
	weatherService := NewWeatherService(locationGateway, weatherGateway)

	weather, err := weatherService.GetWeather("12345678")

	assert.NoError(t, err)
	assert.NotNil(t, weather)
	assert.Equal(t, 25.0, weather.TempC)
	assert.Equal(t, 77.0, weather.TempF)
	assert.Equal(t, 298.0, weather.TempK)
}

func TestWeatherService_GetWeather_InvalidZipcode(t *testing.T) {
	locationGateway := &mockLocationGateway{
		mockGetLocation: func(zipcode string) (*entity.Location, error) {
			return nil, nil
		},
	}
	weatherGateway := &mockWeatherGateway{}
	weatherService := NewWeatherService(locationGateway, weatherGateway)

	weather, err := weatherService.GetWeather("invalid")

	assert.Error(t, err)
	assert.Nil(t, weather)
	assert.Equal(t, "invalid zipcode", err.Error())
}

func TestWeatherService_GetWeather_ZipcodeNotFound(t *testing.T) {
	locationGateway := &mockLocationGateway{
		mockGetLocation: func(zipcode string) (*entity.Location, error) {
			return nil, errors.New("cannot find zipcode")
		},
	}
	weatherGateway := &mockWeatherGateway{}
	weatherService := NewWeatherService(locationGateway, weatherGateway)

	weather, err := weatherService.GetWeather("87654321")

	assert.Error(t, err)
	assert.Nil(t, weather)
	assert.Equal(t, "cannot find zipcode", err.Error())
}

func TestWeatherService_GetWeather_WeatherNotFound(t *testing.T) {
	locationGateway := &mockLocationGateway{
		mockGetLocation: func(zipcode string) (*entity.Location, error) {
			return &entity.Location{Localidade: "Test City"}, nil
		},
	}
	weatherGateway := &mockWeatherGateway{
		mockGetWeather: func(city string) (*entity.Weather, error) {
			return nil, errors.New("cannot get weather information")
		},
	}
	weatherService := NewWeatherService(locationGateway, weatherGateway)

	weather, err := weatherService.GetWeather("12345678")

	assert.Error(t, err)
	assert.Nil(t, weather)
	assert.Equal(t, "cannot get weather information", err.Error())
}
