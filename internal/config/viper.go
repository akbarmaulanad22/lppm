package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"tugasakhir/internal/helper"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	// Tentukan path relatif ke file .env, 2 tingkat di atas
	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(currentFile), "..", "..") // naik 2 level
	envPath := filepath.Join(projectRoot, ".env")

	// Cek apakah file .env ada
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Fatalf("could not find the .env file in: %s", envPath)
	}

	config := viper.New()

	// Setup Viper
	config.SetConfigFile(envPath) // set full path ke .env
	config.SetConfigType("env")

	err := config.ReadInConfig()
	helper.FatalIfErrorWithMessage(err, fmt.Sprintf(".env File no found: %s", err))

	return config
}
