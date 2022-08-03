package main

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

// Создать Go-функцию, осуществляющую примитивную распаковку
// строки, содержащую повторяющиеся символы/руны, например:
// ●
//  "a4bc2d5e" => "aaaabccddddde"
// ●
//  "abcd" => "abcd"
// ●
//  "45" => "" (некорректная строка)
// ●
//  "" => ""

// Дополнительно
// Реализовать поддержку escape-последовательностей.
// Например:
// ●
//  qwe\4\5 => qwe45 (*)
// ●
//  qwe\45 => qwe44444 (*)
// ●
//  qwe\\5 => qwe\\\\\ (*)
// В случае если была передана некорректная строка, функция
// должна возвращать ошибку. Написать unit-тесты.

func main() {
	str := os.Args[1]

	fmt.Println(stringUnpacking(str))
}

func stringUnpacking(str string) (string, error) {
	runeStr := []rune(str)
	lenStr := len(runeStr)
	result := []rune{}
	count := 0

	for i := 0; i < lenStr; i++ {
		if string(runeStr[i]) == `\` {
			i++
			result = append(result, runeStr[i])
		} else if unicode.IsDigit(runeStr[i]) {
			if i < lenStr-1 && unicode.IsDigit(runeStr[i+1]) {
				return "", errors.New("некорректная строка")

			} else {
				count = (int(runeStr[i] - '0'))
				//count, _ = strconv.Atoi(string(runeStr[i]))
				for j := 0; j < count-1; j++ {
					result = append(result, runeStr[i-1])
				}
			}
		} else {
			result = append(result, runeStr[i])
		}
	}

	return string(result), nil
}

// func stringUnpacking2(str string) {
// 	runeStr := []rune(str)
// 	var res []rune
// 	count := 0

// 	for i := 0; i < len(runeStr); i++ {
// 		if i == len(runeStr)-1 {
// 			res = append(res, runeStr[i])
// 			break
// 		}
// 		chek := (runeStr[i] >= '0') && (runeStr[i] <= '9')
// 		chek2 := (runeStr[i+1] >= '0') && (runeStr[i+1] <= '9')
// 		// если рядом стоят два числа, прекращаем работу
// 		if chek && chek2 {
// 			return
// 		}
// 		// если это буква и за ней стоит число
// 		if !chek && chek2 {
// 			count = count*10 + (int(runeStr[i+1] - '0'))findAnagrams
// 			for j := 0; j < count; j++ {
// 				res = append(res, runeStr[i])
// 			}
// 			//fmt.Println("1", string(res))
// 			count = 0
// 		} else if !chek {
// 			res = append(res, runeStr[i])
// 			//fmt.Println("2", string(res))
// 		}
// 	}
// 	fmt.Println(string(res))
// }
