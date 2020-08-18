package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	pb := NewFilePhonebook("store.json")
	commands := MakeCommands(pb)

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