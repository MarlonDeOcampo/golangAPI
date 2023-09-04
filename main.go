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
	router.Run(":4005")
}
