package routes

import (
	"net/http"
	"project_office_monitoring_backend/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type HandleRoutesResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/account/signin", controllers.SignInAccount)
	r.POST("/account/signup", controllers.SignUpAccount)
	r.POST("/account/editprofile", controllers.EditProfile)
	r.POST("/platform/create", controllers.CreatePlatform)
	r.POST("/platform/initialize", controllers.InitializePlatform)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, HandleRoutesResponse{
			Status:  http.StatusNotFound,
			Message: "Page not found",
		})

	})

	return r
}
