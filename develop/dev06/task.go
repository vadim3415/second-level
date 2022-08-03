package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"os"
)

// Реализовать утилиту аналог консольной команды cut (man cut).
// Утилита должна принимать строки через STDIN, разбивать по
// разделителю (TAB) на колонки и выводить запрошенные.
// Реализовать поддержку утилитой следующих ключей:
// -f - "fields" - выбрать поля (колонки)
// -d - "delimiter" - использовать другой разделитель
// -s - "separated" - только строки с разделителем

func main() {
	var fields string
	var delimiter string
	var seperated bool

	flag.StringVar(&fields, "f", "", "set fields")
	flag.StringVar(&delimiter, "d", "\t", "set delimiter")
	flag.BoolVar(&seperated, "s", false, "set seperated")
	flag.Parse()

	a := os.Args[1:]
	// если запуск приложения без аргументов
	if len(a) == 0 {
		fmt.Println("-help\n", "go run main.go -f 1:3 -d ':' -s note.txt\n", "go run main.go -f 1:3 -d ':'")
		return
	}

	if filename := flag.Arg(0); filename != "" {
		openReadFile(filename, fields, delimiter, seperated)

	} else {
		readStdin(fields, delimiter, seperated)
	}
}

func readStdin(fields string, delimiter string, seperated bool) {
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	textInput := buf.Text()

	cutFile([]byte(textInput), fields, delimiter, seperated)
}

func openReadFile(fileName string, fields string, delimiter string, seperated bool) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
	}
	defer file.Close()

	readFile, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("error reading file", err)
	}

	cutFile(readFile, fields, delimiter, seperated)
}

func cutFile(strByte []byte, fields string, delimiter string, seperated bool) {
	var first, second int

	if len(fields) == 3 {
		first, second = intFields(fields)
	}
	newStr := ""
	newStr2 := ""
	// берем строку
	splitLine := strings.Split(string(strByte), "\n")

	for _, v := range splitLine {
		// делим строку на слова
		splitDelimiter := strings.Split(v, delimiter)

		// если в строке нет делителя и sep false, записываем строку без делителя
		if len(splitDelimiter) == 1 && !seperated {
			newStr = strings.Join(splitDelimiter[0:1], delimiter)
			newStr2 += newStr + "\n"

			// если строка без делителя и sep true, то пропускаем строку
		} else if len(splitDelimiter) == 1 && seperated {
			continue

			// если длина строка деленная на слова больше значений fields и флаг fields исаользован
		} else if len(splitDelimiter) > first && len(splitDelimiter) > second && len(fields) == 3 {
			newStr = strings.Join(splitDelimiter[first:second], delimiter)
			newStr2 += newStr + "\n"

			// если флаг fields не использован
		} else if len(fields) == 0 {
			newStr = strings.Join(splitDelimiter, delimiter)
			newStr2 += newStr + "\n"

			// если указан неверный диапозон fields
		} else {
			fmt.Println("wrong set fields")
			os.Exit(1)
		}
	}
	if len(newStr2) < 1 {
		fmt.Println("incorrect use or no data")
	}
	print(newStr2)
}

func intFields(fields string) (int, int) {
	sliceIndex := strings.Split(fields, ":")

	if len(sliceIndex) != 2 {
		fmt.Println("wrong set fields")
		os.Exit(1)
	}

	first, err := strconv.Atoi(sliceIndex[0])
	if err != nil {
		log.Fatal(err)
	}

	second, err := strconv.Atoi(sliceIndex[1])
	if err != nil {
		log.Fatal(err)
	}

	if first > second || first == second || first == 0 || first < 0 || second < 0 {
		fmt.Println("wrong set fields")
		os.Exit(1)
	}

	return first - 1, second - 1
}
