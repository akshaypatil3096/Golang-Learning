package main

import "fmt"

// MapHeap struct has a slice of int
type MaxHeap struct {
	array []int
}

// Insert adds an element in the heap
func (h *MaxHeap) Insert(key int) {
	// append will add the new element at the end of the slice
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)

}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	// this loop will iterate till parent index value is greater than index value
	for h.array[parent(index)] < h.array[index] {
		// swap the value of parent index with the new index
		h.swap(parent(index), index)
		// make the new index as parent index
		index = parent(index)
	}
}

// get the parent index
func parent(index int) int {
	return (index - 1) / 2
}

// get left child index
func left(index int) int {
	return (index * 2) + 1
}

// get right child index
func right(index int) int {
	return (index * 2) + 2
}

// swap keys in the array
func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func main() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 40, 25, 35}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
}
