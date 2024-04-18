package stringx

import (
	"strings"
)

func InArray(x string, s []string) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func Unique(s []string) []string {
	m := make(map[string]struct{}, len(s))
	r := make([]string, 0, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	for k := range m {
		r = append(r, k)
	}
	return r
}

func Filter(s []string) []string {
	r := make([]string, 0, len(s))
	for _, v := range s {
		if strings.TrimSpace(v) != "" {
			r = append(r, strings.TrimSpace(v))
		}
	}
	return r
}

// Intersect 交集
func Intersect(s1, s2 []string) []string {
	m := make(map[string]int)
	n := make([]string, 0)
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
func Union(s1, s2 []string) []string {
	m := make(map[string]int)
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
func Diff(s1, s2 []string) []string {
	m := make(map[string]int)
	n := make([]string, 0)
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
