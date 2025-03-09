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
func GetEnv(filePath, key string) (string, error) {
	envMap, err := LoadEnv(filePath)
	if err != nil {
		return "", err
	}
	value, exists := envMap[key]
	if !exists {
		return "", nil
	}
	return value, nil
}