package judge

import (
	"fmt"
	"reflect"
)

func Judge_type() int {
	var a int
	fmt.Println(reflect.TypeOf(a))
	return a
}
