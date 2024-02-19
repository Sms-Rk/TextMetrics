package helper

import (
    "bufio"
    "os"
    "fmt"
)

func GetLineCount(filename string) (int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineCount := 0
    for scanner.Scan() {
        lineCount++
    }

    if err := scanner.Err(); err != nil {
        return 0, err
    }
    fmt.Println(lineCount)
    return lineCount, nil
}