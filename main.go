package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const GREETING_1_MESSAGE = "Введите математическое выражение (пример: 5 + 2, 10 / 5, VII + III) и нажмите Enter"
const GREETING_2_MESSAGE = "Для выхода из программы введите quit"
const INCORRECT_MATH_OPERATION_MESSAGE = "Введите корректное математическое выражение"
const ROMAN_NEGATIVE_EXCEPTION_MESSAGE = "В римской системе нет отрицательных чисел"
const PROGRAM_EXIT_MESSAGE = "Программа завершена"

var RomanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func calculation(first, second int, operation string) int {
	switch operation {
	case "+":
		return first + second
	case "-":
		return first - second
	case "*":
		return first * second
	case "/":
		return first / second
	default:
		return 0
	}
}

func arabicCalc(inputArray []string) int {
	first, _ := strconv.Atoi(inputArray[0])
	second, _ := strconv.Atoi(inputArray[1])
	operation := inputArray[2]
	return calculation(first, second, operation)
}

func romanCalc(inputArray []string) string {
	first := romanToInteger(inputArray[0])
	second := romanToInteger(inputArray[1])
	operation := inputArray[2]

	if first < 1 || first > 10 || second < 1 || second > 10 {
		panic(INCORRECT_MATH_OPERATION_MESSAGE)
	}

	result := calculation(first, second, operation)

	if result < 1 {
		panic(ROMAN_NEGATIVE_EXCEPTION_MESSAGE)
	}

	return integerToRoman(result)
}

func calc(input string) string {

	inputArray := split(input)

	if len(inputArray) == 3 {
		if arabicMath(inputArray[0]) {
			if arabicMath(inputArray[1]) {
				return strconv.Itoa(arabicCalc(inputArray))
			}
		} else {
			if romanMath(inputArray[1]) {
				return romanCalc(inputArray)
			}
		}
	}
	panic(INCORRECT_MATH_OPERATION_MESSAGE)
}

func split(text string) []string {
	pattern := "[\\-\\+\\*\\/]{1}"
	regex := regexp.MustCompile(pattern)
	result := regex.Split(text, -1)
	index := regex.FindStringIndex(text)
	return append(result, getSubstring(text, index))
}

func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

func arabicMath(input string) bool {

	value, err := strconv.Atoi(input)
	if err != nil {
		//panic(err)
		return false
	}
	if value >= 1 && value <= 10 {
		return true
	}
	return false
}

func romanMath(input string) bool {
	roman := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, n := range roman {
		if input == n {
			return true
		}
	}
	return false
}

func romanToInteger(s string) int {
	sum := 0
	greatest := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := RomanNumerals[rune(letter)]
		if num >= greatest {
			greatest = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}
	return roman.String()
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(GREETING_1_MESSAGE)
	fmt.Println(GREETING_2_MESSAGE)

	for {
		fmt.Println(GREETING_1_MESSAGE)

		scanner.Scan()
		input := scanner.Text()
		if len(input) >= 3 {
			result := calc(strings.ReplaceAll(input, " ", ""))
			fmt.Println(result)
		} else {
			panic(INCORRECT_MATH_OPERATION_MESSAGE)
		}

		if input == "quit" {
			break
		}
	}
	fmt.Println(PROGRAM_EXIT_MESSAGE)
}
