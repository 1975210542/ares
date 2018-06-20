package controllers

import (
	"net/http"

	"fmt"

	"ares/logic"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println("gaoqiankun")
	c.JSON(http.StatusOK, logic.Login())
}
