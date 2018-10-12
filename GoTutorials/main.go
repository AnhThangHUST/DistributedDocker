package main
import "fmt"

func insertNode(value string, d dataStructure){
    d.insertNode(value)
}

func deleteNode(value string, d dataStructure){
    d.deleteNode(value)
}

func main(){
    a := &List{root:nil, queue:nil}
    insertNode("cuong", a)
    insertNode("thang", a)
    insertNode("thu",a )
    insertNode("linh", a)
    insertNode("abc", a)
    insertNode("nga", a)
    fmt.Println("      Original Linked List: ")
    a.printStructure()
    fmt.Printf("       After delete node thang:\n===> ")
    deleteNode("thang", a)
    fmt.Printf("       After delete node linh:\n===> ")
    deleteNode("linh",a )
    fmt.Printf("       After delete root:\n===> ")
    deleteNode("cuong",a)
    fmt.Println()

    b := &UBLtree{root:nil, queue:nil}
    insertNode("cuong", b)
    insertNode("thang", b)
    insertNode("thu", b)
    insertNode("linh", b)
    insertNode("abc", b)
    insertNode("nga", b)
    deleteNode("thang", b)
    deleteNode("linh", b)
    deleteNode("cuong",b)
    fmt.Println()
}
