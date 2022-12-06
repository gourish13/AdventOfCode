package main

import "fmt"

type Void struct{}
var member Void

func checkRepeat(str string) bool {
    var chars = make(map[rune]Void, 14)
    for _, c := range str {
        if _, ok := chars[c]; ok {
            return true
        }
        chars[c] = member
    }
    return false
}

func findUniqueSet(str string) int {
    for i := 0; i < len(str); i++ {
        if !checkRepeat(str[i : i + 14]) {
            return i + 14
        }
    }
    return -1
}

func main() {
    var line string
    fmt.Scanf("%s", &line)
    fmt.Printf("%d\n", findUniqueSet(line))
}
