package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"sinarmas/model"
	"sinarmas/repository"
)

type UserController struct {
	Db *sql.DB
}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &UserController{Db: db}
}

func (m *UserController) RequestOtp(c *gin.Context) {
	var post model.RequestUser
	DB := m.Db

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": err})

		return
	}

	repository := repository.NewUserRepository(DB)
	insert := repository.RequestOtp(post)

	if (insert != model.ResponseUser{}) {
		c.JSON(200, gin.H{"user_id": insert.UserId, "otp": insert.Otp})

		return
	} else {
		c.JSON(200, gin.H{"status": "success", "message": "user not found"})

		return
	}

}

func (m *UserController) CheckOtp(c *gin.Context) {
	var post model.RequestUser
	DB := m.Db

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": err})

		return
	}

	repository := repository.NewUserRepository(DB)
	validateOtp := repository.CheckOtp(post)

	if (validateOtp != model.ResponseUser{}) {
		c.JSON(200, gin.H{"user_id": validateOtp.UserId, "message": "Otp Validated Successfully"})

		return
	} else {
		c.JSON(200, gin.H{"status": "success", "message": "user not found"})

		return
	}

}
