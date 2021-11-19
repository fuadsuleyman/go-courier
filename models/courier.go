package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Courier struct {
	// in first creation
	Id        uint   `json:"id" gorm:"primary_key"`
	Username  string `validate:"required" gorm:"unique; type:varchar(50); not null" json:"username"`
	Usertype  string `validate:"required" gorm:"type:varchar(50); not null"json:"usertype"`
	Firstname string `validate:"required" gorm:"type:varchar(50); not null" json:"firstname"`
	Lastname  string `validate:"required" gorm:"type:varchar(50); not null" json:"lastname"`
	Email     string `validate:"required" gorm:"unique; type:varchar(50); not null" json:"email"`
	// in update
	Patronymic     string          `gorm:"type:varchar(50)" json:"patronymic"`
	Phone          string          `gorm:"type:varchar(12)" json:"phone"`
	Transport      string          `gorm:"type:varchar(150)" json:"transport"`
	WorkExperience int             `json:"work_experience"` // pythonda integer field
	Rating         decimal.Decimal `gorm:"type:decimal(2,1)" json:"rating"`
	IsAvailable    bool            `json:"is_available"`
	Location       string          `gorm:"type:varchar(255)" json:"location"`
	CreatedAt      time.Time       // `gorm:"default:current_timestamp"`
	UpdatedAt      time.Time       // `gorm:"default:current_timestamp"`
}
