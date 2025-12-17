package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestStackBasic(t *testing.T) {
	stack := NewStack(5)
	stack.Push("test")
	val, err := stack.Pop()
	if err != nil || val != "test" {
		t.Errorf("Stack basic test failed")
	}
}

func TestArrayBasic(t *testing.T) {
	arr := NewArray()
	arr.AddToEnd("test")
	val, err := arr.GetIndex(0)
	if err != nil || val != "test" {
		t.Errorf("Array basic test failed")
	}
}

func TestBinaryTreeBasic(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	if tree.IsEmpty() {
		t.Errorf("Binary tree should not be empty")
	}
}

func TestStackCoverage(t *testing.T) {
	stack := NewStack(3)

	// Test empty
	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}

	// Test push/pop
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")

	if stack.Size() != 3 {
		t.Error("Stack size should be 3")
	}

	val, _ := stack.Peek()
	if val != "c" {
		t.Error("Peek should return 'c'")
	}

	val, _ = stack.Pop()
	if val != "c" {
		t.Error("Pop should return 'c'")
	}

	// Test overflow
	err := stack.Push("d")
	if err == nil {
		t.Error("Should get overflow error")
	}

	// Test underflow
	stack.Pop()
	stack.Pop()
	_, err = stack.Pop()
	if err == nil {
		t.Error("Should get underflow error")
	}

	// Test toVector/fromVector
	stack2 := NewStack(5)
	stack2.Push("x")
	stack2.Push("y")
	data := stack2.ToVector()

	stack3 := NewStack(5)
	stack3.FromVector(data)
	if stack3.Size() != 2 {
		t.Error("FromVector should load 2 elements")
	}
}

func TestArrayCoverage(t *testing.T) {
	arr := NewArray()

	// Add elements
	arr.AddToEnd("a")
	arr.AddToEnd("b")
	arr.AddToEnd("c")

	if arr.GetSize() != 3 {
		t.Error("Array size should be 3")
	}

	// Get element
	val, err := arr.GetIndex(1)
	if err != nil || val != "b" {
		t.Error("GetIndex should return 'b'")
	}

	// Insert
	arr.Add(1, "x")
	val, _ = arr.GetIndex(1)
	if val != "x" {
		t.Error("Add should insert at position")
	}

	// Remove
	arr.Remove(2)
	if arr.GetSize() != 3 {
		t.Error("Remove should decrease size")
	}

	// Replace
	arr.Replace(0, "z")
	val, _ = arr.GetIndex(0)
	if val != "z" {
		t.Error("Replace should change value")
	}

	// Error cases
	_, err = arr.GetIndex(10)
	if err == nil {
		t.Error("Should get index out of range")
	}

	// toVector/fromVector
	data := arr.ToVector()
	arr2 := NewArray()
	arr2.FromVector(data)
	if arr2.GetSize() != 3 {
		t.Error("FromVector should load elements")
	}
}

func TestBinaryTreeCoverage(t *testing.T) {
	tree := NewFullBinaryTree()

	// Empty tree
	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}

	// Insert and check order
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)

	// Should not be empty
	if tree.IsEmpty() {
		t.Error("Tree should not be empty after inserts")
	}

	// toVector should return sorted
	elements := tree.ToVector()
	if len(elements) != 5 {
		t.Error("Should have 5 elements")
	}

	// Check sorted
	for i := 1; i < len(elements); i++ {
		if elements[i] < elements[i-1] {
			t.Error("Elements should be sorted")
		}
	}

	// fromVector
	tree2 := NewFullBinaryTree()
	tree2.FromVector(elements)
	if tree2.IsEmpty() {
		t.Error("Tree should not be empty after FromVector")
	}

	// Print methods (just check they don't crash)
	tree.Print()
	tree.PrintZigZag()
}

func TestSinglyLinkedListBasic(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack("test")
	val, err := list.Get(0)
	if err != nil || val != "test" {
		t.Errorf("SinglyLinkedList basic test failed")
	}
}

func TestDoublyLinkedListBasic(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack("test")
	val, err := list.Get(0)
	if err != nil || val != "test" {
		t.Errorf("DoublyLinkedList basic test failed")
	}
}

