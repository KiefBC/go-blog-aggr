package config

import (
	"encoding/json"
	"os"
)

// Config holds the application configuration.
// It includes the database URL and the current user's name.
type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

// Read reads the configuration from the user's home directory.
// It returns a Config struct or an error if the file cannot be read or parsed.
func Read() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// SetUser sets the current user in the configuration and writes it back to the config file.
// It returns an error if the write operation fails.
func (c *Config) SetUser(username string) error {
	c.Current_user_name = username
	return write(c)
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/.gatorconfig.json", nil
}

// write writes the configuration to the user's home directory.
// It returns an error if the file cannot be written.
func write(cfg *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
