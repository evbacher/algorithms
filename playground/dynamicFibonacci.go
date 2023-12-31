/* Fibonacci series */

package main

import (
    "strconv"
    "fmt"
)

func main() {
    // Compare non-dynamic and dynamic
    fibonacciTest()
    dynamicFibTest()
}

// A slice to hold already-calulated Fibonacci numbers.
var fibonacciValues []int64

// Calculates fibonacci number of N and returns it.
// Uses fibonacciValues to store already-calculated numbers.
func fibonacciOnTheFly(N int64) int64 {
    fCount++
    
    // input error
    if N < 0 {
        return -1
    }
    
    // Set up fibonacciValues
    fibonacciValues[0] = 0
    fibonacciValues[1] = 1
    
    if N <= 1 {
        return fibonacciValues[N]
    }
    
    if int64(len(fibonacciValues)) > N {
        return fibonacciValues[N]
    }
    fmt.Println("length of fibonacciValues:", len(fibonacciValues))
    // we need to calculate some values now
    fmt.Println("adding to fibonacciValues starting at", len(fibonacciValues))
    fmt.Println("fibValues:",fibonacciValues)
    for i := int64(len(fibonacciValues))-1; i < N; i++ {
        fmt.Println("calculating fib", i+1, " by adding fibs ", i, i-1)
        fibonacciValues = append(fibonacciValues, 
            fibonacciOnTheFly(i) + fibonacciOnTheFly(i-1))
        if fibonacciValues[i] < 0 {
            fmt.Println("ERROR: int64 overflow at N =", i)
            return -1
        }
        fmt.Println("fibValues:",fibonacciValues)
    }
    
    fmt.Println("N:", N)
    return fibonacciValues[N]
}


// Calculates the Fibonacci number for N and returns it.
// (No use of fibonacciValues)
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


// Tests dynamic fibonacci approaches (not all are implemented here).
func dynamicFibTest() {
    fCount = 0
    // Fill-on-the-fly.
    fibonacciValues = make([]int64, 2)
    fibonacciValues[0] = 0
    fibonacciValues[1] = 1

    // Prefilled.
    //initialize_slice()

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
   
        // Uncomment one of the following.
        fmt.Printf("fibonacciOnTheFly(%d) = %d\n", n, fibonacciOnTheFly(n))
        //fmt.Printf("fibonacci_prefilled(%d) = %d\n", n, fibonacci_prefilled(n))
        //fmt.Printf("fibonacci_bottom_up(%d) = %d\n", n, fibonacci_bottom_up(n))
        fmt.Println(fCount, "calls to fibonacciOnTheFly().")
        fmt.Println()

    }

    // Print out all memoized values just so we can see them.
    for i := 0; i < len(fibonacciValues) ; i++ {
        fmt.Printf("%d: %d\n", i, fibonacciValues[i])
    }
}

// Tests non-dynamic fibonacci().
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

