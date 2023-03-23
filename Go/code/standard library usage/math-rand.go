// https://golang.google.cn/pkg/math/rand/

import "math/rand"

func Seed(seed int64)
// safe for concurrent use
// if not called, default as Seed(1)

func Intn(n int) int
// [0, n)
// It panics if n <= 0.

func Int() int
// non-negative