package models

import "time"

type AccountUserModel struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	UserStamp       string    `json:"user_stamp"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	NoReg           string    `json:"noreg"`
	Phone           string    `json:"phone"`
	ProfilePicture  string    `json:"profile_picture"`
	Jabatan         string    `json:"jabatan"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
	Typeuser        uint      `json:"typeuser"`
	PlatformName    string    `json:"platform_name"`
	CompanyStamp    string    `json:"company_stamp"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
