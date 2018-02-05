package main

import (
	_ "fmt"
	"strings"
)

const (

)
var million int = 1e6

func main()  {
	a, b, c := "a", "b", "c"
	s := strings.Join([]string{ a, b, c }, ",")
	println(s)
}
