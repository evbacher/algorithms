/*
Sorted binary tree.
Builds on Binary tree and associated functions.
*/

package main

import (
    "fmt"
    "strings"
)

type Node struct {
    data    string
    left    *Node
    right   *Node
}

// Item for use in a doubly-linked list of nodes. data is *Node.
type Item struct {
    data    *Node
    prev    *Item
    next    *Item
}

// A doubly-linked list for building queues of Items.
type DoublyLinkedList struct {
    topSentinel     *Item
    bottomSentinel  *Item 
}

func main() {
    // dev testing
    devTest()
    
    // basic tree test
    //treeTest()
    
    // sorted tree test
    //sortedTreeTest()
}

// dev testing
func devTest() {
    // Make a root node to act as sentinel.
    root := Node { "", nil, nil }

    // Add some values.
    root.insertValue("I")
    root.insertValue("G")
    root.insertValue("C")
    root.insertValue("E")
    root.insertValue("B")
    root.insertValue("K")
    root.insertValue("S")
    root.insertValue("Q")
    root.insertValue("M")

    // Add F.
    root.insertValue("F")
    
    // Add 237
    root.insertValue("237")
    
    // Add some more things
    root.insertValue("abc")
    root.insertValue("ABC")
    root.insertValue("abcd")

    // If the tree has any values, they will start at root.right.
    if (root.right != nil) {
        fmt.Printf("Sorted values: %s\n", root.right.inorder())
    } else {
        fmt.Printf("root.right == nil\n")
    }
    
    fmt.Printf("tree: %s\n", root.inorder())
    
    // now, test find
    // Let the user search for values.
    for {
        // Get the target value.
        target := ""
        fmt.Printf("String: ")
        fmt.Scanln(&target)
        if len(target) == 0 { break }

        // Find the value's node.
        node := root.findValue(target)
        if node == nil {
            fmt.Printf("%s not found\n", target)
        } else {
            fmt.Printf("Found value %s\n", target)
        }
    }
}

func sortedTreeTest() {
    // Make a root node to act as sentinel.
    root := Node { "", nil, nil }

    // Add some values.
    root.insertValue("I")
    root.insertValue("G")
    root.insertValue("C")
    root.insertValue("E")
    root.insertValue("B")
    root.insertValue("K")
    root.insertValue("S")
    root.insertValue("Q")
    root.insertValue("M")

    // Add F.
    root.insertValue("F")

    // Display the values in sorted order.
    fmt.Printf("Sorted values: %s\n", root.right.inorder())

    // Let the user search for values.
    for {
        // Get the target value.
        target := ""
        fmt.Printf("String: ")
        fmt.Scanln(&target)
        if len(target) == 0 { break }

        // Find the value's node.
        node := root.findValue(target)
        if node == nil {
            fmt.Printf("%s not found\n", target)
        } else {
            fmt.Printf("Found value %s\n", target)
        }
    }
}


// test basic tree functionality
func treeTest() {
    tree := buildTree()
    fmt.Println(tree.displayIndented("  ", 0))
    fmt.Println("Preorder:     ", tree.preorder())
    fmt.Println("Inorder:      ", tree.inorder())
    fmt.Println("Postorder:    ", tree.postorder())
    fmt.Println("Breadth first:", tree.breadthFirst())
}

// *** Sorted binary tree additions. ***

// Creates a new Node and inserts it at the proper place in the sorted tree.
// If value already exists, no insertion takes place.
func (node *Node) insertValue(value string) {
    
    //fmt.Println("\nTHIS Node:", node)
    //fmt.Println("Trying to insert Node for", value)
    
    // display current tree (before insertion)
    //fmt.Println("current node:", node.data)
    //fmt.Printf("Current tree: %s\n", node.inorder())
    
    newNode := &Node{ value, nil, nil }
    //fmt.Println("newNode:", newNode)
    
    // I'm doing something stupid. I just don't know what it is yet.
    // (it was a } in the wrong place!)
    // Compare current Node data to value.
    current := node.data
    //fmt.Println("current:", current)
    //fmt.Println("node:", node)
    if current == "" {
        //fmt.Println("at root")
        //fmt.Println("value < current:", value < current)
        //fmt.Println("value > current:", value > current)
    }
    if value < current {
        //fmt.Println("node.left:", node.left)
        if node.left == nil {
            node.left = newNode
            //fmt.Println("BEFORE: node.left:", &node.left)
            //fmt.Println("node.left:", &node.left)
            //fmt.Println("AFTER: this node: ", node)
            return
        } else {
            //fmt.Println("\ncalling node.left(", value, ")")
            (node.left).insertValue(value)
            return
        }
    }
    
    if value > current {
        //fmt.Println("node.right:", node.right)
        if node.right == nil {
            node.right = newNode
            //fmt.Println("BEFORE: node.right:", &node.right)
            //fmt.Println("AFTER: this node: ", node)
            return
        } else {
            //fmt.Println("\ncalling node.right(", value, ")")
            (node.right).insertValue(value)
            return
        }
    }
}

// Finds a Node with data = value in the tree and returns the Node.
func (node *Node) findValue(value string) *Node {
    if value == node.data {
        return node
        
    // Check the left side.
    } else if value < node.data {
        if node.left != nil {
            return node.left.findValue(value)
        } else {
            return nil
        }
    } else {
        // value must be > node.data
        // Check the right side.
        if node.right != nil {
            return node.right.findValue(value)
        } else {
            return nil
        }
    }
}

