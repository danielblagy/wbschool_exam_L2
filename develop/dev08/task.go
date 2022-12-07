package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("# ")
loop:
	for scanner.Scan() {
		input := scanner.Text()

		command := strings.Split(input, " ")

		switch command[0] {
		case "exit":
			fmt.Printf("Exiting\n")
			break loop

		case "cd":
			if len(command) < 2 {
				fmt.Printf("invalid command: cd must have [path]")
				break
			}

			os.Chdir(command[1])
			dir, _ := os.Getwd()
			fmt.Printf("Current dir: %s", dir)

		case "pwd":
			dir, _ := os.Getwd()
			fmt.Printf("Current dir: %s", dir)

		case "echo":
			if len(command) < 2 {
				fmt.Printf("invalid command: echo must have [string]")
				break
			}

			fmt.Printf(strings.Join(command[1:], " "))

		case "kill":
			if len(command) < 2 {
				fmt.Printf("invalid command: kill must have [process name]")
				break
			}

			err := exec.Command("kill", command[1]).Run()
			if err != nil {
				fmt.Printf("failed to execute kill command: %v", err)
			}

		case "ps":
			procs, err := ps.Processes()
			if err != nil {
				fmt.Printf("failed to execute ps command: %v", err)
			}

			fmt.Println("current processes: ")
			for _, p := range procs {
				fmt.Println(p.Pid())
			}

		default:
			fmt.Printf("The command '%v' was not recognized\n", command[0])
		}

		fmt.Printf("\n# ")
	}

}
