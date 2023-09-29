package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func Division(a, b *big.Int) string {
	if b.IsInt64() && b.Int64() == 0 {
		return "нельзя делить на ноль"
	}

	div, mod := a.QuoRem(a, b, big.NewInt(0))

	return fmt.Sprintf(
		"Частное:\n%s\nОстаток:\n%s\n",
		div.String(),
		mod.String(),
	)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	fmt.Println("Введите первое большое число")
	scanner.Scan()
	firstNumStr := scanner.Text()

	firstNum, b := big.NewInt(0).SetString(firstNumStr, 10)
	if !b {
		fmt.Println("Это не число")
		return
	}

	fmt.Println("Введите второе большое число")
	scanner.Scan()
	secondNumStr := scanner.Text()

	secondNum, b := big.NewInt(0).SetString(secondNumStr, 10)
	if !b {
		fmt.Println("Это не число")
		return
	}

	fmt.Println("Введите необходимое действие: +, -, * или /")
	scanner.Scan()
	action := scanner.Text()

	result := ""
	switch action {
	case "+":
		result = firstNum.Add(firstNum, secondNum).String()

	case "-":
		result = firstNum.Sub(firstNum, secondNum).String()

	case "*":
		result = firstNum.Mul(firstNum, secondNum).String()

	case "/":
		result = Division(firstNum, secondNum)

	default:
		fmt.Println("Это не действие")
		return
	}

	fmt.Println("Результат:")
	fmt.Println(result)
}
