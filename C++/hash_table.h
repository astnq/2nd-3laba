#ifndef HASH_TABLE_H
#define HASH_TABLE_H

#include "data_structures.h"

class OpenAddressingHashTable {
private:
    vector<pair<char, int>> table;
    vector<bool> occupied;
    int capacity;
    int size;

    int hash(char key, int attempt);

public:
    OpenAddressingHashTable(int cap = 256);
    void insert(char key, int value);
    bool search(char key, int& value);
    bool remove(char key);
    void clear();
    bool isEmpty() const;
    int getSize() const;
    int getCapacity() const;
    
    // Для сериализации
    vector<pair<char, int>> toVector() const;  // Уже есть
    void fromVector(const vector<pair<char, int>>& elements);  // Уже есть
    
    // НОВЫЕ МЕТОДЫ для бинарной сериализации
    void binarySerialize(const string& filename) const;
    void binaryDeserialize(const string& filename);
    
    // НОВЫЕ МЕТОДЫ для текстовой сериализации  
    void textSerialize(const string& filename) const;
    void textDeserialize(const string& filename);
};

#endif