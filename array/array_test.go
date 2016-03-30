package array

import "fmt"

func ExampleArray_Add() {
	array := NewArray()

	array.Add(1)
	array.Add("2")
	array.Add(true)
	array.Print()
	// Output:
	// 1 2 true
}

func ExampleArray_Remove() {
	array := NewArray()
	array.Add(1)
	array.Add("2")
	array.Add(true)
	array.Add(4)
	array.Remove(1)
	array.Remove(true)
	array.Print()
	// Output:
	// 2 4
}

func ExampleArray_Search() {
	array := NewArray()
	array.Add(1)
	array.Add("2")
	array.Add(true)
	fmt.Println(array.Search(true))
	// Output:
	// 2
}

func ExampleArray_GetAtIndex() {
	array := NewArray()
	array.Add(1)
	array.Add("2")
	array.Add(true)
	fmt.Println(array.GetAtIndex(2))
	// Output:
	// true
}

func ExampleArray_Map() {
	array := NewArray()
	array.Add(1)
	array.Add(2)
	array.Add(3)
	array2 := array.Map(func(object interface{}, index int) interface{} {
		return object.(int) * 2
	})
	array2.Print()
	// Output:
	// 2 4 6
}

func ExampleArray_Length() {
	array := NewArray()
	array.Add(1)
	array.Add("2")
	array.Add(true)
	fmt.Println(array.Length())
	// Output:
	// 3
}
