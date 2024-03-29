package data

type GetUserDataResponseModel struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    *GetUserDataModel `json:"Data"`
}

type GetUserDataModel struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	NoReg          string `json:"noreg"`
	Jabatan        string `json:"jabatan"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
}
