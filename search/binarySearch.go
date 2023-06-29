
/*
 Binary sort. Uses quicksort or sort.Ints() to sort the array first.
 There is a good discussion (and animation) in the Wikipedia article:
 https://en.wikipedia.org/wiki/Binary_search_algorithm.
 
 Jon Bentley also has an excellent discussion in Programming Pearls.
 But Joshua Bloch points out a longstanding bug that both he and
 Bentley missed. See
 https://ai.googleblog.com/2006/06/extra-extra-read-all-about-it-nearly.html.
*/

package main

import (
    "fmt"
    "math"
    "math/rand"
    "sort"
    "strconv"
    "time"
)

func main() {
    
    // Get the number of items and maximum item value.
    var num_items, max int
    var useLib string
    fmt.Printf("# Items: ")
    fmt.Scanln(&num_items)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)
    
    useLibSort := false
    fmt.Printf("Use library sort.Ints? (y or n (default): ")
    fmt.Scanln(&useLib)
    if useLib == "y" {
        useLibSort = true
    }
    
    // Make and display the slice of random ints.
    values := makeRandomSlice(num_items, max)
    printSlice(values, 40)
    fmt.Println()
    
    // Sort the array first (necessary for binary search)
    if useLibSort {
        // Use the Golang sort package.
        fmt.Println("Using sort.Ints().")
        sort.Ints(values)
    } else {
        // Use the homegrown quicksort.
        fmt.Println("Using local quicksort.")
        quicksort(values)
    }
    
    printSlice(values, 40)
    fmt.Println("\nN: ", len(values))
    fmt.Println("log2(N): ", math.Log2(float64(len(values))), "\n" )


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
        
        index, numTests = binarySearch(values, itarget)
        
        if index >= 0 {
            fmt.Print("Target found at values[", index, "] after ", numTests, " tests.\n")
        } else {
            fmt.Println("Target not found after", numTests, "tests.")
        }
        
        // reset target value
        target = ""
    }
}

// Searches the (sorted) array values.
// Returns the index of the target or -1 if not found.
// Also returns the number of tests.
func binarySearch(values []int, target int) (mid, numTests int) {
    
    // Set starting index for left and right.
    l := 0
    r := len(values) - 1
    for l <= r {
        numTests++
        // There is an interesting story about a bug in calculating the midpoint at
        // https://ai.googleblog.com/2006/06/extra-extra-read-all-about-it-nearly.html
        // Only a problem when l and r are both very large (2^30).
        // mid = (l + r)/2  //(this is where the bug can show up)
        mid = (l + r) >> 1
        if values[mid] < target {
            l = mid + 1
        } else if values[mid] > target {
            r = mid - 1 
        } else {
            // not < or > the target, so this must be it!
            return mid, numTests
        }
    }
    // l must be > r, meaning the target is not in the values array.
    return -1, numTests
}

// Sorts an array (actually a slice of an array).
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

// Returns a slice of numItems random ints, values from 0 to max.
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
