package main

import (
	"encoding/csv"
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
			os.Exit(0)
		case 1:
			fmt.Println("Monitoring...")
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
			readLog()
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
	sites := readCsvFile("sites.csv")

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

	writeLogCsv(site, response.StatusCode == 200)

	time.Sleep(DELAY_IN_SECONDS * time.Second)
}

func readCsvFile(fileName string) []string {
	var listItens []string

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := csv.NewReader(file)
	itens, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
	}

	itens = itens[1:]
	fmt.Println(itens)

	for _, item := range itens {
		listItens = append(listItens, item[0])
	}

	file.Close()

	return listItens
}

func writeLogCsv(site string, status bool) {
	file, err := os.OpenFile("log.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}
	// write first line if file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error:", err)
	}

	if fileInfo.Size() == 0 {
		file.WriteString("date,site,status\n")
	}

	file.WriteString(time.Now().Format(time.RFC3339) + "," + site + "," + fmt.Sprint(status) + "\n")
	file.Close()
}

func readLog() {
	file, err := os.Open("log.csv")
	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := csv.NewReader(file)
	itens, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, log := range itens {
		fmt.Println(log)
	}

	file.Close()

}
