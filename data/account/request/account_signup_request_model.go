package data

type AccountSignUpRequestModel struct {
	Name            string `gorm:"not null" json:"name"  binding:"required,max=30" validate:"required"`
	Email           string `gorm:"not null" json:"email" binding:"required" validate:"required"`
	NoReg           string `gorm:"not null" json:"noreg" binding:"required" validate:"required"`
	Phone           string `gorm:"not null" json:"phone"  binding:"required,max=14" validate:"required"`
	Password        string `gorm:"not null" json:"password" binding:"required" validate:"required"`
	ConfirmPassword string `gorm:"not null" json:"confirm_password" binding:"required" validate:"required"`
}
