package function

import "fmt"

func Slice() {
	s := make([]int, 5)
	fmt.Println(len(s), cap(s), s)
	s = append(s, 2)
	s = append(s, 3, 4, 5)
	fmt.Println(s)
}
