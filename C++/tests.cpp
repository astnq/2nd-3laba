#include <gtest/gtest.h>
#include "data_structures.h"
#include "stack.h"
#include "array.h"
#include "singly_linked_list.h"
#include "doubly_linked_list.h"
#include "binary_tree.h"
#include "hash_table.h"
#include "serialization.h"

using namespace std;

// ==================== БАЗОВЫЕ ТЕСТЫ ====================

TEST(StackTest, Basic) {
    Stack stack(5);
    stack.push("test");
    EXPECT_EQ(stack.pop(), "test");
}

TEST(ArrayTest, Basic) {
    Array arr;
    arr.addToEnd("test");
    EXPECT_EQ(arr.getIndex(0), "test");
}

// ==================== ТЕСТЫ СЕРИАЛИЗАЦИИ STACK ====================

TEST(SerializationStackTest, BinaryBasic) {
    Stack stack1(5);
    stack1.push("hello");
    stack1.push("world");

    Serializer::binarySerialize(stack1, "test_binary.dat");
    
    Stack stack2(5);
    Serializer::binaryDeserialize(stack2, "test_binary.dat");
    
    EXPECT_EQ(stack2.size(), 2);
    EXPECT_EQ(stack2.pop(), "world");
    EXPECT_EQ(stack2.pop(), "hello");
}

TEST(SerializationStackTest, TextBasic) {
    Stack stack1(5);
    stack1.push("data1");
    stack1.push("data2");

    Serializer::textSerialize(stack1, "test_text.txt");
    
    Stack stack2(5);
    Serializer::textDeserialize(stack2, "test_text.txt");
    
    EXPECT_EQ(stack2.size(), 2);
    EXPECT_EQ(stack2.pop(), "data2");
    EXPECT_EQ(stack2.pop(), "data1");
}

TEST(SerializationStackTest, BinaryEmpty) {
    Stack stack1(5);
    Serializer::binarySerialize(stack1, "test_empty.dat");
    
    Stack stack2(5);
    Serializer::binaryDeserialize(stack2, "test_empty.dat");
    
    EXPECT_TRUE(stack2.isEmpty());
}

TEST(SerializationStackTest, TextEmpty) {
    Stack stack1(5);
    Serializer::textSerialize(stack1, "test_empty.txt");
    
    Stack stack2(5);
    Serializer::textDeserialize(stack2, "test_empty.txt");
    
    EXPECT_TRUE(stack2.isEmpty());
}

TEST(SerializationStackTest, BinaryMultiple) {
    Stack stack1(10);
    for (int i = 0; i < 5; i++) {
        stack1.push("item_" + to_string(i));
    }

    Serializer::binarySerialize(stack1, "test_multiple.dat");
    
    Stack stack2(10);
    Serializer::binaryDeserialize(stack2, "test_multiple.dat");
    
    EXPECT_EQ(stack2.size(), 5);
    for (int i = 4; i >= 0; i--) {
        EXPECT_EQ(stack2.pop(), "item_" + to_string(i));
    }
}

TEST(SerializationStackTest, TextMultiple) {
    Stack stack1(10);
    for (int i = 0; i < 3; i++) {
        stack1.push("elem_" + to_string(i));
    }

    Serializer::textSerialize(stack1, "test_multiple.txt");
    
    Stack stack2(10);
    Serializer::textDeserialize(stack2, "test_multiple.txt");
    
    EXPECT_EQ(stack2.size(), 3);
    for (int i = 2; i >= 0; i--) {
        EXPECT_EQ(stack2.pop(), "elem_" + to_string(i));
    }
}

TEST(SerializationStackTest, BinarySpecialChars) {
    Stack stack1(5);
    stack1.push("hello\nworld");
    stack1.push("tab\tdata");
    stack1.push("quote\"test");

    Serializer::binarySerialize(stack1, "test_special.dat");
    
    Stack stack2(5);
    Serializer::binaryDeserialize(stack2, "test_special.dat");
    
    EXPECT_EQ(stack2.size(), 3);
    EXPECT_EQ(stack2.pop(), "quote\"test");
    EXPECT_EQ(stack2.pop(), "tab\tdata");
    EXPECT_EQ(stack2.pop(), "hello\nworld");
}

