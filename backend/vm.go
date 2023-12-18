package backend

import (
	"log"
	"net"
	"time"

	"github.com/digitalocean/go-libvirt"
)

func Connection(socketPath string) (*libvirt.Libvirt, error) {
	c, err := net.DialTimeout("unix", socketPath, 2*time.Second)
	if err != nil {
		log.Fatalf("[backend.vm] failed to dial libvirt: %v", err)
		return nil, err
	}

	l := libvirt.New(c)

	return l, nil
}


func (a *App) TestConnection(socketPath string) (bool, error) {

	l, err := Connection(socketPath)

	if err != nil {
		return false, err
	}

	err = l.Connect()

	if err != nil {
		return false, err
	}

	log.Println("[backend.vm] Suceessfully connected to libvirt")

	return false, nil
}