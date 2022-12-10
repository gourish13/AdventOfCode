package main

import (
	"fmt"
	"strconv"
)

type void struct{}
type Point struct { x, y int }

func (point *Point) toString() string {
    return strconv.Itoa(point.x) + "," + strconv.Itoa(point.y)
}

func (point *Point) moveBy(p Point) {
    point.x += p.x
    point.y += p.y
}


func abs(number int) int {
    if number < 0 {
        return -number
    }
    return number
}

func gap(H, T *Point) (int, int) {
    return H.x - T.x , H.y - T.y
}

func moveWhere(H, T *Point) Point {
    var x, y = gap(H, T)

    if abs(x) <= 1 && abs(y) <= 1 {
        return Point{}
    }
    if abs(x) == 2 && abs(y) != 2 {
        return Point{x: x/2, y: y}
    } else if abs(y) == 2 && abs(x) != 2{
        return Point{x: x, y: y/2}
    } else {
        return Point{x: x/2, y: y/2}
    }
}

func main() {
    var member void
    var set = make(map[string]void)

    var to = map[string]Point {
        "L": {-1, 0},
        "R": {1, 0},
        "U": {0, 1},
        "D": {0, -1},
    }

    const knotsCount = 10
    var knots = make([]Point, knotsCount)
    var H *Point = &knots[0]
    var T *Point = &knots[knotsCount - 1]
    set[T.toString()] = member

    var motion string
    var times int
    for {
        if _, err := fmt.Scanf("%s %d", &motion, &times); err != nil {
            break
        }

        for i := 0; i < times; i++ {
            H.moveBy(to[motion])
            for j := 1; j < knotsCount; j++ {
                knots[j].moveBy( moveWhere(&knots[j - 1], &knots[j]) )
            }
            set[T.toString()] = member
            fmt.Println(knots)
        }
    }

    fmt.Println(len(set))
}