/*TEST(SerializationStackTest, TextSpecialChars) {
    Stack stack1(5);
    stack1.push("line1\nline2");
    stack1.push("data with spaces");
    stack1.push("end");

    Serializer::textSerialize(stack1, "test_special.txt");
    
    Stack stack2(5);
    Serializer::textDeserialize(stack2, "test_special.txt");
    
    EXPECT_EQ(stack2.size(), 3);
    EXPECT_EQ(stack2.pop(), "end");
    EXPECT_EQ(stack2.pop(), "data with spaces");
    EXPECT_EQ(stack2.pop(), "line1\nline2");
}
*/
TEST(HashTableSerializationTest, BinaryBasic) {
    OpenAddressingHashTable ht1(10);
    ht1.insert('a', 100);
    ht1.insert('b', 200);
    ht1.insert('c', 300);
    
    Serializer::binarySerialize(ht1, "test_hash_binary.dat");
    
    OpenAddressingHashTable ht2;
    Serializer::binaryDeserialize(ht2, "test_hash_binary.dat");
    
    EXPECT_EQ(ht2.getSize(), 3);
    
    int value;
    EXPECT_TRUE(ht2.search('a', value));
    EXPECT_EQ(value, 100);
    EXPECT_TRUE(ht2.search('b', value));
    EXPECT_EQ(value, 200);
    EXPECT_TRUE(ht2.search('c', value));
    EXPECT_EQ(value, 300);
}

TEST(HashTableSerializationTest, TextBasic) {
    OpenAddressingHashTable ht1(5);
    ht1.insert('x', 10);
    ht1.insert('y', 20);
    ht1.insert('z', 30);
    
    Serializer::textSerialize(ht1, "test_hash_text.txt");
    
    OpenAddressingHashTable ht2;
    Serializer::textDeserialize(ht2, "test_hash_text.txt");
    
    EXPECT_EQ(ht2.getSize(), 3);
    
    int value;
    EXPECT_TRUE(ht2.search('x', value));
    EXPECT_EQ(value, 10);
    EXPECT_TRUE(ht2.search('y', value));
    EXPECT_EQ(value, 20);
    EXPECT_TRUE(ht2.search('z', value));
    EXPECT_EQ(value, 30);
}



TEST(SerializationStackTest, BinaryFullStack) {
    Stack stack1(3);
    stack1.push("one");
    stack1.push("two");
    stack1.push("three");

    Serializer::binarySerialize(stack1, "test_full.dat");
    
    Stack stack2(3);
    Serializer::binaryDeserialize(stack2, "test_full.dat");
    
    EXPECT_EQ(stack2.size(), 3);
    EXPECT_EQ(stack2.pop(), "three");
    EXPECT_EQ(stack2.pop(), "two");
    EXPECT_EQ(stack2.pop(), "one");
}

TEST(SerializationStackTest, TextFullStack) {
    Stack stack1(3);
    stack1.push("a");
    stack1.push("b");
    stack1.push("c");

    Serializer::textSerialize(stack1, "test_full.txt");
    
    Stack stack2(3);
    Serializer::textDeserialize(stack2, "test_full.txt");
    
    EXPECT_EQ(stack2.size(), 3);
    EXPECT_EQ(stack2.pop(), "c");
    EXPECT_EQ(stack2.pop(), "b");
    EXPECT_EQ(stack2.pop(), "a");
}