func TestHashTableBasic(t *testing.T) {
	ht := NewOpenAddressingHashTable(10)
	ht.Insert('a', 1)
	val, found := ht.Search('a')
	if !found || val != 1 {
		t.Errorf("HashTable basic test failed")
	}
}

func TestSerializationBasic(t *testing.T) {
	serializer := NewSerializer()
	stack := NewStack(5)
	stack.Push("test")

	// Бинарная
	err := serializer.BinarySerializeStack(stack, "test_binary.dat")
	if err != nil {
		t.Errorf("BinarySerializeStack failed: %v", err)
	}
	os.Remove("test_binary.dat")

	// Текстовая
	err = serializer.TextSerializeStack(stack, "test_text.txt")
	if err != nil {
		t.Errorf("TextSerializeStack failed: %v", err)
	}
	os.Remove("test_text.txt")
}

func TestSinglyLinkedListCoverage(t *testing.T) {
	list := NewSinglyLinkedList()

	// Test empty
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}

	// Test push back
	list.PushBack("a")
	list.PushBack("b")
	list.PushBack("c")

	if list.Size() != 3 {
		t.Error("Size should be 3")
	}

	// Test push front
	list.PushFront("start")
	if list.Size() != 4 {
		t.Error("Size should be 4")
	}

	// Test get
	val, err := list.Get(0)
	if err != nil || val != "start" {
		t.Error("Get should return 'start'")
	}

	// Test insert
	list.Insert(2, "middle")
	if list.Size() != 5 {
		t.Error("Size should be 5")
	}

	// Test remove
	list.Remove(1)
	if list.Size() != 4 {
		t.Error("Size should be 4")
	}

	// Test pop front
	list.PopFront()
	if list.Size() != 3 {
		t.Error("Size should be 3")
	}

	// Test clear
	list.Clear()
	if !list.IsEmpty() {
		t.Error("List should be empty after clear")
	}

	// Test toVector/fromVector
	list2 := NewSinglyLinkedList()
	list2.PushBack("x")
	list2.PushBack("y")
	data := list2.ToVector()

	list3 := NewSinglyLinkedList()
	list3.FromVector(data)
	if list3.Size() != 2 {
		t.Error("FromVector should load 2 elements")
	}
}

func TestDoublyLinkedListCoverage(t *testing.T) {
	list := NewDoublyLinkedList()

	// Test empty
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}

	// Test push back/front
	list.PushBack("b")
	list.PushFront("a")
	list.PushBack("c")

	if list.Size() != 3 {
		t.Error("Size should be 3")
	}

	// Test get front/back
	front, _ := list.GetFront()
	back, _ := list.GetBack()
	if front != "a" || back != "c" {
		t.Error("GetFront/GetBack incorrect")
	}

	// Test insert
	list.Insert(1, "middle")
	if list.Size() != 4 {
		t.Error("Size should be 4")
	}

	// Test remove
	list.Remove(2)
	if list.Size() != 3 {
		t.Error("Size should be 3")
	}

	// Test pop front/back
	list.PopFront()
	list.PopBack()
	if list.Size() != 1 {
		t.Error("Size should be 1")
	}

	// Test clear
	list.Clear()
	if !list.IsEmpty() {
		t.Error("List should be empty after clear")
	}

	// Test toVector/fromVector
	list2 := NewDoublyLinkedList()
	list2.PushBack("x")
	list2.PushBack("y")
	data := list2.ToVector()

	list3 := NewDoublyLinkedList()
	list3.FromVector(data)
	if list3.Size() != 2 {
		t.Error("FromVector should load 2 elements")
	}
}

