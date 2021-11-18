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
	Phone          string          `gorm:"type:varchar(12)" json:"phone"`
	Transport      string          `gorm:"type:varchar(150)" json:"transport"`
	WorkExperience uint            `gorm:"type:varchar(150)" json:"work_experience"` // pythonda integer field
	Rating         decimal.Decimal `gorm:"type:decimal(2,1);" json:"rating"`
	Isavailable    bool            `json:"is_available"`
	Location       string          `gorm:"type:varchar(255)" json:"location"`
	Created_at     time.Time
	Updated_at     time.Time
}
