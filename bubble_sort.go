
package main

import (
    "fmt"
    "math/rand"
    "time")

func main() {
    rand.Seed(time.Now().UnixNano())

    // Get the number of items and maximum item value.
    var num_items, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&num_items)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)

    // Make and display the unsorted array.
    arr := make_random_array(num_items, max)
    print_array(arr, 40)
    fmt.Println()

    // Sort and display the result.
    bubble_sort(arr)
    print_array(arr, 40)

    // Verify that it's sorted.
    check_sorted(arr)
}

func make_random_array(num_items, max int) []int {
    arr := make([]int, num_items)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < num_items; i++ {
        arr[i] = rand.Intn(max)
    }
    return arr
}

func print_array(arr []int, num_items int) {
    if (len(arr) < num_items) {
        num_items = len(arr)
    }
    for i := 0; i < num_items; i++ {
        fmt.Print(arr[i], ", ")
    }
    fmt.Println()
}

func check_sorted(arr []int) {
    if len(arr) == 1 {
        fmt.Println("The array is sorted.")
        return
    }
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] <= arr[i+1] {
            continue
        }
        fmt.Println("The array is NOT sorted!")
        return
    }
    fmt.Println("The array is sorted")
}

func bubble_sort(arr []int) {
    l := len(arr)
    // If we had to swap, the array is not sorted yet.
    swapped := true
    for swapped {
        // Assume sorted to begin with. May change during loop.
        swapped = false
        // Check next element, swap if necessary
        for i := 0; i < l-1; i++ {
            if arr[i] > arr[i+1] {
                arr[i], arr[i+1] = arr[i+1], arr[i]
                swapped = true
            }
        }
    }
}


