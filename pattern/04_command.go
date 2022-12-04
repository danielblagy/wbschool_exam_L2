package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// receiver interface
type device interface {
	on()
	off()
}

// concrete receiver
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// command interface
type command interface {
	execute()
}

// concrete command
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// concrete command
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// invoker
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// invoker
type voiceInput struct {
	command command
}

func (v *voiceInput) say() {
	v.command.execute()
}

func CommandExample() {
	// receiver
	tv := &tv{}

	// command
	onCommand := &onCommand{
		device: tv,
	}
	// command
	offCommand := &offCommand{
		device: tv,
	}

	// invoker
	onButton := &button{
		command: onCommand,
	}
	// invoker
	offButton := &button{
		command: offCommand,
	}

	onVoiceInput := &voiceInput{
		command: onCommand,
	}
	// invoker
	offVoiceInput := &voiceInput{
		command: offCommand,
	}

	// client code

	onButton.press()
	offButton.press()
	onVoiceInput.say()
	offVoiceInput.say()
}
