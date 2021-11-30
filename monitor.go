package main

import (
	"fmt"
	"net/http"
	"os"
)

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
	// This first site give me a random status code every time I access it
	monitoredSites := []string{"https://random-status-code.herokuapp.com/", "https://google.com", "https://github.com"}

	for _, sites := range monitoredSites {
		resp, _ := http.Get(sites)
		if resp.StatusCode == 200 {
			fmt.Println("The site", sites, "was successfully accessed!")
		} else {
			fmt.Println("Couldn't access", sites, "smoothly. Status Code:", resp.StatusCode)
		}
	}
}
