package main

import (
	"fmt"
)

func langSelect(userInput string) string{
	switch userInput {
	case "1":
		return "Hello"
	case "2":
		return "Hola"
	case "3":
		return "Guten Tag"
	case "4":
		return "Bonjour"
	default:
		return "Yo"
	}
}

func main() {
	var greeting, userInput string
	var languages = [4]string{"English", "Espanol", "Deutsche", "Francais"}
	var itera = [4]string{"1", "2", "3", "4"}

	for i:=0;i<4;i++{
		
		fmt.Printf("[" + itera[i] +"] " + languages[i] + " \n")
	}

	fmt.Scanf("%s", &userInput)

	greeting = langSelect(userInput)
	fmt.Printf(greeting + ", Go!\n")
}