func TestHashTableCoverage(t *testing.T) {
	ht := NewOpenAddressingHashTable(10)

	// Test empty
	if !ht.IsEmpty() {
		t.Error("New hash table should be empty")
	}

	// Test insert/search
	ht.Insert('a', 1)
	ht.Insert('b', 2)
	ht.Insert('c', 3)

	if ht.GetSize() != 3 {
		t.Error("Size should be 3")
	}

	val, found := ht.Search('b')
	if !found || val != 2 {
		t.Error("Search should find 'b' with value 2")
	}

	// Test update
	ht.Insert('a', 100)
	val, found = ht.Search('a')
	if !found || val != 100 {
		t.Error("Insert should update existing key")
	}

	// Test remove
	removed := ht.Remove('b')
	if !removed || ht.GetSize() != 2 {
		t.Error("Remove should work")
	}

	// Test search non-existent
	_, found = ht.Search('x')
	if found {
		t.Error("Should not find non-existent key")
	}

	// Test remove non-existent
	removed = ht.Remove('x')
	if removed {
		t.Error("Should not remove non-existent key")
	}

	// Test clear
	ht.Clear()
	if !ht.IsEmpty() {
		t.Error("Hash table should be empty after clear")
	}

	// Test toVector/fromVector
	ht2 := NewOpenAddressingHashTable(10)
	ht2.Insert('x', 10)
	ht2.Insert('y', 20)
	data := ht2.ToVector()

	ht3 := NewOpenAddressingHashTable(10)
	ht3.FromVector(data)
	if ht3.GetSize() != 2 {
		t.Error("FromVector should load 2 elements")
	}
}

func TestSerializationCoverage(t *testing.T) {
	serializer := NewSerializer()

	// Test binary serialization
	stack1 := NewStack(5)
	stack1.Push("hello")
	stack1.Push("world")

	err := serializer.BinarySerializeStack(stack1, "test.dat")
	if err != nil {
		t.Errorf("BinarySerializeStack failed: %v", err)
	}
	defer os.Remove("test.dat")

	stack2 := NewStack(5)
	err = serializer.BinaryDeserializeStack(stack2, "test.dat")
	if err != nil {
		t.Errorf("BinaryDeserializeStack failed: %v", err)
	}

	if stack2.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack2.Size())
	}

	// Test text serialization
	stack3 := NewStack(5)
	stack3.Push("data1")
	stack3.Push("data2")

	err = serializer.TextSerializeStack(stack3, "test.txt")
	if err != nil {
		t.Errorf("TextSerializeStack failed: %v", err)
	}
	defer os.Remove("test.txt")

	stack4 := NewStack(5)
	err = serializer.TextDeserializeStack(stack4, "test.txt")
	if err != nil {
		t.Errorf("TextDeserializeStack failed: %v", err)
	}

	if stack4.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack4.Size())
	}

	// Test error cases
	err = serializer.BinaryDeserializeStack(stack1, "non_existent.dat")
	if err == nil {
		t.Error("Should get error for non-existent file")
	}

	err = serializer.TextDeserializeStack(stack1, "non_existent.txt")
	if err == nil {
		t.Error("Should get error for non-existent text file")
	}
}

func TestSinglyLinkedListEdgeCases(t *testing.T) {
	list := NewSinglyLinkedList()

	// Тест 1: Get с неверным индексом на пустом списке
	_, err := list.Get(0)
	if err == nil {
		t.Error("Should get error for Get on empty list")
	}

	// Тест 2: Remove с неверным индексом
	list.Remove(0) // Не должен падать

	// Тест 3: Insert с индексом больше размера
	list.Insert(100, "test")
	if !list.IsEmpty() {
		t.Error("Insert with large index should not add to empty list")
	}

	// Тест 4: Добавляем, удаляем, проверяем
	list.PushBack("a")
	list.PushBack("b")
	list.PushBack("c")

	// Удаляем из середины
	list.Remove(1)
	if list.Size() != 2 {
		t.Error("Size should be 2 after remove")
	}

	// Удаляем последний
	list.Remove(1)
	if list.Size() != 1 {
		t.Error("Size should be 1")
	}

	// Удаляем первый
	list.Remove(0)
	if !list.IsEmpty() {
		t.Error("List should be empty")
	}

	// Тест 5: Проверка tail после операций
	list.PushBack("x")
	list.PushBack("y")
	list.PushBack("z")

	list.Remove(2) // Удаляем последний
	val, _ := list.Get(list.Size() - 1)
	if val != "y" {
		t.Error("Tail should be 'y' after removing last")
	}

	// Тест 6: Clear и повторное использование
	list.Clear()
	list.PushBack("new")
	if list.Size() != 1 {
		t.Error("Should be able to reuse after clear")
	}

	// Тест 7: ToVector на пустом списке
	list.Clear()
	vec := list.ToVector()
	if len(vec) != 0 {
		t.Error("ToVector on empty list should return empty slice")
	}

	// Тест 8: FromVector с nil/empty
	list.FromVector(nil)
	list.FromVector([]string{})
	if !list.IsEmpty() {
		t.Error("FromVector with empty slice should result in empty list")
	}

	// Тест 9: Много элементов
	for i := 0; i < 100; i++ {
		list.PushBack(fmt.Sprintf("item_%d", i))
	}
	if list.Size() != 100 {
		t.Error("Should handle many elements")
	}
}

