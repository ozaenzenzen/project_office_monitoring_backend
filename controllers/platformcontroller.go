package controllers

import (
	"net/http"
	req "project_office_monitoring_backend/data/platform/request"
	resp "project_office_monitoring_backend/data/platform/response"
	helper "project_office_monitoring_backend/helper"
	platform "project_office_monitoring_backend/models/platform"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func CreatePlatform(c *gin.Context) {
	var reqCreatePlatform req.CreatePlatformRequestModel
	if err := c.ShouldBindJSON(&reqCreatePlatform); err != nil {
		c.JSON(http.StatusBadRequest, resp.CreatePlatformResponseModel{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	validate := validator.New()

	if err := validate.Struct(reqCreatePlatform); err != nil {
		c.JSON(http.StatusBadRequest, resp.CreatePlatformResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Data tidak lengkap",
			Data:    nil,
		})
		return
	}

	stampToken := uuid.New().String()
	stampTokenPN := uuid.New().String()
	stampTokenPS := uuid.New().String()

	platformName := reqCreatePlatform.CompanyName + stampTokenPN
	platformSecret := reqCreatePlatform.CompanyName + reqCreatePlatform.CompanyEmail + stampTokenPS

	hashPw := strconv.FormatUint(uint64(helper.Hash(reqCreatePlatform.Password)), 10)
	hashCpw := strconv.FormatUint(uint64(helper.Hash(reqCreatePlatform.ConfirmPassword)), 10)

	dataHandler := platform.PlatformModel{
		CompanyName:     reqCreatePlatform.CompanyName,
		CompanyEmail:    reqCreatePlatform.CompanyEmail,
		CompanyNoReg:    stampToken,
		CompanyPhone:    reqCreatePlatform.CompanyPhone,
		Password:        hashPw,
		ConfirmPassword: hashCpw,
		PlatformName:    platformName,
		PlatformSecret:  platformSecret,
	}

	db := c.MustGet("db").(*gorm.DB)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, resp.CreatePlatformResponseModel{
			Status:  http.StatusInternalServerError,
			Message: db.Error.Error(),
			Data:    nil,
		})
		return
	}

	result := db.FirstOrCreate(&dataHandler, platform.PlatformModel{CompanyNoReg: stampToken})

	if result.Value == nil && result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, resp.CreatePlatformResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Record found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, resp.CreatePlatformResponseModel{
		Status:  http.StatusCreated,
		Message: "Platform created successfully",
		Data: &resp.CreatePlatformDataModel{
			CompanyName:    dataHandler.CompanyName,
			CompanyEmail:   dataHandler.CompanyEmail,
			CompanyNoReg:   dataHandler.CompanyNoReg,
			CompanyPhone:   dataHandler.CompanyPhone,
			PlatformName:   dataHandler.PlatformName,
			PlatformSecret: dataHandler.PlatformSecret,
		},
	})
}

func InitializePlatform(c *gin.Context) {
	var reqInitPlatform req.InitializePlatformRequestModel

	if err := c.ShouldBindJSON(&reqInitPlatform); err != nil {
		c.JSON(http.StatusBadRequest, resp.InitializePlatformResponseModel{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var table platform.PlatformModel
	db := c.MustGet("db").(*gorm.DB)

	result := db.Where("platform_name = ?", reqInitPlatform.PlatformName).Where("platform_secret = ?", reqInitPlatform.PlatformSecret).First(&table)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, resp.InitializePlatformResponseModel{
			Status:  http.StatusNotFound,
			Message: "Account not match",
			Data:    nil,
		})
		return
	}

	// platformKey, err := helper.GeneratePlatformToken(strconv.FormatUint(uint64(table.ID), 10), reqInitPlatform.PlatformSecret, time.Now().Add(time.Hour*24*365).Unix())
	platformKey, err := helper.GeneratePlatformToken(strconv.FormatUint(uint64(table.ID), 10), reqInitPlatform.PlatformSecret, time.Now().Add(time.Minute*30).Unix())

	if err != nil {
		c.JSON(http.StatusNotFound, resp.InitializePlatformResponseModel{
			Status:  http.StatusInternalServerError,
			Message: "Failed to generate key",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, resp.InitializePlatformResponseModel{
		Status:  200,
		Message: "Initialization Success",
		Data: &resp.InitializePlatformDataModel{
			PlatformName:   reqInitPlatform.PlatformName,
			PlatformSecret: reqInitPlatform.PlatformSecret,
			PlatformKey:    platformKey,
		},
	})
}
