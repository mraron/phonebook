package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name string
	Phone string
}

var people = make([]Person, 0)

var commands = map[string]func([]string){
	"quit": func(args []string) {
		os.Exit(0)
	},
	"help": func(args []string) {
		fmt.Println("@TODO")
	},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("PB> ")
		line, _ := reader.ReadString('\n')

		args := strings.Fields(line)
		command := args[0]
		if _, ok := commands[command]; !ok {
			fmt.Printf("unknown command: %s\n", command)
		}else {
			commands[command](args)
		}
	}
}