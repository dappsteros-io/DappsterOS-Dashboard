package controllers

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"encoding/json"

	"github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/models"
	"github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/requests"
	"github.com/dappsteros-io/DappsterOS-Dashboard/backend/src/responses"
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
		// proxmox.WithCredentials(&credentials),
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
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
	}
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

func DeleteVM(ctx iris.Context) {

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
	t, err := vm.Delete(context.Background())
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", err.Error()).
			Status(iris.StatusInternalServerError))
		return
	}
	ctx.JSON(responses.CommonResponse{Success: true, Data: t, Message: fmt.Sprintf("VM %d was deleted successfully.", vmId), Title: "VM Deleted!"})
}

func CreateVM(ctx iris.Context) {

	if client == nil {
		Init()
	}
	var vmdata requests.VMData

	if err := ctx.ReadJSON(&vmdata); err != nil {
		ctx.StatusCode(400)
		ctx.JSON(responses.CommonResponse{Success: false, Message: "Invalid request", Detail: err.Error(), Title: "Invalid request"})
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
	vms, err := node.VirtualMachines(context.Background())
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	lastVmId := 0
	for _, v := range vms {
		if v.VMID > proxmox.StringOrUint64(lastVmId) {
			lastVmId = int(v.VMID)
		}
	}

	sss, _ := node.StorageISO(context.Background())
	iso, _ := sss.ISO(context.Background(), "jammy-server-cloudimg-amd64-disk-kvm.img")
	println(iso.VolID)

	sss2, _ := node.Storage(context.Background(), "local-lvm")
	iso2, _ := sss2.ISO(context.Background(), "jammy-server-cloudimg-amd64-disk-kvm.img")
	println(iso2)
	// Define VM options
	options := []proxmox.VirtualMachineOption{
		{Name: "agent", Value: "1"},
		{Name: "autostart", Value: 1},
		{Name: "cpu", Value: "x86-64-v2-AES"},
		// {Name: "ide0", Value: "local:iso/seed.iso,media=cdrom"},
		// {Name: "ide2", Value: "local:iso/ubuntu-20.04.6-live-server-amd64.iso,media=cdrom"},
		{Name: "machine", Value: "q35"},
		{Name: "memory", Value: "16535"},
		{Name: "net0", Value: "virtio,bridge=vmbr0,firewall=1"},
		{Name: "numa", Value: 0},
		{Name: "onboot", Value: 1},
		{Name: "ostype", Value: "l26"},
		{Name: "cores", Value: 2},
		{Name: "scsihw", Value: "virtio-scsi-pci"},
		{Name: "scsi0", Value: "local:32,format=qcow2,iothread=on"},
		// {Name: "scsi0", Value: "local:32,import-from=/root/jammy-server-cloudimg-amd64-disk-kvm.img"},
		{Name: "start", Value: 1},
		{Name: "sockets", Value: 2},
		{Name: "boot", Value: "order=scsi0;ide0;ide2;net0"},
	}

	t, err := node.NewVirtualMachine(context.Background(), lastVmId+1, options...)

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error").
			Status(iris.StatusInternalServerError))
		return
	}

	// vm, err := node.VirtualMachine(context.Background(), lastVmId+1)
	// if err != nil {
	// }
	// vm.Start(context.Background())
	// storage, _ :=node.Storage(context.Background(), "local-")
	// storage.ISO()
	ctx.JSON(responses.CommonResponse{Success: true, Data: t, Message: fmt.Sprintf("VM %d was created successfully.", lastVmId+1), Title: "VM Created!"})

}

func CloneVm(ctx iris.Context) {
	user := ctx.Value("user")
	print(user)
	u, ok := user.(*models.User)
	if !ok {
		// Handle the case where the type assertion fails
		fmt.Println("User is not of type *model.User")
		ctx.Problem(iris.NewProblem().
			Key("success", false).
			Key("message", "User is not of type *model.User").
			Status(iris.StatusInternalServerError))
		return
	}
	if client == nil {
		Init()
	}

	node, err := client.Node(context.Background(), os.Getenv("PROXMOX_NODE"))

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error1").
			Status(iris.StatusInternalServerError))
		return
	}

	vms, err := node.VirtualMachines(context.Background())
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	lastVmId := 0
	for _, v := range vms {
		if v.VMID > proxmox.StringOrUint64(lastVmId) {
			lastVmId = int(v.VMID)
		}
	}

	vm, err := node.VirtualMachine(context.Background(), 100)

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error2").
			Status(iris.StatusInternalServerError))
		return
	}
	params := proxmox.VirtualMachineCloneOptions{
		NewID: lastVmId + 1,
		Name:  fmt.Sprintf("%s-VM%d", u.Username, lastVmId+1),
	}

	newId, t, err := vm.Clone(context.Background(), &params)
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error2").
			Status(iris.StatusInternalServerError))
		return
	}
	t.Wait(context.Background(), time.Second*1, time.Second*20)

	newVM, err := node.VirtualMachine(context.Background(), newId)

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error - Creation Failed").
			Status(iris.StatusInternalServerError))
		return
	}
	t, err = newVM.Start(context.Background())
	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error - VM not started").
			Status(iris.StatusInternalServerError))
		return
	}
	t.Wait(context.Background(), time.Second*1, time.Second*5)

	ctx.JSON(responses.CommonResponse{Success: true, Data: newId, Message: "Please a moment. Your Vm is creating now...", Title: "Initalizing...", Detail: fmt.Sprintf("Your vm is creating %d...", newId)})

}

