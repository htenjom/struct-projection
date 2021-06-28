package main

func NewPersonAR(person Person) interface{} {
	return PersonAR{
		Name: person.Name,
		Age:  person.Age,
		AddressAR: AddressAR{
			StreetName: person.StreetName,
		},
	}
}

func NewPersonBR(person Person) interface{} {
	return PersonBR{
		Name: person.Name,
		AddressBR: AddressBR{
			ZipCode: person.ZipCode,
		},
	}
}

var countryModels = map[string]func(Person) interface{}{
	"AR": NewPersonAR,
	"BR": NewPersonBR,
}
