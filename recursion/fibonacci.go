/* Fibonacci series */

package main

import (
    "strconv"
    "fmt")

func main() {
    fibonacciTest()
}

// Calculates the Fibonacci number for N and returns it.
func fibonacci(N int64) int64 {
    fCount++
    
    // input error
    if N < 0 {
        return -1
    }
    
    // base case
    if N == 0 {
        return 0
    }
    if N == 1 {
        return 1
    }
    return fibonacci(N-1) + fibonacci(N-2)
}

// Keep track of calls to fibonacci().
var fCount int 


func fibonacciTest() {
    for {
        fCount = 0
        // Get n as a string.
        var n_string string
        fmt.Printf("N: ")
        fmt.Scanln(&n_string)

        // If the n string is blank, break out of the loop.
        if len(n_string) == 0 { break }

        // Convert to int and calculate the Fibonacci number.
        n, _ := strconv.ParseInt(n_string, 10, 64)
        fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
        fmt.Println(fCount, "calls to fibonacci().")
    }
}