func TestDoublyLinkedListEdgeCases(t *testing.T) {
	list := NewDoublyLinkedList()

	// Тест 1: Все методы на пустом списке
	list.PopFront() // Не должен падать
	list.PopBack()  // Не должен падать

	_, err := list.GetFront()
	if err == nil {
		t.Error("GetFront should error on empty list")
	}

	_, err = list.GetBack()
	if err == nil {
		t.Error("GetBack should error on empty list")
	}

	_, err = list.Get(0)
	if err == nil {
		t.Error("Get should error on empty list")
	}

	// Тест 2: Insert граничные случаи
	list.Insert(0, "first") // Insert в пустой список
	if list.Size() != 1 {
		t.Error("Insert(0) should work on empty list")
	}

	list.Insert(1, "second") // Insert в конец
	if list.Size() != 2 {
		t.Error("Insert(size) should work as PushBack")
	}

	list.Insert(1, "middle") // Insert в середину
	if list.Size() != 3 {
		t.Error("Insert in middle should work")
	}

	// Тест 3: Проверка связей после Insert
	val1, _ := list.Get(0)
	val2, _ := list.Get(1)
	val3, _ := list.Get(2)
	if val1 != "first" || val2 != "middle" || val3 != "second" {
		t.Error("Insert messed up order")
	}

	// Тест 4: Remove граничные случаи
	list.Remove(2) // Удаляем последний
	if list.Size() != 2 {
		t.Error("Remove last should work")
	}

	list.Remove(0) // Удаляем первый
	if list.Size() != 1 {
		t.Error("Remove first should work")
	}

	list.Remove(0) // Удаляем единственный
	if !list.IsEmpty() {
		t.Error("Remove only element should work")
	}

	// Тест 5: Remove с неверным индексом
	list.Remove(100) // Не должен падать

	// Тест 6: Комплексная проверка связей
	list.PushBack("a")
	list.PushBack("b")
	list.PushBack("c")
	list.PushFront("start")

	// Удаляем из середины
	list.Remove(2)

	front, _ := list.GetFront()
	back, _ := list.GetBack()
	if front != "start" || back != "c" {
		t.Error("Front/Back incorrect after remove")
	}

	// Тест 7: ToVector/FromVector цикл
	data := list.ToVector()
	list2 := NewDoublyLinkedList()
	list2.FromVector(data)

	if list2.Size() != list.Size() {
		t.Error("FromVector/ToVector should preserve size")
	}

	// Тест 8: FromVector с пустым срезом
	list2.FromVector([]string{})
	if !list2.IsEmpty() {
		t.Error("FromVector empty should clear list")
	}

	// Тест 9: Много операций
	list.Clear()
	for i := 0; i < 50; i++ {
		list.PushBack(fmt.Sprintf("%d", i))
	}
	for i := 0; i < 25; i++ {
		list.PopFront()
	}
	for i := 0; i < 10; i++ {
		list.PopBack()
	}
	if list.Size() != 15 {
		t.Error("Complex operations failed")
	}
}

