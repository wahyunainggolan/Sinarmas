package app

import (
	"database/sql"
	"fmt"

	"sinarmas/controller"
	"sinarmas/db"

	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConnection() {
	db := db.Connectdb()
	a.DB = db
	fmt.Printf("success connect")
}

func (a *App) Routes() {
	r := gin.Default()
	userController := controller.NewUserController(a.DB)

	r.POST("/request-otp", userController.RequestOtp)
	r.POST("/validate", userController.CheckOtp)

	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":8080")
}
