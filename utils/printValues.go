package utils

import (
	"fmt"
)

// This is a special kind of interface in which we can pass any type of value

func PrintValue(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float64: ", value)
	case string:
		fmt.Println("String: ", value)
	default:
		fmt.Println("Unknown Type: ", value)

	}
	/*
		# Another way

		intVal, ok := value.(int)

		if !ok {
			fmt.Println(ok)
		}

		fmt.Println(intVal+1)
	*/

}

// Generic
func PrintGeneric[T interface{}](v1 T, v2 T) {
	fmt.Println(v1, v2)
}