// Array дополнительные тесты
TEST(ArrayTest, Comprehensive) {
    Array arr;
    
    // Тест множественных операций
    for (int i = 0; i < 20; i++) {
        arr.addToEnd("item_" + to_string(i));
    }
    EXPECT_EQ(arr.getSize(), 20);
    
    // Тест вставки
    arr.add(5, "inserted");
    EXPECT_EQ(arr.getIndex(5), "inserted");
    
    // Тест замены
    arr.replace(10, "replaced");
    EXPECT_EQ(arr.getIndex(10), "replaced");
    
    // Тест удаления
    arr.remove(7);
    EXPECT_EQ(arr.getSize(), 20); // 20 - 1 + 1 вставка = 20
}

// BinaryTree дополнительные тесты
TEST(FullBinaryTreeTest, Comprehensive) {
    FullBinaryTree tree;
    
    // Тест множественных вставок
    vector<int> values = {50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45};
    for (int val : values) {
        tree.insert(val);
    }
    
    // Тест toVector
    vector<int> elements = tree.toVector();
    EXPECT_EQ(elements.size(), values.size());
    
    // Тест что элементы отсортированы
    EXPECT_TRUE(is_sorted(elements.begin(), elements.end()));
}

// SinglyLinkedList дополнительные тесты
TEST(SinglyLinkedListTest, Comprehensive) {
    SinglyLinkedList list;
    
    // Комплексный тест всех операций
    list.push_back("a");
    list.push_front("start");
    list.push_back("b");
    list.insert(2, "middle");
    list.push_back("end");
    
    EXPECT_EQ(list.size(), 5);
    EXPECT_EQ(list.get(0), "start");
    EXPECT_EQ(list.get(2), "middle");
    EXPECT_EQ(list.get(4), "end");
    
    // Тест очистки
    list.clear();
    EXPECT_TRUE(list.isEmpty());
}

// DoublyLinkedList дополнительные тесты
TEST(DoublyLinkedListTest, Comprehensive) {
    DoublyLinkedList list;
    
    // Тест всех операций
    list.push_back("first");
    list.push_front("very_first");
    list.push_back("last");
    list.insert(2, "middle");
    
    EXPECT_EQ(list.size(), 4);
    EXPECT_EQ(list.get_front(), "very_first");
    EXPECT_EQ(list.get_back(), "last");
    EXPECT_EQ(list.get(2), "middle");
    
    // Тест последовательного удаления
    list.pop_front();
    EXPECT_EQ(list.get_front(), "first");
    
    list.pop_back();
    EXPECT_EQ(list.get_back(), "middle");
}

// HashTable дополнительные тесты
TEST(HashTableTest, Comprehensive) {
    OpenAddressingHashTable ht(50);
    
    // Тест множественных операций
    for (char c = 'a'; c <= 'z'; c++) {
        ht.insert(c, int(c));
    }
    
    // Тест поиска всех элементов
    int value;
    for (char c = 'a'; c <= 'z'; c++) {
        EXPECT_TRUE(ht.search(c, value));
        EXPECT_EQ(value, int(c));
    }
    
    // Тест toVector
    vector<pair<char, int>> data = ht.toVector();
    EXPECT_FALSE(data.empty());
}

// Stack дополнительные тесты
TEST(StackTest, Comprehensive) {
    Stack stack(10);
    
    // Тест множественных операций
    for (int i = 0; i < 8; i++) {
        stack.push("val_" + to_string(i));
    }
    
    EXPECT_EQ(stack.size(), 8);
    
    // Тест toVector/fromVector
    vector<string> data = stack.toVector();
    Stack stack2(10);
    stack2.fromVector(data);
    
    EXPECT_EQ(stack2.size(), 8);
    for (int i = 7; i >= 0; i--) {
        EXPECT_EQ(stack2.pop(), "val_" + to_string(i));
    }
}

// ==================== ФИНАЛЬНЫЕ ТЕСТЫ ДЛЯ 85%+ COVERAGE ====================

