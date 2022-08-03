package main

import "fmt"

/*
Команда — это поведенческий паттерн, позволяющий заворачивать запросы или простые операции в отдельные объекты.
Это позволяет откладывать выполнение команд, выстраивать их в очереди, а также хранить историю и делать отмену.
Плюсы:
- Убирается прямая связь между отправителями и исполнителями запросов
- Позволяет удобно реализовывать различные операции: отмена и повтор запросов,
отложенный запуск запросов, выстраивание очереди запросов.
Минусы:
- Усложняет код из-за необходимости реализации дополнительных классов
*/

type command interface {
	execute()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("TV on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("TV off")
}

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

func main() {
	tv := &tv{}

	onCommand := &onCommand{device: tv}
	offCommand := &offCommand{device: tv}

	onButton := &button{command: onCommand}
	offButton := &button{command: offCommand}

	onButton.press()
	offButton.press()
}
