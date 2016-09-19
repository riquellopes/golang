package pingpong

import "strconv"

func ping(number int) string {
    if doubleDivisor(number){
        return "PingPong"
    } else if divisible(number, 3) {
        return "Ping"
    } else if divisible(number, 5){
        return "Pong"
    }

    return strconv.Itoa(number)
}

func divisible(number, divider int) bool {
    if (number % divider) == 0 {
        return true
    }
    return false
}

func doubleDivisor(number int) bool{
    return divisible(number, 3) && divisible(number, 5)
}
