package main

import (
	setup "project_office_monitoring_backend/database"
	account "project_office_monitoring_backend/models/account"
	routes "project_office_monitoring_backend/routes"
)

func main() {
	db := setup.SetupDB()
	sqlDB := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	db.AutoMigrate(&account.AccountUserModel{})

	r := routes.SetupRoutes(db)
	err := r.Run(":8080")

	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
	return

}
