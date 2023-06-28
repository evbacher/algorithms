
/*
 An implementation of a simple linear search (brute force) algorithm.
*/

package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "time"
)

func main() {
    
    // Get the number of items and maximum item value.
    var num_items, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&num_items)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)
    
    // Make and display the slice of random ints.
    values := makeRandomSlice(num_items, max)
    printSlice(values, 40)
    fmt.Println()

    target := ""
    var index, numTests int
    // Loop to find target values.
    for {
        fmt.Printf("Target: ")
        fmt.Scanln(&target)

        // If input is just a newline (empty string) break the loop.
        if target == "" {
            break
        }
        
        itarget, err := strconv.Atoi(target)
        if err != nil {
            fmt.Println("error: ", err)
            fmt.Println("Please enter an integer.")
            target = ""
            continue
        }
        
        index, numTests = linearSearch(values, itarget)
        
        if index >= 0 {
            fmt.Print("Target found at values[", index, "] after ", numTests, " tests.\n")
        } else {
            fmt.Println("Target not found after", numTests, "tests.")
        }
        
        // reset target value
        target = ""
    }
}

// Searches the values slice from the beginning for target.
// Returns the index of target (or -1 if not found) and the number of tests required.
func linearSearch(values []int, target int) (index, numTests int) {
    for i := 0; i < len(values); i++ {
        numTests++
        if values[i] == target {
            index = i
            return index, numTests
        }
    }
    // Failed to find target.
    index = -1
    
    return index, numTests
}

/* Utility functions. */

// Returns a slice of numItems random ints, up to max.
func makeRandomSlice(numItems, max int) []int {
    // We are calling this an array, but it's actually a slice.
    arr := make([]int, numItems)
    
    // Seed the random number generator, otherwise it will generate the same set.
    rand.Seed(time.Now().UnixNano())
    
    for i := 0; i < numItems; i++ {
        arr[i] = rand.Intn(max)
    }
    return arr
}

// Prints the first numItems of the slice arr.
func printSlice(arr []int, numItems int) {
    if (len(arr) < numItems) {
        numItems = len(arr)
    }
    for i := 0; i < numItems; i++ {
        fmt.Print(arr[i], ", ")
    }
    fmt.Println()
}
