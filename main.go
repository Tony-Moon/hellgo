package main

import (
	"fmt"
	"strconv"
)

func langSelect(int locale) string{
	switch locale {
	case 1:
		return = "Hello"
	case 2:
		return = "Hola"
	case 3:
		return = "Guten Tag"
	case 4:
		return = "Bonjour"
	default:
		return = "Yo"
	}
}

func scan(in io.Reader)

func main() {
	var greeting, userInput string
	var selector int
	var languages = [4]string{"English", "Espanol", "Deutsche", "Francais"}

	for i:=0;i<4;{
		fmt.Printf("[" + i +"] " + languages[i] + " \n")
	}

	fmt.Scanf("%s", &userInput)
	selector := strconv.Atoi(userInput)

	greeting = langSelect(selector)
	fmt.Printf(greeting + ", Go!\n")
}
