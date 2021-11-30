package main

import (
	"fmt"
	"net/http"
	"os"
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
	// This first site give me a random status code every time I access it
	monitoredSites := []string{"https://random-status-code.herokuapp.com/", "https://google.com", "https://github.com"}

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
	resp, _ := http.Get(sites)
	if resp.StatusCode == 200 {
		fmt.Println("The site", sites, "was successfully accessed!")
	} else {
		fmt.Println("Couldn't access", sites, "smoothly. Status Code:", resp.StatusCode)
	}
	fmt.Println()
}
