// map

// 声明
var m map[int]int // [keyType]valueType
m := map[int]int{}
m := map[int]int{
	1 : 2,
	2 : 3,
	3 : 4,
}

// make
m := make(map[int]int, 10)

// comma ok idiom
v, ok = m[3] // if 3 is not in m, m won't record m[3] = 0

// delete
delete(m, 3)

// use map as set
set := map[int]bool{}	// [setType]bool
// or
set := map[int]struct{}{}	// smaller, but complicated to write
set[1] = struct{}{}