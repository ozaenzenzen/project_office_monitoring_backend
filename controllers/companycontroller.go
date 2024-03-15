package controllers

import (
	"net/http"
	helper "project_office_monitoring_backend/helper"
	account "project_office_monitoring_backend/models/account"
	company "project_office_monitoring_backend/models/monitor"

	req "project_office_monitoring_backend/data/company/request"
	resp "project_office_monitoring_backend/data/company/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func CreateCompany(c *gin.Context) {
	var reqCreateCompany req.CreateCompantRequestModel
	if err := c.ShouldBindJSON(&reqCreateCompany); err != nil {
		c.JSON(http.StatusBadRequest, resp.CreateCompanyResponseModel{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(reqCreateCompany); err != nil {
		c.JSON(http.StatusBadRequest, resp.CreateCompanyResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Data tidak lengkap",
			Data:    nil,
		})
		return
	}

	header_platformkey := c.Request.Header.Get("platformkey")
	if header_platformkey == "" {
		c.JSON(http.StatusBadRequest, resp.CreateCompanyResponseModel{
			Status:  http.StatusBadRequest,
			Message: "invalid platformkey",
			Data:    nil,
		})
		return
	}

	isValid, err := helper.VerifyPlatformToken(header_platformkey)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.CreateCompanyResponseModel{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if isValid {
		hashPw, errPw := helper.HashPassword(reqCreateCompany.Password)
		if errPw != nil {
			c.JSON(http.StatusInternalServerError, resp.CreateCompanyResponseModel{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		hashCpw, errCpw := helper.HashPassword(reqCreateCompany.ConfirmPassword)
		if errCpw != nil {
			c.JSON(http.StatusInternalServerError, resp.CreateCompanyResponseModel{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		stampToken := uuid.New().String()
		companyStamp := reqCreateCompany.CompanyName + stampToken

		createCompanyResponsePayload := company.CompanyModel{
			CompanyStamp:    companyStamp,
			CompanyName:     reqCreateCompany.CompanyName,
			CompanyNoReg:    stampToken,
			CompanyEmail:    reqCreateCompany.CompanyEmail,
			CompanyPhone:    reqCreateCompany.CompanyPhone,
			Picture:         reqCreateCompany.Picture,
			City:            reqCreateCompany.City,
			State:           reqCreateCompany.State,
			Zip:             reqCreateCompany.Zip,
			Address:         reqCreateCompany.Address,
			Information:     reqCreateCompany.Information,
			Password:        hashPw,
			ConfirmPassword: hashCpw,
			CompanyTypeuser: 0,
			PlatformName:    header_platformkey,
		}

		db := c.MustGet("db").(*gorm.DB)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, resp.CreateCompanyResponseModel{
				Status:  http.StatusInternalServerError,
				Message: db.Error.Error(),
				Data:    nil,
			})
			return
		}

		result := db.FirstOrCreate(&createCompanyResponsePayload, account.AccountUserModel{Email: reqCreateCompany.CompanyEmail})

		if result.Value == nil && result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, resp.CreateCompanyResponseModel{
				Status:  http.StatusBadRequest,
				Message: "Record found",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, resp.CreateCompanyResponseModel{
			Status:  http.StatusCreated,
			Message: "Company account created successfully",
			Data: &resp.CreateCompanyDataModel{
				CompanyStamp:    createCompanyResponsePayload.CompanyStamp,
				CompanyName:     createCompanyResponsePayload.CompanyName,
				CompanyEmail:    createCompanyResponsePayload.CompanyEmail,
				CompanyPhone:    createCompanyResponsePayload.CompanyPhone,
				CompanyTypeuser: createCompanyResponsePayload.CompanyTypeuser,
				Picture:         createCompanyResponsePayload.Picture,
				City:            createCompanyResponsePayload.City,
				State:           createCompanyResponsePayload.State,
				Zip:             createCompanyResponsePayload.Zip,
				Address:         createCompanyResponsePayload.Address,
				Information:     createCompanyResponsePayload.Information,
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, resp.CreateCompanyResponseModel{
			Status:  http.StatusBadRequest,
			Message: "invalid platformkey",
			Data:    nil,
		})
		return
	}
}

func GetCompanyData(c *gin.Context) {
	//
}
