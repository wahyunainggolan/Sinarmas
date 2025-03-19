package controller

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	RequestOtp(*gin.Context)
	CheckOtp(*gin.Context)
}
