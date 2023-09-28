package main

import (
	"errors"
	"fmt"
	"strconv"
)

func SetIbit(number *int64, position int, bit int) error {

	// будем менять только биты, которые есть, с 0 по 63
	if position > 63 || position < 0 {
		return errors.New("position must be between 0 and 63")
	}

	if bit == 0 {
		// Установка нуля
		//
		// 1) число "-1" в битовом представлении - это 64 единицы
		// 2) число "(1 << position)" в битовом представлении -
		//    это единица в бите под номером position и остальные нули
		//    "100..."
		// 3) побитовое исключающее ИЛИ (^) этих двух чисел даст число numberToAnd
		//    педставленное всеми единицами и нулём в бите под номером position
		numberToAnd := -1 ^ (1 << position)

		// 4) побитовое И с числом *number (&=) поставит 0 в бит под номером position и
		//    оставит неизменными все остальные биты
		*number &= int64(numberToAnd)
	} else {
		// Установка нуля
		//
		// 1) число "(1 << position)" в битовом представлении -
		//    это единица в бите под номером position и остальные нули
		//    "100..."
		// 2) побитовое ИЛИ с числом *number (|=) поставит 1 в бит под номером position и
		//    оставит неизменными все остальные биты
		*number |= (1 << position)
	}

	return nil
}

func main() {
	number := int64(0b101000)
	fmt.Printf("number: %v\n", strconv.FormatInt(number, 2))

	SetIbit(&number, 4, 1)
	SetIbit(&number, 3, 0)

	fmt.Printf("number: %v\n", strconv.FormatInt(number, 2))

	fmt.Printf("number == 0b110000: %v\n", (number == int64(0b110000)))
}
