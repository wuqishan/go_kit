package stringx

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
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

// String converts `any` to string.
// It's most commonly used converting function.
func String(any interface{}) string {
	if any == nil {
		return ""
	}
	switch value := any.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	default:
		// Empty checks.
		if value == nil {
			return ""
		}
		// Reflect checks.
		var (
			rv   = reflect.ValueOf(value)
			kind = rv.Kind()
		)
		switch kind {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		case reflect.String:
			return rv.String()
		}
		if kind == reflect.Ptr {
			return String(rv.Elem().Interface())
		}
		// Finally, we use json.Marshal to convert.
		if jsonContent, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(jsonContent)
		}
	}
}
