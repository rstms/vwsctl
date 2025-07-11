package vmx

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"runtime"
	"strings"
)

type VM struct {
	Name string
	Id   string

	Running   bool
	IpAddress string

	CpuCount   int
	RamSizeMb  int
	DiskSizeMb int
	MacAddress string

	IsoAttached bool
	IsoPath     string

	SerialAttached bool
	SerialPath     string

	VncEnabled bool
	VncPort    int
	VncAddress string

	VmxPath string
}

type CreateOptions struct {
	CpuCount          int
	MemorySize        int64
	DiskType          VDiskType
	DiskSize          int64
	EFIBoot           bool
	HostTimeSync      bool
	GuestTimeZone     string
	EnableDragAndDrop bool
	EnableClipboard   bool
	MacAddress        string
	IsoSource         string
	soAttached        bool
	Start             *StartOptions
}

type DestroyOptions struct {
	Force bool
	Wait  bool
}

type StartOptions struct {
	GUI        bool
	FullScreen bool
	Wait       bool
}

type StopOptions struct {
	PowerOff bool
	Wait     bool
}

type Controller interface {
	List(string, bool, bool) ([]VM, error)
	Create(string, CreateOptions) (VM, error)
	Destroy(string, DestroyOptions) error
	Start(string, StartOptions) error
	Stop(string, StopOptions) error
	Get(string, string) (any, error)
	Set(string, string, any) error
	Close() error
}

type vmctl struct {
	Hostname string
	Username string
	KeyFile  string
	Path     string
	api      *APIClient
	relay    *Relay
	Shell    string
	Local    string
	Remote   string
}

func isLocal() (bool, error) {
	remote := viper.GetString("hostname")
	if remote == "" || remote == "localhost" || remote == "127.0.0.1" {
		return true, nil
	}
	host, err := os.Hostname()
	if err != nil {
		return false, err
	}
	if host == remote {
		return true, nil
	}
	return false, nil
}

func detectRemoteOS() (string, error) {
	vars, _, err := Run("ssh", "windows", "env")
	if err != nil {
		return "", err
	}
	for _, line := range vars {
		if strings.HasPrefix(line, "OS=Windows") {
			return "windows", nil
		}
	}
	olines, _, err := Run("ssh", "", "uname")
	if err != nil {
		return "", err
	}
	if len(olines) != 1 {
		return "", fmt.Errorf("unexpected uname response: %v", olines)
	}
	return strings.ToLower(olines[0]), nil
}

func NewController() (Controller, error) {

	_, keyfile, err := GetViperPath("key")
	if err != nil {
		return nil, err
	}

	v := vmctl{
		Hostname: viper.GetString("hostname"),
		Username: viper.GetString("username"),
		KeyFile:  keyfile,
		Path:     viper.GetString("path"),
	}

	relayConfig := viper.GetString("relay")
	if relayConfig != "" {
		r, err := NewRelay(relayConfig)
		if err != nil {
			return nil, err
		}
		v.relay = r
	}

	client, err := newVMRestClient()
	if err != nil {
		return nil, err
	}
	v.api = client

	v.Local = runtime.GOOS
	local, err := isLocal()
	if err != nil {
		return nil, err
	}
	if local {
		v.Remote = v.Local
		switch v.Local {
		case "windows":
			v.Shell = "cmd"
		default:
			v.Shell = "sh"
		}
	} else {
		v.Shell = "ssh"
		remote, err := detectRemoteOS()
		if err != nil {
			return nil, err
		}
		v.Remote = remote
	}
	return &v, nil
}

func (v *vmctl) Close() error {
	if v.relay != nil {
		return v.relay.Close()
	}
	return nil
}

func (v *vmctl) List(name string, detail, all bool) ([]VM, error) {
	vms := []VM{}
        vmxFiles := make(map[string]bool)

        if !all {
	olines, _, err := v.vmrun("list")
	if err != nil {
		return []VM{}, err
	}
	for i, line := range olines {
		fmt.Printf("oline[%d] %s\n", i, line)
	}
	panic("howdy")
	vmids, err := v.api.GetVMs()
	if err != nil {
		return vms, err
	}
	for _, vmid := range vmids {
		vm, err := v.api.GetVM(vmid.Id)
		if err != nil {
			return vms, err
		}
		vms = append(vms, vm)
	}
	return vms, nil
}

func (v *vmctl) Create(name string, options CreateOptions) (VM, error) {
	log.Printf("Create: %s %+v\n", name, options)
	return VM{}, fmt.Errorf("Error: %s", "unimplemented")
}

func (v *vmctl) Destroy(name string, options DestroyOptions) error {
	log.Printf("Destroy: %s %+v\n", name, options)
	return fmt.Errorf("Error: %s", "unimplemented")
}

func (v *vmctl) Start(name string, options StartOptions) error {
	log.Printf("Start: %s %+v\n", name, options)
	return fmt.Errorf("Error: %s", "unimplemented")
}

func (v *vmctl) Stop(name string, options StopOptions) error {
	log.Printf("Stop: %s %+v\n", name, options)
	return fmt.Errorf("Error: %s", "unimplemented")
}

func (v *vmctl) Get(name, property string) (any, error) {
	log.Printf("Get: %s %+v\n", name, property)
	return nil, fmt.Errorf("Error: %s", "unimplemented")
}

func (v *vmctl) Set(name, property string, value any) error {
	log.Printf("Set: %s %s %+v\n", name, property, value)
	return fmt.Errorf("Error: %s", "unimplemented")
}

func (v *vmctl) vmrun(command string) ([]string, []string, error) {
	return Run(v.Shell, v.Remote, "vmrun "+command)
}
