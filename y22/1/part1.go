package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func max(prev, cur int) int {
    if cur > prev {
        return cur
    } else {
        return prev
    }
}

func main()  {
    var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
    var (
        maxCal int 
        curCalSum int
        calorie int
    )
    
    for scanner.Scan() {
        var text string = scanner.Text()

        if text == "" {
            maxCal = max(maxCal, curCalSum)
            curCalSum = 0
        } else {
            calorie, _ = strconv.Atoi(text)
            curCalSum += calorie
        }
    }

    maxCal = max(maxCal, curCalSum)
    fmt.Println(maxCal)
}
