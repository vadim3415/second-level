package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"

	"strings"
)

// Утилита sort
// Отсортировать строки в файле по аналогии с консольной
// утилитой sort (man sort — смотрим описание и основные
// параметры): на входе подается файл из несортированными
// строками, на выходе — файл с отсортированными.
// Реализовать поддержку утилитой следующих ключей:
// -k — указание колонки для сортировки (слова в строке могут
// выступать в качестве колонок, по умолчанию разделитель —
// пробел)
// -n — сортировать по числовому значению
// -r — сортировать в обратном порядке
// -u — не выводить повторяющиеся строки

func main() {
	var k int
	var r bool
	var u bool
	var n bool
	flag.IntVar(&k, "k", 1, "set column")
	flag.BoolVar(&r, "r", false, "set reverse")
	flag.BoolVar(&u, "u", false, "set no duplicate lines")
	flag.BoolVar(&n, "n", false, "set delete space")
	flag.Parse()

	a := os.Args[1:]
	if len(a) == 0 {
		fmt.Println("-help\n go run main.go -k 1 -r note.txt")
		return
	}

	if fileName := flag.Arg(0); fileName != "" {
		res := sorted(fileName, k, r, u, n)

		for _, v := range res {
			fmt.Println(v)
		}
	}
}

func sorted(fileName string, k int, r bool, u bool, n bool) []string {
	file := openReadFile(fileName)
	sMap := make(map[string][]string)

	splitLine := strings.Split(string(file), "\n")

	if u {
		splitLine = checkDuplicate(splitLine)
	}

	column := []string{}

	for i, v := range splitLine {
		var checkLine int

		if n {
			checkLine = strings.Count(deletSpace(splitLine[i]), " ")
			v = deletSpace(v)
		} else {
			checkLine = strings.Count(splitLine[i], " ")
		}

		if k-1 > checkLine || k-1 < 0 {
			continue
		}
		word := strings.Split(v, " ")[k-1]
		wordToLover := strings.ToLower(word)

		sMap[wordToLover] = append(sMap[wordToLover], v)
		column = append(column, wordToLover)
	}

	checkColumns := checkDuplicate(column)

	if len(checkColumns) < 1 {
		fmt.Println("no columns to sort")
		os.Exit(1)
	}
	if r {
		sort.Sort(sort.Reverse(sort.StringSlice(checkColumns)))
	} else {
		sort.Strings(checkColumns)
	}

	result := []string{}
	for _, v := range checkColumns {
		if vel, ok := sMap[v]; ok {
			result = append(result, vel...)
		}
	}

	return result
}

func openReadFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	readFile, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return readFile
}

func checkDuplicate(column []string) []string {
	m := make(map[string]struct{})
	var output []string

	for _, v := range column {
		m[v] = struct{}{}
	}
	for key := range m {
		output = append(output, key)
	}
	return output
}

func deletSpace(str string) string {
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(str, " ")
	result := strings.TrimSpace(replace)

	return result
}