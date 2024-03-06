package data

type AccountSignInRequestModel struct {
	Email    string `gorm:"not null" json:"email"  binding:"required" validate:"required"`
	Password string `gorm:"not null" json:"password"  binding:"required" validate:"required"`
}
