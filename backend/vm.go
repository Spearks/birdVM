package backend

import (
	"log"
	"github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-libvirt/socket/dialers"
	"time"
)

var STATES = map[libvirt.ConnectListAllDomainsFlags]string{libvirt.ConnectListDomainsRunning: "Running", libvirt.ConnectListDomainsPaused: "Paused", libvirt.ConnectListDomainsShutoff: "Shutoff"}

func Connect() (*libvirt.Libvirt, error) {
	// Connect to Libvirt unix domain socket
	localSocket := dialers.NewLocal()
	libvirtconn := libvirt.NewWithDialer(localSocket)
	err := libvirtconn.Connect()
	if err != nil {
		log.Fatalf("Failed to dial libvirt socket: %v", err)
		return nil, err
	}

	return libvirtconn, nil
}

func (a * App) GetAllVmsStats() []Vm {
	
	log.Println("[backend.vm] Getting domains")
	
	conn, _ := Connect()
	vms := []Vm{}

	for state := range STATES {
		flags := state
		domains, _, _ := conn.ConnectListAllDomains(1, flags)
		for _, vm := range domains {

			domainInfo, err := conn.DomainMemoryStats(vm, 8, 0)

			if err == nil {
				log.Println("[backend.vm] Getting domain info for", vm.Name)
				// https://github.com/runyutech/go-cpulimiter/blob/master/pkg/drivers/libvirt-kvm.go
				var cpucount uint16
				var cputime2 uint64

				_, _, _, _, cputime1, _ := conn.DomainGetInfo(vm)

				time.Sleep(1 * time.Second)
	
				_, _, _, cpucount, cputime2, _ = conn.DomainGetInfo(vm)
	
				var cputimepercent = 100 * (cputime2 - cputime1) / (1000000000 * uint64(cpucount))

				// Convert memory from KBytes to GB
				memory := int32(domainInfo[2].Val / 1024)

				vmInfo := Vm{
					Name:   vm.Name,
					Os:     "", 
					Memory: memory,
					CpuUsage:    cputimepercent, 
				}
				vms = append(vms, vmInfo)
			}
		}
	}

	return vms
}

func (a *App) TestConnection(socketPath string) (bool, error) {

	l, err := Connect()

	if err != nil {
		return false, err
	}

	err = l.Connect()

	if err != nil {
		return false, err
	}

	log.Println("[backend.vm] Suceessfully connected to libvirt")

	return true, nil
}