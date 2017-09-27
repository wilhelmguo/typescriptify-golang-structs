package main

import (
	"github.com/wilhelmguo/typescriptify-golang-structs/typescriptify"
)

type TypeString string

type Address struct {
	// Used in html
	City    string  `json:"city"`
	Number  float64 `json:"number"`
	Country string  `json:"country,omitempty"`
}

type PersonalInfo struct {
	Hobbies []string `json:"hobby"`
	PetName string   `json:"pet_name"`
}

type Person struct {
	Address `json:",inline"`
	PersonalInfo        `json:"metadata,omitempty"`
	AddressMeta Address `json:"address,omitempty"`
}

func main() {
	converter := typescriptify.New()
	converter.CreateFromMethod = true
	converter.Indent = "    "

	converter.Add(Person{})

	err := converter.ConvertToFile("browser_test/example_output.ts")
	if err != nil {
		panic(err.Error())
	}
}
