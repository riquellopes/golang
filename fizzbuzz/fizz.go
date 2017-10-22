package fizzbuzz

import "strconv"

// Fizz -
func Fizz(numeric int) bool{
    return numeric % 3 == 0
}

// Buzz -
func Buzz(numeric int) bool {
    return numeric % 5 == 0
}

// FizzBuzz - Just improving my golang skill
func FizzBuzz(numeric int) string {
    if Fizz(numeric) && Buzz(numeric) {
        return "fizzbuzz"
    } else if Fizz(numeric) {
        return "fizz"
    } else if Buzz(numeric) {
        return "buzz"
    }
    return strconv.Itoa(numeric)
}
