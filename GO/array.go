package main

import "fmt"

type Array struct {
	arr    []string
	volume int
}

func NewArray() *Array {
	return &Array{
		arr:    make([]string, 0, 10),
		volume: 10,
	}
}

func (a *Array) ShowArray() {
	for i := 0; i < len(a.arr); i++ {
		fmt.Print(a.arr[i], " ")
	}
	fmt.Println()
}

func (a *Array) AddToEnd(value string) {
	if len(a.arr) >= a.volume {
		a.volume *= 2
		newArr := make([]string, len(a.arr), a.volume)
		copy(newArr, a.arr)
		a.arr = newArr
	}
	a.arr = append(a.arr, value)
}

func (a *Array) Add(index int, value string) {
	if index > len(a.arr) {
		return
	}
	a.arr = append(a.arr[:index], append([]string{value}, a.arr[index:]...)...)
}

func (a *Array) GetIndex(index int) (string, error) {
	if index >= len(a.arr) {
		return "", fmt.Errorf("index out of range")
	}
	return a.arr[index], nil
}

func (a *Array) Remove(index int) {
	if index >= len(a.arr) {
		return
	}
	a.arr = append(a.arr[:index], a.arr[index+1:]...)
}

func (a *Array) Replace(index int, value string) {
	if index >= len(a.arr) {
		return
	}
	a.arr[index] = value
}

func (a *Array) GetSize() int {
	return len(a.arr)
}

func (a *Array) ToVector() []string {
	result := make([]string, len(a.arr))
	copy(result, a.arr)
	return result
}

func (a *Array) FromVector(elements []string) {
	a.volume = len(elements) + 10
	a.arr = make([]string, len(elements))
	copy(a.arr, elements)
}
