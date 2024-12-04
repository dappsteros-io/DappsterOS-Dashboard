package routes

import (
	"github.com/dappster-io/DappsterOS-Dashboard/backend/src/controllers"
	middleware "github.com/dappster-io/DappsterOS-Dashboard/backend/src/middlewares"
	"github.com/kataras/iris/v12"
)

func ProxmoxAPIRoutes(v1 iris.Party) {

	proxmoxAPI := v1.Party("/proxmox")
	{
		proxmoxAPI.Use(middleware.AuthMiddleware)
		proxmoxAPI.Get("/", controllers.Version)
		proxmoxAPI.Get("/nodes", controllers.GetNodes)
		proxmoxAPI.Get("/vms", controllers.GetVms)
		proxmoxAPI.Get("/vm/{vmid}", controllers.GetVM)
		proxmoxAPI.Post("/vm/{vmid}/start", controllers.StartVM)
		proxmoxAPI.Post("/vm/{vmid}/stop", controllers.StopVM)
	}
}
