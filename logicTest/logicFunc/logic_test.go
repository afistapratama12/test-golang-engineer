package logicFunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	testTable := []struct {
		name   string
		input  int
		expect []string
	}{
		{"test fizzbuzz with input 3", 3, []string{"1", "2", "Fizz"}},
		{"test fizzbuzz with input 5", 5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{"test fizzbuzz with input 15", 15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			test := FizzBuzz(tc.input)
			assert.Equal(t, tc.expect, test)
		})
	}
}

func TestLeapYear(t *testing.T) {
	testTable := []struct {
		name           string
		inputA, inputB int
		expect         []int
	}{
		{"test leapyear with input 1900, 2020", 1900, 2020, []int{1904, 1908, 1912, 1916, 1920, 1924, 1928, 1932, 1936, 1940, 1944, 1948, 1952, 1956, 1960, 1964, 1968, 1972, 1976, 1980, 1984, 1988, 1992, 1996, 2000, 2004, 2008, 2012, 2016, 2020}},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			test := LeapYear(tc.inputA, tc.inputB)
			assert.Equal(t, tc.expect, test)
		})
	}
}

func TestPalindrome(t *testing.T) {
	testTable := []struct {
		name   string
		input  string
		expect string
	}{
		{"test palindorme with input Radar", "Radar", "palindrome"},
		{"test palindorme with input Malam", "Malam", "palindrome"},
		{"test palindorme with input Kasur ini rusak", "Kasur ini rusak", "palindrome"},
		{"test palindorme with input Ibu Ratna antar ubi", "Ibu Ratna antar ubi", "palindrome"},
		{"test palindorme with input Malas ", "Malas", "not palindrome"},
		{"test palindorme with input Makan nasi goreng", "Makan nasi goreng", "not palindrome"},
		{"test palindorme with input Balonku ada lima", "Balonku ada lima", "not palindrome"},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			test := Palindrome(tc.input)
			assert.Equal(t, tc.expect, test)
		})
	}
}

func TestReverseWord(t *testing.T) {
	testTable := []struct {
		name   string
		input  string
		expect string
	}{
		{"test ReverseWord with input I am A Great human", "I am A Great human", "I ma A Taerg namuh"},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			test := ReverseWord(tc.input)
			assert.Equal(t, tc.expect, test)
		})
	}
}

func TestNearestFibonacci(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		expect float64
	}{
		{"test NearestFibonacci with input [15,1,3]", []int{15, 1, 3}, 2},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			test := NearestFibonacci(tc.input)
			assert.Equal(t, tc.expect, test)
		})
	}
}
