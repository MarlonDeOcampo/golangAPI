package route

import (
	controller "payment-service/controllers"

	"github.com/gin-gonic/gin"
)

func GetPayment(router *gin.Engine) {
	router.GET("/payment/getpayments", controller.GetPayments)
}
