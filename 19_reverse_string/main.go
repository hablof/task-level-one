package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func reverseUnicodeString(s string) string {

	// Поскольку строка может содержать символы Юникод,
	// было бы ошибкой преобразовывать строку в слайс байт
	stringRunes := []rune(s)

	// В версии GO 1.21 добавили небольшой, но удобный пакет slices
	// благодоря нему разместим руны в слайсе в обратном порядке
	slices.Reverse(stringRunes)

	// Вернём строку полученную из массива рун
	return string(stringRunes)
}

func main() {
	fmt.Println("Введите строку")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	userString := scanner.Text()

	fmt.Printf("Развёрнутая строка: %s\n", reverseUnicodeString(userString))
}
