package main

import (
	"bufio"
	"fmt"
	"os"
)

type void struct{}
var member void

func fillSet(s string, set map[rune]void) {
    for _, c := range s {
        if _, ok := set[c]; !ok {
            set[c] = member
        }
    }
}

func getRunePriority(c rune) rune {
    if c < 97 {
        return c - 64 + 26
    } else {
        return c - 96
    }
}

func getCommonPriority(comp1, comp2 string) int {
    set := make(map[rune]void)
    fillSet(comp1, set)

    var pri rune
    
    for _, c := range comp2 {
        if _, ok := set[c]; ok {
            pri += getRunePriority(c)
            break
        }
    }
    return int(pri)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var pri int
    
    for scanner.Scan() {
        items := scanner.Text()
        mid := len(items) / 2
        comp1, comp2 := items[:mid], items[mid:]
        pri += getCommonPriority(comp1, comp2)
    }

    fmt.Println(pri)
}
