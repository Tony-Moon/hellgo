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
	var languages = [4]string{
		"Press    [1] for English", 
		"Presione [2] para Espanol", 
		"Presse   [3] für Deutsche",
		"Presse   [4] Pour Francais" }

	for i:=0;i<4;i++{
		fmt.Printf(languages[i] + " \n")
	}

	fmt.Scanf("%s", &userInput)

	greeting = langSelect(userInput)
	fmt.Printf(greeting + ", Go!\n")
}
