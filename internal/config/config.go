package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Username string `json:"current_user_name,omitempty"`
	DbUrl    string `json:"db_url"`
}

const configFileName = ".gatorconfig.json"

func (c *Config) ToString() string {
	return fmt.Sprintf("Username: %v\nDB URL: %v\n", c.Username, c.DbUrl)
}
func (c *Config) SetUser(usr string) error {
	c.Username = usr
	if err := write(*c); err != nil {
		return fmt.Errorf("Cannot write config: %v", err)
	}
	return nil
}
func getConfigFilePath() (string, error) {
	dir, dirErr := os.UserHomeDir()
	if dirErr != nil {
		return "", dirErr
	}
	return filepath.Join(dir, configFileName), nil
}
func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	config, err := getConfigFilePath()
	if err != nil {
		return err
	}

	if err := os.WriteFile(config, jsonData, 0777); err != nil {
		return err
	}
	return nil

}
func Read() Config {
	configFile, err := getConfigFilePath()
	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error opening config: %v", err)
	}
	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		log.Fatalf("Error unmarshaling config: %v\n", err)
	}
	return config
}
