/* Recursive factorial function. */

package main
import "fmt"

func main() {
    factorialTest()
}

func factorialTest() {
    var n int64
    for n = 0; n <= 21; n++ {
        fmt.Printf("%3d! = %20d\n", n, factorial(n))
    }
    fmt.Println()
}

// Recursively calculates N! and returns it.
func factorial(N int64) int64 {
    
    // input check
    if N < 0 {
        fmt.Println("Number must >= 0")
        return -1
    }
    
    // base case
    if N == 0 {
        return 1
    }
    // not done yet
    return N * factorial(N - 1)
}
