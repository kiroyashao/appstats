package appstats

import (
	"html/template"
	"reflect"
	"strconv"
	"strings"
)

// eq reports whether the first argument is equal to
// any of the remaining arguments.
func eq(args ...interface{}) bool {
	if len(args) == 0 {
		return false
	}
	x := args[0]
	switch x := x.(type) {
	case string, int, int64, byte, float32, float64:
		for _, y := range args[1:] {
			if x == y {
				return true
			}
		}
		return false
	}

	for _, y := range args[1:] {
		if reflect.DeepEqual(x, y) {
			return true
		}
	}
	return false
}

func add(a, b int) int {
	return a + b
}

func rjust(i, count int) string {
	s := strconv.Itoa(i)
	return strings.Repeat(" ", count-len(s)) + s
}

var funcs = template.FuncMap{
	"add":   add,
	"eq":    eq,
	"rjust": rjust,
}