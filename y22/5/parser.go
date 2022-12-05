package main

import (
	"bufio"
	"regexp"
	"strings"
)

func getLine(scanner *bufio.Scanner) {
    scanner.Scan()
    var line string = scanner.Text()

    if line == "" {
        return
    }
    getLine(scanner)

    if strings.HasPrefix(line, " 1") {
        var num int = len(strings.Fields(line[1:]))
        stack = make([]Stack, num)
        return
    }

    parseToStack(line)
}

func parseToStack(line string) {
    var re *regexp.Regexp = regexp.MustCompile(`\[([A-Z])\]`)
    for _, loc := range re.FindAllStringIndex(line, -1) {
        i := loc[0] / 4
        stack[i].push(line[ loc[0] + 1 : loc[1] - 1 ])
    }
}
