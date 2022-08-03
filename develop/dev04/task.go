package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func findAnagrams(str []string) map[string][]string {
	result := make(map[string][]string)
	unicKey := make(map[string]struct{})
	outputResult := make(map[string][]string)

	for _, v := range str {

		toLovelKey := strings.ToLower(v)

		anagram := reverseWord(toLovelKey)

		// если слово уже есть в ключе карты uniKey, то пропускаем
		// иначе записываем по ключу-анаграмме в карту result
		if _, ok := unicKey[toLovelKey]; !ok {
			result[anagram] = append(result[anagram], toLovelKey)
			unicKey[toLovelKey] = struct{}{}
		}
	}

	for _, v := range result {
		if len(v) > 1 {
			outputResult[v[0]] = v
		}
	}
	return outputResult
}

func reverseWord(str string) string {
	wordByLetter := strings.Split(str, "")
	sort.Strings(wordByLetter)
	result := strings.Join(wordByLetter, "")

	return result
}

func main() {

	str := []string{
		"кино",
		"кони",
		"Пятак",
		"пятка",
		"слиток",
		"листок",
		"столик",
		"листок",
		"Порт",
		"рог",
		"тяпка",
	}

	anagrams := findAnagrams(str)

	for k, v := range anagrams {
		fmt.Printf("k: %s, v : %q\n", k, v)
	}

}
