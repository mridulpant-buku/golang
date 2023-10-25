package algorithms

//what is a package name?
import "fmt"

//define  LinkedListNode

// define DoubleLinkedList and the head pointer.
// also have the following methods defined like insertAtEnd, removeFromBeginning ,linkedListSize and printDll.
type LinkedListNode struct {
	data int
	prev *LinkedListNode
	next *LinkedListNode
} //what will it return? memory pointer ,right?

type DoubleLinkedList struct {
	Head *LinkedListNode //head stores the memory location of the first element of DoubleLinkedList and is of type LinkedListNode.
	Len  int
}

//initialize an emptyConstructor

func NewDoublyLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

// why a pointer here?
func (self *DoubleLinkedList) InsertAtEnd(data int) {
	temp := self.Head
	newNode := LinkedListNode{data, nil, nil}
	self.Len += 1
	if self.Head == nil {
		self.Head = &newNode //why & is used here
	} else {
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &newNode
		newNode.prev = temp
	}
	//return self.head
}

func (self *DoubleLinkedList) Display() {
	temp := self.Head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func (self *DoubleLinkedList) RemoveFromBeginning() int {
	temp := self.Head
	if temp != nil {
		self.Head = self.Head.next
		self.Head.prev = nil
	}
	self.Len -= 1
	return temp.data
}
