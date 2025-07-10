package vmx

type VM struct {
	Name string
	ID   string

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
	List(string, bool) ([]VM, error)
	Create(string, CreateOptions) (VM, error)
	Destroy(string, bool) error
	Start(string, StartOptions) error
	Stop(string, StopOptions) error
	Get(string, string) (any, error)
	Set(string, string, any) error
	ReadConfig(string) ([]bytes, error)
	WriteConfig(string, []bytes) error
}

type ControllerOptions struct {
	Remote   bool
	Hostname string
	Username string
	KeyFile  string
	Path     string
}

type vmctl struct {
	cxn ControllerOptions
}

func NewController(options ControllerOptions) Controller {
	return &vmctl{cxn: options}
}

func (v *vmctl) List(name string, running bool) ([]VM, error) {
	vms := []VM{}
	log.Printf("List: %s %+v\n", name, running)
	return vms, fmt.Errorf("Error: %s", "unimplemented")
}

func (v *vmctl) Create(name string, options CreateOptions) (VM, error) {
	log.Printf("Create: %s %+v\n", name, options)
	return nil, fmt.Errorf("Error: %s", "unimplemented")
}

func (v *vmctl) Destroy(name string, options DestroyOptions) error {
	log.Printf("Destroy: %s %+v\n", name, options)
	return nil, fmt.Errorf("Error: %s", "unimplemented")
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
	return fmt.Errorf("Error: %s", "unimplemented")
}

func (v *vmctl) Set(name, property string, value any) error {
	log.Printf("Set: %s %s %+v\n", name, property, value)
	return fmt.Errorf("Error: %s", "unimplemented")
}

func (v *vmctl) ReadConfig(name string) ([]byte, error) {
	log.Printf("ReadConfig: %s\n", name)
	return []byte{}, fmt.Errorf("Error: %s", "unimplemented")
}

func (v *vmctl) WriteConfig(name string, data []byte) error {
	log.Printf("WriteConfig: %s (%d bytes)\n", name, len(data))
	return fmt.Errorf("Error: %s", "unimplemented")
}
