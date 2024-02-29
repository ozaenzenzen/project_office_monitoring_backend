package data

type AccountSignInResponseModel struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *AccountUserDataSignInModel `json:"Data"`
}

type AccountUserDataSignInModel struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}
