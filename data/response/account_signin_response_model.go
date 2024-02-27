package data

type AccountUserSignInResponse struct {
	Status   int                  `json:"status"`
	Message  string               `json:"message"`
	UserData *UserDataModelSignIn `json:"userdata"`
}

type UserDataModelSignIn struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}
