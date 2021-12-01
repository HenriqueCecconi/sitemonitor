package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const repetitions = 3
const delay = 5

func main() {

	showIntroduction()

	for {
		showMenu()

		command := readInput()

		switch command {
		case 0:
			fmt.Println("Exiting the program")
			os.Exit(0)
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing the logs...")
			fmt.Println()
		default:
			fmt.Println("I don't know this command. Please enter a valid one.")
		}
	}
}

func showIntroduction() {
	name := "Harry"
	version := 1.2
	fmt.Println("Hello,", name)
	fmt.Println("This is the version", version, "of the program")
}

func showMenu() {
	fmt.Println("---------------------------------------------")
	fmt.Println("0 - Exit this program")
	fmt.Println("1 - Start monitoring the specified sites")
	fmt.Println("2 - Show the logs for the previous attempts")
	fmt.Println("---------------------------------------------")
	fmt.Print("Insert your command: ")
}

func readInput() int {
	var inputRecieved int
	fmt.Scan(&inputRecieved)
	fmt.Println("You chose the command:", inputRecieved)
	fmt.Println()
	return inputRecieved
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	fmt.Println()

	monitoredSites := readSitesList()

	for i := 0; i < repetitions; i++ {
		for _, sites := range monitoredSites {
			fmt.Println("Testing the site:", sites)
			testConnection(sites)
		}
		// This is set to Seconds because it's just a demonstration, made to be quick
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
}

func testConnection(sites string) {
	resp, err := http.Get(sites)

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("The site", sites, "was successfully accessed!")
	} else {
		fmt.Println("Couldn't access", sites, "smoothly. Status Code:", resp.StatusCode)
	}
	fmt.Println()
}

func readSitesList() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	reader := bufio.NewReader(arquivo)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	fmt.Println("Testing the following sites:", sites)

	return sites
}
