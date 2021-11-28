package main

import "fmt"

func main() {
	name := "Harry"
	version := 1.1

	fmt.Println("Hello,", name)
	fmt.Println("This is the version", version, "of the program")

	fmt.Println("0 - Exit")
	fmt.Println("1 - Start monitoring sites")
	fmt.Println("2 - Show logs")

	var command int
	fmt.Scan(&command)
	fmt.Println("The chosen command is:", command)

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
