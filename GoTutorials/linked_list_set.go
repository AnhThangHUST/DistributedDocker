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

func (l *List) insertNode(value string){
    /*******Your Code Here*********/
}

func (l *List) searchNode(value string, node *linkedListNode) (*linkedListNode){
    /*******Your Code Here*********/
    return nil
}

/* deleteNode will use searchNode */
func (l *List) deleteNode(value string){
    /*******Your Code Here*********/
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