// BinaryTree - дополнительные тесты для print методов
TEST(FullBinaryTreeTest, PrintCoverage) {
    FullBinaryTree tree;
    
    // Тест пустого дерева
    EXPECT_NO_THROW(tree.print());
    EXPECT_NO_THROW(tree.printZigZag());
    
    // Тест с данными
    tree.insert(5);
    tree.insert(3);
    tree.insert(7);
    
    EXPECT_NO_THROW(tree.print());
    EXPECT_NO_THROW(tree.printZigZag());
}

// Array - тесты граничных случаев
TEST(ArrayTest, EdgeCases) {
    Array arr;
    
    // Тест пустого массива
    EXPECT_NO_THROW(arr.ShowArray());
    EXPECT_EQ(arr.getSize(), 0);
    
    // Тест добавления в начало пустого массива
    arr.add(0, "first");
    EXPECT_EQ(arr.getSize(), 1);
    EXPECT_EQ(arr.getIndex(0), "first");
    
    // Тест fromVector с пустым вектором
    vector<string> empty_vec;
    arr.fromVector(empty_vec);
    EXPECT_EQ(arr.getSize(), 0);
}

// HashTable - тесты граничных случаев и всех методов
TEST(HashTableTest, EdgeCases) {
    OpenAddressingHashTable ht(5);
    
    // Тест очистки пустой таблицы
    ht.clear();
    EXPECT_TRUE(ht.isEmpty());
    
    // Тест поиска в пустой таблице
    int value;
    EXPECT_FALSE(ht.search('x', value));
    
    // Тест удаления из пустой таблицы
    EXPECT_FALSE(ht.remove('x'));
    
    // Тест переполнения
    ht.insert('a', 1);
    ht.insert('b', 2);
    ht.insert('c', 3);
    ht.insert('d', 4);
    ht.insert('e', 5);
    ht.insert('f', 6); // должно обработать переполнение
    
    // Тест toVector/fromVector
    vector<pair<char, int>> data = ht.toVector();
    OpenAddressingHashTable ht2;
    ht2.fromVector(data);
    
    EXPECT_TRUE(ht2.search('a', value));
    EXPECT_EQ(value, 1);
}

// SinglyLinkedList - тесты граничных случаев
TEST(SinglyLinkedListTest, EdgeCases) {
    SinglyLinkedList list;
    
    // Тест toVector на пустом списке
    vector<string> empty = list.toVector();
    EXPECT_TRUE(empty.empty());
    
    // Тест fromVector
    vector<string> data = {"from", "vector", "test"};
    list.fromVector(data);
    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.get(1), "vector");
    
    // Тест удаления из пустого списка
    SinglyLinkedList empty_list;
    EXPECT_NO_THROW(empty_list.remove(0));
    EXPECT_NO_THROW(empty_list.pop_front());
}

// DoublyLinkedList - тесты граничных случаев
TEST(DoublyLinkedListTest, EdgeCases) {
    DoublyLinkedList list;
    
    // Тест toVector на пустом списке
    vector<string> empty = list.toVector();
    EXPECT_TRUE(empty.empty());
    
    // Тест fromVector
    vector<string> data = {"p", "q", "r"};
    list.fromVector(data);
    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.get_front(), "p");
    EXPECT_EQ(list.get_back(), "r");
    
    // Тест граничных индексов
    EXPECT_THROW(list.get(10), out_of_range);
    EXPECT_THROW(list.get(-1), out_of_range);
}

// Stack - финальные тесты
TEST(StackTest, EdgeCases) {
    Stack stack(2);
    
    // Тест toVector на пустом стеке
    vector<string> empty = stack.toVector();
    EXPECT_TRUE(empty.empty());
    
    // Тест fromVector с пустым вектором
    stack.fromVector(empty);
    EXPECT_TRUE(stack.isEmpty());
    
    // Тест циклических операций
    vector<string> data = {"cycle", "test"};
    stack.fromVector(data);
    vector<string> result = stack.toVector();
    EXPECT_EQ(data, result);
}

// ==================== ФИНАЛЬНЫЕ ТЕСТЫ ДЛЯ 85%+ ====================

