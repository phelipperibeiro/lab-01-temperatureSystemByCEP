package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/phelipperibeiro/lab-01-temperatureSystemByCEP/internal/application/service"

	"github.com/phelipperibeiro/lab-01-temperatureSystemByCEP/internal/adapter/gateway"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("Starting server...")
	port := "8080"

	locationGateway := gateway.NewViaCepGateway()
	weatherGateway := gateway.NewWeatherAPIGateway()
	weatherService := service.NewWeatherService(locationGateway, weatherGateway)

	router := mux.NewRouter()
	router.HandleFunc("/weather/{zipcode}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		zipcode := vars["zipcode"]

		weather, err := weatherService.GetWeather(zipcode)
		if err != nil {
			switch err.Error() {
			case "invalid zipcode":
				http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			case "cannot find zipcode":
				http.Error(w, "cannot find zipcode", http.StatusNotFound)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weather)
	}).Methods("GET")

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

// go run cmd/server/main.go
