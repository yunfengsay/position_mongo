package tools

import "fmt"

func PanicError(err error) {
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
}
