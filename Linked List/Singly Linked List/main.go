package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	// Add To Head
	linkedList.AddToHead(2)
	linkedList.AddToHead(1)
	// Add To End
	linkedList.AddToEnd(5)
	linkedList.AddToEnd(10)
	linkedList.IterateList()
	// Search
	linkedList.Search(10)
	// Delete number
	linkedList.DeleteByProperty(5)
	linkedList.IterateList()
	// Delete Head
	linkedList.DeleteHead()
	linkedList.IterateList()
	// Delete Last
	linkedList.DeleteLast()
	linkedList.IterateList()
}

func (linkedList *LinkedList) AddToHead(newProperty int) {
	var node = Node{property: newProperty, nextNode: linkedList.headNode}
	linkedList.headNode = &node
}

func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Printf("%v -> ", node.property)
	}
	fmt.Println()
}

func (linkedList *LinkedList) LastNode() *Node {
	if linkedList.headNode == nil {
		return nil
	}
	var node *Node
	for node = linkedList.headNode; node.nextNode != nil; node = node.nextNode {
		continue
	}
	return node
}

func (linkedList *LinkedList) AddToEnd(newProperty int) {
	var node = Node{property: newProperty}
	if linkedList.headNode == nil {
		linkedList.headNode = &node
	} else {
		var lastNode *Node = linkedList.LastNode()
		lastNode.nextNode = &node
	}
}

func (linkedList *LinkedList) Search(lookupProperty int) {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == lookupProperty {
			fmt.Printf("Property %v is found\n", lookupProperty)
			return
		}
	}
	fmt.Printf("Property %v is not found\n", lookupProperty)
}

func (linkedList *LinkedList) DeleteByProperty(deleteProperty int) {
	if linkedList.headNode.property == deleteProperty {
		linkedList.headNode = linkedList.headNode.nextNode
		return
	}
	var prevNode, currNode *Node
	prevNode, currNode = linkedList.headNode, linkedList.headNode.nextNode
	for ; currNode != nil; currNode = currNode.nextNode {
		if currNode.property == deleteProperty {
			prevNode.nextNode = currNode.nextNode
			return
		}
		prevNode = currNode
	}
	fmt.Printf("%v property is not found\n", deleteProperty)
}

func (linkedList *LinkedList) DeleteHead() {
	if linkedList.headNode == nil {
		fmt.Println("No property to delete")
		return
	}
	linkedList.headNode = linkedList.headNode.nextNode
}

func (linkedList *LinkedList) DeleteLast() {
	if linkedList.headNode == nil {
		fmt.Println("No property to delete")
		return
	}
	if linkedList.headNode.nextNode == nil {
		linkedList.DeleteHead()
	}
	var prevNode, currNode *Node
	prevNode, currNode = linkedList.headNode, linkedList.headNode.nextNode
	for ; currNode.nextNode != nil; currNode = currNode.nextNode {
		prevNode = currNode
	}
	prevNode.nextNode = nil
}
