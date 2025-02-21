package routes

import (
	"github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/controllers"
	middleware "github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/middlewares"
	"github.com/kataras/iris/v12"
)

func ProxmoxAPIRoutes(v1 iris.Party) {

	proxmoxAPI := v1.Party("/proxmox")
	{
		proxmoxAPI.Use(middleware.AuthMiddleware)
		proxmoxAPI.Get("/", controllers.Version)
		proxmoxAPI.Get("/nodes", controllers.GetNodes)
		proxmoxAPI.Get("/vms", controllers.GetVms)
		// proxmoxAPI.Post("/vms", controllers.CreateVM)
		proxmoxAPI.Post("/vms", controllers.CloneVm)
		proxmoxAPI.Get("/vms/{vmid}", controllers.GetVM)
		proxmoxAPI.Delete("/vms/{vmid}", controllers.DeleteVM)
		proxmoxAPI.Post("/vms/{vmid}/start", controllers.StartVM)
		proxmoxAPI.Post("/vms/{vmid}/stop", controllers.StopVM)
		proxmoxAPI.Post("/vms/{vmid}/install", controllers.InstallDappster)
		proxmoxAPI.Post("/vms/{vmid}/status", controllers.GetDappsterStatus)
	}
}
