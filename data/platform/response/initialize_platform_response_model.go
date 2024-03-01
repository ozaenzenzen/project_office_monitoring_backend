package data

type InitializePlatformResponseModel struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *InitializePlatformDataModel `json:"Data"`
}

type InitializePlatformDataModel struct {
	PlatformName   string `gorm:"not null" json:"platform_name" binding:"required"`
	PlatformSecret string `gorm:"not null" json:"platform_secret" binding:"required"`
	PlatformKey    string `gorm:"not null" json:"platform_key" binding:"required"`
}
