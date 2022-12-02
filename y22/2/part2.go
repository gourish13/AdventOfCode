package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

type Move struct {
    win string
    draw string
    loss string
}

func main() {
    var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
    var totalScore int

    var winPoint = map[string]int{ "A": 1, "B": 2, "C": 3 }
    var outcomeScore = map[string]int{ "X": 0, "Y": 3, "Z": 6 }

    var moves = map[string]Move{
        "A": {"B", "A", "C"}, 
        "B": {"C", "B", "A"}, 
        "C": {"A", "C", "B"},
    }

    var expectedMove string

    for scanner.Scan() {
        var move []string = strings.Fields(scanner.Text())

        switch move[1] {
            case "X":
                expectedMove = moves[move[0]].loss
            case "Y":
                expectedMove = moves[move[0]].draw
            default:
                expectedMove = moves[move[0]].win
        }

        totalScore += outcomeScore[move[1]] + winPoint[expectedMove]
    }

    fmt.Println(totalScore)
}
