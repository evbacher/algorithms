/*
 Singly linked list data structure used to implement a stack.
 This verson also adds functions to detect a loop in the list.
*/

package main

import "fmt"

func main() {

    // milestone test for basic list functionality   
    //smallListTest()
    
    // milestone test for loop detection
    loopTest()
    
    // development testing
    //testList()
}

// Milestone test for loop detection.
func loopTest () {
    // Make a list from a slice of values.
    values := []string {
        "0", "1", "2", "3", "4", "5",
    }
    list := makeLinkedList()
    list.addRange(values)

    fmt.Println(list.toString(" "))
    if list.hasLoop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
    fmt.Println()

    // Make cell 5 point to cell 2.
    list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

    fmt.Println(list.toStringMax(" ", 10))
    if list.hasLoop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
    fmt.Println()

    // Make cell 4 point to cell 2.
    list.sentinel.next.next.next.next.next = list.sentinel.next.next

    fmt.Println(list.toStringMax(" ", 10))
    if list.hasLoop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
}

func smallListTest() {
    
   // Make a list from a slice of values.
   greek_letters := []string {
       "α", "β", "γ", "δ", "ε",
   }
   list := makeLinkedList()
   list.addRange(greek_letters)
   fmt.Println(list.toString(" "))
   fmt.Println()

   // Demonstrate a stack.
   stack := makeLinkedList()
   stack.push("Apple")
   stack.push("Banana")
   stack.push("Coconut")
   stack.push("Date")
   for !stack.isEmpty() {
       fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
           stack.pop(),
           stack.length(),
           stack.toString(" "))
   }  
}

// for development testing of basic list functionality...
func testList() {
    // Create some Items.
    aItem := Item { "Apple", nil }
    bItem := Item { data: "Banana" }
    aItem.next = &bItem
    cItem := Item { "Cherry", nil }
    dItem := Item {"Date", nil}
    
    fmt.Println("using makeLinkedList now")
    list := makeLinkedList()
    fmt.Println("is list empty?", list.isEmpty())
    list.sentinel.addAfter(&aItem)
    fmt.Println("is list empty?", list.isEmpty())
    aItem.addAfter(&bItem)
    bItem.addAfter(&cItem)
    cItem.addAfter(&dItem)
    fmt.Println(list.toString(" "))
    
    fmt.Println("Try to add nil.")
    bItem.addAfter(nil)
    fmt.Println(list.toString(" "))
    
    // add a copy of aItem  (not a reference to the original aItem)
    anotherItem := aItem
    fmt.Println("Adding a new copy of ",aItem,": ",anotherItem,"after",cItem)
    cItem.addAfter(&anotherItem)
    
    fmt.Println("\nlist before addRange:")
    fmt.Println(list.toString(" "))
    fmt.Println("length", list.length())
    
    // Create an array of new string data.
    newData := []string{"kiwis!", "prunes!", "dates"}
    list.addRange(newData)
    
    // Examine the list.
    fmt.Println("\nlist AFTER addRange:")
    fmt.Println(list.toString(" "))
    fmt.Println("length", list.length())
    
    mango := Item{"mango", nil}
    fmt.Println("pushing", mango)
    list.push("mango")
    fmt.Println(list.toString(" "))
    
    fmt.Println("\nTry pushing nil.")
    list.push("")
    
    top := list.pop()
    fmt.Println("top:", top)
    fmt.Println(list.toString(" "))
}

type Item struct {
    data    string
    next    *Item
}

type LinkedList struct {
    sentinel    *Item
}

// Some LinkedList functions.

// Creates a new LinkedList and initializes its sentinel.
func makeLinkedList() LinkedList {
    list := LinkedList {}
    list.sentinel = &Item {"SENTINEL", nil}
    return list
}

// Adds an Item after me.
func (me *Item) addAfter(after *Item) {
    if after == nil {
        return
    }
    // Order matters here -- set after.next first.
    after.next = me.next
    me.next = after
}

// Deletes the Item after me and returns the deleted Item
// (or nil if none exists, which means we are at the last Item).
func (me *Item) deleteAfter() *Item {
    if me.next == nil {
        //panic("me.next is nil")
        // Could panic or just return nil if this is the last item?
        return nil
    }
    after := me.next
    me.next = after.next
    return after
}

// Creates Items for the slice of strings and adds them to the list.
func (list *LinkedList) addRange(values []string) {
    
    // First, find the last item.
    // lastItem is a reference to the last item.
    // Correct version of the loop to find the last item
    // (let item be a pointer to an Item, not a copy of an Item):
    var lastItem *Item
    for item := list.sentinel; ; item = item.next {
        if item.next == nil {
            lastItem = item
            break
        }
    }
    // Now, add new Items to the end of the list.
	for _, v := range values {
	    // Create a brand new Item to add to the end of the list.
        newItem := Item{v, nil}
        lastItem.addAfter(&newItem)
        lastItem = &newItem
    }
}

// Returns a string with the data contained in each Item, separated by separator.
func (list *LinkedList) toString(separator string) string {
    listString := ""
    for item := list.sentinel.next; item != nil; item = item.next {
        listString += item.data
        // Option: add item.next to the string.
        //listString  += "," + fmt.Sprint(item.next)
        if item.next != nil {
            listString += separator
        }
    }
    return listString
}

// Returns a string with the data contained in each Item, separated by separator,
// up to max Items.
func (list *LinkedList) toStringMax(separator string, max int) string {
    
    listString := ""
    count := 0
    for item := list.sentinel.next; item != nil; item = item.next {
        count++
        if count > max {
            break
        }
        
        listString += item.data
        if item.next != nil {
            listString += separator
        }
    }
    return listString
}


// Returns the length of the list (not counting the sentinel).
func (list *LinkedList) length() int {
    length := 0
    for item := list.sentinel.next; item != nil; item = item.next {
        length++
    }
    return length
}

// Returns true if the list is empty.
func (list *LinkedList) isEmpty() bool {
    return list.sentinel.next == nil
}

// Adds a new Item to the front of the list (top of the stack).
func (list *LinkedList) push(data string) {
    // Create a new Item using data.
    item := Item{data, nil}
    list.sentinel.addAfter(&item)
}

// Removes the Item at the front of the list and returns its string value.
func (list *LinkedList) pop() string {
    topItem := list.sentinel.deleteAfter()
    return topItem.data
}

// Returns true if list has a loop, false otherwise.
func (list *LinkedList) hasLoop() bool {
    
    fast, slow := list.sentinel, list.sentinel
    

    for {
        
        // Note that we need to make sure we're not using a nil pointer.
        if slow != nil {
            slow = slow.next
        } else {
            return false
        }
        
        if fast != nil && fast.next !=nil {
            fast = fast.next.next
        } else {
            return false
        }
        
        if fast == slow {
            return true
        }
    }
}

// There are several other useful functions you can add for lists:
// contains(), find(), remove(), removeAt(), append(), addList(), 
// toSlice(), clone(), clear().
