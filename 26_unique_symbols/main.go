package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isUinqueSymbolsCaseInsensitive(s string) bool {

	storage := map[rune]struct{}{}

	// Будем записывать в мапу ключи-руны из строки, приведённой к нижнему регистру
	for _, r := range strings.ToLower(s) {

		// если по текущему ключу уже имеется запись, значит символ не уникален, возвращаем false
		if _, ok := storage[r]; ok {
			return false
		}
		storage[r] = struct{}{}
	}

	// если смогли пройти все символы в строке и не наткнулись на дубликаты
	return true
}

func main() {
	fmt.Println("Введите строку для проверки её на уникальность символов")
	fmt.Println("Регистр не будет учитываться")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	userString := scanner.Text()

	if isUinqueSymbolsCaseInsensitive(userString) {
		fmt.Println("✅ Символы в строке уникальны!")
		return
	}

	fmt.Println("❌ Символы в строке не уникальны :(")
}
