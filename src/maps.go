package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]

	statics := pickStatics(words)

	imprint(statics)
}

func pickStatics(words []string) map[string]int {
	statics := make(map[string]int)

	for _, word := range words {
		ini := strings.ToUpper(string(word[0]))
		cont, found := statics[ini]
		if found{
			statics[ini] = cont + 1
		} else {
			statics[ini] = 1
		}
	}
	return statics
}

func imprint(statics map[string]int){
	fmt.Println("word count starting at each letter:")

	for initial, cont := range statics{
		fmt.Printf("%s = %d\n", initial, cont)
	}
}