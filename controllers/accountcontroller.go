package controllers

import (
	"net/http"
	req "project_office_monitoring_backend/data/account/request"
	resp "project_office_monitoring_backend/data/account/response"
	helper "project_office_monitoring_backend/helper"
	account "project_office_monitoring_backend/models/account"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

func SignUpAccount(c *gin.Context) {
	var accountInput req.AccountSignUpRequestModel
	if err := c.ShouldBindJSON(&accountInput); err != nil {
		c.JSON(http.StatusBadRequest, resp.AccountSignUpResponseModel{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(accountInput); err != nil {
		// log.Println(fmt.Sprintf("error log2: %s", err))
		c.JSON(http.StatusBadRequest, resp.AccountSignUpResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Data tidak lengkap",
			Data:    nil,
		})
		return
	}

	header_platformkey := c.Request.Header.Get("platformkey")
	if header_platformkey == "" {
		c.JSON(http.StatusBadRequest, resp.AccountSignUpResponseModel{
			Status:  http.StatusBadRequest,
			Message: "invalid platformkey",
			Data:    nil,
		})
		return
	}

	isValid, err := helper.VerifyPlatformToken(header_platformkey)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.AccountSignUpResponseModel{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if isValid {
		hashPw, errPw := helper.HashPassword(accountInput.Password)
		if errPw != nil {
			c.JSON(http.StatusInternalServerError, resp.AccountSignUpResponseModel{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		hashCpw, errCpw := helper.HashPassword(accountInput.ConfirmPassword)
		if errCpw != nil {
			c.JSON(http.StatusInternalServerError, resp.AccountSignUpResponseModel{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		accountResponsePayload := account.AccountUserModel{
			Name:            accountInput.Name,
			Email:           accountInput.Email,
			NoReg:           accountInput.NoReg,
			Phone:           accountInput.Phone,
			Password:        hashPw,
			ConfirmPassword: hashCpw,
		}

		db := c.MustGet("db").(*gorm.DB)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, resp.AccountSignUpResponseModel{
				Status:  http.StatusInternalServerError,
				Message: db.Error.Error(),
				Data:    nil,
			})
			return
		}

		result := db.FirstOrCreate(&accountResponsePayload, account.AccountUserModel{Email: accountInput.Email})

		if result.Value == nil && result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, resp.AccountSignUpResponseModel{
				Status:  http.StatusBadRequest,
				Message: "Record found",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, resp.AccountSignUpResponseModel{
			Status:  http.StatusCreated,
			Message: "Account created successfully",
			Data: &resp.AccountUserDataSignUpModel{
				ID:      accountResponsePayload.ID,
				Name:    accountInput.Name,
				Email:   accountInput.Email,
				NoReg:   accountResponsePayload.NoReg,
				Jabatan: accountResponsePayload.Jabatan,
				Phone:   accountInput.Phone,
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, resp.AccountSignUpResponseModel{
			Status:  http.StatusBadRequest,
			Message: "invalid platformkey",
			Data:    nil,
		})
		return
	}

}

func SignInAccount(c *gin.Context) {
	var table account.AccountUserModel
	var dataUser req.AccountSignInRequestModel
	if err := c.ShouldBindJSON(&dataUser); err != nil {
		c.JSON(http.StatusBadRequest, resp.AccountSignInResponseModel{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	header_platformkey := c.Request.Header.Get("platformkey")
	if header_platformkey == "" {
		c.JSON(http.StatusBadRequest, resp.AccountSignInResponseModel{
			Status:  http.StatusBadRequest,
			Message: "invalid platformkey",
			Data:    nil,
		})
		return
	}

	isValid, err := helper.VerifyPlatformToken(header_platformkey)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.AccountSignUpResponseModel{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if isValid {
		db := c.MustGet("db").(*gorm.DB)

		result := db.Where("email = ?", dataUser.Email).First(&table)

		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, resp.AccountSignInResponseModel{
				Status:  http.StatusUnauthorized,
				Message: "Invalid user email or password",
				Data:    nil,
			})
			return
		}

		checkHashPw := helper.CheckPasswordHash(dataUser.Password, table.Password)
		if !checkHashPw {
			c.JSON(http.StatusUnauthorized, resp.AccountSignInResponseModel{
				Status:  http.StatusUnauthorized,
				Message: "Invalid user email or password",
			})
			return
		}

		tokenString, err := helper.GenerateJWTToken(strconv.FormatUint(uint64(table.ID), 10), dataUser.Email)

		if err != nil {
			c.JSON(http.StatusNotFound, resp.AccountSignInResponseModel{
				Status:  http.StatusInternalServerError,
				Message: "Failed to generate token",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, resp.AccountSignInResponseModel{
			Status:  200,
			Message: "Account SignIn Successfully",
			Data: &resp.AccountUserDataSignInModel{
				ID:       table.ID,
				Name:     table.Name,
				Email:    dataUser.Email,
				NoReg:    table.NoReg,
				Jabatan:  table.Jabatan,
				Phone:    table.Phone,
				Typeuser: table.Typeuser,
				Token:    tokenString,
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, resp.AccountSignInResponseModel{
			Status:  http.StatusBadRequest,
			Message: "invalid platformkey",
			Data:    nil,
		})
		return
	}

}

func GetUserData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	headertoken := c.Request.Header.Get("token")

	if headertoken == "" {
		c.JSON(http.StatusBadRequest, resp.GetUserDataResponseModel{
			Status:  http.StatusBadRequest,
			Message: "token empty",
			Data:    nil,
		})
		return
	}
	isValid, err := helper.VerifyToken(headertoken)

	if isValid {
		if err != nil {
			c.JSON(http.StatusBadRequest, resp.GetUserDataResponseModel{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		var userData account.AccountUserModel

		tokenRaw, err := helper.DecodeJWTToken(headertoken)
		// fmt.Printf("\ntoken raw %v", tokenRaw)
		if err != nil {
			c.JSON(http.StatusBadRequest, resp.GetUserDataResponseModel{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		emails := tokenRaw["email"].(string)

		// if err := db.Where("id = ?", c.Param("id")).First(&userData).Error; err != nil {
		if err := db.Where("email = ?", emails).First(&userData).Error; err != nil {
			c.JSON(http.StatusBadRequest, resp.GetUserDataResponseModel{
				Status:  http.StatusBadRequest,
				Message: "User Data Not Found",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, resp.GetUserDataResponseModel{
			Status:  200,
			Message: "get user data success",
			Data: &resp.GetUserDataModel{
				ID:             userData.ID,
				Name:           userData.Name,
				Email:          userData.Email,
				NoReg:          userData.NoReg,
				Jabatan:        userData.Jabatan,
				Phone:          userData.Phone,
				ProfilePicture: userData.ProfilePicture,
			},
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, resp.GetUserDataResponseModel{
			Status:  http.StatusBadRequest,
			Message: "invalid token",
			Data:    nil,
		})
		return
	}

}

// func checkIDHelper(c *gin.Context, db *gorm.DB, ids string, out interface{}) error {
// 	//--------check id--------check id--------check id--------
// 	iduint64, err := strconv.ParseUint(ids, 10, 32)

// 	if err != nil {
// 		return err
// 	}
// 	iduint := uint(iduint64)

// 	checkID := db.Table("account_user_models").Where("id = ?", ids).Find(&account.AccountUserModel{
// 		ID: iduint,
// 	})

// 	if checkID.Error != nil {

// 		return checkID.Error
// 	}
// 	//--------check id--------check id--------check id--------
// 	return nil
// }

func EditProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	headertoken := c.Request.Header.Get("token")
	if headertoken == "" {
		c.JSON(http.StatusBadRequest, resp.EditProfileResponseModel{
			Status:  http.StatusBadRequest,
			Message: "token empty",
		})
		return
	}
	isValid, err := helper.VerifyToken(headertoken)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.EditProfileResponseModel{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if isValid {
		var editProfileRequest req.EditProfileRequestModel
		if err := c.ShouldBindJSON(&editProfileRequest); err != nil {
			c.JSON(http.StatusBadRequest, resp.EditProfileResponseModel{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		tokenRaw, err := helper.DecodeJWTToken(headertoken)
		// fmt.Printf("\ntoken raw %v", tokenRaw)
		if err != nil {
			c.JSON(http.StatusBadRequest, resp.EditProfileResponseModel{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		ids := tokenRaw["uid"].(string)

		//--------check id--------check id--------check id--------

		iduint64, err := strconv.ParseUint(ids, 10, 32)

		if err != nil {
			c.JSON(http.StatusInternalServerError, resp.EditProfileResponseModel{
				Status:  500,
				Message: "error parsing",
			})
			return
		}
		iduint := uint(iduint64)

		checkID := db.Table("account_user_models").Where("id = ?", ids).Find(&account.AccountUserModel{
			ID: iduint,
		})

		if checkID.Error != nil {
			c.JSON(http.StatusBadRequest, resp.EditProfileResponseModel{
				Status:  http.StatusBadRequest,
				Message: checkID.Error.Error(),
			})
			return
		}

		//--------check id--------check id--------check id--------

		if db.Error != nil {
			c.JSON(http.StatusBadRequest, resp.EditProfileResponseModel{
				Status:  http.StatusBadRequest,
				Message: db.Error.Error(),
			})
			return
		}
		// result := db.Create(&vehicleDataOutput)
		result := db.Table("account_user_models").Where("id = ?", ids).Update(&editProfileRequest)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, resp.EditProfileResponseModel{
				Status:  http.StatusBadRequest,
				Message: result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, resp.EditProfileResponseModel{
			Status:  http.StatusAccepted,
			Message: "Edit profile success",
		})
		return

	} else {
		c.JSON(http.StatusBadRequest, resp.EditProfileResponseModel{
			Status:  http.StatusBadRequest,
			Message: "invalid token",
		})
		return
	}
}
