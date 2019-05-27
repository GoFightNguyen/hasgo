// code generated by hasgo. [DO NOT EDIT!]
package types

// =============== filter.go =================

func (s persons) Filter(f func(person) bool) (out persons) {
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
func (s persons) Head() (out person) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}

// =============== init.go =================

// Take n-1 elements from a slice, where n = len(list)
func (s persons) Init() (out persons) {
	slicecopy := append([]person(nil), s...)
	return slicecopy[:len(s)-1]
}

// =============== last.go =================

// Returns the last element in the slice
// If no element is found, returns the zero-value of the type
func (s persons) Last() (out person) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}

// =============== length.go =================

func (s persons) Length() int {
	return len(s)
}

// =============== map.go =================

// Return a new slice with the map operation applied to each element.
func (s persons) Map(f func(person) person) (out persons) {
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
func (s persons) Null() bool {
	return len(s) == 0
}

// =============== reverse.go =================

// Returns the reversed slice
func (s persons) Reverse() (out persons) {
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return
}

// =============== tail.go =================

// Take [1 -> n] elements from a slice, where n = len(list)
// Returns an empty slice if there are less than 2 elements in slice
func (s persons) Tail() (out persons) {
	if len(s) <= 1 {
		return
	}
	slicecopy := append([]person(nil), s...)
	return slicecopy[1:]
}

// =============== uncons.go =================

func (s persons) Uncons() (head person, tail persons) {
	return s.Head(), s.Tail()
}
