package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func loadEnvFile() {
	rootDir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Error getting root directory: %v", err)
	}

	envPath := filepath.Join(rootDir, ".env")

	file, err := os.Open(envPath)

	if err != nil {
		log.Fatalf("Error opening .env file: %v", err)
	}

	defer file.Close()

	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		line := scaner.Text()

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		os.Setenv(key, value)
	}

	if err := scaner.Err(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}
}
