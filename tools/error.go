package tools

import "fmt"

func PanicError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
		recover()
		fmt.Println("recover")
	}
}
