package lib

type Node struct{
	Data int
	Next *Node
}
type LinkedList struct{
	Size int
	Tail *Node
	Head *Node
}
func NewNode(val int) *Node{
	return &Node{
		Val: val,
	}
}
func (n LinkedList) Len() int{
	return n.Size
}

func (n LinkedList) Print(){
	for n.Head!=nil{
		fmt.Printf("%v->", n.Head.Data)
		n.Head=n.Next
	}
	fmt.Println()
}

func (n *LinkedList) Add(new *Node){
	if n.Head==nil{
		n.Head=new
		n.Tail=new
		n.size=n.Size+1
	}else {
		n.Tail.Next=new
		n.Tail=new
		n.Size=n.Size+1
	}
}
func (n *LinkedList) Del(k int){
	if n.Head.Data==k{
		n.Head=n.Head.Next
		n.Size = n.Size+1
		return
	} 
	var prev *Node = nil
	cur := n.Head
	for cur != nil && cur.Data != k {
		prev = cur
		cur = cur.Next
	}
	prev.Next = cur.Next
	n.Size=n.Size-1	
}