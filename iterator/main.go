package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type SliceIterator struct {
	currentPosition int
	items           []interface{}
}

func NewSliceIterator(items []interface{}) *SliceIterator {
	return &SliceIterator{currentPosition: 0, items: items}
}

func (s *SliceIterator) HasNext() bool {
	return s.currentPosition <= len(s.items)-1
}

func (s *SliceIterator) Next() interface{} {
	item := s.items[s.currentPosition]
	s.currentPosition++
	return item
}

type MapIterator struct {
	currentPosition int
	m               map[interface{}]interface{}
	values          []interface{}
}

func NewMapIterator(m map[interface{}]interface{}) *MapIterator {
	values := []interface{}{}
	for _, v := range m {
		values = append(values, v)
	}
	return &MapIterator{currentPosition: 0, m: m, values: values}
}

func (s *MapIterator) HasNext() bool {
	return s.currentPosition <= len(s.values)-1
}

func (s *MapIterator) Next() interface{} {
	item := s.values[s.currentPosition]
	s.currentPosition++
	return item
}

func main() {
	slice := []interface{}{1, 2, 3, 4, 5}
	_map := map[interface{}]interface{}{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
	}

	fmt.Println("Slice iterator:")
	sliceIterator := NewSliceIterator(slice)
	for sliceIterator.HasNext() {
		fmt.Printf("- %d\n", sliceIterator.Next())
	}

	fmt.Println()

	fmt.Println("Map iterator:")
	mapIterator := NewMapIterator(_map)
	for mapIterator.HasNext() {
		fmt.Printf("- %d\n", mapIterator.Next())
	}
}
