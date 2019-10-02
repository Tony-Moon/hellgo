package main

import (
	"fmt"
	"os"
)

func translate(inp string) string {
	var translations = make(map[string]string)
	translations["en"] = "Hello"
	translations["es"] = "Hola"
	translations["de"] = "Guten tag"
	translations["fr"] = "Bonjour"
	return translations[inp]
}

func argswitch(inp []string) string {
	var locale string
	if len(inp) == 1 {
		fmt.Println("Enter a language code:")
		fmt.Scanf("%s", &locale)
	} else {
		locale = inp[1]
	}
	return locale
}

func main() {
	locale := argswitch(os.Args)
	output := translate(locale)
	if output == "" {
		output = "Yo"
	}

	fmt.Printf(output + ", Go!\n")
}