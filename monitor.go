package main

import "fmt"

func main() {

	showIntroduction()

	showMenu()

	command := readInput()

	switch command {
	case 0:
		fmt.Println("Exiting the program")
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Exhibiting logs...")
	default:
		fmt.Println("I don't know this command")
	}
}

func showIntroduction() {
	name := "Harry"
	version := 1.1
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
