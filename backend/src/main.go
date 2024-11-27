package main

import (
	"log"
	"os"

	"github.com/dappster-io/DappsterOS-Dashboard/backend/src/db"
	"github.com/dappster-io/DappsterOS-Dashboard/backend/src/routes"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {
	db := &db.Conn{}
	db.Init()

	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.Use(crs)
	api := app.Party("/api")
	api.Use(iris.Compression)

	routes.Routes(api)

	port := (map[bool]string{true: "8000", false: os.Getenv("PORT")})[os.Getenv("PORT") == ""]

	if port == "" {
		log.Fatal("PORT environment variable not specified")
	}

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
