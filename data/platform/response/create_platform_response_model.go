package data

type CreatePlatformResponseModel struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *CreatePlatformDataModel `json:"Data"`
}

type CreatePlatformDataModel struct {
	CompanyName    string `gorm:"not null" json:"company_name" binding:"required"`
	CompanyEmail   string `gorm:"not null" json:"company_email" binding:"required"`
	CompanyNoReg   string `gorm:"not null" json:"company_noreg" binding:"required"`
	CompanyPhone   string `gorm:"not null" json:"company_phone" binding:"required"`
	PlatformName   string `gorm:"not null" json:"platform_name" binding:"required"`
	PlatformSecret string `gorm:"not null" json:"platform_secret" binding:"required"`
}
