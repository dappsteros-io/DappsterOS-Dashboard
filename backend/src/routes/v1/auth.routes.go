package routes

import (
	"github.com/dappster-io/DappsterOS-Dashboard/backend/src/controllers"
	"github.com/kataras/iris/v12"
)

func AuthRoutes(v1 iris.Party) {

	authAPI := v1.Party("/auth")
	{
		authAPI.Post("/signin", controllers.SignIn)
		authAPI.Post("/signup", controllers.SignUp)
	}
}
