package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const DELAY_IN_SECONDS time.Duration = 0

var STOP_MENU_COMMAND bool = false

func main() {
	showInstroduction()
	for {
		showMenu()

		command := readCommand()
		switch command {
		case 0:
			fmt.Println("Exiting program...")
			break
		case 1:
			fmt.Println("Monitoring...")
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}
}

func showInstroduction() {
	var name string = "Hiago"
	var age int
	var version float32 = 1.1
	fmt.Println("Hello Sr.", name, "your age is", age)
	fmt.Println("Program version is: ", version)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var readedCommand int
	fmt.Scan(&readedCommand)
	fmt.Println("The command chosen was", readedCommand)
	return readedCommand
}

func startMonitoring() {
	sites := []string{"https://httpstat.us/200", "https://httpstat.us/500", "https://httpstat.us/200", "https://httpstat.us/404"}
	for _, site := range sites {
		checkSite(site)
	}

}

func checkSite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "was successfully loaded!")
	} else {
		fmt.Println("Site:", site, "is with problems. Status code:", response.StatusCode)
	}

	time.Sleep(DELAY_IN_SECONDS * time.Second)
}
