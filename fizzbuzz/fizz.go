package fizzbuzz

import "strconv"

// FizzBuzz - Just improving my golang skill
func FizzBuzz(numeric int) string {
    if numeric % 3 == 0 && numeric % 5 == 0 {
        return "fizzbuzz"
    } else if numeric % 3 == 0 {
        return "fizz"
    } else if numeric % 5 == 0 {
        return "buzz"
    }
    return strconv.Itoa(numeric)
}
