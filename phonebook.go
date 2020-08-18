package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Phone string
}

//Phonebook interface to make phonebook more testable
//Default implementation is FilePhonebook which saves a json into a file
//@TODO tests and mock implementation
type Phonebook interface {
	Add(Person) error
	Delete(string) error
	Find(string) (*Person, error)
	People() ([]*Person, error)
	Save() error
	Load() error
}

type FilePhonebook struct {
	people []*Person
	file string
}

func NewFilePhonebook(file string) *FilePhonebook {
	return &FilePhonebook{make([]*Person, 0), file}
}

func (fp *FilePhonebook) findIndexOf(name string) int {
	index := -1
	for idx, curr := range fp.people {
		if name == curr.Name {
			index = idx
		}
	}

	return index
}

func (fp *FilePhonebook) Add(p Person) error {
	fp.people = append(fp.people, &p)
	return nil
}

func (fp *FilePhonebook) Delete(name string) error {
	index := fp.findIndexOf(name)
	if index == -1 {
		return fmt.Errorf("%s not found, can't delete", name)
	}

	return nil
}

func (fp *FilePhonebook) Find(name string) (*Person, error) {
	index := fp.findIndexOf(name)
	if index == -1 {
		return nil, fmt.Errorf("%s not found, can't find", name)
	}

	return fp.people[index], nil
}

func (fp *FilePhonebook) People() ([]*Person, error) {
	return fp.people, nil
}

func (fp *FilePhonebook) Save() error {
	f, err := os.Create(fp.file) //overwrites the actual contents of file
	if err != nil {
		return err
	}
	defer f.Close() //runs after return value so ok

	enc := json.NewEncoder(f)
	return enc.Encode(fp.people)
}

func (fp *FilePhonebook) Load() error {
	f, err := os.OpenFile(fp.file, os.O_RDWR|os.O_APPEND, 0660) //os.Open: read-only os.Create: truncates file
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	return dec.Decode(&fp.people) //dangerous, destroys the content of people
}