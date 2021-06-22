package function

import (
	"fmt"
)

func Iter() {
	sum := 0
	//Original for
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum_while := 1
	//for like while
	for sum_while < 100 {
		sum_while *= 2
	}
	fmt.Println(sum_while)

	names := []string{"a", "b", "c"}

	for index, name := range names {
		fmt.Println(index, name)
	}
}
