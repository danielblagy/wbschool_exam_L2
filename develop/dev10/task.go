package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"

	"github.com/danielblagy/hurlean"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

// using https://github.com/danielblagy/hurlean

type MyClientFunctionalityProvider struct {
	scanner          *bufio.Scanner
	disconnectSignal chan os.Signal
}

func (fp MyClientFunctionalityProvider) OnServerMessage(clientInstance *hurlean.ClientInstance, message hurlean.Message) {

	if message.Type == "message" {
		fmt.Printf("%v\n\n", message.Body)
	}
}

func (fp MyClientFunctionalityProvider) OnClientInit(clientInstance *hurlean.ClientInstance) {
	signal.Notify(fp.disconnectSignal, os.Interrupt, os.Kill) // Ctrl + D functionality
	fmt.Printf("Welcome to the client!\n\n")
}

func (fp MyClientFunctionalityProvider) OnClientUpdate(clientInstance *hurlean.ClientInstance) {
	select {
	case <-fp.disconnectSignal:
		clientInstance.Disconnect()
	default:
	}

	if fp.scanner.Scan() {
		input := fp.scanner.Text()

		message := hurlean.Message{
			Type: "message",
			Body: input,
		}
		clientInstance.Send(message)
	}
}

func main() {
	// set the app-specific client's state
	var myClientFunctionalityProvider MyClientFunctionalityProvider = MyClientFunctionalityProvider{
		scanner:          bufio.NewScanner(os.Stdin),
		disconnectSignal: make(chan os.Signal),
	}

	if err := hurlean.ConnectToServer("localhost", "8080", myClientFunctionalityProvider); err != nil {
		fmt.Println(err)
	}
}
