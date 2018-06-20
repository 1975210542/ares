package logic

import (
	"ares/models"

	"github.com/gin-gonic/gin"
)

func Login() gin.H {

	user := models.User{}
	user.UserName = "gaoqiankun"
	return gin.H{"status": user}
}
