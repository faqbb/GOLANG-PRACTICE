package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Spendings struct {
	ID     string  `json: "id"`
	Name   string  `json: "name"`
	Amount float64 `json: "amount"`
	Date   string  `json: "date"`
}

var spendings []Spendings

func getSpendings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spendings)
}

func getSpendingsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Buscamos id
	for _, item := range spendings {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode((&Spendings{}))
}

func main() {
	newSpendings := []Spendings{
		{
			ID:     "1",
			Name:   "Groceries",
			Amount: 50.75,
			Date:   "2024-01-01",
		},
		{
			ID:     "2",
			Name:   "Rent",
			Amount: 1200.00,
			Date:   "2024-02-01",
		},
		{
			ID:     "3",
			Name:   "Utilities",
			Amount: 150.30,
			Date:   "2024-03-01",
		},
		{
			ID:     "4",
			Name:   "Transportation",
			Amount: 75.20,
			Date:   "2024-04-01",
		},
		{
			ID:     "5",
			Name:   "Entertainment",
			Amount: 200.50,
			Date:   "2024-05-01",
		},
	}
	spendings = append(spendings, newSpendings...)

	// Definimos router
	router := mux.NewRouter()

	// Definimos endpoints
	router.HandleFunc("/getSpendings", getSpendings).Methods("GET")
	router.HandleFunc("/getSpendings/{id}", getSpendingsById).Methods("GET")
	// router.HandleFunc("/newSpending", newSpending).Methods("POST")
	// router.HandleFunc("/getSpendings/{id}", updateSpendings).Methods("PUT")
	// router.HandleFunc("/getSpendings/{id}", deleteSpendings).Methods("DELETE")

	// Define el server
	address := ":8080"
	server := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Inicia el server
	fmt.Printf("Escuchando en %s. Presiona CTRL + C para salir", address)
	log.Fatal(server.ListenAndServe())
}
