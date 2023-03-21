// how to convert between string and number

import (
	"strconv"
)

// string -> int
// func Atoi(s string) (int, error)
n, err := strconv.Atoi(s)

// int -> string
// func Itoa(i int) string
s := strconv.Itoa(i)