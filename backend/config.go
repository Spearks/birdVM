package backend

import (
    "encoding/json"
    "log"
    "os"
    "path/filepath"
)

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