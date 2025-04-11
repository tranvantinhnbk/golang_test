package main

import "fmt"

type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Square[T Numeric](x T) T {
	return x * x
}

type MyInt int

func main() {
	fmt.Println(Square(3))
	fmt.Println(Square(3.2))
	fmt.Println(Square(2))
	fmt.Println(Square(MyInt(4)))
}
