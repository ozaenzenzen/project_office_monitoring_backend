package data

type AccountUserSignInRequest struct {
	Email    string `gorm:"not null" json:"email"  binding:"required"`
	Password string `gorm:"not null" json:"password"  binding:"required"`
}
