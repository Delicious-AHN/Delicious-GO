/*
	bool
	string ->immutable type
	int, int8, int16, int32, int64
	uint, uint8, uint16, uint32, uint64, uintptr
	float32, float64, complex64, complex128
	byte -> uint8과 동일, 바이트 코드에 사용
	rune -> int32와 동일, 유니코드 코드포인터에 사용
*/

package function

import "fmt"

func Datatyp() {
	interPreted := "abc\n acbc\n"

	fmt.Println(interPreted)
}
