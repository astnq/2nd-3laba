package main

import (
	"fmt"
	"os"
)

func demoStack() {
	fmt.Println("=== Stack Demo ===")
	stack := NewStack(5)

	stack.Push("first")
	stack.Push("second")
	stack.Push("third")

	fmt.Printf("Stack size: %d\n", stack.Size())
	top, _ := stack.Peek()
	fmt.Printf("Top element: %s\n", top)

	for !stack.IsEmpty() {
		popped, _ := stack.Pop()
		fmt.Printf("Popped: %s\n", popped)
	}
	fmt.Println()
}

func demoArray() {
	fmt.Println("=== Array Demo ===")
	arr := NewArray()

	arr.AddToEnd("apple")
	arr.AddToEnd("banana")
	arr.AddToEnd("cherry")
	arr.Add(1, "blueberry")

	fmt.Print("Array contents: ")
	arr.ShowArray()
	fmt.Printf("Array size: %d\n\n", arr.GetSize())
}

func demoSinglyLinkedList() {
	fmt.Println("=== Singly Linked List Demo ===")
	list := NewSinglyLinkedList()

	list.PushBack("first")
	list.PushBack("third")
	list.Insert(1, "second")

	fmt.Print("List contents: ")
	for i := 0; i < list.Size(); i++ {
		val, _ := list.Get(i)
		fmt.Printf("%s ", val)
	}
	fmt.Printf("\nList size: %d\n\n", list.Size())
}

func demoDoublyLinkedList() {
	fmt.Println("=== Doubly Linked List Demo ===")
	list := NewDoublyLinkedList()

	list.PushBack("first")
	list.PushFront("very_first")
	list.PushBack("last")
	list.Insert(2, "middle")

	fmt.Print("List contents: ")
	for i := 0; i < list.Size(); i++ {
		val, _ := list.Get(i)
		fmt.Printf("%s ", val)
	}
	fmt.Printf("\nList size: %d\n", list.Size())
	
	front, _ := list.GetFront()
	back, _ := list.GetBack()
	fmt.Printf("Front: %s, Back: %s\n\n", front, back)
}

func demoBinaryTree() {
	fmt.Println("=== Binary Tree Demo ===")
	tree := NewFullBinaryTree()

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)

	fmt.Println("Tree structure:")
	tree.Print()

	fmt.Println("Zigzag traversal:")
	tree.PrintZigZag()
	fmt.Println()
}

func demoHashTable() {
	fmt.Println("=== Hash Table Demo ===")
	ht := NewOpenAddressingHashTable(10)

	ht.Insert('a', 1)
	ht.Insert('b', 2)
	ht.Insert('c', 3)

	if value, found := ht.Search('b'); found {
		fmt.Printf("Found 'b': %d\n", value)
	}

	ht.Remove('a')
	fmt.Printf("Hash table size: %d\n\n", ht.GetSize())
}

func demoSerialization() {
	fmt.Println("=== Serialization Demo ===")
	stack := NewStack(5)
	stack.Push("hello")
	stack.Push("world")

	serializer := NewSerializer()
	
	// Бинарная сериализация
	serializer.BinarySerializeStack(stack, "test_binary.dat")
	stack2 := NewStack(5)
	serializer.BinaryDeserializeStack(stack2, "test_binary.dat")
	
	fmt.Printf("Binary deserialized stack size: %d\n", stack2.Size())
	
	// Текстовая сериализация
	serializer.TextSerializeStack(stack, "test_text.txt")
	stack3 := NewStack(5)
	serializer.TextDeserializeStack(stack3, "test_text.txt")
	
	fmt.Printf("Text deserialized stack size: %d\n\n", stack3.Size())
	
	// Очистка
	os.Remove("test_binary.dat")
	os.Remove("test_text.txt")
}

func main() {
	fmt.Println("🚀 Go Data Structures Demo")

	demoStack()
	demoArray()
	demoSinglyLinkedList()
	demoDoublyLinkedList()
	demoBinaryTree()
	demoHashTable()
	demoSerialization()

	fmt.Println("✅ All demos completed successfully!")
}
