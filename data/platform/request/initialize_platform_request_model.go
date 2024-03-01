package data

type InitializePlatformRequestModel struct {
	PlatformName   string `gorm:"not null" json:"platform_name" binding:"required"`
	PlatformSecret string `gorm:"not null" json:"platform_secret" binding:"required"`
}
