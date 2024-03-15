package models

import "time"

type PlatformModel struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	PlatformName    string    `json:"platform_name"`
	PlatformSecret  string    `json:"platform_secret"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
	PlatformType    string    `json:"platform_type"`
	MaxCompany      uint      `json:"max_company"`
	MaxUser         uint      `json:"max_user"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
