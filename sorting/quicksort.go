
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

// For testing only, to see how deep the quicksort stack is.
var qcount = 0

func main() {
    
    // Get the number of items and maximum item value.
    var num_items, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&num_items)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)
    
    /*
    // Already sorted array for testing.
    values := make([]int, 1000)
    for i := 0; i < 1000; i++ {
        values[i] = i
    }
    */

    // Make and display the unsorted slice.
    values := makeRandomSlice(num_items, max)
    printSlice(values, 40)
    fmt.Println()

    // Sort and display the result.
    quicksort(values)
    printSlice(values, 40)

    // Verify that it's sorted.
    checkSorted(values)
    
    fmt.Println("qcount: ", qcount)
}


// Sorts an array.
// See Programming Pearls by Jon Bentley for a great discussion of
// this algorithm.
func quicksort(arr []int) {
    
    qcount++

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
// When you're done, move the pivot value to a position between
// the small and large value partitions.
// Returns the final position of the pivot (middle).
func partition(arr []int) int {
    lo, hi := 0, len(arr)-1
    
    // Instead of always starting the pivot at the first element,
    // we can use the median of the first, last, and middle element
    // (as suggested by Sedgewick).
    middle := (lo + hi) / 2
    if arr[middle] > arr[hi] {
        // middle <= hi after swap.
        arr[middle], arr[hi] = arr[hi], arr[middle]
    }
    if arr[lo] > arr[hi] {
        // middle <= hi  and lo <= hi after swap.
        arr[lo], arr[hi] = arr[hi], arr[lo]
    }
    if arr[middle] > arr[lo] {
        // Put the median-of-three value into arr[lo]
        // lo >= middle and lo <= hi after swap.
        arr[middle], arr[lo] = arr[lo], arr[middle]
    }

    // Even though we're calling this mid, it starts out
    // as the first position in the slice, where the pivot is.
    mid := lo
    pivot := arr[mid]
    for i := mid+1; i < hi+1; i++{
        if arr[i] < pivot {
            
            // Increment mid index, and swap the element now at the
            // mid index with a[i], which we just tested. Note that,
            // for the first iteration, the swap will just keep the element
            // where it is (at arr[1]). Also note that we are NOT moving the
            // pivot element yet.
            mid++
            arr[mid], arr[i] = arr[i], arr[mid]
        }
    }
    
    // Now that all the partitioning is done we can move the pivot to the mid index.
    arr[mid], arr[lo] = arr[lo], arr[mid]
    return mid
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

