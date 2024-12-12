package requests

type VMData struct {
	VMID    int64    `json:"vmid"`
	Tags    []string `json:"tags"`
	Sockets int16    `json:"sockets"`
	SCSI0   string   `json:"scsi0"`
	OSType  string   `json:"ostype"`
	OnBoot  int16    `json:"onboot"`
	Numa    int16    `json:"numa"`
	Memory  int64    `json:"memory"`
	CPU     string   `json:"cpu"`
	Cores   int16    `json:"cores"`
	Net0    string   `json:"net0"`
	IDE2    string   `json:"ide2"`
	Agent   int16    `json:"agent"`
}
