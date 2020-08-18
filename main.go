package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name string
	Phone string
}

var people = make([]Person, 0)

func findIndexOf(name string) int {
	index := -1
	for idx, curr := range people {
		if name == curr.Name {
			index = idx
		}
	}

	return index
}

func savePeople(file string) error {
	f, err := os.Create(file) //overwrites the actual contents of file
	if err != nil {
		return err
	}
	defer f.Close() //runs after return value so ok

	enc := json.NewEncoder(f)
	return enc.Encode(people)
}

func loadPeople(file string) error {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0660) //os.Open: read-only os.Create: truncates file
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	return dec.Decode(&people) //dangerous, destroys the content of people
}

var commands = map[string]func([]string)int {
	"add": func(args []string) int { //@TODO: support multiple names
		if len(args) != 3 {
			fmt.Println("usage: add <name> <phone>")
			return 1
		}

		people = append(people, Person{args[1], args[2]})
		return 0
	},
	"del": func(args []string) int {
		if len(args) != 2 {
			fmt.Println("usage: del <name>")
			return 1
		}

		index := findIndexOf(args[1])
		if index == -1 {
			fmt.Printf("%s not found, can't delete\n", args[1]);
			return 1
		}

		people = append(people[:index], people[index+1:]...)
		return 0
	},
	"print": func(args []string) int {
		if len(args) != 2 {
			fmt.Println("usage: print <name>")
			return 1
		}

		index := findIndexOf(args[1])
		if index == -1 {
			fmt.Printf("%s not found, can't print\n", args[1])
			return 1
		}

		fmt.Println(people[index].Phone)
		return 0
	},
	"list": func(args []string) int {
		for _, curr := range people {
			fmt.Println(curr.Name, curr.Phone)
		}

		fmt.Printf("\nok, listed %d name-phone pairs\n", len(people))
		return 0
	},
	"save": func(args []string) int {
		err := savePeople("store.json")
		if err != nil {
			fmt.Println("error while saving", err)
			return 1
		}
		return 0
	},
	"load": func(args []string) int {
		err := loadPeople("store.json")
		if err != nil {
			fmt.Println("error while loading", err)
			return 1
		}
		return 0
	},
	"quit": func(args []string) int {
		os.Exit(0)
		return 0
	},
	"help": func(args []string) int {
		fmt.Println("@TODO")
		return 0
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