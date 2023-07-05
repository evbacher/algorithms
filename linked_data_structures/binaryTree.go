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
