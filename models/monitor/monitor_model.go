package models

import "time"

type MonitorModel struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	Name            string    `json:"name"`
	Picture         string    `json:"picture"`
	Email           string    `json:"email"`
	NoReg           string    `json:"noreg"`
	Jabatan         string    `json:"jabatan"`
	Phone           string    `json:"phone"`
	Typeuser        uint      `json:"typeuser"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
