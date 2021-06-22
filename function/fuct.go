package function

import "fmt"

func Func_example() {
	msg := "CODING IS JUST LIKE THE SEX"
	say1(msg)
	say2(&msg)
	say3("YOUR", "DICK")
	say3("맛있어")
	sum, word := summing(msg)
	fmt.Println(sum, word)

	add := func(i int, j int) int {
		return i + j
	}

	r1 := calc(add, 10, 20)
	fmt.Println(r1)

	r2 := calc(func(x int, y int) int { return x * y }, 10, 20)
	fmt.Println(r2)
}
func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func say1(msg string) {
	fmt.Println(msg)
}

func say2(msg *string) {
	*msg = "FUCK!!!"
	fmt.Println(*msg)
}

func say3(msg ...string) {
	for _, s := range msg {
		fmt.Println(s)
	}
}

func summing(msg ...string) (sum int, say string) {
	for i, s := range msg {
		sum += (i + 10) * 10
		fmt.Println(s)
		say += s
	}
	return sum, say
}
