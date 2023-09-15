package route

import (
	"encoding/json"
	"fmt"
	controller "payment-service/controllers"
	producer "payment-service/kafka/producer"

	"github.com/gin-gonic/gin"
)

func GetPayment(router *gin.Engine) {
	response := router.GET("/payment/getpayments", controller.GetPayments)
	topicName := "payments"
	stringJson, _ := json.Marshal(response)
	fmt.Print(response)
	producer.Produce(string(stringJson), topicName)

}

func UpdatePayment(router *gin.Engine) {
	router.GET("/payment/updatepayments", controller.UpdatePayments)
}
