package routes

import (
	"github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/routes/v1"
	"github.com/kataras/iris/v12/core/router"
)

func Routes(api router.Party) {

	v1 := api.Party("/v1")
	routes.AuthRoutes(v1)
	routes.ProfileAPIRoutes(v1)
	routes.ProxmoxAPIRoutes(v1)
}
