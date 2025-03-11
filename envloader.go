package envloader

import (
	"bufio"
	"os"
	"strings"
)

// LoadEnv reads a .env file and returns a map of key-value pairs
func LoadEnv(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	envMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			envMap[key] = value
		}
	}
	return envMap, scanner.Err()
}

// GetEnv fetches a specific key from the .env file
func GetEnv(filePath, key string, defaultValue ...string) (string, error) {
	// Load ENV file
	envMap, err := LoadEnv(filePath)
	if err != nil {
		return "", err
	}

	// Fetch key value from env content
	value, exists := envMap[key]
	if !exists {
		// If a default value was provided, return it
		if len(defaultValue) > 0  {
			return defaultValue[0], nil
		}
		return "", nil
	}
	return value, nil
}