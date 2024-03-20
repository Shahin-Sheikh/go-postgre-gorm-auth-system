package routes

import (
	"github.com/Shahin-Sheikh/go-postgre-gorm-auth-system/controllers"
	"github.com/Shahin-Sheikh/go-postgre-gorm-auth-system/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
}
