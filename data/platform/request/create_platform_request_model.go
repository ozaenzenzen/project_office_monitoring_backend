package data

type CreatePlatformRequestModel struct {
	Name            string `gorm:"not null" json:"name" binding:"required"`
	Email           string `gorm:"not null" json:"email" binding:"required"`
	Password        string `gorm:"not null" json:"password" binding:"required"`
	ConfirmPassword string `gorm:"not null" json:"confirmPassword" binding:"required"`
}
