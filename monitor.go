package main

import (
	"fmt"
	"net/http"
)

func main() {

	showIntroduction()

	showMenu()

	command := readInput()

	switch command {
	case 0:
		fmt.Println("Exiting the program")
	case 1:
		starMonitoring()
	case 2:
		fmt.Println("Exhibiting logs...")
	default:
		fmt.Println("I don't know this command")
	}
}

func showIntroduction() {
	name := "Harry"
	version := 1.2
	fmt.Println("Hello,", name)
	fmt.Println("This is the version", version, "of the program")
}

func showMenu() {
	fmt.Println("0 - Exit")
	fmt.Println("1 - Start monitoring sites")
	fmt.Println("2 - Show logs")
}

func readInput() int {
	var inputRecieved int
	fmt.Scan(&inputRecieved)
	fmt.Println("The chosen command is:", inputRecieved)
	return inputRecieved
}

func starMonitoring() {
	fmt.Println("Monitoring...")
	// This site give me a random status code every time I access it
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("The site", site, "was successfully accessed!")
	} else {
		fmt.Println("Couldn't access", site, "smoothly. Status Code:", resp.StatusCode)
	}
}
