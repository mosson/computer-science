package array

import "fmt"

func ExampleUint8Array_Add() {
	array := NewUint8Array()

	array.Add(1)
	array.Add(2)
	array.Add(3)
	array.Add(4)
	array.Print() // => 1 2 3 4
	// Output:
	// 1 2 3 4
}

func ExampleUint8Array_Search() {
	array := NewUint8Array()
	array.Add(1)
	array.Add(2)
	array.Add(3)
	array.Add(4)
	fmt.Println(array.Search(3))
	// Output:
	// 2
}

func ExampleUint8Array_GetAtIndex() {
	array := NewUint8Array()
	array.Add(1)
	array.Add(2)
	array.Add(3)
	array.Add(4)
	fmt.Println(array.GetAtIndex(2))
	// Output:
	// 3
}

func ExampleUint8Array_Remove() {
	array := NewUint8Array()
	array.Add(1)
	array.Add(2)
	array.Add(3)
	array.Add(4)
	array.Remove(3)
	array.Print()
	// Output:
	// 1 2 4
}
