package internals

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func NewUUID() string {
	return uuid.NewString()
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(strings.TrimSpace(u))
	return err == nil
}

func InitUUID() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".gocp")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	path := filepath.Join(dir, "uuid")

	if data, err := os.ReadFile(path); err == nil {
		str := strings.TrimSpace(string(data))
		if isValidUUID(str) {
			return str, nil
		}
	}

	newUUID := NewUUID()

	if err := os.WriteFile(path, []byte(newUUID), 0644); err != nil {
		return "", err
	}

	return newUUID, nil
}
