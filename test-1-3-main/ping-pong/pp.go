package main

import "fmt"

func main() {
	fmt.Println("The program prints Ping, Pong or a number.")
	fmt.Println("Enter the number.")
	var num int
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Print("Ping")
	}
	if num%7 == 0 {
		fmt.Print("Pong")
	}
	if num%2 == 0 && num%7 == 0 {
		fmt.Print(num)
	}
	fmt.Println()
}