func TestArrayEdgeCases(t *testing.T) {
	arr := NewArray()

	// Тест 1: Все методы на пустом массиве
	arr.ShowArray() // Не должен падать

	_, err := arr.GetIndex(0)
	if err == nil {
		t.Error("GetIndex should error on empty array")
	}

	arr.Remove(0)          // Не должен падать
	arr.Replace(0, "test") // Не должен падать
	arr.Add(100, "test")   // Не должен падать

	// Тест 2: Добавление большого количества элементов (проверка перераспределения)
	for i := 0; i < 100; i++ {
		arr.AddToEnd(fmt.Sprintf("item_%d", i))
	}

	if arr.GetSize() != 100 {
		t.Error("Should handle many elements")
	}

	// Проверяем что все элементы на месте
	for i := 0; i < 100; i++ {
		val, err := arr.GetIndex(i)
		if err != nil || val != fmt.Sprintf("item_%d", i) {
			t.Errorf("Element %d incorrect", i)
		}
	}

	// Тест 3: Вставка в начало
	arr2 := NewArray()
	for i := 0; i < 5; i++ {
		arr2.AddToEnd(fmt.Sprintf("base_%d", i))
	}
	arr2.Add(0, "new_first")

	val, _ := arr2.GetIndex(0)
	if val != "new_first" {
		t.Error("Add to beginning failed")
	}

	// Тест 4: Вставка когда нужно перераспределение
	// Сначала заполним до capacity
	arr3 := NewArray()
	for i := 0; i < 15; i++ { // Изначально capacity = 10
		arr3.AddToEnd(fmt.Sprintf("fill_%d", i))
	}

	// Вставляем когда capacity должно увеличиться
	arr3.Add(5, "inserted_during_resize")

	val, _ = arr3.GetIndex(5)
	if val != "inserted_during_resize" {
		t.Error("Insert during resize failed")
	}

	// Тест 5: Удаление всех элементов
	for i := 0; i < arr3.GetSize(); i++ {
		arr3.Remove(0)
	}
	if arr3.GetSize() != 0 {
		t.Error("Should be empty after removing all")
	}

	// Тест 6: FromVector с большим срезом
	bigSlice := make([]string, 200)
	for i := range bigSlice {
		bigSlice[i] = fmt.Sprintf("from_vec_%d", i)
	}

	arr4 := NewArray()
	arr4.FromVector(bigSlice)

	if arr4.GetSize() != 200 {
		t.Error("FromVector should handle large slices")
	}

	// Тест 7: ToVector/FromVector цикл
	original := []string{"a", "b", "c", "d", "e"}
	arr5 := NewArray()
	arr5.FromVector(original)

	copied := arr5.ToVector()
	if len(copied) != len(original) {
		t.Error("ToVector should return all elements")
	}

	for i := range original {
		if copied[i] != original[i] {
			t.Error("ToVector/FromVector mismatch")
		}
	}

	// Тест 8: Замена несуществующего индекса
	arr5.Replace(100, "should_not_work")
	// Не должен падать

	// Тест 9: Получение несуществующего индекса
	_, err = arr5.GetIndex(100)
	if err == nil {
		t.Error("Should get error for out of bounds")
	}

	// Тест 10: Добавление с индексом > размера
	arr5.Add(100, "ignore")
	if arr5.GetSize() != 5 {
		t.Error("Add with large index should not add")
	}
}

