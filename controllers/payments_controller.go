package controller

import (
	"payment-service/db_client"
	model "payment-service/model"

	"github.com/gin-gonic/gin"
)

func GetPayments(c *gin.Context) {
	payment := []model.Payment{}
	db_client.DB.Find(&payment)
	c.JSON(200, &payment)
}
