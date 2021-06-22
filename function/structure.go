package function

import "fmt"

type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}

func Stru() {
	p := person{}

	p.name = "Fuck"
	p.age = 18

	p_new := new(person)
	p_new.name = "It's new!"
	p_new.age = 0

	dic := newDict()
	dic.data[1] = "A"

	fmt.Println(p)
	fmt.Println(*p_new)
	fmt.Println(dic)
}
