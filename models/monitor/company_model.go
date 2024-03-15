package models

import "time"

type CompanyModel struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	CompanyStamp    string    `json:"company_stamp"`
	CompanyName     string    `json:"company_name"`
	CompanyEmail    string    `json:"company_email"`
	CompanyNoReg    string    `json:"company_noreg"`
	CompanyPhone    string    `json:"company_phone"`
	Picture         string    `json:"picture"`
	City            string    `json:"city"`
	State           string    `json:"state"`
	Zip             string    `json:"zip"`
	Address         string    `json:"address"`
	Information     string    `json:"information"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
	CompanyTypeuser uint      `json:"company_typeuser"`
	PlatformName    string    `json:"platform_name"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
