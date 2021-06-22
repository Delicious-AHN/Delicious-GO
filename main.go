package main

import (
	"runtime"

	"github.io/Delicious-Ahn/Learning/function"
)

func main() {
	println("Hello World")
	function.Datatyp()
	function.Varcons()
	function.IfSwitch()
	function.Iter()
	function.Func_example()
	function.Slice()
	function.Map()
	function.Stru()
	function.Method()
	runtime.GOMAXPROCS(3)
	function.Go_routine()
	function.Go_channel()
}
