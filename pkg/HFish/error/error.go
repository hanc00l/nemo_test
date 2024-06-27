package error

import (
	"fmt"
)

func Check(e error, tips string) {
	if e != nil {
		fmt.Println(tips)
		//panic(e)
	}
}
