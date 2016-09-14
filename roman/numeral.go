package roman

import "strings"

func ArabicToRoman(n int) string {

    // romanNumerals := map[int] string {
    //     1: "I",
    //     5: "V",
    //     10: "X",
    //     50: "L",
    //     100: "C",
    //     500: "D",
    //     1000: "M",
    // }

    if n == 5 {
        return "V"
    } else if n == 4 {
        return "IV"
    }
    return strings.Repeat("I", n)
}
