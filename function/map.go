package function

import "fmt"

func Map() {
	idMap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	idMap[20] = "ABC"
	idMap[434] = "FUCKYOU"

	for i, v := range idMap {
		fmt.Println(i, v)
	}
}
