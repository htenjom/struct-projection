package main

import (
	"fmt"
)

//Choices:
//1. Pure reflection R//Expensive and possibly hard
//2. Map struct to JSON and operate over the JSON string R// Great and useful but overkill
//3. Map struct to map -> filter -> map to JSON R// First approach but difficult for nested structures
//4. Use libs with before behavior:
// 	 - copier || deepcopier: useful for same structure, but no more complex if required https://github.com/jinzhu/copier
//   - mapstruct: https://github.com/mitchellh/mapstructure useful but for other purposes when knowing fields is required
//5. Manual mapping with functions R//It´s easier, doesn't require additional libs, faster than json manipulation and flexible.
func main() {
	testCopier()
}

func testCopier() {
	person := buildPerson()
	brModel := countryModels["BR"](person)
	arModel := countryModels["AR"](person)

	fmt.Printf("::: Person %v\n", person)
	fmt.Printf("::: Person BR %v\n", brModel)
	fmt.Printf("::: Person AR %v\n", arModel)
}

func buildPerson() Person {
	return Person{
		Name:       "First Name",
		Age:        30,
		SomeNumber: 15,
		Address: Address{
			StreetName: "My StreetName",
			ZipCode:    "ZIP-12345",
		},
	}
}
