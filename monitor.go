package main

import "fmt"

func main() {
	name := "Harry"
	version := 1.1

	fmt.Println("Hello,", name)
	fmt.Println("This is the version", version, "of the program")

	fmt.Println("1 - Start monitoring sites")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")

	var command int
	fmt.Scanf("%d", &command)
}
