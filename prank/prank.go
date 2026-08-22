package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"math"
)

//go:embed booba.jpeg
var imageBytes []byte

func showLoading(duration time.Duration) error {
	if !validDimension(float64(duration)) {
		return fmt.Errorf("Некорректное число: %v", duration)
	}

	symbols := `-\|/`
	delay := 25 * time.Millisecond
	stop := time.Now().Add(duration)

	for time.Now().Before(stop) {
		for _, symb := range symbols {
			if !time.Now().Before(stop) {
				break
			}

			fmt.Printf("\b%c", symb)
			time.Sleep(delay)
		}
	}

	fmt.Print("\r\033[KГотово!\n")
	return nil
}

func openFile(filePath string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", filePath)
	case "darwin":
		cmd = exec.Command("open", filePath)
	case "linux":
		cmd = exec.Command("xdg-open", filePath)
	default:
		return fmt.Errorf("почему ты такой тупой")
	}
	return cmd.Run()
}

func main() {
	tempFile, err := os.CreateTemp("", "sys_update_*.jpg")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	tempFile.Close()

	err = os.WriteFile(tempFile.Name(), imageBytes, 0644)
	if err != nil {
		fmt.Println("ты курочка", err)
		return
	}

	fmt.Println("НА ВАШЕМ КОПЬЮТЕРЕ ЕСТЬ ВИРУС!!")
	time.Sleep(time.Second)
	fmt.Print("ПРОИЗОВЖУ ЕГО УДАЛЕНИЕ  ")
	showLoading(4 * time.Second)
	
	err = openFile(tempFile.Name())
	if err != nil {
		fmt.Println("у тебя попка", err)
	}

}


func validDimension(value float64) bool {
	return value >= 0 &&
		!math.IsNaN(value) &&
		!math.IsInf(value, 0)
}
