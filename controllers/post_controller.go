package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

var DBClient *sql.DB

type Payment struct {
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Payment")
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Payment")
}

func GetAllPayments(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get All Payment")
}

func GetDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Details Payment")
}
