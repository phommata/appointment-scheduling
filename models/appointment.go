package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Appointment struct {
	gorm.Model
	TrainerID   int
	UserID   	int
	StartsAt   	time.Time `gorm:"type:datetime"` // atom format time.stdLongYear-time.stdZeroMonth-time.stdZeroDayTstdHour:stdZeroMinute:stdZeroSecondstdNumColonSecondsTZ
	EndsAt   	time.Time `gorm:"type:datetime"` // atom format time.stdLongYear-time.stdZeroMonth-time.stdZeroDayTstdHour:stdZeroMinute:stdZeroSecondstdNumColonSecondsTZ
}