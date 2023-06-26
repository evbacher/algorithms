
/* An implementation of counting sort. This works for a large number of items
that have a relatively small range of values.
*/

package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "time"
)

type Customer struct {
    id              string
    numPurchases    int
}

func main() {
    
    // Get the number of items and maximum item value.

    var numItems, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&numItems)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)

    // Make and display the unsorted slice.
    values := makeRandomSlice(numItems, max)
    fmt.Println("unsorted:")
    printSlice(values, 40)
    fmt.Println()

    // Sort and display the result.
    sorted := countingSort(values, max)
    fmt.Println("sorted:")
    printSlice(sorted, 40)

    // Verify that it's sorted.
    checkSorted(sorted)
}

// Count the numPurchases for each Customer and sort by numPurchases.
func countingSort(customers []Customer, max int) []Customer {
    counts := make([]int, max+1)

    l := len(customers)
    for i := 0; i < l; i++ {
        counts[customers[i].numPurchases]++
    }

    // Adjust counts slice so that each index contains the number of items <= to the
    // value at the current index. So we basically add all the previous values to
    // get the count at the current index. This gives us information about the order.
    for i := 1; i < len(counts); i++ {
        counts[i] += counts[i-1]
    }

    lenCustomers := len(customers)
    sorted := make([]Customer, lenCustomers)
    
    // Rearrange sorted slice. Start at the back, so we preserve the relative order
    // for duplicates.
    for j := lenCustomers-1; j >= 0; j-- {
        // Check the counts array for numPurchases of current Customer.
        countIndex := customers[j].numPurchases

        // The number in counts tells us where to put this Customer
        // in the sorted slice (once we decrement it).
        counts[countIndex]--
        sorted[counts[countIndex]] = customers[j]
    }

    return sorted
}

/* Utility functions. */

// Returns a slice of numItems Customers with random numPurchases, up to max.
func makeRandomSlice(numItems, max int) []Customer {
    // We are calling this an array, but it's actually a slice.
    customers := make([]Customer, numItems)
    
    // Seed the random number generator, otherwise it will generate the same set.
    rand.Seed(time.Now().UnixNano())
    
    for i := 0; i < numItems; i++ {
        customers[i].id = "c" + strconv.Itoa(i)
        customers[i].numPurchases = rand.Intn(max)
    }
    return customers
}

// Prints the first numItems of the Customer slice.
func printSlice(arr []Customer, numItems int) {
    if (len(arr) < numItems) {
        numItems = len(arr)
    }
    for i := 0; i < numItems; i++ {
        fmt.Print(arr[i], ", ")
    }
    fmt.Println()
}


// Checks to see if arr of Customers is sorted.
func checkSorted(arr []Customer) {
    for i := 0; i < len(arr)-1; i++ {
        if arr[i].numPurchases <= arr[i+1].numPurchases {
            continue
        }
        fmt.Println("The array is NOT sorted!")
        return
    }
    fmt.Println("The array is sorted")
}
