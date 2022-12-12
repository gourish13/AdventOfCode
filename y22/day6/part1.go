package main

import "fmt"

type Void struct{}
var member Void

func checkRepeat(str string) bool {
    var chars = make(map[rune]Void, 4)
    for _, c := range str {
        if _, ok := chars[c]; ok {
            return true
        }
        chars[c] = member
    }
    return false
}

func findUniqueSet(str string) int {
    for i := 3; i < len(str); i++ {
        if !checkRepeat(str[i - 3 : i + 1]) {
            return i + 1
        }
    }
    return -1
}

func main() {
    var line string
    fmt.Scanf("%s", &line)
    fmt.Printf("%d\n", findUniqueSet(line))
}