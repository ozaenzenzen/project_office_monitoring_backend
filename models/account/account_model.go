package models

import "time"

type AccountUserModel struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	Name            string    `json:"name"`
	ProfilePicture  string    `json:"profile_picture"`
	Email           string    `json:"email"`
	NoReg           string    `json:"noreg"`
	UserStamp       string    `json:"user_stamp"`
	Jabatan         string    `json:"jabatan"`
	Phone           string    `json:"phone"`
	Typeuser        uint      `json:"typeuser"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
