package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindConfigFile searches for sqlcquash.yaml in the current directory and its subdirectories
func FindConfigFile() (string, error) {
	const configFileName = "sqlcquash.yaml"

	// Get current working directory
	startDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var configPath string
	err = filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden directories
		if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		// Check if current file is the config file
		if !info.IsDir() && info.Name() == configFileName {
			configPath = path
			return filepath.SkipAll // Stop walking once found
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	if configPath == "" {
		return "", fmt.Errorf("config file %s not found in %s or its subdirectories", configFileName, startDir)
	}

	return configPath, nil
}

// FindConfigDir searches for the directory containing the config file
func FindConfigDir(cfgPath string) (string, error) {
	configDir := filepath.Dir(cfgPath)
	if configDir == "." {
		return "", fmt.Errorf("config file %s not found in current directory", cfgPath)
	}
	return configDir, nil
}
