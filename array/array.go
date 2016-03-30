package array

import "fmt"

// Array is magic slice
type Array struct {
	data []interface{}
}

// NewArray returns new Array
func NewArray() *Array {
	return &Array{data: []interface{}{}}
}

// Add adds object to the array
func (array *Array) Add(object interface{}) {
	array.data = append(array.data, object)
}

// Remove removes object from the array
func (array *Array) Remove(object interface{}) {
	var data []interface{}

	for _, content := range array.data {
		if object != content {
			data = append(data, content)
		}
	}

	array.data = data
}

// Search returns the index of object in the array
func (array *Array) Search(object interface{}) int {
	for index, content := range array.data {
		if object == content {
			return index
		}
	}
	return -1
}

// GetAtIndex returns the data of index in array buffer
func (array *Array) GetAtIndex(index int) interface{} {
	return array.data[index]
}

// Length returns the length of the array
func (array *Array) Length() int {
	return len(array.data)
}

// Join joins contents of the array, and returns concatenated string with separator
func (array *Array) Join(separator string) string {
	str := ""
	for _, object := range array.data {
		str = fmt.Sprintf("%s%s%v", str, separator, object)
	}
	return str
}

// Map returns new array that manipulated with arguments
func (array *Array) Map(f func(interface{}, int) interface{}) *Array {
	newArray := NewArray()
	for index, object := range array.data {
		newArray.Add(f(object, index))
	}
	return newArray
}

// Print prints the array to std out
func (array *Array) Print() {
	fmt.Println(array.Join(" "))
}
