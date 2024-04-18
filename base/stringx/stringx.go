package stringx

import (
	"github.com/gogf/gf/v2/util/gconv"
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

// Strings 转换至[]int
func Strings(i interface{}) []string {
	return gconv.Strings(i)
}
