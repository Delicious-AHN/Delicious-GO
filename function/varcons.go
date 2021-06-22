package function

func varcons() {
	var a int = 2
	a += 2
	b := 3
	b += 1

	const c int = 10
	const s string = "hi"
	const x = 20
	const y = "hello"
	const (
		i = "abc"
		j = "xyz"
	)
	const (
		i0 = iota
		i1
		i2
	)
	println(i0, i1, i2)
}
