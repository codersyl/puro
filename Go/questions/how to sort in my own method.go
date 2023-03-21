// how to sort in my own method

import (
	"fmt"
	"sort"
)

sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
// 1st parameter
// the slice who needs to sort
// 2nd parameter
// a less func