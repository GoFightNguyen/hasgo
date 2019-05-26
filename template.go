// Code generated by go generate; DO NOT EDIT.
package main

var hasgoTemplates = map[string]string{
	"abs.go": `
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
	"filter.go": `
func (s SliceType) Filter(f func(ElementType) bool) (out SliceType) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}
`,
	"head.go": `
// Returns the first element in the slice
// If no element is found, returns the zero-value of the type
func (s SliceType) Head() (out ElementType) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}
`,
	"init.go": `
// Take n-1 elements from a slice, where n = len(list)
func (s SliceType) Init() (out SliceType) {
	slicecopy := append([]ElementType(nil), s...)
	return slicecopy[:len(s)-1]
}
`,
	"last.go": `
// Returns the last element in the slice
// If no element is found, returns the zero-value of the type
func (s SliceType) Last() (out ElementType) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}
`,
	"length.go": `
func (s SliceType) Length() int {
	return len(s)
}
`,
	"map.go": `
// Return a new slice with the map operation applied to each element.
func (s SliceType) Map(f func(ElementType) ElementType) (out SliceType) {
	if f == nil {
		return s
	}
	for _, v := range s {
		out = append(out, f(v))
	}
	return
}
`,
	"maximum.go": `
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
	"minimum.go": `
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
	"reverse.go": `
// Returns the reversed slice
func (s SliceType) Reverse() (out SliceType) {
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return
}
`,
	"sort.go": `
import (
	"sort"
)

// wrapper around go sort functions
func (s SliceType) Sort() SliceType {
	out := make(SliceType, len(s))
	copy(out, s)
	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return out
}
`,
	"sum.go": `
func (s SliceType) Sum() ElementType {
	var sum ElementType
	for _, v := range s {
		sum += v
	}
	return sum
}
`,
	"tail.go": `
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
	"uncons.go": `
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
	"abs.go":     []string{ForNumbers},
	"filter.go":  []string{ForNumbers, ForStrings, ForStructs},
	"head.go":    []string{ForNumbers, ForStrings, ForStructs},
	"init.go":    []string{ForNumbers, ForStrings, ForStructs},
	"last.go":    []string{ForNumbers, ForStrings, ForStructs},
	"length.go":  []string{ForNumbers, ForStrings, ForStructs},
	"map.go":     []string{ForNumbers, ForStrings, ForStructs},
	"maximum.go": []string{ForNumbers},
	"minimum.go": []string{ForNumbers},
	"reverse.go": []string{ForNumbers, ForStrings, ForStructs},
	"sort.go":    []string{ForNumbers, ForStrings},
	"sum.go":     []string{ForNumbers, ForStrings},
	"tail.go":    []string{ForNumbers, ForStrings, ForStructs},
	"uncons.go":  []string{ForNumbers, ForStrings, ForStructs},
}
