package middleware

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	// ctx.Header("Access-Control-Allow-Credentials", "true")
	fmt.Print(ctx.Method())
	if ctx.Method() == iris.MethodOptions {
		ctx.Header("Access-Control-Methods",
			"POST, PUT, PATCH, DELETE")

		ctx.Header("Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin,Content-Type")

		ctx.Header("Access-Control-Max-Age",
			"86400")

		ctx.StatusCode(iris.StatusNoContent)
		return
	}

	ctx.Next()
}
