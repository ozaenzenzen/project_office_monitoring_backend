package data

type CreateCompanyResponseModel struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *CreateCompanyDataModel `json:"Data"`
}

type CreateCompanyDataModel struct {
	CompanyStamp    string `json:"company_stamp"`
	CompanyName     string `json:"company_name"`
	CompanyEmail    string `json:"company_email"`
	CompanyPhone    string `json:"company_phone"`
	CompanyTypeuser uint   `json:"company_typeuser"`
	Picture         string `json:"picture"`
	City            string `json:"company_city"`
	State           string `json:"company_state"`
	Zip             string `json:"company_zip"`
	Address         string `json:"company_address"`
	Information     string `json:"company_information"`
}
