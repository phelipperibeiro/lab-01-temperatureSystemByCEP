package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/phelipperibeiro/lab-01-temperatureSystemByCEP/internal/domain/entity"
)

type ViaCepGateway struct{}

func (v *ViaCepGateway) GetLocation(zipcode string) (*entity.Location, error) {

	resp, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", zipcode))
	if err != nil {
		fmt.Println("Erro ao buscar dados:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot find zipcode")
	}

	var location entity.Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	if location.Localidade == "" {
		return nil, fmt.Errorf("cannot find zipcode")
	}

	return &location, nil
}

func NewViaCepGateway() *ViaCepGateway {
	return &ViaCepGateway{}
}
