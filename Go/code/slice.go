// 声明
s := []int
var s []int
var s = []int{}
var s = []int{10, 20, 30}
var s = []int{1, 5:6, 10: 11}   // 用于初始化稀疏切片

// append
s = append(s, 10)
s = append(s, 10, 20, 30) // 多值
s = append(s, s1)

// len and cap
len(s)
cap(s)

// make
s = make([]int , len, cap) // at least 2 parameters

// slice of slice
s1 = s[0: 3] // [start, end)
s1 = s[0: 3: 3] // [start, end), 3rd num means the limit of append for subslice

// good advece
// 1. never use append to slice or
// 2. use full slice expression to produce subslice

// copy
// from s2 to s1
copy(s1, s2)
// n is the size of this copy
// depend on min len of s1 and s2
n = copy(s1, s2)    

// clear
s = []int{}
s = make([]int, 0)