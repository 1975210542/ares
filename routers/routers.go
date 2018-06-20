package routers

import (
	//	"ares/common"
	"ares/controllers"
	//	"ares/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	//	common.InitGlbConf()
}

func InitRouter() {
	router := gin.Default()

	initRouters(router)

	println("server listening port on", "8080", "...")
	http.ListenAndServe(":8080", router)
}

func initRouters(router *gin.Engine) {
	homeMain := Home(router)

	homeMain.POST("/login", controllers.Login)

}

func Home(router *gin.Engine) *gin.RouterGroup {
	homeMain := router.Group("/user")
	return homeMain
}
