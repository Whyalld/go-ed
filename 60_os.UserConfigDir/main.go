package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func config() (string, error) {
	confDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(confDir, "name"), err
}

func main() {
	configPath, err := config()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
	}
	fmt.Println(configPath)
}