func InstallDappster(ctx iris.Context) {

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
			Key("message", "Internal Server Error1").
			Status(iris.StatusInternalServerError))
		return
	}

	vm, err := node.VirtualMachine(context.Background(), vmId)

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", "Internal Server Error2").
			Status(iris.StatusInternalServerError))
		return
	}

	// pid, err := vm.AgentExec(context.Background(), []string{"bash"}, "curl 10.69.12.99:8080 | sudo bash")
	pid2, err := vm.AgentExec(context.Background(), []string{"bash"}, "curl -fsSL http://get.dappster.io:8080/ | sudo bash")

	s2, err := vm.WaitForAgentExecExit(context.Background(), pid2, 3600)
	fmt.Println(s2)
	fmt.Println(pid2)
	fmt.Println(err)
	if err != nil {

		pid, err := vm.AgentExec(context.Background(), []string{"bash"}, "sudo mv /etc/apt/sources.list /etc/apt/sources.list.backup && sudo touch /etc/apt/sources.list")
		s, err := vm.WaitForAgentExecExit(context.Background(), pid, 1800)
		fmt.Println(s)

		pid, err = vm.AgentExec(context.Background(), []string{"bash"}, "sudo rm -f /var/lib/dpkg/lock-frontend")
		s, err = vm.WaitForAgentExecExit(context.Background(), pid, 1800)
		fmt.Println(s)

		pid, err = vm.AgentExec(context.Background(), []string{"bash"}, "sudo rm -f /var/lib/dpkg/lock")
		s, err = vm.WaitForAgentExecExit(context.Background(), pid, 1800)

		pid, err = vm.AgentExec(context.Background(), []string{"bash"}, "sudo dpkg --configure -a")
		s, err = vm.WaitForAgentExecExit(context.Background(), pid, 1800)
		fmt.Println(s)

		fmt.Println(s)
		pid, err = vm.AgentExec(context.Background(), []string{"bash"}, "sudo apt-get update")
		s, err = vm.WaitForAgentExecExit(context.Background(), pid, 1800)
		fmt.Println(s)
		fmt.Println(err)
		// ctx.JSON(responses.CommonResponse{Success: true, Data: s, Message: "Success", Title: "Installing DappsterOS", Detail: fmt.Sprintf("Installing DappsterOS on %s...", vm.Name)})
		// if err != nil {

		// 	ctx.Problem(iris.NewProblem().
		// 		Detail(err.Error()).
		// 		Key("success", false).
		// 		Key("message", "Internal Server Error2").
		// 		Status(iris.StatusInternalServerError))
		// 	return
		// }
		// pid, err := vm.AgentExec(context.Background(), []string{"bash"}, "curl 10.69.12.99:8080 | sudo bash")
		pid2, err := vm.AgentExec(context.Background(), []string{"bash"}, "sudo rm -rf /tmp/dappster-installer/")

		pid2, err = vm.AgentExec(context.Background(), []string{"bash"}, "wget -qO- http://get.dappster.io:8080/ | sudo bash")

		// s2, err := vm.WaitForAgentExecExit(context.Background(), pid2, 1800)
		// fmt.Println(s2)
		// fmt.Println(pid2)
		ctx.JSON(responses.CommonResponse{Success: true, Data: pid2, Message: "Installing DappsterOS... Please take a few mins. You can check the status of running of dappsteros.", Title: "Installing...", Detail: fmt.Sprintf("Installing DappsterOS on %s...", vm.Name)})
		return

	}
	ctx.JSON(responses.CommonResponse{Success: true, Data: pid2, Message: "Installing DappsterOS... Please take a few mins. You can check the status of running of dappsteros.", Title: "Installing...", Detail: fmt.Sprintf("Installing DappsterOS on %s...", vm.Name)})
}

func GetDappsterStatus(ctx iris.Context) {
	if client == nil {
		Init()
	}

	vmId, err := strconv.Atoi(ctx.Params().Get("vmid"))
	// pid, err := strconv.Atoi(ctx.URLParam("pid"))
	// fmt.Println(vmId, pid)
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

	pid, err := vm.AgentExec(context.Background(), []string{"curl", "-s", "http://localhost/v1/users/status"}, "")
	s, err := vm.AgentExecStatus(context.Background(), pid)
	fmt.Println(s)
	fmt.Println(pid)

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", fmt.Sprintf("Cannot find the DappsterOS on %s", vm.Name)).
			Status(iris.StatusInternalServerError))
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(s.OutData), &result)

	if err != nil {
		ctx.Problem(iris.NewProblem().
			Detail(err.Error()).
			Key("success", false).
			Key("message", fmt.Sprintf("Cannot find the DappsterOS on %s", vm.Name)).
			Status(iris.StatusInternalServerError))
		return
	}
	ctx.JSON(responses.CommonResponse{Success: true, Data: result, Message: fmt.Sprintf("DappsterOS is running on %s", vm.Name), Title: "DappsterOS Running", Detail: fmt.Sprintf("DappsterOS is running on %s", vm.Name)})

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
