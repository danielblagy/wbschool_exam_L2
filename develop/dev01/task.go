package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	printExactCurrentTime()
}

func printExactCurrentTime() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal("Error when querying\n", err)
	}
	if err := response.Validate(); err != nil {
		log.Fatal("response data is suitable for synchronization purposes\n", err)
	}
	time := time.Now().Add(response.ClockOffset)
	fmt.Println(time)
}
