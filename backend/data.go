package backend

type Playbook struct {
	Path string
}

type Vm struct {
	Name string
	Os string
	Playbook Playbook
}