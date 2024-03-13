package data

type AccountSignInResponseModel struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *AccountUserDataSignInModel `json:"Data"`
}

type AccountUserDataSignInModel struct {
	UserStamp string `json:"user_stamp"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	NoReg     string `json:"noreg"`
	Jabatan   string `json:"jabatan"`
	Phone     string `json:"phone"`
	Typeuser  uint   `json:"typeuser"`
	Token     string `json:"token"`
}
