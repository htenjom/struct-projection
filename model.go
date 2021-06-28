package main

type Address struct {
	StreetName string
	ZipCode    string
}

type Person struct {
	Name       string
	Age        int
	SomeNumber int
	Address
}

type AddressBR struct {
	ZipCode string
}

type AddressAR struct {
	StreetName string
}

type PersonBR struct {
	Name string
	AddressBR
}

type PersonAR struct {
	Name string
	Age  int
	AddressAR
}
