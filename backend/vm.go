package backend

import "log"

func (a *App) Test(name string) string {

	log.Println("test")
	return name
}