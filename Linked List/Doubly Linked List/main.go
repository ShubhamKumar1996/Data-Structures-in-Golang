package main

import "fmt"

type Node struct {
	Property int
	PrevNode *Node
	NextNode *Node
}

type DoublyLinkedList struct {
	HeadNode *Node
	TailNode *Node
}

func main() {
	var doublyLinkedList DoublyLinkedList
	doublyLinkedList.AddToHead(10)
	doublyLinkedList.AddToHead(9)
	doublyLinkedList.IterateDoublyLinkedList()
	doublyLinkedList.AddToTail(18)
	doublyLinkedList.AddToTail(19)
	doublyLinkedList.AddToTail(20)
	doublyLinkedList.IterateDoublyLinkedList()
	doublyLinkedList.Search(19)
	doublyLinkedList.DeleteByProperty(18)
	doublyLinkedList.IterateDoublyLinkedList()
	doublyLinkedList.DeleteFromHead()
	doublyLinkedList.IterateDoublyLinkedList()
	doublyLinkedList.DeleteFromTail()
	doublyLinkedList.IterateDoublyLinkedList()
}

func (doublyLinkedList *DoublyLinkedList) AddToHead(newProperty int) {
	var node = Node{newProperty, nil, nil}
	if doublyLinkedList.HeadNode == nil && doublyLinkedList.TailNode == nil {
		doublyLinkedList.HeadNode = &node
		doublyLinkedList.TailNode = &node
		return
	}
	doublyLinkedList.HeadNode.PrevNode = &node
	node.NextNode = doublyLinkedList.HeadNode
	doublyLinkedList.HeadNode = &node
}

func (doublyLinkedList *DoublyLinkedList) IterateDoublyLinkedList() {
	var node *Node = doublyLinkedList.HeadNode
	for node != nil {
		fmt.Printf("%v -> ", node.Property)
		node = node.NextNode
	}
	fmt.Println()
}

func (doublyLinkedList *DoublyLinkedList) AddToTail(newProperty int) {
	var node = Node{newProperty, nil, nil}
	if doublyLinkedList.HeadNode == nil && doublyLinkedList.TailNode == nil {
		doublyLinkedList.HeadNode, doublyLinkedList.TailNode = &node, &node
		return
	}
	doublyLinkedList.TailNode.NextNode = &node
	node.PrevNode = doublyLinkedList.TailNode
	doublyLinkedList.TailNode = &node
}

func (doublyLinkedList *DoublyLinkedList) Search(lookupProperty int) {
	var node *Node = doublyLinkedList.HeadNode
	for node != nil {
		if node.Property == lookupProperty {
			fmt.Printf("Property %v is found\n", lookupProperty)
			return
		}
		node = node.NextNode
	}
	fmt.Printf("Property %v is not found\n", lookupProperty)
}

func (doublyLinkedList *DoublyLinkedList) DeleteByProperty(deleteProperty int) {
	if doublyLinkedList.HeadNode == nil {
		fmt.Println("No property found to be deleted")
		return
	}
	node := doublyLinkedList.HeadNode
	for node != nil {
		if node.Property == deleteProperty {
			if node.PrevNode == nil {
				doublyLinkedList.HeadNode = node.NextNode
				if doublyLinkedList.HeadNode != nil {
					doublyLinkedList.HeadNode.PrevNode = nil
				}
			} else {
				node.PrevNode.NextNode = node.NextNode

			}
			if node.NextNode == nil {
				doublyLinkedList.TailNode = node.PrevNode
				if doublyLinkedList.TailNode != nil {
					doublyLinkedList.TailNode.NextNode = nil
				}
			} else {
				node.NextNode.PrevNode = node.PrevNode
			}
			fmt.Printf("Property %v is deleted\n", deleteProperty)
			return
		}
		node = node.NextNode
	}
	fmt.Printf("Property %v is not found\n", deleteProperty)
}

func (doublyLinkedList *DoublyLinkedList) DeleteFromHead() {
	if doublyLinkedList.HeadNode == nil {
		fmt.Println("No property found to delete")
		return
	}
	nextNode := doublyLinkedList.HeadNode.NextNode
	doublyLinkedList.HeadNode = nextNode
	if nextNode != nil {
		nextNode.PrevNode = nil
	} else {
		doublyLinkedList.TailNode = nil
	}
}

func (doublyLinkedList *DoublyLinkedList) DeleteFromTail() {
	if doublyLinkedList.TailNode == nil {
		fmt.Println("No property found to delete")
		return
	}
	prevNode := doublyLinkedList.TailNode.PrevNode
	doublyLinkedList.TailNode = prevNode
	if prevNode != nil {
		prevNode.NextNode = nil
	} else {
		doublyLinkedList.HeadNode = nil
	}
}
