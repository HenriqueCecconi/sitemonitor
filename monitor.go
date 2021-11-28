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
	//fmt.Scanf("%d", &command)
	fmt.Println("The chosen command is:", command)

	if command == 0 {
		fmt.Println("Exiting the program")
	} else if command == 1 {
		fmt.Println("Monitoring...")
	} else if command == 2 {
		fmt.Println("Exhibiting logs...")
	} else {
		fmt.Println("I don't know this command")
	}
}
