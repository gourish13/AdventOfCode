package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func max3s(max1, max2, max3, cur int) (int, int, int) {
    if cur >= max1 {
        return cur, max1, max2
    } else if cur >= max2 {
        return max1, cur, max2
    } else if cur >= max3 {
        return max1, max2, cur
    } else {
        return max1, max2, max3
    }
}

func main()  {
    var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
    var (
        maxCal int 
        secdMaxCal int
        thrdMaxCal int
        curCalSum int
        calorie int
    )
    
    for scanner.Scan() {
        var text string = scanner.Text()

        if text == "" {
            maxCal, secdMaxCal, thrdMaxCal = max3s(maxCal, secdMaxCal, thrdMaxCal, curCalSum)
            curCalSum = 0
        } else {
            calorie, _ = strconv.Atoi(text)
            curCalSum += calorie
        }
    }

    maxCal, secdMaxCal, thrdMaxCal = max3s(maxCal, secdMaxCal, thrdMaxCal, curCalSum)
    fmt.Println(maxCal + secdMaxCal + thrdMaxCal)
}

