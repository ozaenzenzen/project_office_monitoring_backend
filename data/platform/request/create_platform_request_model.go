package data

type CreatePlatformRequestModel struct {
	CompanyName     string `gorm:"not null" json:"company_name" binding:"required"`
	CompanyEmail    string `gorm:"not null" json:"company_email" binding:"required"`
	CompanyNoReg    string `gorm:"not null" json:"company_noreg" binding:"required"`
	CompanyPhone    string `gorm:"not null" json:"company_phone" binding:"required"`
	Password        string `gorm:"not null" json:"password" binding:"required"`
	ConfirmPassword string `gorm:"not null" json:"confirmPassword" binding:"required"`
}
