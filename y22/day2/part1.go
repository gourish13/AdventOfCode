package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func outcomeScore(player, you string) int {
    const (
        Loss = iota * 3
        Draw = iota * 3
        Win = iota * 3
    )

    if ((player == "A" && you == "X") ||
            (player == "B" && you == "Y") ||
            ( player == "C" && you == "Z")) {
        return Draw
    } else if ((player == "A" && you == "Y") ||
            (player == "B" && you == "Z") ||
            ( player == "C" && you == "X")) {
        return Win
    } else {
        return Loss
    }
}


func main() {
    var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
    var totalScore int

    var winPoint = map[string]int{ "X": 1, "Y": 2, "Z": 3 }

    for scanner.Scan() {
        var move []string = strings.Fields(scanner.Text())
        totalScore += outcomeScore(move[0], move[1]) + winPoint[move[1]]
    }

    fmt.Println(totalScore)
}
