package main

import (
	"fmt"
	"os"
)

func conf() error {
	confDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	fmt.Println(confDir)
	return err
}

func main() {
	err := conf()
	if err != nil {
		panic(err)
	}
}
