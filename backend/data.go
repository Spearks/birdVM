package backend

type Vm struct {
	Name string
	Os string
	Memory int32
	CpuUsage uint64
}

type Provider struct {
	Name string
}

type Config struct {
	Provider string
	Vms []Vm
}