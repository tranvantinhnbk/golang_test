package main

import "fmt"

type Rectangle struct {
	width  float32
	height float32
}

func (rec Rectangle) Area() float32 {
	return rec.width * rec.height
}

func (rec *Rectangle) Scale(k float32) {
	rec.width *= k
	rec.height *= k
}

func main() {
	fmt.Println("Hello world")
	rec := Rectangle{width: 3, height: 2}
	fmt.Println(rec)
	fmt.Println("Area:", rec.Area())
	fmt.Println("Scale by 2")
	rec.Scale(2)
	fmt.Println(rec)
	fmt.Println("Area now:", rec.Area())
}
