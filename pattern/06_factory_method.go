package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
The client only interacts with a factory struct and tells the kind of instances that needs
to be created. The factory class interacts with the corresponding concrete structs and
returns the correct instance back.
*/

type iGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) GetName() string {
	return g.name
}

func (g *Gun) SetPower(power int) {
	g.power = power
}

func (g *Gun) GetPower() int {
	return g.power
}

type Pistol9mm struct {
	Gun
}

func NewPistol9mm() iGun {
	return &Pistol9mm{
		Gun: Gun{
			name:  "9mm Pistol",
			power: 4,
		},
	}
}

type DoubleShotgun struct {
	Gun
}

func NewDoubleShotgun() iGun {
	return &DoubleShotgun{
		Gun: Gun{
			name:  "Double-barreled shotgun",
			power: 28,
		},
	}
}

// factory
func GetGun(gunType string) (iGun, error) {
	if gunType == "9mm Pistol" {
		return NewPistol9mm(), nil
	}
	if gunType == "Double-barreled shotgun" {
		return NewDoubleShotgun(), nil
	}
	return nil, fmt.Errorf("No such gun type")
}

func FactoryMethodExample() {
	pistol, _ := GetGun("9mm Pistol")
	fmt.Printf("name: %s\npower: %d\n\n", pistol.GetName(), pistol.GetPower())

	shotgun, _ := GetGun("Double-barreled shotgun")
	fmt.Printf("name: %s\npower: %d\n\n", shotgun.GetName(), shotgun.GetPower())
}
