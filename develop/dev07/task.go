package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		if len(channels) == 0 {
			return nil
		}
		if len(channels) == 1 {
			return channels[0]
		}

		signal := make(chan interface{})
		go func() {
			defer close(signal)
			<-channels[0]
			// call or recursively
			<-or(channels[1:]...)

		}()
		return signal
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(20*time.Millisecond),
		sig(3000*time.Millisecond),
		sig(1*time.Second),
		sig(15*time.Millisecond),
		sig(85*time.Millisecond),
	)

	fmt.Printf("done after %v", time.Since(start))
}
