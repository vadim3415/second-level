package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// Создать программу печатающую точное время с использованием
// NTP -библиотеки. Инициализировать как go module. Использовать
// библиотеку github.com/beevik/ntp. Написать программу
// печатающую текущее время / точное время с использованием этой
// библиотеки.
// Требования:
// 1. Программа должна быть оформлена как go module
// 2. Программа должна корректно обрабатывать ошибки
// библиотеки: выводить их в STDERR и возвращать ненулевой
// код выхода в OS

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = response.Validate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	time := time.Now().Add(response.ClockOffset).Format("15:04:05")
	fmt.Println(time)
}
