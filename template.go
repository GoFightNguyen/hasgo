// Code generated by go generate; DO NOT EDIT.
package main

var hasgoTemplates = map[string]string{
	"Abs.go": `
import (
	"math"
)

func (s SliceType) Abs() (out SliceType) {
	for _, v := range s {
		out = append(out, ElementType(math.Abs(float64(v))))
	}
	return
}
`,
	"Filter.go": `
func (s SliceType) Filter(f func(ElementType) bool) (out SliceType) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}
`,
	"Head.go": `
// Returns the first element in the slice
// If no element is found, returns the zero-value of the type
func (s SliceType) Head() (out ElementType) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}
`,
	"Init.go": `
// Take n-1 elements from a slice, where n = len(list)
func (s SliceType) Init() (out SliceType) {
	slicecopy := append([]ElementType(nil), s...)
	return slicecopy[:len(s)-1]
}
`,
	"Last.go": `
// Returns the last element in the slice
// If no element is found, returns the zero-value of the type
func (s SliceType) Last() (out ElementType) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}
`,
	"Maximum.go": `
func (s SliceType) Maximum() (out ElementType) {
	if len(s) == 0 {
		return
	}
	for _, i := range s {
		if i > out {
			out = i
		}
	}
	return
}
`,
	"Minimum.go": `
func (s SliceType) Minimum() ElementType {
	if len(s) == 0 {
		return 0 // bit strange?
	}
	min := s[0]
	for _, i := range s {
		if i < min {
			min = i
		}
	}
	return min
}
`,
	"Reverse.go": `
// Returns the reversed slice
func (s SliceType) Reverse() (out SliceType) {
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return
}
`,
	"Sum.go": `
func (s SliceType) Sum() ElementType {
	var sum ElementType
	for _, v := range s {
		sum += v
	}
	return sum
}
`,
	"Tail.go": `
// Take [1 -> n] elements from a slice, where n = len(list)
// Returns an empty slice if there are less than 2 elements in slice
func (s SliceType) Tail() (out SliceType) {
	if len(s) <= 1 {
		return
	}
	slicecopy := append([]ElementType(nil), s...)
	return slicecopy[1:]
}
`,
	"Uncons.go": `
func (s SliceType) Uncons() (head ElementType, tail SliceType) {
	return s.Head(), s.Tail()
}
`,
}

const (
	ForNumbers = "ForNumbers"
	ForStrings = "ForStrings"
	ForStructs = "ForStructs"
)

var funcDomains = map[string][]string{
	"Abs.go":     []string{ForNumbers},
	"Filter.go":  []string{ForNumbers, ForStrings, ForStructs},
	"Head.go":    []string{ForNumbers, ForStrings, ForStructs},
	"Init.go":    []string{ForNumbers, ForStrings, ForStructs},
	"Last.go":    []string{ForNumbers, ForStrings, ForStructs},
	"Maximum.go": []string{ForNumbers},
	"Minimum.go": []string{ForNumbers},
	"Reverse.go": []string{ForNumbers, ForStrings, ForStructs},
	"Sum.go":     []string{ForNumbers, ForStrings},
	"Tail.go":    []string{ForNumbers, ForStrings, ForStructs},
	"Uncons.go":  []string{ForNumbers, ForStrings, ForStructs},
}
