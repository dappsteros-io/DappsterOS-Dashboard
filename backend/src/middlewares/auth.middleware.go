package middleware

import (
	"strings"

	repo "github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/repos"
	"github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/utils"
	"github.com/kataras/iris/v12"
)

func AuthMiddleware(ctx iris.Context) {

	access_token := ""
	cookie := ctx.GetCookie("access_token")

	authorizationHeader := ctx.Request().Header.Get("Authorization")
	if authorizationHeader != "" {
		fields := strings.Fields(authorizationHeader)
		if len(fields) > 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		}
	} else {
		access_token = cookie
	}

	if access_token == "" {
		ctx.Problem(iris.NewProblem().
			Key("success", false).
			Key("message", "The session is expired or invalid.").
			Status(iris.StatusUnauthorized))
		return
	}

	sub, err := utils.NewJWT().ValidateToken(access_token)

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Key("success", false).
			Key("message", "The session is expired or invalid.").
			Detail(err.Error()).
			Status(iris.StatusForbidden))
		return
	}
	data := sub.(map[string]any)

	user, err := repo.FindUserByID(data["id"].(string))

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Key("success", false).
			Key("message", "The session is expired or invalid.").
			Detail(err.Error()).
			Status(iris.StatusBadRequest))
		return
	}

	// if user.Token != access_token {
	// 	ctx.Problem(iris.NewProblem().
	// 		Key("success", false).
	// 		Detail("The user belonging to this token no logger exists").
	// 		Status(iris.StatusBadRequest))
	// 	return
	// }
	ctx.Values().Set("user", user)
	ctx.SetUser(user)
	ctx.Next()
}
