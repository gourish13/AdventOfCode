package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func main() {
    var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
    var cycles = [6]int{20, 60, 100, 140, 180, 220}
    var cycle, X = 0, 1
    var index, strength int

    var line []string
    for scanner.Scan() && index < len(cycles) {
        line = strings.Fields(scanner.Text())
        var V int

        if len(line) == 1 {
            cycle++
        } else {
            V, _ = strconv.Atoi(line[1])
            X += V
            cycle += 2
        }

        if (cycle == cycles[index] && len(line) == 2) || cycle - cycles[index] == 1 {
            strength += cycles[index] * (X - V)
            index++
        } else if cycle == cycles[index] {
            strength += cycles[index] * X
            index++
        }

    }
    fmt.Println(strength)
}