func TestSerializationEdgeCases(t *testing.T) {
	serializer := NewSerializer()

	// Тест 1: Бинарная сериализация пустого стека
	stack1 := NewStack(5)
	err := serializer.BinarySerializeStack(stack1, "empty.dat")
	if err != nil {
		t.Errorf("Binary serialize empty failed: %v", err)
	}
	defer os.Remove("empty.dat")

	stack2 := NewStack(5)
	err = serializer.BinaryDeserializeStack(stack2, "empty.dat")
	if err != nil {
		t.Errorf("Binary deserialize empty failed: %v", err)
	}

	if !stack2.IsEmpty() {
		t.Error("Deserialized empty stack should be empty")
	}

	// Тест 2: Текстовая сериализация пустого стека
	err = serializer.TextSerializeStack(stack1, "empty.txt")
	if err != nil {
		t.Errorf("Text serialize empty failed: %v", err)
	}
	defer os.Remove("empty.txt")

	stack3 := NewStack(5)
	err = serializer.TextDeserializeStack(stack3, "empty.txt")
	if err != nil {
		t.Errorf("Text deserialize empty failed: %v", err)
	}

	if !stack3.IsEmpty() {
		t.Error("Deserialized empty stack should be empty")
	}

	// Тест 3: Стек с одним элементом
	stack4 := NewStack(5)
	stack4.Push("single")

	err = serializer.BinarySerializeStack(stack4, "single.dat")
	if err != nil {
		t.Errorf("Binary serialize single failed: %v", err)
	}
	defer os.Remove("single.dat")

	stack5 := NewStack(5)
	err = serializer.BinaryDeserializeStack(stack5, "single.dat")
	if err != nil {
		t.Errorf("Binary deserialize single failed: %v", err)
	}

	if stack5.Size() != 1 {
		t.Error("Should have one element")
	}

	val, _ := stack5.Pop()
	if val != "single" {
		t.Error("Element should be 'single'")
	}

	// Тест 4: Длинные строки
	stack6 := NewStack(5)
	longStr := strings.Repeat("very_long_string_", 100)
	stack6.Push(longStr)
	stack6.Push("short")

	err = serializer.BinarySerializeStack(stack6, "long.dat")
	if err != nil {
		t.Errorf("Binary serialize long string failed: %v", err)
	}
	defer os.Remove("long.dat")

	stack7 := NewStack(5)
	err = serializer.BinaryDeserializeStack(stack7, "long.dat")
	if err != nil {
		t.Errorf("Binary deserialize long string failed: %v", err)
	}

	if stack7.Size() != 2 {
		t.Error("Should have 2 elements")
	}

	// Тест 5: Ошибки файлов
	stack8 := NewStack(5)
	err = serializer.BinaryDeserializeStack(stack8, "non_existent_file_12345.dat")
	if err == nil {
		t.Error("Should get error for non-existent file")
	}

	err = serializer.TextDeserializeStack(stack8, "non_existent_file_12345.txt")
	if err == nil {
		t.Error("Should get error for non-existent text file")
	}

	// Тест 6: Многократная сериализация
	stack9 := NewStack(10)
	for i := 0; i < 5; i++ {
		stack9.Push(fmt.Sprintf("item_%d", i))
	}

	// Сериализуем несколько раз
	for i := 0; i < 3; i++ {
		filename := fmt.Sprintf("multi_%d.dat", i)
		err = serializer.BinarySerializeStack(stack9, filename)
		if err != nil {
			t.Errorf("Multi serialize %d failed: %v", i, err)
		}
		defer os.Remove(filename)

		stack10 := NewStack(10)
		err = serializer.BinaryDeserializeStack(stack10, filename)
		if err != nil {
			t.Errorf("Multi deserialize %d failed: %v", i, err)
		}

		if stack10.Size() != 5 {
			t.Errorf("Iteration %d: expected size 5, got %d", i, stack10.Size())
		}
	}

	// Тест 7: Проверка что файлы создаются
	testFile := "test_create.dat"
	stack11 := NewStack(3)
	stack11.Push("test")

	err = serializer.BinarySerializeStack(stack11, testFile)
	if err != nil {
		t.Errorf("Create test failed: %v", err)
	}
	defer os.Remove(testFile)

	// Проверяем что файл существует и не пустой
	if info, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Error("File was not created")
	} else if info.Size() == 0 {
		t.Error("File should not be empty")
	}
}

