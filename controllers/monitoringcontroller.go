package controllers

import (
	"net/http"
	req "project_office_monitoring_backend/data/monitor/request"
	resp "project_office_monitoring_backend/data/monitor/response"
	helper "project_office_monitoring_backend/helper"
	monitor "project_office_monitoring_backend/models/monitor"
	platform "project_office_monitoring_backend/models/platform"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CaptureLocation(c *gin.Context) {
	header_platformkey := c.Request.Header.Get("platformkey")
	if header_platformkey == "" {
		c.JSON(http.StatusBadRequest, resp.CaptureLocationResponseModel{
			Status:  http.StatusBadRequest,
			Message: "invalid credential",
			Data:    nil,
		})
		return
	}

	isValidPlatformKey, err := helper.VerifyPlatformToken(header_platformkey)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.CaptureLocationResponseModel{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if isValidPlatformKey {
		headertoken := c.Request.Header.Get("token")
		if headertoken == "" {
			c.JSON(http.StatusBadRequest, resp.CaptureLocationResponseModel{
				Status:  http.StatusBadRequest,
				Message: "invalid credential",
				Data:    nil,
			})
			return
		}

		isValid, err := helper.VerifyToken(headertoken)
		if isValid {
			if err != nil {
				c.JSON(http.StatusBadRequest, resp.CaptureLocationResponseModel{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Data:    nil,
				})
				return
			}

			db := c.MustGet("db").(*gorm.DB)
			if db.Error != nil {
				c.JSON(http.StatusInternalServerError, resp.CaptureLocationResponseModel{
					Status:  http.StatusInternalServerError,
					Message: db.Error.Error(),
					Data:    nil,
				})
				return
			}

			platformTokenRaw, err := helper.DecodeJWTToken(header_platformkey)
			// fmt.Printf("\ntoken raw %v", tokenRaw)
			if err != nil {
				c.JSON(http.StatusBadRequest, resp.CaptureLocationResponseModel{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Data:    nil,
				})
				return
			}
			platform_secret := platformTokenRaw["platform_secret"].(string)

			//--------check id--------check id--------check id--------

			var dataPlatform platform.PlatformModel

			if err != nil {
				c.JSON(http.StatusInternalServerError, resp.CaptureLocationResponseModel{
					Status:  500,
					Message: "error parsing",
				})
				return
			}

			if err := db.Table("platform_models").Where("platform_secret = ?", platform_secret).First(&dataPlatform).Error; err != nil {
				c.JSON(http.StatusBadRequest, resp.CaptureLocationResponseModel{
					Status:  http.StatusBadRequest,
					Message: "User Data Not Found",
					Data:    nil,
				})
				return
			}

			//--------check id--------check id--------check id--------

			tokenRaw, err := helper.DecodeJWTToken(headertoken)
			// fmt.Printf("\ntoken raw %v", tokenRaw)
			if err != nil {
				c.JSON(http.StatusBadRequest, resp.CaptureLocationResponseModel{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Data:    nil,
				})
				return
			}

			userStamp := tokenRaw["userStamp"].(string)

			var reqData req.CaptureLocationRequestModel

			payloadDataInput := monitor.MonitorModel{
				UserStamp:           userStamp,
				CompanyName:         dataPlatform.CompanyName,
				CompanyPlatformName: dataPlatform.PlatformName,
				Picture:             reqData.Picture,
				Location:            reqData.Location,
				Address:             reqData.Address,
				Information:         reqData.Information,
				Latitude:            reqData.Latitude,
				Longitude:           reqData.Longitude,
			}

			result := db.Create(&payloadDataInput)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, resp.CaptureLocationResponseModel{
					Status:  http.StatusBadRequest,
					Message: result.Error.Error(),
					Data:    nil,
				})
				return
			}

			c.JSON(http.StatusOK, resp.CaptureLocationResponseModel{
				Status:  200,
				Message: "get user data success",
				Data: &resp.CaptureLocationDataModel{
					Location:    payloadDataInput.Location,
					Address:     payloadDataInput.Address,
					Information: payloadDataInput.Information,
				},
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, resp.CaptureLocationResponseModel{
				Status:  http.StatusBadRequest,
				Message: "invalid credential",
				Data:    nil,
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, resp.CaptureLocationResponseModel{
			Status:  http.StatusBadRequest,
			Message: "invalid credential",
			Data:    nil,
		})
		return
	}
}

func GetListLocation(c *gin.Context) {
	//
}

func GetListLog(c *gin.Context) {
	//
}
