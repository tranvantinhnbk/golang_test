package models

import (
	"time"

	"gorm.io/gorm"
)

// Account represents a user account with balance
type Account struct {
	ID        uint    `gorm:"primarykey"`
	Username  string  `gorm:"uniqueIndex;not null"`
	Balance   float64 `gorm:"not null;default:0"`
	Version   int     `gorm:"not null;default:1"` // For optimistic locking
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeUpdate is a GORM hook that increments version
func (a *Account) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("Balance") {
		a.Version++
	}
	return nil
}
