package main

import (
    "fmt"
    "strings"
    "strconv"
)

type View struct {
    left int
    right int
    up int
    down int
}

func (view View) getScenicView() int {
    return view.left * view.right * view.up * view.down
}

func mapToInt8s(list []string) []int8 {
    var newlist = make([]int8, len(list))

    for i, v := range list {
        var val, _ = strconv.Atoi(v)
        newlist[i] = int8(val)
    }
    return newlist
}

func findScenicView(forest [][]int8) [][]View {
    var l = len(forest)
    var views = make([][]View, l)
    for i := 0; i < l; i++ {
        views[i] = make([]View, l)
    }

    for i := 0; i < l; i++ {
        for j := 0; j < l; j++ {
            // Right
            for k := j + 1; k < l; k++ {
                views[i][j].right++
                if forest[i][j] <= forest[i][k] {
                    break
                }
            } 

            // Left
            for k := j - 1; k >= 0; k-- {
                views[i][j].left++
                if forest[i][j] <= forest[i][k] {
                    break
                }
            }

            // Down
            for k := i + 1; k < l; k++ {
                views[i][j].down++
                if forest[i][j] <= forest[k][j] {
                    break
                }
            }

            // Up
            for k := i - 1; k >= 0; k-- {
                views[i][j].up++
                if forest[i][j] <= forest[k][j] {
                    break
                }
            }
            
        }
    }

    return views
}

func maxScenicView(views [][]View) int {
    var max int
    for _, row := range views {
        for _, view := range row {
            if scenicView := view.getScenicView(); max < scenicView {
                max = scenicView
            }
        }
    }
    return max
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

    var views = findScenicView(forest)
    var maxView = maxScenicView(views)

    fmt.Println(maxView)
}