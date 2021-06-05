package logicFunc

import (
	"math"
	"strconv"
	"strings"
)

// var (
// 	val int
// )

func FizzBuzz(n int) []string {
	var fizzbuzz []string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fizzbuzz = append(fizzbuzz, "FizzBuzz")
		} else if i%3 == 0 {
			fizzbuzz = append(fizzbuzz, "Fizz")
		} else if i%5 == 0 {
			fizzbuzz = append(fizzbuzz, "Buzz")
		} else {
			fizzbuzz = append(fizzbuzz, strconv.Itoa(i))
		}
	}

	return fizzbuzz
}

func LeapYear(startYear, endYear int) []int {
	var leapData = []int{}

	for {
		startYear++
		if startYear%4 == 0 {
			for startYear <= endYear {
				leapData = append(leapData, startYear)
				startYear += 4
			}
			break
		}
	}

	return leapData
}

func NearestFibonacci(num []int) float64 {
	var total int
	var nearFIrst int
	var nearLast int
	var output float64

	for _, n := range num {
		total += n
	}

	// fmt.Println("total", total)
	if total == 0 {
		return 0
	}

	first := 0
	second := 1

	third := first + second

	for third <= total {
		first = second
		second = third
		third = first + second
		// fmt.Println(first, second, third)
	}

	if third-total >= second-total {
		nearFIrst = second
		nearLast = third
	}

	if math.Abs(float64(nearLast-total)) <= math.Abs(float64(nearFIrst-total)) {
		output = math.Abs(float64(nearLast - total))
	} else {
		output = math.Abs(float64(nearFIrst - total))
	}

	return output
}

func Palindrome(val string) string {
	var changeLower = strings.ToLower(val)

	var reverseInput string

	for i := len(changeLower); i > 0; i-- {
		reverseInput += string(changeLower[i-1])
	}

	if reverseInput == changeLower {
		return "palindrome"
	}

	return "not palindrome"
}

func ReverseWord(word string) string {
	var splitWord = strings.Split(word, " ")

	for idx, sw := range splitWord {
		var reverseInput string

		for i := len(sw); i > 0; i-- {
			if strings.ToUpper(string(sw[len(sw)-i])) == string(sw[len(sw)-i]) {
				reverseInput += strings.ToUpper(string(sw[i-1]))
				continue
			}

			reverseInput += strings.ToLower(string(sw[i-1]))
		}

		splitWord[idx] = reverseInput
	}

	var reverse = strings.Join(splitWord, " ")
	return reverse
}
