package task22

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Start() {
	slice := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}

	family := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}

	sort.Ints(slice)
	fmt.Println("slice: ", slice)

	sort.Sort(ByAge(family))
	fmt.Println("Sort Person by Age: ", family)

	sort.Sort(ByName(family))
	fmt.Println("Sort Person by Age: ", family)
}
