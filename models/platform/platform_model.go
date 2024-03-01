package models

import "time"

type PlatformModel struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	CompanyName     string    `json:"company_name"`
	CompanyEmail    string    `json:"company_email"`
	CompanyNoReg    string    `json:"company_noreg"`
	CompanyPhone    string    `json:"company_phone"`
	ProfilePicture  string    `json:"profile_picture"`
	PlatformName    string    `json:"platform_name"`
	PlatformSecret  string    `json:"platform_secret"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
