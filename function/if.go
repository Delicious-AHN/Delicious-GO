package function

import "fmt"

func IfSwitch() {
	k := 1
	if k == 1 {
		fmt.Println("k is 1")
	} else {
		fmt.Println("k isn't 1")
	}

	if val := k * 2; val < k*5 {
		fmt.Println(val - 1)
		val++
	}

	var name string
	var category = 2

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "EBOOK"
		fallthrough
	case 3:
		name = "FUCKYOU"
		fallthrough
	default:
		name = "FXXK YOUR SELF"
	}
	println(name)

}
