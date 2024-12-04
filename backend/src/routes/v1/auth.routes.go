package routes

import (
	"github.com/dappster-io/DappsterOS-Dashboard/backend/src/controllers"
	middleware "github.com/dappster-io/DappsterOS-Dashboard/backend/src/middlewares"
	"github.com/kataras/iris/v12"
)

func AuthRoutes(v1 iris.Party) {

	authAPI := v1.Party("/auth")
	{
		authAPI.Post("/signin", controllers.SignIn)
		authAPI.Post("/signup", controllers.SignUp)

		authAPI.Use(middleware.AuthMiddleware)
		authAPI.Get("/me", controllers.GetMe)
	}
}
