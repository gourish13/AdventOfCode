package main

import (
    "os"
	"fmt"
    "bufio"
    "strings"
	"strconv"
)

type Worry int
type Queue []Worry
type Monkey struct {
    inspectCount int
    items Queue
    op []string
    divideBy int
    trueM int
    falseM int
}

func (w *Worry) operate(op []string) {
    var val int 

    if op[1] == "old" {
        val = int(*w)
    } else {
        val, _ = strconv.Atoi(op[1])
    }

    if op[0] == "+" {
        *w += Worry(val)
    } else {
        *w *= Worry(val)
    }
}

func (q *Queue) insert(w Worry) {
    *q = append(*q, w)
}

func (q *Queue) remove() Worry {
    var w = (*q)[0]
    *q = (*q)[1:]
    return w
}

var monkeys []Monkey
var modulus int

func doMonkeyInspection(m *Monkey) {
    for len(m.items) != 0 {
        var item = m.items.remove()
        m.inspectCount++

        item.operate(m.op)
        item %= Worry(modulus)

        if int(item) % m.divideBy == 0 {
            monkeys[ m.trueM ].items.insert(item)
        } else {
            monkeys[ m.falseM ].items.insert(item)
        }
    }
}

func find2Maxima() (int, int) {
    var max, smax int

    for _, m := range monkeys {
        if m.inspectCount >= max {
            smax = max
            max = m.inspectCount
        } else if m.inspectCount >= smax {
            smax = m.inspectCount
        }
    }

    return max, smax
}

func findModulas() int {
    var divisible bool
    var max int = monkeys[0].divideBy

    for _, m := range monkeys {
        if max < m.divideBy {
            max = m.divideBy
        }
    }

    for i := max ;; i++ {
        divisible = true
        for _, m := range monkeys {
            divisible = divisible && i % m.divideBy == 0
        }

        if divisible {
            return i
        }
    }
}

func main() {
    var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        scanner.Text()

        var m Monkey

        // items
        scanner.Scan()
        line := strings.Fields(scanner.Text())[2:]
        for _, v := range line {
            item, _ := strconv.Atoi( strings.Replace(v, ",", "", 1) )
            m.items.insert(Worry(item))
        }

        // operation
        scanner.Scan()
        m.op = strings.Fields(scanner.Text())[4:]

        // Test
        scanner.Scan()
        line = strings.Fields(scanner.Text())[3:]
        val, _ := strconv.Atoi(line[0])
        m.divideBy = val

        // True case
        scanner.Scan()
        line = strings.Fields(scanner.Text())[5:]
        val, _ = strconv.Atoi(line[0])
        m.trueM = val

        // False case
        scanner.Scan()
        line = strings.Fields(scanner.Text())[5:]
        val, _ = strconv.Atoi(line[0])
        m.falseM = val

        monkeys = append(monkeys, m)

        if !scanner.Scan() {
            break
        }
        scanner.Text()
    }

    modulus = findModulas()

    for i := 0; i < 10000; i++ {
        for m := range monkeys {
            doMonkeyInspection(&monkeys[m])
        }
    }

    var max, smax = find2Maxima()
    fmt.Println(max * smax)
}
