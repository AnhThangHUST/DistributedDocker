package main
import "fmt"

type linkedListNode struct {
    Value string
    next_ptr *linkedListNode
}

type List struct{
    root *linkedListNode
    queue []*linkedListNode
}

func (l *List) createNode (value string) *linkedListNode{
    node := &linkedListNode{
        Value: value,
        next_ptr :nil,
    }
    return node
}

// if l.root didn't existed, we create new root
func (l *List) createRoot (value string){
    l.root = l.createNode(value)
}

// justfor interface
func (l *List) insertNode(value string){
    l.insertRecursive(value, l.root)
}

func (l *List) insertRecursive(value string, node *linkedListNode){
    if l.root == nil{
        l.createRoot(value)
        return
    }
    if node.next_ptr == nil{
        node.next_ptr = l.createNode(value)
    }else{
        if node.Value == value{
            fmt.Println("Node value have existed, no need to insert")
        }else{
            l.insertRecursive(value, node.next_ptr)
        }
    }
}
// implement recursive search
func (l *List) searchNode(value string, node *linkedListNode) (*linkedListNode){
    if l.queue != nil{
        l.queue = nil
    }
    for temp := node; temp !=nil; temp=temp.next_ptr{
        l.queue = append(l.queue, node)
        if temp.Value == value{
            return temp
        }
    }
    fmt.Println("No such node with this value exist")
    return nil
}

// deleteNode 
func (l *List) deleteNode(value string){
    node := l.searchNode(value, l.root)
    if node == nil{
        fmt.Println("this node even not exist so we can't delete it")
    }else{
        if node == l.root{
            l.root = l.root.next_ptr
        }else{
            parent_node := l.queue[len(l.queue)-2]
            parent_node.next_ptr = node.next_ptr
            node = nil
        }
    }
    l.printLinkedList(l.root)
}

func (l *List) printStructure(){
    l.printLinkedList(l.root)
}

func (l *List) printLinkedList(node *linkedListNode){
    if l.root == nil{
        fmt.Println("Our linked list has no member")
    }
    if node !=nil{
        if node.next_ptr == nil{
            fmt.Println(node.Value)
        }else{
            fmt.Printf("%s--->", node.Value)
        }
        l.printLinkedList(node.next_ptr)
    }
}

//
//func main(){
//    l := List{root:nil, queue:nil}
//    l.insertRecursive("mat", l.root)
//    l.insertRecursive("thang", l.root)
//    l.insertRecursive("thu", l.root)
//    l.insertRecursive("linh", l.root)
//    l.insertRecursive("abc", l.root)
//    l.insertRecursive("nga", l.root)
//    l.deleteNode("thang")
//    l.deleteNode("linh")
//    l.deleteNode("mat")
//}
