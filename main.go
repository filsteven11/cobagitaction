package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Endpoint untuk pengguna
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")

	// Endpoint untuk produk
	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")

	// Endpoint untuk transaksi
	router.HandleFunc("/transactions", controllers.GetAllTransactions).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
