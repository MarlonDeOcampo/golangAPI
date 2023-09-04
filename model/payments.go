package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Paymentid *int    `json:"paymentid" gorm:"primary_key"`
	Userid    *int    `json:"userid"`
	Lastname  *string `json:"lastname"`
	Firstname *string `json:"firstname"`
}
