package main

import (
	"fmt"
	"os"
	"strings"
)

func MakeCommands(pb Phonebook) map[string]func([]string)int {
	var commands = map[string]func([]string)int {
		"add": func(args []string) int {
			if len(args) < 3 {
				fmt.Println("usage: add <name> <phone>")
				return 1
			}

			if err := pb.Add(Person{strings.Join(args[1:len(args)-1], " "), args[len(args)-1]}); err != nil {
				fmt.Println(err)
				return 1
			}

			return 0
		},
		"del": func(args []string) int {
			if len(args) < 2 {
				fmt.Println("usage: del <name>")
				return 1
			}

			name := strings.Join(args[1:len(args)], " ")
			if err := pb.Delete(name); err != nil {
				fmt.Println(err)
				return 1
			}
			return 0
		},
		"print": func(args []string) int {
			if len(args) < 2 {
				fmt.Println("usage: print <name>")
				return 1
			}

			name := strings.Join(args[1:len(args)], " ")
			p, err := pb.Find(name)
			if err != nil {
				fmt.Println(err)
				return 1
			}

			fmt.Println(p.Phone)
			return 0
		},
		"list": func(args []string) int {
			people, err := pb.People()
			if err != nil {
				fmt.Println(err)
				return 1
			}

			for _, curr := range people {
				fmt.Println(curr.Name, curr.Phone)
			}

			fmt.Printf("\nok, listed %d name-phone pairs\n", len(people))
			return 0
		},
		"save": func(args []string) int {
			if err := pb.Save(); err != nil {
				fmt.Println(err)
				return 1
			}
			return 0
		},
		"load": func(args []string) int {
			if err := pb.Load(); err != nil {
				fmt.Println(err)
				return 1
			}
			return 0
		},
		"quit": func(args []string) int {
			os.Exit(0)
			return 0
		},
		"help": func(args []string) int {
			fmt.Println("add <name> <phone>: adds <name> with <phone> to the store (for example: add Kiss Béla +6423432)")
			fmt.Println("del <name>: deletes <name> from to the store (for example: del Kiss Béla)")
			fmt.Println("print <name>: prints <name>'s phone number (for example: print Kiss Béla)")
			fmt.Println("list: lists all numbers")
			fmt.Println("save: save the list")
			fmt.Println("load: load the list")
			fmt.Println("help: displays this message")
			fmt.Println("quit: quits from the program")
			return 0
		},
	}

	return commands
}