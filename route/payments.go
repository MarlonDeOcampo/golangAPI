package route

import (
	"encoding/json"
	controller "payment-service/controllers"
	producer "payment-service/kafka/producer"

	"github.com/gin-gonic/gin"
)

func GetPayment(router *gin.Engine) {
	var response = router.GET("/payment/getpayments", controller.GetPayments)
	topicName := "test"
	stringJson, _ := json.Marshal(response)
	producer.Produce(string(stringJson), topicName)
}

func UpdatePayment(router *gin.Engine) {
	router.GET("/payment/updatepayments", controller.UpdatePayments)
}
