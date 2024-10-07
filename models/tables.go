package models

import (
	"time"

	"github.com/google/uuid"
)

type Debt struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreditorID        string    `gorm:"not null" json:"creditor_id"`
	IndebtedID        string    `gorm:"not null" json:"indebted_id"`
	CreditorConfirmed *bool     `gorm:"not null" json:"creditor_confirmed"`
	IndebtedConfirmed *bool     `gorm:"not null" json:"indebted_confirmed"`
	Amount            uint      `gorm:"not null" json:"amount"`
	CreatedAt         time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt         time.Time `gorm:"not null" json:"updated_at"`
}

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username    string    `gorm:"not null" json:"username"`
	Password    string    `gorm:"not null" json:"password"`
	CardNumbers *[]uint
	TotalDebts  uint
}
