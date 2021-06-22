package function

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func Method() {
	rect := Rect{10, 20}
	area := rect.area()
	println(area)
}
