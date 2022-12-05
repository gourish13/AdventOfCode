package main

/*
parser.go parses the input stack.
stack.go contains stack implementation.
*/

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

var stack []Stack

func performMove(count, from, to int) {
    var i int
    var element string

    var tempStack Stack

    for ;i < count; i++ {
        element, _ = stack[from].pop()
        tempStack.push(element)
    }

    for element, ok := tempStack.pop(); ok; 
            element, ok = tempStack.pop() {
        stack[to].push(element)
    }
}

func main() {
    var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
    getLine(scanner)

    var move []string
    var count, from, to int

    for scanner.Scan() {
        move = strings.Fields(scanner.Text())

        count, _  = strconv.Atoi(move[1])
        from, _ = strconv.Atoi(move[3])
        to, _ = strconv.Atoi(move[5])

        performMove(count, from - 1, to - 1)
    }

    var tops string = ""
    for _, s := range stack {
        element, _ := s.peek()
        tops += element
    }
    fmt.Println(tops)
}

