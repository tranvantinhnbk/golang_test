package main

import (
	"fmt"
	"reflect"
)

func main() {
	var pi float64 = 3.14

	// 1. REFLECTION: get type and value
	fmt.Println("1. REFLECTION: get type and value")
	t := reflect.TypeOf(pi)
	v := reflect.ValueOf(pi)

	fmt.Println("Type:", t)        // float64
	fmt.Println("Value:", v)       // 3.14
	fmt.Println("Kind:", t.Kind()) // float64

	// 2. REFLECTION: modify value using reflection:
	fmt.Println("2. REFLECTION: modify value using reflection:")
	v = reflect.ValueOf(&pi).Elem()
	v.SetFloat(6.28)
	fmt.Println("Current Value: ", pi)

	//3. REFLECTION: struct introspection
	fmt.Println("3. REFLECTION: struct introspection")
	p := Person{"Alice", 30}
	t = reflect.TypeOf(p)
	v = reflect.ValueOf(p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("%s (%s) = %v\n", field.Name, field.Type, value)
	}

	//4. REFLECTION: reflect on method
	fmt.Println("4. REFLECTION: reflect on method")
	m := Math{}
	v = reflect.ValueOf(m)
	method := v.MethodByName("Add")
	args := []reflect.Value{reflect.ValueOf(3), reflect.ValueOf(4)}
	result := method.Call(args)
	fmt.Println("Result:", result[0].Int()) // 7

}

type Person struct {
	Name string
	Age  int
}

type Math struct{}

func (m Math) Add(a, b int) int {
	return a + b
}
