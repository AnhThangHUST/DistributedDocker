package main
import "fmt"

// This tree is unbalanced
// ==> we named node : UBLnode
type UBLtree struct{
    root *UBLnode
    queue []*UBLnode // This queue for save the link from root -> a leaf
}

type UBLnode struct{
    Value string
    Left *UBLnode
    Right *UBLnode
}


func (t *UBLtree)createNode(value string) *UBLnode{
    leaf := &UBLnode{
        Value: value,
        Left :nil,
        Right :nil,
    }
    return leaf
}

func (t *UBLtree) createRoot(value string) {
    t.root = t.createNode(value)
}

// test interface
func (t *UBLtree) insertNode(value string){
    /*Your Code Here
    
    End Of Your Code*/
}

// queue save the path from root to value Node
func (t *UBLtree) searchNode (value string, leaf *UBLnode) (*UBLnode){
    /*Your Code Here
    
    End Of Your Code*/
}

// first we search if value exist in binary tree
// then we assign chau chat cua leaf.Left se la parent cua leaf.Right
func (t *UBLtree) deleteNode (value string){
    var parent_leaf *UBLnode
    leaf := t.searchNode(value, t.root)
    if leaf == nil{
        fmt.Printf("this leaf we search even not exist")
    }else{
        if leaf == t.root{
            fmt.Printf("\n\nOriginal Tree: ")
            t.printTree(t.root)
            fmt.Printf("\n       We delete root : %s      \n", t.root.Value)
            /*Your Code Here

            End Of Your Code*/
            fmt.Println("============> New root is: "+ t.root.Value)
            fmt.Printf("============> New tree: ")
            t.printTree(t.root)
        }else{
            fmt.Printf("\n\nOriginal Tree: ")
            t.printTree(t.root)
            fmt.Printf("\n       We delete leaf: %s (not root)     \n", leaf.Value)
            /*Your Code Here

            End Of Your Code*/
            fmt.Printf("============> New Tree: ")
            t.printTree(t.root)
        }
    }
}

// Please dont care about leaf's parent
// If leaf is remove, make child of leaf as an alternate
func swapLeaf (leaf *UBLnode) *UBLnode{
    /*Your Code Here

    End Of Your Code */
}

func (t *UBLtree) printStructure (){
    t.printTree(t.root)
}
func (t *UBLtree) printTree(leaf *UBLnode){
    if t.root == nil{
        fmt.Println("This tree has no leaf")
    }else{
        if leaf!=nil{
            t.printTree(leaf.Left)
            fmt.Printf("--%s--", leaf.Value)
            t.printTree(leaf.Right)
        }
    }
}

func (t *UBLtree) printQueue(){
    fmt.Println(len(t.queue))
    for i:=0; i<len(t.queue); i++{
        fmt.Printf("-%s-", t.queue[i].Value)
    }
}
