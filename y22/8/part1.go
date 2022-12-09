package main

import (
    "fmt"
    "strings"
    "strconv"
)


func mapToInt8s(list []string) []int8 {
    var newlist = make([]int8, len(list))

    for i, v := range list {
        var val, _ = strconv.Atoi(v)
        newlist[i] = int8(val)
    }
    return newlist
}

func buildSeenTreeMap(forest [][]int8) [][]bool {
    var l = len(forest)
    var seen = make([][]bool, l)
    for i := 0; i < l; i++ {
        seen[i] = make([]bool, l)
    }

    // Prepopulate seen with defaults for edges
    for i := 0; i < l; i++ {
        seen[0][i] = true
        seen[i][0] = true
        seen[l-1][i] = true
        seen[i][l-1] = true
    }

    // Goiing forwards in grid
    for i := 1; i < l-1; i++ {
        var hrow, hcol = forest[0][i], forest[i][0]

        for j := 1; j < l-1; j++ {
            // Columnwise
            if hcol < forest[i][j] {
                seen[i][j] = true
                hcol = forest[i][j]
            } 

            // Rowwise
            if hrow < forest[j][i] {
                seen[j][i] = true
                hrow = forest[j][i]
            } 
        }
    }

    // Goiing backwards in grid 
    for i := l-2; i > 0; i-- {
        var hrow, hcol = forest[l-1][i], forest[i][l-1]

        for j := l-2; j > 0; j-- {
            // Columnwise
            if hcol < forest[i][j] {
                seen[i][j] = true
                hcol = forest[i][j]
            } 

            // Rowwise
            if hrow < forest[j][i] {
                seen[j][i] = true
                hrow = forest[j][i]
            } 
        }
    }

    return seen
}

func countVisibleTrees(seen [][]bool) int {
    var count int
    for _, row := range seen {
        for _, visible := range row {
            if visible {
                count++
            }
        }
    }
    return count
}

func main() {
    var line string
    var forest = [][]int8{}
    for {
        if _, err := fmt.Scanf("%s", &line); err != nil {
            break
        }
        forest = append(forest, mapToInt8s(strings.Split(line, "")))
    }

    var seen = buildSeenTreeMap(forest)
    var count = countVisibleTrees(seen)

    fmt.Println(count)
}
