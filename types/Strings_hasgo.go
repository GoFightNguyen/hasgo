// code generated by hasgo. [DO NOT EDIT!]
package types

import (
	"fmt"
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

// =============== any.go =================

// Returns true if any of the elements satisfy the predicate
func (s Strings) Any(f func(string) bool) bool {
	if f == nil {
		return false
	}
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
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
	if len(s) == 0 {
		return
	}
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

// =============== intersperse.go =================

func (s Strings) Intersperse(value string) (out Strings) {
	for i, el := range s {
		out = append(out, el)
		if i == len(s)-1 {
			break
		}
		out = append(out, value)
	}
	return
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

// =============== modes.go =================

func (s Strings) Modes() (out Strings) {
	if len(s) == 0 {
		return
	}

	counts := make(map[string]int)
	for _, v := range s {
		counts[v] += 1
	}

	var max int
	for _, v := range counts {
		if v > max {
			max = v
		}
	}

	for k, v := range counts {
		if v == max {
			out = append(out, k)
		}
	}

	return
}

// =============== nub.go =================

func (s Strings) Nub() (out Strings) {
	if len(s) == 0 {
		return
	}

	contains := make(map[string]struct{})
	for _, v := range s {
		if _, ok := contains[v]; !ok {
			contains[v] = struct{}{}
			out = append(out, v)
		}
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

// =============== take.go =================

func (s Strings) Take(n uint64) (out Strings) {
	if len(s) == 0 {
		return
	}
	out = make(Strings, len(s))
	copy(out, s)
	if n < uint64(len(s)) {
		return out[:n]
	}
	return
}

// =============== uncons.go =================

func (s Strings) Uncons() (head string, tail Strings) {
	return s.Head(), s.Tail()
}

// =============== unlines.go =================

// Joins together the string representation of the slice
// With newlines after each element.
func (s Strings) Unlines() (out string) {
	for i, v := range s {
		out += fmt.Sprintf("%v", v)
		if i != len(s)-1 {
			out += "\n"
		}
	}
	return
}
