// code generated by hasgo. [DO NOT EDIT!]
package types

import (
	"sort"
)

// =============== all.go =================

func (s Strings) All(f func(string) bool) bool {
	if f == nil {
		return false
	}
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

// =============== filter.go =================

func (s Strings) Filter(f func(string) bool) (out Strings) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}

// =============== head.go =================

// Returns the first element in the slice
// If no element is found, returns the zero-value of the type
func (s Strings) Head() (out string) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}

// =============== init.go =================

// Take n-1 elements from a slice, where n = len(list)
func (s Strings) Init() (out Strings) {
	slicecopy := append([]string(nil), s...)
	return slicecopy[:len(s)-1]
}

// =============== intercalate.go =================

// inserts the method receiver slice into the function slice at each step
func (s Strings) Intercalate(ss [][]string) (out Strings) {
	for i, slice := range ss {
		for _, el := range slice {
			out = append(out, el)
		}
		if i == len(ss)-1 {
			break
		}
		for _, el := range s {
			out = append(out, el)
		}
	}
	return out
}

// =============== last.go =================

// Returns the last element in the slice
// If no element is found, returns the zero-value of the type
func (s Strings) Last() (out string) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}

// =============== length.go =================

func (s Strings) Length() int {
	return len(s)
}

// =============== map.go =================

// Return a new slice with the map operation applied to each element.
func (s Strings) Map(f func(string) string) (out Strings) {
	if f == nil {
		return s
	}
	for _, v := range s {
		out = append(out, f(v))
	}
	return
}

// =============== null.go =================

// tests if the slice is empty
func (s Strings) Null() bool {
	return len(s) == 0
}

// =============== reverse.go =================

// Returns the reversed slice
func (s Strings) Reverse() (out Strings) {
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return
}

// =============== sort.go =================

// wrapper around go sort functions
func (s Strings) Sort() Strings {
	out := make(Strings, len(s))
	copy(out, s)
	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return out
}

// =============== sum.go =================

func (s Strings) Sum() string {
	var sum string
	for _, v := range s {
		sum += v
	}
	return sum
}

// =============== tail.go =================

// Take [1 -> n] elements from a slice, where n = len(list)
// Returns an empty slice if there are less than 2 elements in slice
func (s Strings) Tail() (out Strings) {
	if len(s) <= 1 {
		return
	}
	slicecopy := append([]string(nil), s...)
	return slicecopy[1:]
}

// =============== uncons.go =================

func (s Strings) Uncons() (head string, tail Strings) {
	return s.Head(), s.Tail()
}
