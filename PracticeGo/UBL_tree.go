package main
import "fmt"

// This tree is unbalanced
// ==> we named node : UBLnode
type UBLtree struct{
    root *UBLnode
    queue []*UBLnode
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
    t.insertRecursive(value, t.root)
}

// if value < Value of leaf, make it as leaf's left child
// if value > Value of leaf, make it as leaf's right child
// if value == Value of leaf, no need to insert in tree
func (t *UBLtree) insertRecursive (value string, leaf *UBLnode){
    // YOUR CODE HERE
    if t.root == nil{
        t.createRoot(value)
        return
    }
    if value < leaf.Value{
        if leaf.Left != nil{
            t.insertRecursive(value, leaf.Left)
        }else{
            leaf.Left = t.createNode(value)
        }
    }else if value > leaf.Value{
        if leaf.Right != nil{
            t.insertRecursive(value, leaf.Right)
        }else{
            leaf.Right = t.createNode(value)
        }
    }else{
        fmt.Println("The value we want to insert have existed")
    }
    // END OF YOUR CODE
}

// recursive search
// queue save the path from root to value Node
func (t *UBLtree) searchNode (value string, leaf *UBLnode) (*UBLnode){
    // you should be sure queue is empty when save path
    if t.queue != nil{
        t.queue = nil
    }
    var recurSearch func(string, *UBLnode) (*UBLnode)
    // create anonymous function
    recurSearch = func(value string, leaf *UBLnode) (*UBLnode){
        //YOUR CODE HERE
        if leaf != nil{
            t.queue = append(t.queue, leaf)
            if value == leaf.Value{
                return leaf
            }
            if value<leaf.Value{
                return recurSearch(value, leaf.Left)
            }else{
                return recurSearch(value, leaf.Right)            }
        }else{
            fmt.Println("No such value exist")
            //reset queue
            return nil
        }
        //END OF YOUR CODE
    }
    return recurSearch(value, leaf)
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
            // YOUR CODE HERE
            t.root = swapLeaf(t.root)
            // END OF CODE
            fmt.Println("============> New root is: "+ t.root.Value)
            fmt.Printf("============> New tree: ")
            t.printTree(t.root)
        }else{
            fmt.Printf("\n\nOriginal Tree: ")
            t.printTree(t.root)
            fmt.Printf("\n       We delete leaf: %s (not root)     \n", leaf.Value)
            //YOUR CODE HERE
            parent_leaf = t.queue[len(t.queue) - 2]
            if parent_leaf.Left.Value == leaf.Value{
                parent_leaf.Left = swapLeaf(leaf)
            }
            if parent_leaf.Right.Value == leaf.Value{
                parent_leaf.Right = swapLeaf(leaf)
            }
            //END OF YOUR CODE
            fmt.Printf("============> New Tree: ")
            t.printTree(t.root)
        }
    }
}

// Please dont care about leaf's parent
// If leaf is remove, make child of leaf as an alternate
func swapLeaf (leaf *UBLnode) *UBLnode{
    //YOUR CODE HERE
    var temp_leaf *UBLnode
    if leaf.Left != nil{
        temp_leaf = leaf.Left
        for temp_leaf.Right != nil{
            temp_leaf = temp_leaf.Right
        }
        temp_leaf.Right = leaf.Right
    }else{
        if leaf.Right != nil{
            temp_leaf = leaf.Right
        }else{
            temp_leaf = nil
        }
    }
    leaf = nil
    return  temp_leaf
    //END OF YOUR CODE
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
//
//func main(){
//    t := UBLtree{root:nil, queue:nil}
//    t.insertRecursive("mat", t.root)
//    t.insertRecursive("thang", t.root)
//    t.insertRecursive("thu", t.root)
//    t.insertRecursive("linh", t.root)
//    t.insertRecursive("abc", t.root)
//    t.insertRecursive("nga", t.root)
//    t.deleteNode("thang")
//    t.deleteNode("linh")
//    t.deleteNode("mat")
//    fmt.Println()
//}
