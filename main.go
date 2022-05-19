package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	getBikesEndPoint = "/bikes"
	getBikeEndPoint  = "/bike/{model}"
)
var bikes []Bike

type Bike struct {
	Model        string `json:"model"`
	Displacement string `json:"displacement"`
	Brand        *Brand `json:"brand"`
}

type Brand struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func GetBikes(responseWrite http.ResponseWriter, request *http.Request) {
	responseWrite.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(responseWrite).Encode(bikes)
	if err != nil {
		log.Fatal(err)
	}
}

func GetBike(responseWrite http.ResponseWriter, request *http.Request) {
	responseWrite.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(request)
	for _, bike := range bikes {
		if bike.Model != parameters["model"] {
			continue
		}
		err := json.NewEncoder(responseWrite).Encode(bike)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
}

func DeleteBike(responseWrite http.ResponseWriter, request *http.Request) {
	responseWrite.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(request)

	for i, bike := range bikes {
		if bike.Model != parameters["model"] {
			continue
		}
		bikes = append(bikes[:i], bikes[i+1:]...)
	}
	err := json.NewEncoder(responseWrite).Encode(bikes)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateBike(responseWrite http.ResponseWriter, request *http.Request) {
	responseWrite.Header().Set("Content-Type", "application/json")
	var bike Bike
	err := json.NewDecoder(request.Body).Decode(&bike)
	if err != nil {
		log.Fatal(err)
	}
	bikes = append(bikes, bike)

	err = json.NewEncoder(responseWrite).Encode(bikes)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateBike(responseWrite http.ResponseWriter, request *http.Request) {
	responseWrite.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(request)

	for i, bike := range bikes {
		if bike.Model != parameters["model"] {
			continue
		}
		bikes = append(bikes[:i], bikes[i+1:]...)
		var bike Bike
		err := json.NewDecoder(request.Body).Decode(&bike)
		if err != nil {
			log.Fatal(err)
		}
		bike.Model = parameters["model"]
		bikes = append(bikes, bike)
	}
	err := json.NewEncoder(responseWrite).Encode(bikes)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := mux.NewRouter()
	bikes = append(bikes,
		Bike{Model: "VFR800V-TEC", Displacement: "800", Brand: &Brand{Name: "Honda", Country: "Japan"}},
		Bike{Model: "mt-03", Displacement: "300", Brand: &Brand{Name: "Yamaha", Country: "Japan"}})

	router.HandleFunc(getBikesEndPoint, GetBikes).Methods("GET")
	router.HandleFunc(getBikeEndPoint, GetBike).Methods("GET")
	router.HandleFunc(getBikesEndPoint, CreateBike).Methods("POST")
	router.HandleFunc(getBikeEndPoint, UpdateBike).Methods("PUT")
	router.HandleFunc(getBikeEndPoint, DeleteBike).Methods("DELETE")

	log.Printf("Strarting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
