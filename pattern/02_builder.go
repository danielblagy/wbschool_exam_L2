package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Person struct {
	// personal info
	name, address, pin string
	// job info
	workAddress, company, position string
	salary                         int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

func (b *PersonBuilder) BuildPersonalInfo() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) BuildJobInfo() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

//At adds address to person
func (a *PersonAddressBuilder) At(address string) *PersonAddressBuilder {
	a.person.address = address
	return a
}

//WithPostalCode adds postal code to person
func (a *PersonAddressBuilder) WithPostalCode(pin string) *PersonAddressBuilder {
	a.person.pin = pin
	return a
}

type PersonJobBuilder struct {
	PersonBuilder
}

//As adds position to person
func (j *PersonJobBuilder) As(position string) *PersonJobBuilder {
	j.person.position = position
	return j
}

//For adds company to person
func (j *PersonJobBuilder) For(company string) *PersonJobBuilder {
	j.person.company = company
	return j
}

//In adds company address to person
func (j *PersonJobBuilder) In(companyAddress string) *PersonJobBuilder {
	j.person.workAddress = companyAddress
	return j
}

//WithSalary adds salary to person
func (j *PersonJobBuilder) WithSalary(salary int) *PersonJobBuilder {
	j.person.salary = salary
	return j
}

// for example
func BuildPerson() {
	pb := NewPersonBuilder()
	pb.BuildPersonalInfo().
		At("Moscow").
		WithPostalCode("000000").
		BuildJobInfo().
		As("Intern").
		For("WB").
		In("Moscow").
		WithSalary(0)

	person := pb.Build()

	fmt.Println(person)
}
