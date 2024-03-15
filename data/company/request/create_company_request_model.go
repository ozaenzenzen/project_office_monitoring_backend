package data

type CreateCompantRequestModel struct {
	CompanyName     string `gorm:"not null" json:"company_name" binding:"required"`
	CompanyEmail    string `gorm:"not null" json:"company_email" binding:"required"`
	CompanyPhone    string `gorm:"not null" json:"company_phone" binding:"required"`
	Picture         string `json:"picture"`
	City            string `gorm:"not null" json:"company_city" binding:"required"`
	State           string `gorm:"not null" json:"company_state" binding:"required"`
	Zip             string `gorm:"not null" json:"company_zip" binding:"required"`
	Address         string `gorm:"not null" json:"company_address" binding:"required"`
	Information     string `gorm:"not null" json:"company_information" binding:"required"`
	Password        string `gorm:"not null" json:"password" binding:"required"`
	ConfirmPassword string `gorm:"not null" json:"confirmPassword" binding:"required"`
}
