package backend 

import (
    "log"
    "os"
    "path/filepath"
)

func (a *App) CheckIfIsInstalled() bool {
    log.Println("[backend.CheckIfIsInstalled] Running verification")

    homeDir, _ := os.UserHomeDir()
    configPath := filepath.Join(homeDir, ".birdVM", "config")

    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        log.Println("[backend.CheckIfIsInstalled] File does not exist")
        return false
    } else {
        log.Println("[backend.CheckIfIsInstalled] File exists")
        return true
    }
}