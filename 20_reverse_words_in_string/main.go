package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func reverseWords(s string) string {

	// Разделим входящую строку на массив слов
	// разделителем будем считать пробел
	words := strings.Split(s, " ")

	// в версии GO 1.21 добавили небольшой, но удобный пакет slices
	// благодоря нему разместим слова в слайсе в обратном порядке
	slices.Reverse(words)

	// объединим слова в новую строку
	return strings.Join(words, " ")
}

func main() {
	fmt.Println("Введите строку")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	userString := scanner.Text()

	fmt.Printf("Развёрнутая строка: %s\n", reverseWords(userString))
}
