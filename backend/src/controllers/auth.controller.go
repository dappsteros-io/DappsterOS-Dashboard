package controllers

import (
	"log"
	"time"

	"github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/models"
	repo "github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/repos"
	"github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/requests"
	"github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/responses"
	"github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/utils"
	"github.com/kamva/mgm/v3"
	"github.com/kataras/iris/v12"
)

func SignIn(ctx iris.Context) {
	var cred requests.Credential

	if err := ctx.ReadJSON(&cred); err != nil {
		ctx.Problem(iris.NewProblem().
			// Title("Bad Request").
			Key("success", false).
			Detail("Invalid Credential").
			Status(iris.StatusBadRequest),
		)
		return
	}

	user, err := repo.FindUserByEmail(cred.Email)

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "The current user is not existing in our database.").
			Status(iris.StatusNotFound))
		return
	}

	if err := user.CheckPassword(cred.Password); err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "The credential is mismatched.").
			Status(iris.StatusUnauthorized))
		return
	}
	jwtToken := utils.NewJWT()
	accessToken, err := jwtToken.CreateToken(time.Hour*24, iris.Map{"email": user.Email, "id": user.ID.Hex(), "username": user.Username})
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal server error was occurred.").
			Status(iris.StatusInternalServerError))
		return
	}
	user.Token = accessToken
	user.Put()
	user.Password = ""
	// user.Token = ""
	ctx.JSON(responses.CommonResponse{Success: true, Data: user, Title: "Welcome!", Detail: "You logged in successfully.", Message: "You logged in successfully."})
}

func SignUp(ctx iris.Context) {
	var info requests.UserRequest

	if err := ctx.ReadJSON(&info); err != nil {
		ctx.Problem(iris.NewProblem().
			// Title("Bad Request").
			Key("success", false).
			Detail("Invalid Credential").
			Status(iris.StatusBadRequest),
		)
		return
	}

	isExistingUser, _ := repo.FindUserByEmail(info.Email)

	if !isExistingUser.ID.IsZero() {
		ctx.Problem(iris.NewProblem().
			// Title("Bad Request").
			Detail("This email is already existing").
			Key("success", false).
			Status(iris.StatusConflict))
		return
	}

	user, err := (&models.User{Username: info.Username, Email: info.Email, Password: info.Password}).Register()

	if err != nil {
		ctx.Problem(iris.NewProblem().
			// Title("Bad Request").
			Detail(err.Error()).
			Key("success", false).
			Status(iris.StatusInternalServerError))
		return
	}
	user.Password = ""
	ctx.JSON(iris.Map{"success": true, "user": user})
}

func GetUsers(ctx iris.Context) {
	var users []responses.UserResponse
	mgm.Coll(&models.User{}).SimpleFind(&users, iris.Map{})
	log.Println(users)

	ctx.JSON(iris.Map{"success": true, "users": users})
}