func TestSerializationEdgeCasesForCoverage(t *testing.T) {
	serializer := NewSerializer()

	// Тест 1: Ошибки записи (попытка записи в защищенную директорию)
	// Пропускаем так как требует прав суперпользователя

	// Тест 2: Стек с capacity 0 или 1
	stack1 := NewStack(1)
	stack1.Push("only")

	err := serializer.BinarySerializeStack(stack1, "tiny.dat")
	if err != nil {
		t.Errorf("Binary serialize tiny stack failed: %v", err)
	}
	defer os.Remove("tiny.dat")

	// Тест 3: Десериализация с поврежденным файлом
	// Создаем файл с неправильным форматом
	badFile := "bad_format.dat"
	f, err := os.Create(badFile)
	if err != nil {
		t.Fatalf("Cannot create bad file: %v", err)
	}
	// Пишем мусор
	f.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF})
	f.Close()
	defer os.Remove(badFile)

	stack2 := NewStack(5)
	err = serializer.BinaryDeserializeStack(stack2, badFile)
	// Может вернуть ошибку или проигнорировать - главное не упасть

	// Тест 4: Файл с правильным заголовком но без данных
	emptyDataFile := "empty_data.dat"
	f, err = os.Create(emptyDataFile)
	if err != nil {
		t.Fatalf("Cannot create empty data file: %v", err)
	}
	// Пишем размер 1, но не пишем данные
	binary.Write(f, binary.LittleEndian, uint64(1))
	f.Close()
	defer os.Remove(emptyDataFile)

	stack3 := NewStack(5)
	err = serializer.BinaryDeserializeStack(stack3, emptyDataFile)
	// Может вернуть ошибку EOF - это нормально

	// Тест 5: Текстовый файл с неправильным форматом
	badTextFile := "bad_text.txt"
	f, err = os.Create(badTextFile)
	if err != nil {
		t.Fatalf("Cannot create bad text file: %v", err)
	}
	f.WriteString("not_a_number\n") // Вместо размера - текст
	f.Close()
	defer os.Remove(badTextFile)

	stack4 := NewStack(5)
	err = serializer.TextDeserializeStack(stack4, badTextFile)
	// Должна быть ошибка парсинга

	// Тест 6: Файл с отрицательным размером (в текстовом формате)
	negSizeFile := "neg_size.txt"
	f, err = os.Create(negSizeFile)
	if err != nil {
		t.Fatalf("Cannot create negative size file: %v", err)
	}
	f.WriteString("-5\n") // Отрицательный размер
	f.Close()
	defer os.Remove(negSizeFile)

	stack5 := NewStack(5)
	err = serializer.TextDeserializeStack(stack5, negSizeFile)
	// Может обработать или вернуть ошибку

	// Тест 7: Большой размер в заголовке но маленький файл
	bigSizeFile := "big_size.dat"
	f, err = os.Create(bigSizeFile)
	if err != nil {
		t.Fatalf("Cannot create big size file: %v", err)
	}
	// Пишем очень большой размер
	binary.Write(f, binary.LittleEndian, uint64(1000000))
	f.Close()
	defer os.Remove(bigSizeFile)

	stack6 := NewStack(5)
	err = serializer.BinaryDeserializeStack(stack6, bigSizeFile)
	// Должен получить EOF ошибку

	// Тест 8: Последовательная сериализация разных стеков
	stacks := []*Stack{
		NewStack(3),
		NewStack(10),
		NewStack(1),
	}

	stacks[0].Push("a")
	stacks[1].Push("b")
	stacks[1].Push("c")
	// stacks[2] пустой

	for i, stack := range stacks {
		filename := fmt.Sprintf("multi_stack_%d.dat", i)
		err = serializer.BinarySerializeStack(stack, filename)
		if err != nil {
			t.Errorf("Multi stack %d binary serialize failed: %v", i, err)
		}
		defer os.Remove(filename)

		filenameTxt := fmt.Sprintf("multi_stack_%d.txt", i)
		err = serializer.TextSerializeStack(stack, filenameTxt)
		if err != nil {
			t.Errorf("Multi stack %d text serialize failed: %v", i, err)
		}
		defer os.Remove(filenameTxt)
	}

	// Тест 9: Сравнение бинарного и текстового форматов
	compareStack := NewStack(10)
	compareStack.Push("test1")
	compareStack.Push("test2")
	compareStack.Push("test3")

	// Сериализуем в оба формата
	err = serializer.BinarySerializeStack(compareStack, "compare.dat")
	if err != nil {
		t.Errorf("Compare binary serialize failed: %v", err)
	}
	defer os.Remove("compare.dat")

	err = serializer.TextSerializeStack(compareStack, "compare.txt")
	if err != nil {
		t.Errorf("Compare text serialize failed: %v", err)
	}
	defer os.Remove("compare.txt")

	// Десериализуем оба
	binaryStack := NewStack(10)
	err = serializer.BinaryDeserializeStack(binaryStack, "compare.dat")
	if err != nil {
		t.Errorf("Compare binary deserialize failed: %v", err)
	}

	textStack := NewStack(10)
	err = serializer.TextDeserializeStack(textStack, "compare.txt")
	if err != nil {
		t.Errorf("Compare text deserialize failed: %v", err)
	}

	// Должны быть одинаковыми
	if binaryStack.Size() != textStack.Size() {
		t.Error("Binary and text stacks should have same size")
	}

	// Тест 10: Проверка что сериализация не изменяет оригинальный стек
	original := NewStack(5)
	original.Push("original1")
	original.Push("original2")

	originalSize := original.Size()

	err = serializer.BinarySerializeStack(original, "preserve.dat")
	if err != nil {
		t.Errorf("Preserve serialize failed: %v", err)
	}
	defer os.Remove("preserve.dat")

	if original.Size() != originalSize {
		t.Error("Serialization should not change original stack size")
	}

	// Тест 11: Сериализация после FromVector
	vectorStack := NewStack(5)
	vectorStack.FromVector([]string{"from", "vector", "test"})

	err = serializer.BinarySerializeStack(vectorStack, "from_vector.dat")
	if err != nil {
		t.Errorf("FromVector serialize failed: %v", err)
	}
	defer os.Remove("from_vector.dat")

	// Тест 12: Десериализация в стек с недостаточной capacity
	smallStack := NewStack(2) // capacity меньше чем элементов
	err = serializer.BinaryDeserializeStack(smallStack, "from_vector.dat")
	// Должен увеличить capacity или вернуть ошибку
	// Главное - не упасть

	// Тест 13: Циклическая проверка toVector -> serialize -> deserialize -> toVector
	cycleStack := NewStack(10)
	for i := 0; i < 5; i++ {
		cycleStack.Push(fmt.Sprintf("cycle_%d", i))
	}

	originalVector := cycleStack.ToVector()

	err = serializer.BinarySerializeStack(cycleStack, "cycle.dat")
	if err != nil {
		t.Errorf("Cycle serialize failed: %v", err)
	}
	defer os.Remove("cycle.dat")

	newStack := NewStack(10)
	err = serializer.BinaryDeserializeStack(newStack, "cycle.dat")
	if err != nil {
		t.Errorf("Cycle deserialize failed: %v", err)
	}

	newVector := newStack.ToVector()

	if len(originalVector) != len(newVector) {
		t.Error("Cycle test: vectors should have same length")
	}

	for i := range originalVector {
		if originalVector[i] != newVector[i] {
			t.Errorf("Cycle test: element %d mismatch", i)
		}
	}

	// Тест 14: Проверка обработки ошибок в TextDeserializeStack
	// Создаем файл где указан размер но строк меньше
	shortFile := "short.txt"
	f, err = os.Create(shortFile)
	if err != nil {
		t.Fatalf("Cannot create short file: %v", err)
	}
	f.WriteString("3\n") // Размер 3
	f.WriteString("line1\n")
	f.WriteString("line2\n")
	// Нет третьей строки!
	f.Close()
	defer os.Remove(shortFile)

	shortStack := NewStack(5)
	err = serializer.TextDeserializeStack(shortStack, shortFile)
	// Может вернуть ошибку или прочитать только 2 строки

	// Тест 15: Файл с лишними строками
	extraFile := "extra.txt"
	f, err = os.Create(extraFile)
	if err != nil {
		t.Fatalf("Cannot create extra file: %v", err)
	}
	f.WriteString("2\n") // Размер 2
	f.WriteString("line1\n")
	f.WriteString("line2\n")
	f.WriteString("extra_line\n") // Лишняя строка
	f.Close()
	defer os.Remove(extraFile)

	extraStack := NewStack(5)
	err = serializer.TextDeserializeStack(extraStack, extraFile)
	if err != nil {
		t.Errorf("Text deserialize with extra lines failed: %v", err)
	}

	// Должен прочитать только 2 строки
	if extraStack.Size() != 2 {
		t.Errorf("Should read only 2 lines, got %d", extraStack.Size())
	}
}
