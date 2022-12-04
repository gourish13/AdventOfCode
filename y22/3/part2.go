package main

import (
    "os"
    "fmt"
    "bufio"
)

func getItems(scanner *bufio.Scanner) (string, string, string) {
    var items1, items2, items3 string
    items1 = scanner.Text()
    scanner.Scan()
    items2 = scanner.Text()
    scanner.Scan()
    items3 = scanner.Text()
    return items1, items2, items3
}

func getRunePriority(c rune) int {
    i := int(c)
    if i < 97 {
        return i - 64 + 26
    } else {
        return i - 96
    }
}

func fillPriTable(itemNum int, items string, pt []int) {
    var i int

    for _, c := range items {
        i = getRunePriority(c)
        if pt[i] == itemNum {
            pt[i]++
        }
    }
}

func getCommonItemPriority(items1, items2, items3 string) int {
    var pt = make([]int, 53)

    fillPriTable(0, items1, pt)
    fillPriTable(1, items2, pt)
    fillPriTable(2, items3, pt)

    for i, p := range pt {
        if p == 3 {
            return i
        }
    }

    return 0 // Probably won't react according to test cases
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var pri int

    for scanner.Scan() {
        var items1, items2, items3 string = getItems(scanner)
        pri += getCommonItemPriority(items1, items2, items3)
    }

    fmt.Println(pri)
}