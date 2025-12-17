#ifndef ARRAY_H
#define ARRAY_H

#include "data_structures.h"

class Array {
private:
    string *arr;
    size_t volume;
    size_t arr_size;

public:
    Array();
    ~Array();
    void ShowArray() const;
    void addToEnd(string value);
    void add(size_t index, string value);
    string getIndex(size_t index);
    void remove(size_t index);
    void replace(size_t index, string value);
    size_t getSize() const;
    
    // Для сериализации
    vector<string> toVector() const;
    void fromVector(const vector<string>& elements);
};

#endif