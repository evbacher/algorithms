/*
 Doubly linked list data structure.
 Used to implement queue and dequeue.
 Repurposing code from singly linked list, with many mods.
*/

package main

import (
    "fmt"
    "reflect" // for reflect.TypeOf()
)

func main() {

    // milestone test for basic list functionality   
    //smallListTest()
    
    // milestone test for loop detection (in doubly-linked lists)
    loopTest()
    
    // development testing
    //testList()
    
    // doubly-linked list testing
    dlltest()
    
    // queue and dequeue testing
    queueTests()
}

func testList() {
    list := makeDoublyLinkedList()
    fmt.Println(reflect.TypeOf(list))
    newItem := Item{"testing",nil,nil}
    list.topSentinel.addAfter(&newItem)
    newItem.addAfter(&Item{"another",nil,nil})
    newItem.addBefore(&Item{"kiwi",nil,nil})
    fmt.Println(list.toString(" "))
    
    newItem.delete()
    fmt.Println(list.toString(" "))
    
    newData := []string{"cats","dogs","iguanas"}
    list.addRange(newData)
    fmt.Println(list.toString(" "))
}

// Tests basic dll functionality.
func dlltest() {
    fmt.Println("\n*** Doubly-linked list functions ***")
    // Make a list from a slice of values.
    list := makeDoublyLinkedList()
    animals := []string {
        "Ant",
        "Bat",
        "Cat",
        "Dog",
        "Elk",
        "Fox",
    }
    list.addRange(animals)
    fmt.Println(list.toString(" "))
}

// Milestone test for loop detection.
func loopTest () {
    // Make a list from a slice of values.
    values := []string {
        "0", "1", "2", "3", "4", "5",
    }
    list := makeDoublyLinkedList()
    list.addRange(values)

    fmt.Println(list.toString(" "))
    if list.hasLoop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
    fmt.Println()

    // Make cell 5 point to cell 2.
    list.topSentinel.next.next.next.next.next.next = list.topSentinel.next.next

    fmt.Println(list.toStringMax(" ", 10))
    if list.hasLoop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
    fmt.Println()

    // Make cell 4 point to cell 2.
    list.topSentinel.next.next.next.next.next = list.topSentinel.next.next

    fmt.Println(list.toStringMax(" ", 10))
    if list.hasLoop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
}

func queueTests() {
    // Test queue functions.
    fmt.Printf("\n*** Queue Functions ***\n")
    queue := makeDoublyLinkedList()
    queue.enqueue("Agate")
    queue.enqueue("Beryl")
    fmt.Printf("%s ", queue.dequeue())
    queue.enqueue("Citrine")
    fmt.Printf("%s ", queue.dequeue())
    fmt.Printf("%s ", queue.dequeue())
    queue.enqueue("Diamond")
    queue.enqueue("Emerald")
    for !queue.isEmpty() {
        fmt.Printf("%s ", queue.dequeue())
    }
    fmt.Printf("\n\n")

    // Test deque functions. Names starting
    // with F have a fast pass.
    fmt.Printf("*** Deque Functions ***\n")
    deque := makeDoublyLinkedList()
    deque.pushTop("Ann")
    deque.pushTop("Ben")
    fmt.Printf("%s ", deque.popBottom())
    deque.pushBottom("F-Cat")
    fmt.Printf("%s ", deque.popBottom())
    fmt.Printf("%s ", deque.popBottom())
    deque.pushBottom("F-Dan")
    deque.pushTop("Eva")
    for !deque.isEmpty() {
        fmt.Printf("%s ", deque.popBottom())
    }
    fmt.Printf("\n")
}

// Item for use in a doubly-linked list.
type Item struct {
    data    string
    prev    *Item
    next    *Item
}

// A doubly-linked list starts off with two sentinels.
type DoublyLinkedList struct {
    topSentinel     *Item
    bottomSentinel  *Item 
}

// Some DoublyLinkedList functions.

// Creates a new DoublyLinkedList and initializes its sentinels to point to each other.
// Note that a doubly linked list should have no nil pointers (no end).
func makeDoublyLinkedList() DoublyLinkedList {
    // Create list and empty sentinels.
    list := DoublyLinkedList {}
    list.topSentinel = &Item {"topSentinel", nil, nil}
    list.bottomSentinel = &Item{"bottomSentinel", nil, nil}
    
    // Point the sentinels to each other.
    (list.topSentinel).prev, (list.topSentinel).next = 
        list.bottomSentinel, list.bottomSentinel
    (list.bottomSentinel).prev, (list.bottomSentinel).next = 
        list.topSentinel, list.topSentinel

    return list
}

