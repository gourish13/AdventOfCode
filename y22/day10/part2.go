package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
    "strings"
)

func drawPixel(pixels []rune, cycle, pixel, X int) {
    var offset = (cycle / 40) * 40
    if X == pixel || X - 1 == pixel || X + 1 == pixel {
        pixels[offset + pixel] = '#'
    }
}

func crtDisplay(pixels []rune) {
    for i := 0; i < 240; i += 40 {
        fmt.Println(string(pixels[i:i+40]))
    }
}

func main() {
    var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
    var pixel, cycle, X = 0, 0, 1


    var pixels = make([]rune, 240)
    for i := range pixels {
        pixels[i] = '.'
    }

    var line []string
    for scanner.Scan() && cycle < len(pixels) && X < len(pixels) {
        line = strings.Fields(scanner.Text())

        if len(line) == 1 {
            cycle++
            drawPixel(pixels, cycle - 1, pixel, X)
            pixel = (pixel + 1) % 40
        } else {
            V, _ := strconv.Atoi(line[1])
            X += V
            cycle += 2
            drawPixel(pixels, cycle - 2, pixel, X - V)
            pixel = (pixel + 1) % 40
            drawPixel(pixels, cycle - 1, pixel, X - V)
            pixel = (pixel + 1) % 40
        }
    }
    crtDisplay(pixels)
}
