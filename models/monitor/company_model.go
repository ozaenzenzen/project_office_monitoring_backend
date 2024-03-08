package models

import "time"

type CompanyModel struct {
	ID                  uint      `json:"id" gorm:"primary_key"`
	UserNoReg           string    `json:"user_id"`
	CompanyName         string    `json:"company_name"`
	CompanyPlatformName string    `json:"company_platform_name"`
	Picture             string    `json:"picture"`
	Location            string    `json:"location"`
	Address             string    `json:"address"`
	Information         string    `json:"information"`
	Latitude            string    `json:"latitude"`
	Longitude           string    `json:"longitude"`
	CreatedAt           time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt           time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
