package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"indication/internal/config"

	"github.com/stretchr/testify/assert"
)

// func TestMain(m *testing.M) {
// 	code := m.Run()
// 	os.Exit(code)
// }

func TestValidConfig(t *testing.T) {
	tempDir := t.TempDir()

	err := os.MkdirAll(tempDir, os.ModePerm)
	assert.NoError(t, err)

	configFile := filepath.Join(tempDir, "Config.yaml")
	configContent := `
	Server
		port: 8888
		name: "TestServer"
	db:
		port: 123456789
		db: "testdb"
		host: "localhost"
		username: "testuser"
		password: "testpswd"
		sslmode: "disable"
	`
	err = os.WriteFile(configFile, []byte(configContent), 0777)
	assert.NoError(t, err)

	cfg, err := config.New(tempDir, "Config.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	assert.Equal(t, 8888, cfg.Server.Port)
}

func TestNew_EmptyConfig(t *testing.T) {
	tempDir := t.TempDir()
	configFile := "config.yaml"
	err := os.WriteFile(filepath.Join(tempDir, configFile), []byte(""), 0644)
	if err != nil {
		t.Fatalf("Failed to create config file %v", err)
	}
	cfg, err := config.New(tempDir, configFile)

	assert.Error(t, err)
	assert.Nil(t, cfg)
	//assert.Contains(t, err.Error(), "unmarshal")
}
func TestNew_FileDoesNotExist(t *testing.T) {
	tempDir := t.TempDir()
	nonExistentFile := "non_existent_config"

	cfg, err := config.New(tempDir, nonExistentFile)

	assert.Error(t, err)
	assert.Nil(t, cfg)
	assert.Contains(t, err.Error(), "Config File \"non_existent_config\" Not Found in")
}
