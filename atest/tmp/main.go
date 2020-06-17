package main

import "fmt"

//func a() {
//	x := []int{}
//	x = append(x, 0)
//	x = append(x, 1)  // commonTags := labelsToTags(app.Labels)
//	y := append(x, 2) // Tags: append(commonTags, labelsToTags(d.Labels)...)
//	z := append(x, 3) // Tags: append(commonTags, labelsToTags(d.Labels)...)
//	fmt.Println(y, z)
//}
//
func b() {
	x := []int{}
	fmt.Printf("%p %d %d\n", x, len(x), cap(x))
	x = append(x, 0)
	fmt.Printf("%p %d %d\n", x, len(x), cap(x))
	x = append(x, 1)
	fmt.Printf("%p %d %d\n", x, len(x), cap(x))
	x = append(x, 2) // commonTags := labelsToTags(app.Labels)
	fmt.Printf("%p %d %d\n", x, len(x), cap(x))
	y := make([]int, len(x))
	copy(y, x)
	//y := append([]int{}, x...) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	y = append(y, 3)
	fmt.Printf("%p %d %d\n", x, len(x), cap(x))
	z := append([]int{}, x...) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	z = append(z, 4)
	fmt.Printf("%p\n", z)
	fmt.Println(y, z)
}

func main() {
	//a()
	b()
}