// Builds a tree and returns the root Node.
func buildTree() *Node {
    
    // aNode
    aNode := Node { "A", nil, nil }
    bNode := Node { "B", nil, nil }
    cNode := Node { "C", nil, nil }
    aNode.left  = &bNode
    aNode.right = &cNode
    
    // bNode
    dNode := Node { "D", nil, nil }
    eNode := Node { "E", nil, nil }
    bNode.left =  &dNode
    bNode.right = &eNode
    
    // cNode
    fNode := Node { "F", nil, nil }
    cNode.right = &fNode
    
    // dNode is a leaf.
    
    // eNode
    gNode := Node { "G", nil, nil }
    eNode.left =  &gNode
    
    // fNode
    hNode := Node { "H", nil, nil }
    fNode.left =  &hNode
    
    // gNode is a leaf.
    
    // hNode
    iNode := Node { "I", nil, nil }
    jNode := Node { "J", nil, nil }
    hNode.left =  &iNode
    hNode.right =  &jNode
    
    // iNode is a leaf
    // jNode is a leaf
    
    // For this tree, aNode is the root.
    return &aNode
}

// *** Node traversal functions.***

// Traverses from the node to children in preorder (parent, left, right).
// Returns string representation of the tree from this node.
func (node *Node) displayIndented(indent string, depth int) string {
    result := strings.Repeat(indent, depth) + node.data + "\n"
    
    // Do the children.
    depth++
    if node.left != nil {
        result += (node.left).displayIndented(indent, depth)
    }
    if node.right != nil {
        result += (node.right).displayIndented(indent, depth)
    }
    
    return result
}

// Traverses from the node to children in preorder (parent, left, right).
// Returns string representation of the tree from this node.
// Test output: Preorder:      A B D E G C F H I J
func (node *Node) preorder() string {
    result := node.data

    // Do the children.
    if node.left != nil {
        result += " " + (node.left).preorder()
    }
    if node.right != nil {
        result += " " + (node.right).preorder()
    }
    
    return result
}

// Traverses from the node to children in order (left, parent, right).
// Returns string representation of the tree from this node.
// Test output: Inorder:       D B G E A C I H J F
func (node *Node) inorder() string {
    result := ""

    // Do left, parent, right.
    if node.left != nil {
        result += (node.left).inorder() + " "
    }
    result += node.data
    if node.right != nil {
        result += " " + (node.right).inorder()
    }
    
    return result
}

// Traverses from the node to children in post order (left, right, parent).
// Returns string representation of the tree from this node.
// Test output: Postorder:     D G E B I J H F C A
func (node *Node) postorder() string {
    result := ""

    // Do left, right, parent.
    if node.left != nil {
        result += (node.left).postorder() + " "
    }
    if node.right != nil {
        result += (node.right).postorder() + " "
    }
    result += node.data
    
    return result
}

// Traverses the nodes breadth-first. Each generation before the next.
// Returns string representation of the tree from this node.
// Test output: Breadth first: A B C D E F G H I J
func (node *Node) breadthFirst() string {
    
    // We'll use a DoublyLinkedList as a queue.
    // Think of q.enqueue as joining the back of the line.
    // q.dequeue is getting to the front of the line and leaving.
    // A queue is a FIFO (first in, first out) list.
    q := makeDoublyLinkedList()
    result := ""
    
    // Do parent node first. Add to the back of the queue.
    q.enqueue(node)
    
    // Now, take nodes from the front of the q if present, then add children, left to right.
    for !q.isEmpty() {
        node = q.dequeue()
        result += node.data
        // check child nodes and add to q if present
        if node.left != nil {
            q.enqueue(node.left)
        }
        if node.right != nil {
            q.enqueue(node.right)
        }
        if !q.isEmpty() {
            //fmt.Println("result:", result)
            result += " "
        }
    }
    return result
}

// *** We need a queue to do breadthFirst traverse. ***
// Will repurpose some other queue code, but should probably
// make a generic Queue to handle this. Exercise left for the reader (me).

// *** Some DoublyLinkedList functions, repurposed to hold *Node.***

// Creates a new DoublyLinkedList and initializes its sentinels to point to each other.
// Note that a doubly linked list should have no nil pointers (no end).
func makeDoublyLinkedList() DoublyLinkedList {
    // Create list and empty sentinels.
    list := DoublyLinkedList {}
    list.topSentinel = &Item {nil, nil, nil}
    list.bottomSentinel = &Item{nil, nil, nil}
    
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
func (list *DoublyLinkedList) push(node *Node) {
    // Create a new Item using inorder.
    item := Item{data: node}
    list.topSentinel.addAfter(&item)
}

// Removes the Item at the front of the list and returns its data.
func (list *DoublyLinkedList) pop() *Node {
    if list.isEmpty() {
        return nil
    }
    topItem := (list.topSentinel.next).delete()
    return topItem.data
}

// *** Queue functions (using a DoublyLinked List as the underlying data structure). ***

// Uses push() to add an Item to the queue.
func (list *DoublyLinkedList) enqueue(node *Node) {
    list.push(node)
}

// Removes the next Item ready to leave the queue
// (from the bottom of the underlying list).
func (list* DoublyLinkedList) dequeue() *Node {
    if list.isEmpty() {
        return nil
    }
    return (list.bottomSentinel).prev.delete().data
}
