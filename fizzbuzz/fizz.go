package fizzbuzz

import "strconv"

// Fizz -
func Fizz(numeric int) string {
    if numeric % 3 == 0 {
        return "fizz"
    }
    return ""
}

// Buzz -
func Buzz(numeric int) string {
    if numeric % 5 == 0 {
        return "buzz"
    }
    return ""
}

// FizzBuzz - Just improving my golang skill
func FizzBuzz(numeric int) string {
    FizzBuzz := Fizz(numeric) + Buzz(numeric)

    if len(FizzBuzz) > 0 {
        return FizzBuzz
    }

    return strconv.Itoa(numeric)
}
