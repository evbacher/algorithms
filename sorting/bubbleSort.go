
// Implementation of bubble sort algorithm and a couple
// of related utility functions.

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {

    // Get the number of items and maximum item value.
    var numItems, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&numItems)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)

    // Make and display the unsorted array.
    arr := makeRandomSlice(numItems, max)
    printSlice(arr, 40)
    fmt.Println()

    // Sort and display the result.
    bubbleSort(arr)
    printSlice(arr, 40)

    // Verify that it's sorted.
    checkSorted(arr)
}

// Sort an array of integers.
func bubbleSort(arr []int) {
    // If we had to swap, the array is not sorted yet.
    swapped := true
    for swapped {
        // Assume sorted to begin with. May change during loop.
        swapped = false
        // Check next element, swap if necessary.
        for i,l := 0,len(arr); i < l-1; i++ {
            if arr[i] > arr[i+1] {
                arr[i], arr[i+1] = arr[i+1], arr[i]
                swapped = true
            }
        }
    }
}

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
    closer := "...]\n"
    if (len(arr) < numItems) {
        numItems = len(arr)
        closer = "]\n"
    }
    fmt.Print("[ ")
    for i := 0; i < numItems; i++ {
        fmt.Print(arr[i], " ")
    }
    fmt.Print(closer)
}

// Checks to see if arr is sorted.
func checkSorted(arr []int) {
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] <= arr[i+1] {
            continue
        }
        fmt.Println("The array is NOT sorted!")
        return
    }
    fmt.Println("The array is sorted")
}