TEST(DoublyLinkedListTest, FullCoverage) {
    DoublyLinkedList list;
    
    cout << "=== DEBUG START ===" << endl;
    
    // Тест сложной последовательности операций
    list.push_front("start");
    cout << "1. push_front('start'): ";
    for (size_t i = 0; i < list.size(); i++) cout << "[" << i << "]=" << list.get(i) << " ";
    cout << endl;
    
    list.push_back("end");
    cout << "2. push_back('end'): ";
    for (size_t i = 0; i < list.size(); i++) cout << "[" << i << "]=" << list.get(i) << " ";
    cout << endl;
    
    list.insert(1, "middle1");
    cout << "3. insert(1, 'middle1'): ";
    for (size_t i = 0; i < list.size(); i++) cout << "[" << i << "]=" << list.get(i) << " ";
    cout << endl;
    
    list.insert(2, "middle2");
    cout << "4. insert(2, 'middle2'): ";
    for (size_t i = 0; i < list.size(); i++) cout << "[" << i << "]=" << list.get(i) << " ";
    cout << endl;
    
    list.push_front("new_start");
    cout << "5. push_front('new_start'): ";
    for (size_t i = 0; i < list.size(); i++) cout << "[" << i << "]=" << list.get(i) << " ";
    cout << endl;
    
    list.push_back("new_end");
    cout << "6. push_back('new_end'): ";
    for (size_t i = 0; i < list.size(); i++) cout << "[" << i << "]=" << list.get(i) << " ";
    cout << endl;
    
    cout << "7. pop_front() удаляет 'new_start'" << endl;
    list.pop_front();
    cout << "   После pop_front: ";
    for (size_t i = 0; i < list.size(); i++) cout << "[" << i << "]=" << list.get(i) << " ";
    cout << endl;
    
    cout << "8. pop_back() удаляет 'new_end'" << endl;
    list.pop_back();
    cout << "   После pop_back: ";
    for (size_t i = 0; i < list.size(); i++) cout << "[" << i << "]=" << list.get(i) << " ";
    cout << endl;
    
    cout << "9. remove(2) должен удалить 'middle1'" << endl;
    list.remove(2);
    cout << "   После remove(2): ";
    for (size_t i = 0; i < list.size(); i++) cout << "[" << i << "]=" << list.get(i) << " ";
    cout << endl;
    
    cout << "=== DEBUG END ===" << endl;
    
    // Оставь оригинальные проверки
    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.get_front(), "start");
    EXPECT_EQ(list.get_back(), "end");
    EXPECT_EQ(list.get(1), "middle2");
}

// SinglyLinkedList - добиваем до 85%+
TEST(SinglyLinkedListTest, FullCoverage) {
    SinglyLinkedList list;
    
    // Комплексный тест всех возможных операций
    list.push_back("a");
    list.push_front("before_a");
    list.push_back("c");
    list.insert(2, "b");
    list.push_front("very_first");
    list.push_back("very_last");
    
    EXPECT_EQ(list.size(), 6);
    EXPECT_EQ(list.get(0), "very_first");
    EXPECT_EQ(list.get(2), "a");
    EXPECT_EQ(list.get(4), "c");
    EXPECT_EQ(list.get(5), "very_last");
    
    // Тест последовательного удаления
    list.pop_front(); // very_first
    EXPECT_EQ(list.get(0), "before_a");
    
    list.remove(3); // c
    EXPECT_EQ(list.get(3), "very_last");
    
    list.remove(0); // before_a
    EXPECT_EQ(list.get(0), "a");
    
    // Тест граничных случаев удаления
    while (!list.isEmpty()) {
        list.pop_front();
    }
    EXPECT_TRUE(list.isEmpty());
    
    // Тест fromVector с большим количеством элементов
    vector<string> big_data;
    for (int i = 0; i < 20; i++) {
        big_data.push_back("big_" + to_string(i));
    }
    list.fromVector(big_data);
    EXPECT_EQ(list.size(), 20);
    EXPECT_EQ(list.get(15), "big_15");
}

