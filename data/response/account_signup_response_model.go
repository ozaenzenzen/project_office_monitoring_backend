package data

type AccountSignUpResponseModel struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *AccountUserDataSignUpModel `json:"Data"`
}

type AccountUserDataSignUpModel struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}
