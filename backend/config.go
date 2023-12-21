package backend

import (
    "encoding/json"
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

func (a *App) GenerateConfig(provider string) (Config, error) {
	log.Println("[backend.GenerateConfig] Attempting to generate config")

	config := Config{
		Provider: provider,
		Vms: []Vm{},
	}	

	err := WriteConfig(config)

	return config, err
}

func WriteConfig( config Config ) (error) {
	log.Println("[backend.WriteConfig] Attempting to write config")

	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".birdVM", "config")

	configJSON, err := json.Marshal(config)
    if err != nil {
        return err
    }

    err = os.WriteFile(configPath, configJSON, 0644)
    if err != nil {
        return err
    }

    log.Println("[backend.WriteConfig] Config written successfully")
    return nil
}