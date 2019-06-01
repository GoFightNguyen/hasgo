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
	"all.go": `
func (s SliceType) All(f func(ElementType) bool) bool {
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
`,
	"any.go": `
// Returns true if any of the elements satisfy the predicate
func (s SliceType) Any(f func(ElementType) bool) bool {
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
	"intercalate.go": `
// inserts the method receiver slice into the function slice at each step
func (s SliceType) Intercalate(ss SliceSliceType) (out SliceType) {
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
	"null.go": `
// tests if the slice is empty
func (s SliceType) Null() bool {
	return len(s) == 0
}
`,
	"product.go": `
func (s SliceType) Product() ElementType {
	var prod ElementType
	for _, v := range s {
		prod += v
	}
	return prod
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
	"take.go": `
func (s SliceType) Take(n uint64) (out SliceType) {
	out = make(SliceType, len(s))
	copy(out, s)
	return out[:n]
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
	"abs.go":         []string{ForNumbers},
	"all.go":         []string{ForNumbers, ForStrings, ForStructs},
	"any.go":         []string{ForNumbers, ForStrings, ForStructs},
	"filter.go":      []string{ForNumbers, ForStrings, ForStructs},
	"head.go":        []string{ForNumbers, ForStrings, ForStructs},
	"init.go":        []string{ForNumbers, ForStrings, ForStructs},
	"intercalate.go": []string{ForNumbers, ForStrings, ForStructs},
	"last.go":        []string{ForNumbers, ForStrings, ForStructs},
	"length.go":      []string{ForNumbers, ForStrings, ForStructs},
	"map.go":         []string{ForNumbers, ForStrings, ForStructs},
	"maximum.go":     []string{ForNumbers},
	"minimum.go":     []string{ForNumbers},
	"null.go":        []string{ForNumbers, ForStrings, ForStructs},
	"product.go":     []string{ForNumbers},
	"reverse.go":     []string{ForNumbers, ForStrings, ForStructs},
	"sort.go":        []string{ForNumbers, ForStrings},
	"sum.go":         []string{ForNumbers, ForStrings},
	"tail.go":        []string{ForNumbers, ForStrings, ForStructs},
	"take.go":        []string{ForNumbers, ForStrings, ForStructs},
	"uncons.go":      []string{ForNumbers, ForStrings, ForStructs},
}
