package main

import (
	"fmt"
	"time"
)

func showLoading(stopChan chan bool) {
	symbols := `-\|/`
	delay := 100 * time.Millisecond

	for {
			for _, symb := range symbols {
			select {
			case <- stopChan:
				fmt.Println("\033[1m\rГотово!           \033[0m")
				return
			default:	
				fmt.Printf("\033[1m\r%c Загрузка...\033[0m", symb)
				time.Sleep(delay)
			}
		}
	}
}

// Имметация работы некой работы
func doNothing() {
	time.Sleep(4 * time.Second)
}

func main() {
	stopChan := make(chan bool)
	
	go showLoading(stopChan)
	
	doNothing()

	stopChan <- true
}