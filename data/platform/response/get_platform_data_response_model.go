package data

type GetPlatformDataResponseModel struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *GetPlatformDataModel `json:"Data"`
}

type GetPlatformDataModel struct {
	PlatformName   string `gorm:"not null" json:"platform_name" binding:"required"`
	PlatformSecret string `gorm:"not null" json:"platform_secret" binding:"required"`
	PlatformType   string `gorm:"not null" json:"platform_type" binding:"required"`
	MaxCompany     uint   `gorm:"not null" json:"max_company" binding:"required"`
	MaxUser        uint   `gorm:"not null" json:"max_user" binding:"required"`
}