// Stack - добиваем до 85%+
TEST(StackTest, FullCoverage) {
    Stack stack(5);
    
    // Тест полного цикла операций
    stack.push("first");
    stack.push("second");
    stack.push("third");
    
    EXPECT_EQ(stack.size(), 3);
    EXPECT_EQ(stack.peek(), "third");
    
    // Тест toVector/fromVector полного цикла
    vector<string> data1 = stack.toVector();
    Stack stack2(5);
    stack2.fromVector(data1);
    
    EXPECT_EQ(stack2.size(), 3);
    EXPECT_EQ(stack2.peek(), "third");
    
    // Тест последовательного извлечения
    EXPECT_EQ(stack2.pop(), "third");
    EXPECT_EQ(stack2.pop(), "second");
    EXPECT_EQ(stack2.pop(), "first");
    EXPECT_TRUE(stack2.isEmpty());
    
    // Тест edge cases
    Stack empty_stack(3);
    vector<string> empty_data = empty_stack.toVector();
    EXPECT_TRUE(empty_data.empty());
    
    empty_stack.fromVector(empty_data);
    EXPECT_TRUE(empty_stack.isEmpty());
    
    // Тест переполнения
    Stack small_stack(2);
    small_stack.push("one");
    small_stack.push("two");
    EXPECT_THROW(small_stack.push("three"), overflow_error);
}

// BinaryTree - финальные тесты для 90%+
TEST(FullBinaryTreeTest, PrintAndEdgeCases) {
    FullBinaryTree tree;
    
    // Тест всех print методов на пустом дереве
    EXPECT_NO_THROW(tree.print());
    EXPECT_NO_THROW(tree.printZigZag());
    
    // Тест с одним элементом
    tree.insert(42);
    EXPECT_NO_THROW(tree.print());
    EXPECT_NO_THROW(tree.printZigZag());
    
    // Тест с несколькими элементами
    tree.insert(20);
    tree.insert(60);
    tree.insert(10);
    tree.insert(30);
    tree.insert(50);
    tree.insert(70);
    
    // Тест всех методов
    EXPECT_FALSE(tree.isEmpty());
    vector<int> elements = tree.toVector();
    EXPECT_EQ(elements.size(), 7);
    EXPECT_TRUE(is_sorted(elements.begin(), elements.end()));
    
    // Тест print методов на заполненном дереве
    EXPECT_NO_THROW(tree.print());
    EXPECT_NO_THROW(tree.printZigZag());
}

// HashTable - финальные тесты для 90%+
TEST(HashTableTest, StressAndEdgeCases) {
    OpenAddressingHashTable ht(50);
    
    // Тест множественных операций с коллизиями
    for (int i = 0; i < 26; i++) {
        char key = 'a' + i;
        ht.insert(key, i * 10);
    }
    
    // Проверка всех элементов
    int value;
    for (int i = 0; i < 26; i++) {
        char key = 'a' + i;
        EXPECT_TRUE(ht.search(key, value));
        EXPECT_EQ(value, i * 10);
    }
    
    // Тест обновления значений
    ht.insert('a', 999);
    EXPECT_TRUE(ht.search('a', value));
    EXPECT_EQ(value, 999);
    
    // Тест удаления и повторной вставки
    EXPECT_TRUE(ht.remove('b'));
    EXPECT_FALSE(ht.search('b', value));
    
    ht.insert('b', 888);
    EXPECT_TRUE(ht.search('b', value));
    EXPECT_EQ(value, 888);
    
    // Тест toVector/fromVector
    vector<pair<char, int>> data = ht.toVector();
    OpenAddressingHashTable ht2;
    ht2.fromVector(data);
    
    EXPECT_TRUE(ht2.search('z', value));
    EXPECT_EQ(value, 25 * 10);
}


int main(int argc, char** argv) {
    testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
