package routes

import (
	"github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/controllers"
	middleware "github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/middlewares"
	"github.com/kataras/iris/v12"
)

func ProfileAPIRoutes(v1 iris.Party) {

	profileAPI := v1.Party("/profile")
	{
		profileAPI.Use(middleware.AuthMiddleware)
		profileAPI.Get("/", controllers.GetMe)
	}
}
