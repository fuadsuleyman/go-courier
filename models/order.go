package models

import (
	"time"
)

type Order struct {
	// gorm.Model
	ID       int64 `json:"Id" gorm:"primary_key"`
	ClientID int   `json:"client_id"`
	// Client        ClientID
	CookId int64 `json:"cook_id"`
	// Cook      []Cook
	CourierId int64 `json:"courier_id"`
	// Courier       []Courier
	DeliveryInformationId int64  `json:"delivery_information_id"`
	Status                string `gorm:"type:varchar(50); default:'order was placed'" json:"status"`
	CourierStatus         string `gorm:"type:varchar(50); default:'no courier'" json:"courier_status"`
	IsActive              bool   `json:"is_active" gorm:"not null"`
	// IsRejected            bool      `json:"is_rejected"`
	RejectReason string `json:"reject_reason" gorm:"type:varchar(250)"`
	RejectReasonCourier string `json:"reject_reason" gorm:"type:varchar(250)"`
	// Items                 []Item    `json:"items" gorm:"foreignkey:ID"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

func (Order) TableName() string {
	return "orders_order"
}
