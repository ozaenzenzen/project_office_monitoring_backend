package controllers

import (
	"net/http"
	req "project_office_monitoring_backend/data/request"
	resp "project_office_monitoring_backend/data/response"
	jwthelper "project_office_monitoring_backend/helper"
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
	accountResponsePayload := account.AccountUserModel{
		Name:            accountInput.Name,
		Email:           accountInput.Email,
		Phone:           accountInput.Phone,
		Password:        accountInput.Password,
		ConfirmPassword: accountInput.ConfirmPassword,
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
			Status:  400,
			Message: "Record found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, resp.AccountSignUpResponseModel{
		Status:  http.StatusCreated,
		Message: "Account created successfully",
		Data: &resp.AccountUserDataSignUpModel{
			UserId: accountResponsePayload.ID,
			Name:   accountInput.Name,
			Email:  accountInput.Email,
			Phone:  accountInput.Phone,
		},
	})
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
	db := c.MustGet("db").(*gorm.DB)

	result := db.Where("email = ?", dataUser.Email).Where("password = ?", dataUser.Password).First(&table)
	// log.Println(fmt.Sprintf("log signin Value: %s", result.Value))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, resp.AccountSignInResponseModel{
			Status:  http.StatusNotFound,
			Message: "Account not match",
			Data:    nil,
		})
		return
	}

	tokenString, err := jwthelper.GenerateJWTToken(strconv.FormatUint(uint64(table.ID), 10), dataUser.Email)

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
		// Typeuser: &dataUser.Typeuser,
		Data: &resp.AccountUserDataSignInModel{
			ID:    table.ID,
			Name:  table.Name,
			Email: dataUser.Email,
			Phone: table.Phone,
			Token: tokenString,
		},
	})
}

type GetUserDataModel struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
}

type AccountUserGetUserResponse struct {
	Status   int               `json:"status"`
	Message  string            `json:"message"`
	UserData *GetUserDataModel `json:"userdata"`
}

func GetUserData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	headertoken := c.Request.Header.Get("token")

	if headertoken == "" {
		c.JSON(http.StatusBadRequest, AccountUserGetUserResponse{
			Status:  400,
			Message: "token empty",
		})
		return
	}
	isValid, err := jwthelper.VerifyToken(headertoken)

	if isValid == true {
		if err != nil {
			c.JSON(http.StatusBadRequest, AccountUserGetUserResponse{
				Status:   http.StatusBadRequest,
				Message:  err.Error(),
				UserData: nil,
			})
			return
		}

		var userData account.AccountUserModel

		tokenRaw, err := jwthelper.DecodeJWTToken(headertoken)
		// fmt.Printf("\ntoken raw %v", tokenRaw)
		if err != nil {
			c.JSON(http.StatusBadRequest, AccountUserGetUserResponse{
				Status:   http.StatusBadRequest,
				Message:  err.Error(),
				UserData: nil,
			})
			return
		}

		emails := tokenRaw["email"].(string)

		// if err := db.Where("id = ?", c.Param("id")).First(&userData).Error; err != nil {
		if err := db.Where("email = ?", emails).First(&userData).Error; err != nil {
			c.JSON(http.StatusBadRequest, AccountUserGetUserResponse{
				Status:   400,
				Message:  "User Data Not Found",
				UserData: nil,
			})
			return
		}

		c.JSON(http.StatusOK, AccountUserGetUserResponse{
			Status:  200,
			Message: "get user data success",
			UserData: &GetUserDataModel{
				ID:             userData.ID,
				Name:           userData.Name,
				Email:          userData.Email,
				Phone:          userData.Phone,
				ProfilePicture: userData.ProfilePicture,
			},
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, AccountUserGetUserResponse{
			Status:   http.StatusBadRequest,
			Message:  "invalid token",
			UserData: nil,
		})
		return
	}

}

type EditProfileRequest struct {
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
}

type EditProfileResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func checkIDHelper(c *gin.Context, db *gorm.DB, ids string, out interface{}) error {
	//--------check id--------check id--------check id--------
	iduint64, err := strconv.ParseUint(ids, 10, 32)

	if err != nil {
		return err
	}
	iduint := uint(iduint64)

	checkID := db.Table("account_user_models").Where("id = ?", ids).Find(&account.AccountUserModel{
		ID: iduint,
	})

	if checkID.Error != nil {

		return checkID.Error
	}
	//--------check id--------check id--------check id--------
	return nil
}

func EditProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	headertoken := c.Request.Header.Get("token")
	if headertoken == "" {
		c.JSON(http.StatusBadRequest, EditProfileResponse{
			Status:  400,
			Message: "token empty",
		})
		return
	}
	isValid, err := jwthelper.VerifyToken(headertoken)
	if err != nil {
		c.JSON(http.StatusBadRequest, EditProfileResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if isValid == true {
		var editProfileRequest EditProfileRequest
		if err := c.ShouldBindJSON(&editProfileRequest); err != nil {
			c.JSON(http.StatusBadRequest, EditProfileResponse{
				Status:  500,
				Message: err.Error(),
			})
			return
		}
		tokenRaw, err := jwthelper.DecodeJWTToken(headertoken)
		// fmt.Printf("\ntoken raw %v", tokenRaw)
		if err != nil {
			c.JSON(http.StatusBadRequest, EditProfileResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		ids := tokenRaw["uid"].(string)

		//--------check id--------check id--------check id--------

		iduint64, err := strconv.ParseUint(ids, 10, 32)

		if err != nil {
			c.JSON(http.StatusBadRequest, EditProfileResponse{
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
			c.JSON(http.StatusBadRequest, EditProfileResponse{
				Status:  400,
				Message: checkID.Error.Error(),
			})
			return
		}

		//--------check id--------check id--------check id--------

		if db.Error != nil {
			c.JSON(http.StatusBadRequest, EditProfileResponse{
				Status:  400,
				Message: db.Error.Error(),
			})
			return
		}
		// result := db.Create(&vehicleDataOutput)
		result := db.Table("account_user_models").Where("id = ?", ids).Update(&editProfileRequest)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, EditProfileResponse{
				Status:  400,
				Message: result.Error.Error(),
			})
			return
		}

		editProfileResponse := EditProfileResponse{
			Status:  http.StatusAccepted,
			Message: "Edit profile success",
		}

		c.JSON(http.StatusOK, editProfileResponse)

	} else {
		c.JSON(http.StatusBadRequest, EditProfileResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid token",
		})
		return
	}
}
