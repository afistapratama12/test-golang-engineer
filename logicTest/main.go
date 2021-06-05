package main

import (
	"fmt"
	"logicTest/logicFunc"
)

func main() {

	// logic 1
	fmt.Println(logicFunc.Palindrome("Radar"))
	fmt.Println(logicFunc.Palindrome("Malam"))
	fmt.Println(logicFunc.Palindrome("Kasur ini rusak"))
	fmt.Println(logicFunc.Palindrome("Ibu Ratna antar ubi"))
	fmt.Println(logicFunc.Palindrome("Malas"))
	fmt.Println(logicFunc.Palindrome("Makan nasi goreng"))
	fmt.Println(logicFunc.Palindrome("Balonku ada lima"))

	//logic2
	fmt.Println(logicFunc.LeapYear(1900, 2020))

	//logic3
	fmt.Println(logicFunc.ReverseWord("I am A Great human"))

	//logic4
	var num = []int{15, 1, 3}
	fmt.Println(logicFunc.NearestFibonacci(num))

	//logic5
	fmt.Println(logicFunc.FizzBuzz(3))
	fmt.Println(logicFunc.FizzBuzz(5))
	fmt.Println(logicFunc.FizzBuzz(15))
}
