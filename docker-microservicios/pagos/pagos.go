package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Payment struct {
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	Status          string  `json:"status"`
	CreationDate    string  `json:"creation_date"`
	TransactionID   string  `json:"transaction_id"`
	Source          string  `json:"source"`
	SSN             string  `json:"ssn"` // <- Añadido
}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/records", GetPayments)
	log.Println("Servicio de pagos escuchando en http://localhost:8003")
	log.Fatal(http.ListenAndServe(":8003", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
	// Leer el archivo JSON desde una ruta relativa
	data, err := ioutil.ReadFile("/data/payment_records.json")
	if err != nil {
		http.Error(w, "No se pudieron leer los pagos", http.StatusInternalServerError)
		log.Printf("Error leyendo payment_records.json: %v", err)
		return
	}

	// Decodificar el mapa de clientes a sus pagos
	var rawData map[string][]Payment
	if err := json.Unmarshal(data, &rawData); err != nil {
		http.Error(w, "Error al parsear el archivo de pagos", http.StatusInternalServerError)
		log.Printf("Error de formato JSON: %v", err)
		return
	}

	// Añadir el SSN a cada pago (modificando directamente rawData)
	for ssn, payments := range rawData {
		for i := range payments {
			payments[i].SSN = ssn
		}
		rawData[ssn] = payments
	}

	// Enviar el mapa completo como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rawData)
}


