package envloader

import (
	"os"
	"testing"
)

func createTestEnvFile() (string, error) {
	fileName := "test.env"
	content := "TEST_KEY=TestValue\nANOTHER_KEY=AnotherValue\nEMPTY_KEY=\n"

	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func TestLoadEnv(t *testing.T) {
	fileName, err := createTestEnvFile()
	if err != nil {
		t.Fatalf("Failed to create test.env file: %v", err)
	}
	defer os.Remove(fileName)

	envMap, err := LoadEnv(fileName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tests := map[string]string{
		"TEST_KEY":    "TestValue",
		"ANOTHER_KEY": "AnotherValue",
		"EMPTY_KEY":   "",
	}

	for key, expected := range tests {
		actual, exists := envMap[key]
		if !exists {
			t.Errorf("Expected key %s to exist, but it was missing", key)
		}
		if actual != expected {
			t.Errorf("For key %s: expected %s, got %s", key, expected, actual)
		}
	}
}

func TestGetEnv(t *testing.T) {
	fileName, err := createTestEnvFile()
	if err != nil {
		t.Fatalf("Failed to create test.env file: %v", err)
	}
	defer os.Remove(fileName)

	tests := []struct {
		key      string
		expected string
	}{
		{"ANOTHER_KEY", "AnotherValue"},
		{"TEST_KEY", "TestValue"},
		{"EMPTY_KEY", ""}, // Edge case: key exists but is empty
		{"NON_EXISTENT", ""}, // Edge case: key does not exist
	}

	for _, test := range tests {
		value, err := GetEnv(fileName, test.key)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if value != test.expected {
			t.Errorf("For key %s: expected %q, got %q", test.key, test.expected, value)
		}
	}
}
