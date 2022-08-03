package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

// Необходимо реализовать свою собственную UNIX-шелл-утилиту с
// поддержкой ряда простейших команд:
// - cd <args> - смена директории (в качестве аргумента могут
// быть то-то и то)
// - pwd - показать путь до текущего каталога
// - echo <args> - вывод аргумента в STDOUT
// - kill <args> - "убить" процесс, переданный в качесте
// аргумента (пример: такой-то пример)
// - ps - выводит общую информацию по запущенным процессам в
// формате *такой-то формат*
// Так же требуется поддерживать функционал fork/exec-команд
// Дополнительно необходимо поддерживать конвейер на пайпах
// (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
// *Шелл — это обычная консольная программа, которая будучи
// запущенной, в интерактивном сеансе выводит некое приглашение
// в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись
// ввода, обрабатывает команду согласно своей логике
// и при необходимости выводит результат на экран. Интерактивный
// сеанс поддерживается до тех пор, пока не будет введена
// команда выхода (например \quit).

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		comandStdin(strings.Fields(scanner.Text()))
	}
}
func comandStdin(commands []string) {
	command := commands[0]
	args := commands[1:]
	
	switch command {
	case `\quit`:
		fmt.Println("exit")
		os.Exit(1)
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(dir)
	case "ps":
		res, err := ps.Processes()
		if err != nil {
			fmt.Println(err)
		}
		for _, process := range res {
			fmt.Printf("Name p: %v Pid: %v\n", process.Executable(), process.Pid())
		}
	case "cd":
		dir := strings.Join(args, "")
		fmt.Println(dir)
		os.Chdir(dir)
		currentDir, _ := os.Getwd()
		fmt.Println(currentDir)
	case "echo":
		fmt.Println(strings.Join(args, " "))
	case "kill":
		pid, err := strconv.Atoi(strings.Join(args, ""))
		if err != nil {
			fmt.Println(err)
			break
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Println(err)
			break
		}
		err = proc.Kill()
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("invalid or unsupported command")
	}
}
