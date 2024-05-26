package port

import "github.com/phelipperibeiro/lab-01-temperatureSystemByCEP/internal/domain/entity"

type LocationGateway interface {
	GetLocation(zipcode string) (*entity.Location, error)
}
