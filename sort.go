package main

import (
	"container/heap"
	"fmt"
)

// BubbleSort (Public) - for each item in array, make pass
// through array, swapping pairs of adjacent items if
// second item bigger than first
func BubbleSort(w *[]string) {
	var sorted bool
	for i := 0; i < len(*w); i++ {
		if sorted {
			break
		}
		sorted = true
		for j := 0; j < len(*w)-1-i; j++ {
			if (*w)[j] > (*w)[j+1] {
				(*w)[j], (*w)[j+1] = (*w)[j+1], (*w)[j]
				sorted = false
			}
		}
	}
}

// SelectionSort (Public) - for each position in array,
// find item that goes there and swap with item in that spot
func SelectionSort(w *[]string) {
	for i := 0; i < len(*w)-1; i++ {
		smallest := i
		for j := i + 1; j < len(*w); j++ {
			if (*w)[j] < (*w)[smallest] {
				smallest = j
			}
		}
		if smallest > i {
			(*w)[i], (*w)[smallest] = (*w)[smallest], (*w)[i]
		}
	}
}

// word (Private) - stores a value for the heap node,
type word struct {
	value string
}

// words (Private) - creates a datatype of a slice of pointers of words
type words []*word

// Len (len) - Needed to be implemented for the sort interface, returns the length of the array.
func (w words) Len() int {
	return len(w)
}

// Less (Public) - Needed to be implement for the sort interface, compares and determines if left value is less than right value
func (w words) Less(i, j int) bool {
	return w[i].value < w[j].value
}

// Pop (Public) - Needed to be implement for the sort interface for heap (Deletes).
func (w *words) Pop() interface{} {
	old := *w
	n := len(old)
	item := old[n-1]
	*w = old[0 : n-1]
	return item
}

// Push (Public) - Needed to be implement for the sort interface for a heap (Adds).
func (w words) Push(x interface{}) {
	item := x.(*word)
	w = append(w, item)
}

// Swap (Public) - Needed to be impletment for the sort interface, changes the position of elements based
// on index
func (w words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

// String (Public) - returns a string representation for the structure of words
func (w words) String() string {
	s := "["
	for _, word := range w {
		s += word.value + " "
	}
	return s[:len(s)-1] + "]"
}

// InsertionSort (Public) - Still not very fast but another way to sort an array by inserting in sorted order a temporary array
func InsertionSort(w *[]string) {
	temp := []string{(*w)[0]}
	for pos := 1; pos < len(*w); pos++ {
		for x := 0; x < len(temp); x++ {
			if (*w)[pos] < temp[x] {
				// Maybe I should explain this one line trickery
				// append method takes in a slice and appends items of the same type.
				// When you include ... it iterates over the slice and is a shortcut so you don't have to append each item individually
				// Basically I want insert one item inside of a slice at a given position x.
				// I start by taking all the items in temp up until the position I want to insert
				// then append a second slice starting with the item to be inserted + the rest of the temp slice
				// Since in this example append will only take items of slice string, I add the ... to add them individually.
				temp = append(temp[:x], append([]string{(*w)[pos]}, temp[x:]...)...)
				break
			}
			if x == len(temp)-1 {
				temp = append(temp, (*w)[pos])
				break
			}
		}
	}
	*w = temp
}

// HeapSort (Public) - stores the heap in sorted order. When you pop them out of the heap, the items are
// return as a min sorted heap.
func HeapSort(w *[]string) {
	var wd words
	for _, x := range *w {
		wd = append(wd, &word{value: x})
	}
	heap.Init(&wd)
	n := len(wd)
	for x := 0; x < n; x++ {
		item := heap.Pop(&wd).(*word).value
		(*w)[x] = item
	}
}

// MergeSort (Public) - recusive sort where you divide and conquer the slice until you can't anymore
// then you merge back until you are in sorted order.
func MergeSort(w *[]string) {
	if len(*w) > 1 {
		half := len(*w) / 2
		left := (*w)[:half]
		right := (*w)[half:]
		MergeSort(&left)
		MergeSort(&right)
		*w = merge(left, right)
	}
}

// merge (Private) - helper function that takes two arrays and merges them into
// sorted order.
func merge(left, right []string) []string {
	var items []string
	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			items = append(items, left[0])
			left = left[1:]
		} else {
			items = append(items, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		items = append(items, left...)
	}
	if len(right) > 0 {
		items = append(items, right...)
	}
	return items
}

func main() {
	w := []string{
		"banana", "grapes", "coconut",
		"grapes", "plum", "apple",
		"cherries",
	}

	fmt.Println("*************************************************")
	fmt.Printf("Before sorting\n%v\n", w)
	BubbleSort(&w) // completed
	fmt.Printf("After Bubble sorting\n%v\n", w)
	w = []string{"banana", "grapes", "coconut", "grapes", "plum", "apple", "cherries"}
	fmt.Println("*************************************************")
	fmt.Println("")
	fmt.Println("\n*************************************************")
	fmt.Printf("Before sorting\n%v\n", w)
	SelectionSort(&w) // completed
	fmt.Printf("After Selection sorting\n%v\n", w)
	w = []string{"banana", "grapes", "coconut", "grapes", "plum", "apple", "cherries"}
	fmt.Println("*************************************************")
	fmt.Println("")
	fmt.Println("\n*************************************************")
	fmt.Printf("Before sorting\n%v\n", w)
	InsertionSort(&w) // To be implemented
	fmt.Printf("After Insertion sorting\n%v\n", w)
	w = []string{"banana", "grapes", "coconut", "grapes", "plum", "apple", "cherries"}
	fmt.Println("*************************************************")
	fmt.Println("")
	fmt.Println("\n*************************************************")
	fmt.Printf("Before sorting\n%v\n", w)
	HeapSort(&w) // completed
	fmt.Printf("After Heap sorting\n%v\n", w)
	w = []string{"banana", "grapes", "coconut", "grapes", "plum", "apple", "cherries"}
	fmt.Println("*************************************************")
	fmt.Println("")
	fmt.Println("\n*************************************************")
	fmt.Printf("Before sorting\n%v\n", w)
	MergeSort(&w) // completed
	fmt.Printf("After Merge sorting\n%v\n", w)
	fmt.Println("*************************************************")
}
