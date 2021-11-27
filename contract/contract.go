package contract

import (
	"gorm.io/gorm"
)

//response
type JSONResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Massage string      `json:"massage"`
	Data    interface{} `json:"data"`
}

//entity
type Users struct {
	gorm.Model
	Username string `json:"username"`
}

//request
type Chat struct {
	gorm.Model
	User   string `json:"user"`
	SentTo string `json:"sent_to"`
	Text   string `json:"text"`
}
