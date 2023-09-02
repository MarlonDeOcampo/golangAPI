package main

import (
	"log"
	"net/http"
	"payment-service/controllers"
	"payment-service/db_client"
)

func main() {
	db_client.InitializeDBConnection()

	http.HandleFunc("/payment/create", controllers.CreatePayment)

	http.HandleFunc("/payment/update", controllers.UpdatePayment)

	http.HandleFunc("/payment/get-payments", controllers.GetAllPayments)

	http.HandleFunc("/payment/details", controllers.GetAllPayments)

	log.Fatal(http.ListenAndServe(":4001", nil))
}
