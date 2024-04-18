package intx

import "github.com/gogf/gf/v2/util/gconv"

// InArray 包含
func InArray(x int, a []int) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

// Unique 唯一
func Unique(a []int) []int {
	r := make([]int, 0, len(a))
	m := make(map[int]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Filter 过滤0
func Filter(a []int) []int {
	r := make([]int, 0, len(a))
	for _, v := range a {
		if v != 0 {
			r = append(r, v)
		}
	}
	return r
}

// Intersect 交集
func Intersect(s1, s2 []int) []int {
	m := make(map[int]int)
	n := make([]int, 0)
	for _, v := range s1 {
		m[v]++
	}
	for _, v := range s2 {
		times, _ := m[v]
		if times > 0 {
			n = append(n, v)
		}
	}
	return n
}

// Union 并集
func Union(s1, s2 []int) []int {
	m := make(map[int]int)
	for _, v := range s1 {
		m[v]++
	}
	for _, v := range s2 {
		times, _ := m[v]
		if times == 0 {
			s1 = append(s1, v)
		}
	}
	return s1
}

// Diff 差集
func Diff(s1, s2 []int) []int {
	m := make(map[int]int)
	n := make([]int, 0)
	inter := Intersect(s1, s2)
	for _, v := range inter {
		m[v]++
	}
	for _, v := range s1 {
		times, _ := m[v]
		if times == 0 {
			n = append(n, v)
		}
	}
	return n
}

// Ints 转换至[]int
func Ints(i interface{}) []int {
	return gconv.Ints(i)
}
