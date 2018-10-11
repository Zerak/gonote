package main

import "fmt"

// nil check
// 'res' is not 'nil', but its value is 'nil'

func main() {
	doit := func(arg int) interface{} {
		var result struct{}

		if arg > 0 {
			result = struct{}{}
		}

		return result
	}

	if res := doit(-1); res != nil {
		fmt.Println("good result:", res)
	}
}
