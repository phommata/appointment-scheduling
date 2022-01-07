package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/phommata/appointment-scheduling/config"
	"github.com/phommata/appointment-scheduling/controllers"
)

func main(){
	r := gin.Default()

	appConf := config.AppConfig()
	server := fmt.Sprintf(":%d", appConf.Server.Port)

	// Setup service repository
	appointmentRepo := controllers.New(appConf)

	// Routes
	r.GET("/appointments", appointmentRepo.GetAppointments)
	r.GET("/available-appointments", appointmentRepo.ListAvailableAppointments)
	r.POST("/appointment", appointmentRepo.CreateAppointment)

	// Run the server
	r.Run(server)
}