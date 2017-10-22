package fizzbuzz

import "strconv"

// Calculate -
func Calculate(numeric, divisor int, text string) string {
    if numeric > 0 && numeric % divisor == 0 {
        return text
    }
    return ""
}

// FizzBuzz - Just improving my golang skill
func FizzBuzz(numeric int) string {
    FizzBuzz := Calculate(numeric, 3, "fizz") + Calculate(numeric, 5, "buzz")

    if len(FizzBuzz) > 0 {
        return FizzBuzz
    }

    return strconv.Itoa(numeric)
}
