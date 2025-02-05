package main

import (
	"fmt"

	"github.com/danielblagy/wbschool_exam_L2/pattern"
)

func main() {
	fmt.Println("Facade Pattern")
	computer := pattern.ComputerFacade{}
	computer.Start()
	fmt.Println(computer)

	fmt.Println("Builder Pattern")
	pattern.BuildPerson()

	fmt.Println("Visitor Pattern")
	pattern.VisitorExample()

	fmt.Println("Command Pattern")
	pattern.CommandExample()

	fmt.Println("Chain of responsibility Pattern")
	pattern.ChainOfRespExample()

	fmt.Println("Factory Method Pattern")
	pattern.FactoryMethodExample()

	fmt.Println("Strategy Pattern")
	pattern.StrategyExample()

	fmt.Println("State Pattern")
	pattern.StateExample()
}
