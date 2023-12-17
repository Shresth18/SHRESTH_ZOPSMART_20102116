package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var cars []Car = []Car{}

// Defined CarType struct
type Car struct {
	LicensePlate string `json:"licenseplate"`
	Model        string `json:"model"`
	Status       string `json:"status"` // e.g., parked, repair_in_progress, repaired
	Carowner     User   `json:"carowner"`
}

// Defined user struct
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

// ADDCARS FUNCTION FOR ENTRY OF CAR IN GARAGE
func addCars(w http.ResponseWriter, r *http.Request) {
	var newCar Car
	json.NewDecoder(r.Body).Decode(&newCar)

	w.Header().Set("Content-Type", "application/json")

	cars = append(cars, newCar)
	json.NewEncoder(w).Encode(cars)
}

// GETALLCAR FUNCTION FOR LIST OF ALL CARS
func getallcars(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(cars)
}

// GETCAR FUNCTION FOR INFO OF SPECIFIC CAR
func getcar(q http.ResponseWriter, r *http.Request) {
	var idparam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("id could not be converted to integer"))
		return
	}
	if id >= len(cars) {
		q.WriteHeader(404)
		q.Write([]byte("no profile found with specific id"))
		return
	}
	car := cars[id]
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(car)
}

// UPDDATE CAR FUNCTION FOR UPDATE ROUTE
func Updatecar(q http.ResponseWriter, r *http.Request) {
	var idparam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("id could not be converted to integer"))
		return
	}
	if id >= len(cars) {
		q.WriteHeader(404)
		q.Write([]byte("no profile found with specific id"))
		return
	}
	var updatecar Car
	json.NewDecoder(r.Body).Decode(&updatecar)
	cars[id] = updatecar
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(updatecar)
}

// REMOVECAR FUNCTION TO REMOVE THE CAR FROM GARAGE WHEN SERVICE DONE
func removecar(q http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idparam := vars["id"]
	id, err := strconv.Atoi(idparam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("id could not be converted to integer"))
		return
	}
	if id < 0 || id >= len(cars) {
		q.WriteHeader(404)
		q.Write([]byte("no profile found with specific id"))
		return
	}

	cars = append(cars[:id], cars[id+1:]...)
	q.WriteHeader(200)
}

// MAIN FUNCTION
func main() {
	router := mux.NewRouter()

	// ROUTE TO ADD CAR IN THE GARAGE
	router.HandleFunc("/cars", addCars).Methods("POST")

	// ROUTE TO GET LIST OF ALL THE CARS
	router.HandleFunc("/cars", getallcars).Methods("GET")

	// ROUTE TO GET INFO ABOUT SPECIFIC CAR IN GARAGE
	router.HandleFunc("/cars/{id}", getcar).Methods("GET")

	// ROUTE TO UPDATE THE STATUS OF CAR
	router.HandleFunc("/cars/{id}", Updatecar).Methods("PUT")

	// ROUTE TO REMOVE THE CAR FROM GARAGE
	router.HandleFunc("/cars/{id}", removecar).Methods("DELETE")

	http.ListenAndServe(":5000", router)
}

// SHRESTH AGARWAL
// 20102116
// Shresth984@gmail.com