// Adds an Item after me.
func (me *Item) addAfter(after *Item) {
    if after == nil {
        return
    }
    // Order matters. Set after pointers first.
    after.prev, after.next = me, me.next
    (me.next).prev, me.next = after, after
}

// Adds an item before me.
func (me *Item) addBefore(before *Item) {
    if before == nil {
        return
    }
    // Set pointers for before, then use addAfter to put things in the correct order.
    before.prev, before.next = me.prev, me
    (me.prev).addAfter(before)
}

// Deletes the current Item.
func (me *Item) delete() *Item {
    // Deal me out.
    (me.prev).next = me.next
    (me.next).prev = me.prev
    return me
}

// Creates Items for the slice of strings and adds them to the bottom of the list.
func (list *DoublyLinkedList) addRange(values []string) {
    // Add new Items to the bottom of the list.
	for _, v := range values {
        newItem := Item{v, nil, nil}
        (list.bottomSentinel).addBefore(&newItem)
    }
}

// Returns a string with the data contained in each Item, separated by separator.
func (list *DoublyLinkedList) toString(separator string) string {
    listString := ""
    for item := list.topSentinel.next; item.data != "bottomSentinel"; item = item.next {
        listString += item.data
        // Option: add item.next to the string.
        //listString  += "," + fmt.Sprint(item.next)
        if item.next != list.bottomSentinel {
            listString += separator
        }
    }
    return listString
}

// Returns a string with the data contained in each Item, separated by separator,
// up to max Items.
func (list *DoublyLinkedList) toStringMax(separator string, max int) string {
    
    listString := ""
    count := 0
    for item := list.topSentinel.next; item != list.bottomSentinel ; item = item.next {
        count++
        if count > max {
            break
        }
        
        listString += item.data
        if item.next != list.bottomSentinel {
            listString += separator
        }
    }
    return listString
}


// Returns the length of the list (not counting the sentinel).
func (list *DoublyLinkedList) length() int {
    length := 0
    for item := list.topSentinel.next; item != list.bottomSentinel; item = item.next {
        length++
    }
    return length
}

// Returns true if the list is empty.
func (list *DoublyLinkedList) isEmpty() bool {
    return list.topSentinel.next == list.bottomSentinel
}

// List function: push() and pop().

// Adds a new Item to the front of the list (top of the stack).
func (list *DoublyLinkedList) push(data string) {
    // Create a new Item using data.
    item := Item{data, nil, nil}
    list.topSentinel.addAfter(&item)
}

// Removes the Item at the front of the list and returns its string value.
func (list *DoublyLinkedList) pop() string {
    topItem := (list.topSentinel.next).delete()
    return topItem.data
}

// Queue functions (for the bottom of the list)

// Uses push() to add an Item to the front of the list.
func (list *DoublyLinkedList) enqueue(data string) {
    list.push(data)
}

// Removes the Item at the bottom of the list and returns its string value.
func (list* DoublyLinkedList) dequeue() string {
    //last := (list.bottomSentinel).prev.delete()
    //return last.data
    
    // More concise, but maybe less readable.
    return (list.bottomSentinel).prev.delete().data
}

// Dequeue functions (both top and bottom of the list).

// Adds a data Item to the bottom of the dequeue.
func (list *DoublyLinkedList) pushBottom(data string) {
    (list.bottomSentinel).addBefore(&Item{data, nil, nil})
}

// Removes an Item from the bottom of the dequeue and returns the data.
func (list *DoublyLinkedList) popBottom() string {
    return list.dequeue()
}

// Adds a data Item to the top of the dequeue.
func (list *DoublyLinkedList) pushTop(data string) {
    list.enqueue(data)
}

// Removes an Item from the top of the dequeue and returns the data.
func (list *DoublyLinkedList) popTop() string {
    return list.pop()
}

// Tests for loop in doubly-linked list.
// Returns true if list has a loop, false otherwise.
func (list *DoublyLinkedList) hasLoop() bool {
    
    fast, slow := list.topSentinel, list.topSentinel
    for {
        // Check for end of list first.
        if slow == list.bottomSentinel || fast == list.bottomSentinel || fast.next == list.bottomSentinel {
            return false
        }
        // Not at end: advance pointers.
        slow = slow.next
        fast = fast.next.next
        
        // fast looped around.
        if fast == slow {
            return true
        }
    }
}


// There are several other useful functions you can add for lists:
// contains(), find(), remove(), removeAt(), append(), addList(), 
// toSlice(), clone(), clear().
