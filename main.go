package main

import (
	"payment-service/db_client"
	"payment-service/route"

	"github.com/gin-gonic/gin"
)

func main() {
	db_client.InitializeDBConnection()
	router := gin.New()
	route.GetPayment(router)

	route.UpdatePayment(router)
	router.Run(":4001")
}
