package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

func inRange(num, lwr, upr int) bool {
    return num >= lwr && num <= upr
} 

func isOverlapping(s1, s2 []string) bool {
    l1, _ := strconv.Atoi(s1[0])
    u1, _ := strconv.Atoi(s1[1])
    l2, _ := strconv.Atoi(s2[0])
    u2, _ := strconv.Atoi(s2[1])

    return ( inRange(l1, l2, u2) ||
             inRange(u1, l2, u2) ||
             inRange(l2, l1, u1) ||
             inRange(u2, l1, u1)  )

}

func main() {
    var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

    var result int

    for scanner.Scan() {
        var line string = scanner.Text()
        var sections []string = strings.Split(line, ",")
        var sec1 []string = strings.Split(sections[0], "-")
        var sec2 []string = strings.Split(sections[1], "-")

        if (isOverlapping(sec1, sec2)) {
            result++;
        }
    }
    fmt.Println(result)
}
