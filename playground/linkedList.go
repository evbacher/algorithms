/*
 Singly linked list data structure used to implement a stack.
*/

package main

import "fmt"

/*
func main() {
   // small_list_test()

   // Make a list from a slice of values.
   greek_letters := []string {
       "α", "β", "γ", "δ", "ε",
   }
   list := makeLinkedList()
   list.addRange(greek_letters)
   fmt.Println(list.to_string(" "))
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
*/

func main() {
    test()
}

// for testing...
func test() {
    // Create some Items.
    aItem = Item { "Apple", nil }
    bItem = Item { data: "Banana" }
    aItem.next = &bItem
    //top := &aItem

    cItem = Item { "Cherry", nil }
    dItem = Item {"Date", nil}

    /*
    // Print the list.
    for item := top; item != nil; item = item.next {
        fmt.Printf("%s ", item.data)
    }
    fmt.Println()
    
    
    //aItem.deleteAfter()
    cItem.deleteAfter()
    for item := top; item != nil; item = item.next {
        fmt.Printf("%s ", item.data)
    }
    fmt.Println()
    */
    

    fmt.Println("using makeLinkedList now")
    list := makeLinkedList()
    list.sentinel.addAfter(&aItem)
    fmt.Println(list.toString(" "))
    aItem.addAfter(&bItem)
    fmt.Println(list.toString(" "))
    bItem.addAfter(&cItem)
    fmt.Println(list.toString(" "))
    cItem.addAfter(&dItem)
    fmt.Println(list.toString(" "))

    // add a copy of aItem  (not a reference to the original aItem)
    anotherItem := aItem
    fmt.Println(anotherItem)
    cItem.addAfter(&anotherItem)
    
    fmt.Println("\nlist before addRange:")
    fmt.Println(list.toString(" "))
    
    // Create an array of new string data.
    newData := []string{"kiwis!", "prunes!", "dates"}
    //newData := []string{"kiwis!"}
    list.addRange(newData)
    
    // Just printing in a loop.
    fmt.Println("\nlist AFTER addRange:")
    fmt.Println(list.toString(" "))

}

// These probably don't really need to be global scope.
var aItem, bItem, cItem, dItem Item

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
    // Order matters here -- set after.next first.
    fmt.Println("\n",me,"calling addAfter(",after,")")
    fmt.Println("before:",me,after)
    after.next = me.next
    me.next = after
    fmt.Println("after:",me,after)
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

// Create items for the slice of strings and add them to the list.
func (list *LinkedList) addRange(values []string) {
    // First, find the last item.
    // lastItem is a reference to the last item.
    /*
    var lastItem *Item
    for item := *list.sentinel; ; item = *item.next {
        if item.next == nil {
            lastItem = &item
            break
        }
    }
    */
    
    // First, find the last item.
    // lastItem is a reference to the last item.
    // Correct version of the loop to find the last item
    // (let item be a pointer to an Item):
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
        // Point to the new Item.
        lastItem.addAfter(&newItem)
        // Make the new Item the last Item. 
        lastItem = &newItem
        fmt.Println("lastItem: ", lastItem)
    }
}

func (list *LinkedList) toString(separator string) string {
    listString := ""
    for item := list.sentinel.next; item != nil; item = item.next {
        listString += item.data
        //listString  += "," + fmt.Sprint(item.next)
        if item.next != nil {
            listString += separator
        }
    }
    return listString
}
