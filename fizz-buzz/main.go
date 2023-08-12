package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'fizzBuzz' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func fizzBuzz(n int32) {
    // Write your code here

    for i:=int32(1);i<n+1;i++ {
        if i%3==0 && i%5==0 {
            fmt.Println("FizzBuzz")
        } else if i%5==0 {
            fmt.Println("Fizz")
        } else if i%15==0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    fizzBuzz(n)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
