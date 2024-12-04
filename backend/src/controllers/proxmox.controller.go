package controllers

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dappster-io/DappsterOS-Dashboard/backend/src/responses"
	"github.com/kataras/iris/v12"
	"github.com/luthermonson/go-proxmox"
)

var client *proxmox.Client

type ProxmoxController struct {
	/* dependencies */
}

func Init() {
	insecureHTTPClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	// credentials := proxmox.Credentials{
	// 	Username: os.Getenv("PROXMOX_USERNAME"),
	// 	Password: os.Getenv("PROXMOX_PASSWORD"),
	// }
	tokenID := os.Getenv("PROXMOX_TOKEN")
	secret := os.Getenv("PROXMOX_SECRET")
	proxmoxURL := os.Getenv("PROXMOX_URL")
	client = proxmox.NewClient(proxmoxURL,
		proxmox.WithHTTPClient(&insecureHTTPClient),
		proxmox.WithAPIToken(tokenID, secret),
		// proxmox.WithCredentials(&credentials)
	)
}
func Version(ctx iris.Context) {

	if client == nil {
		Init()
	}

	version, err := client.Version(context.Background())
	if err != nil {
		// panic(err)
		fmt.Println(err.Error())
		return
	}
	ctx.JSON(iris.Map{"version": version})
}

func GetNodes(ctx iris.Context) {

	if client == nil {
		Init()
	}

	nodes, err := client.Nodes(context.Background())
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
		return
	}
	fmt.Println(nodes)
	ctx.JSON(responses.CommonResponse{Status: iris.StatusOK, Success: true, Data: nodes})
}

func GetVms(ctx iris.Context) {

	if client == nil {
		Init()
	}

	node, err := client.Node(context.Background(), os.Getenv("PROXMOX_NODE"))
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
		return
	}

	// vmId, err := strconv.Atoi(os.Getenv("PROXMOX_VM"))
	// if err != nil {
	// 	ctx.Problem(iris.NewProblem().
	// 		Detail(err.Error()).
	// 		Key("success", false).
	// 		Key("message", "Internal Server Error").
	// 		Status(iris.StatusInternalServerError))
	// }

	// vm, err := node.VirtualMachine(context.Background(), vmId)
	// if err != nil {
	// 	ctx.Problem(iris.NewProblem().
	// 		Detail(err.Error()).
	// 		Key("success", false).
	// 		Key("message", "Internal Server Error").
	// 		Status(iris.StatusInternalServerError))
	// }
	cons, err := node.VirtualMachines(context.Background())
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
		return
	}
	ctx.JSON(responses.CommonResponse{Success: true, Data: cons})
}
func GetVM(ctx iris.Context) {

	if client == nil {
		Init()
	}

	node, err := client.Node(context.Background(), os.Getenv("PROXMOX_NODE"))
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
		return
	}

	vmId, err := strconv.Atoi(ctx.Params().Get("vmid"))
	vm, err := node.VirtualMachine(context.Background(), vmId)
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
	}
	ctx.JSON(responses.CommonResponse{Success: true, Data: vm})
}

func StartVM(ctx iris.Context) {

	if client == nil {
		Init()
	}

	vmId, err := strconv.Atoi(ctx.Params().Get("vmid"))
	fmt.Println(vmId)
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Invalid VM Id").
			Status(iris.StatusBadRequest))
		return
	}

	node, err := client.Node(context.Background(), os.Getenv("PROXMOX_NODE"))

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
		return
	}

	vm, err := node.VirtualMachine(context.Background(), vmId)
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
		return
	}
	t, err := vm.Start(context.Background())
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
		return
	}

	ctx.JSON(responses.CommonResponse{Success: true, Data: t, Title: "Success", Message: fmt.Sprintf("VM %d started successfully", vmId)})
}

func StopVM(ctx iris.Context) {

	if client == nil {
		Init()
	}

	vmId, err := strconv.Atoi(ctx.Params().Get("vmid"))

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Invalid VM Id").
			Status(iris.StatusBadRequest))
		return
	}

	node, err := client.Node(context.Background(), os.Getenv("PROXMOX_NODE"))

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
		return
	}

	vm, err := node.VirtualMachine(context.Background(), vmId)
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
		return
	}
	t, err := vm.Stop(context.Background())
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
		return
	}

	ctx.JSON(responses.CommonResponse{Success: true, Data: t, Title: "Success", Message: fmt.Sprintf("VM %d stopped successfully", vmId)})
}
