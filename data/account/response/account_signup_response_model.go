package data

type AccountSignUpResponseModel struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *AccountUserDataSignUpModel `json:"Data"`
}

type AccountUserDataSignUpModel struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	NoReg     string `json:"noreg"`
	UserStamp string `json:"user_stamp"`
	Jabatan   string `json:"jabatan"`
	Phone     string `json:"phone"`
	Typeuser  uint   `json:"typeuser"`
}
