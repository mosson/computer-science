package array

import "fmt"

// Uint8Array is array of Uint8
type Uint8Array struct {
	array []uint8
}

// NewUint8Array returns new Uint8Array
func NewUint8Array() *Uint8Array {
	return &Uint8Array{array: []uint8{}}
}

// Add adds data to array buffer
func (a *Uint8Array) Add(data uint8) {
	a.array = append(a.array, data)
}

// Remove removes data from array buffer
func (a *Uint8Array) Remove(data uint8) {
	var array []uint8

	for _, i := range a.array {
		if i != data {
			array = append(array, i)
		}
	}

	a.array = array
}

// IndexOf returns the index of data from array buffer
func (a *Uint8Array) IndexOf(data uint8) int {
	for i, o := range a.array {
		if o == data {
			return i
		}
	}
	return -1
}

// Search is alias of Indexof
func (a *Uint8Array) Search(data uint8) int {
	i := a.IndexOf(data)

	if i >= 0 {
		return i
	}

	return -1
}

// GetAtIndex returns the data of index in array buffer
func (a *Uint8Array) GetAtIndex(index int) uint8 {
	return a.array[index]
}

// Length returns length of array buffer
func (a *Uint8Array) Length() int {
	return len(a.array)
}

// Join returns concatenated string
func (a *Uint8Array) Join(sep string) string {
	str := ""
	for _, i := range a.array {
		str = fmt.Sprintf("%s%s%d", str, sep, i)
	}
	return str
}

// Print prints StdOut
func (a *Uint8Array) Print() {
	fmt.Println(a.Join(" "))
}
