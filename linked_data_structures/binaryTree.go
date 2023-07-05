/*
Binary tree and associated functions.
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
    test()
}

// dev testing
func test() {
    tree := buildTree()
    fmt.Println(tree.displayIndented("  ", 0))
    fmt.Println("Preorder:     ", tree.preorder())
    fmt.Println("Inorder:      ", tree.inorder())
    fmt.Println("Postorder:    ", tree.postorder())
    fmt.Println("Breadth first:", tree.breadthFirst())
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
    
    q := makeDoublyLinkedList()
    result := ""
    
    // Do parent node first.
    q.enqueue(node)
    for !q.isEmpty() {
        node = q.dequeue()
        result += node.data
        // check child nodes
        if node.left != nil {
            q.enqueue(node.left)
        }
        if node.right != nil {
            q.enqueue(node.right)
        }
        if !q.isEmpty() {
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

// *** Queue functions (for the bottom of the list). ***

// Uses push() to add an Item to the front of the list.
func (list *DoublyLinkedList) enqueue(node *Node) {
    list.push(node)
}

// Removes the Item at the bottom of the list and returns its string value.
func (list* DoublyLinkedList) dequeue() *Node {
    if list.isEmpty() {
        return nil
    }
    return (list.bottomSentinel).prev.delete().data
}

