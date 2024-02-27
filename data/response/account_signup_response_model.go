package data

type AccountSingUpResponse struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *AccountUserDataResponseModel `json:"account_data"`
}

type AccountUserDataResponseModel struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}
