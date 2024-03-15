package data

type GetPlatformDataRequestModel struct {
	PlatformName string `gorm:"not null" json:"platform_name" binding:"required"`
	Password     string `gorm:"not null" json:"password" binding:"required"`
}
