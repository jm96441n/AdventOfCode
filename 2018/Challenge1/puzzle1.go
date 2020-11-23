package main

import (
        "fmt"
        "bufio"
        "os"
        "log"
        "strconv"
)

func main() {
        sum := 0
        file, err := os.Open("1_input.txt")
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
                num, _ := strconv.Atoi(scanner.Text())
                sum += num
        }

        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }
        fmt.Println(sum)

}
