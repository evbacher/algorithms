
/*
Quicksort: Sort an array of ints. Note that we refer to an array in
common usage, but we are using a Golang slice, since slices are really
the way to work with arrays in Golang. Slices are built on top of arrays
and make them easy to work with (and they are passed by reference). for
more on slices, see https://go.dev/tour/moretypes/7.
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    
    // Test array (from Programming Pearls, by Jon Bentley):
    //arr := []int{55,41,59,26,53,58,97,93}
    
    // Create a random array and display up to 40 elements.
    arr := makeRandomSlice(3, 1000000)
    printSlice(arr, 40)
    fmt.Println()

    // Sort and display the result.
    quicksort(arr)
    printSlice(arr, 40)

    // Verify that it's sorted.
    checkSorted(arr)
}

// Sorts an array.
// See Programming Pearls by Jon Bentley for a great discussion of
// this algorithm.
func quicksort(arr []int) {

    // If array has 0 or 1 elements, it's already sorted.
    if len(arr) < 2 {
        return
    }
    
    // Partition the array and get the new midpoint index.
    mid := partition(arr)
    
    // Call this function recursively on the two partitions
    // (not including the mid values).
    quicksort(arr[:mid])
    quicksort(arr[mid+1:])
}

// Partitions an array (initial pivot is the first element arr[0]).
// The goal of partitioning is to move all the values less than
// or equal to the pivot value to the left side of the array.
// When you're done, move the pivot value between the small and
// large values (relative to the pivot).
// Returns the final position of the pivot (middle).
func partition(arr []int) int {
    lo, hi := 0, len(arr)-1

    middle := lo
    pivot := arr[middle]
    for i := middle+1; i < hi+1; i++{
        if arr[i] < pivot {
            
            // Increment middle index, and swap the element now at the
            // middle index with a[i], which we just tested. Note that,
            // for the first iteration, the swap will just keep the element
            // where it is (at arr[1]). Also note that we are NOT moving the
            // pivot element yet.
            middle++
            arr[middle], arr[i] = arr[i], arr[middle]
        }
    }
    
    // Now that all the partitioning is done we can move the pivot to the middle index.
    arr[middle], arr[lo] = arr[lo], arr[middle]
    return middle
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
    if (len(arr) < numItems) {
        numItems = len(arr)
    }
    for i := 0; i < numItems; i++ {
        fmt.Print(arr[i], ", ")
    }
    fmt.Println()
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
