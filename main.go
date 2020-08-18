package main

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
)

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "add", Description: "Adds a phone number"},
		{Text: "del", Description: "Deletes a phone number"},
		{Text: "print", Description: "Prints the phone number"},
		{Text: "list", Description: "list all phone numbers"},
		{Text: "save", Description: "saves all phone numbers"},
		{Text: "help", Description: "help!"},
		{Text: "quit", Description: "quit."},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func main() {
	pb := NewFilePhonebook("store.json")
	commands := MakeCommands(pb)
	history := []string{}

	for {
		line := prompt.Input("PB> ", completer, prompt.OptionHistory(history))
		history = append(history, line)

		args := strings.Fields(line)
		command := args[0]
		if _, ok := commands[command]; !ok {
			fmt.Printf("unknown command: %s\n", command)
		}else {
			commands[command](args)
		}
	}
}