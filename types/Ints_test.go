package types

import (
	"testing"
)

// unit testing the Ints
var (
	intsSumTests = []struct {
		input  Ints
		output int64
	}{
		{
			nil,
			0,
		},
		{
			Ints([]int64{}),
			0,
		},
		{
			Ints([]int64{1, 2, 3}),
			6,
		},
	}

	intsFilterTests = []struct {
		input  Ints
		filter func(int64) bool
		output Ints
	}{
		{
			nil,
			nil,
			Ints{},
		},
		{
			Ints{},
			nil,
			Ints{},
		},
		{
			Ints{1, 2, 3, 4},
			func(i int64) bool { return i%2 == 0 },
			Ints{2, 4},
		},
	}

	intsRangeTests = []struct {
		start  int64
		stop   int64
		output Ints
	}{
		{0, 10, Ints{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{-10, 1, Ints{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1}},
	}

	// tests that take a subset of the slice
	intsTakeTests = []struct {
		input Ints
		init  Ints
		tail  Ints
		head  int64
		last  int64
	}{
		{
			IntRange(0, 10),
			IntRange(0, 9),
			IntRange(1, 10),
			0,
			10,
		},
	}
)

func Test_IntsSum(t *testing.T) {
	for _, test := range intsSumTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Sum(); res != test.output {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsFilter(t *testing.T) {
	for _, test := range intsFilterTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Filter(test.filter); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsRange(t *testing.T) {
	for _, test := range intsRangeTests {
		t.Run("", func(t *testing.T) {
			if res := IntRange(test.start, test.stop); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsInit(t *testing.T) {
	for _, test := range intsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Init(); !res.Equals(test.init) {
				t.Errorf("expected %v but got %v", test.init, res)
			}
		})
	}
}

func Test_IntsTail(t *testing.T) {
	for _, test := range intsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Tail(); !res.Equals(test.tail) {
				t.Errorf("expected %v but got %v", test.tail, res)
			}
		})
	}
}

func Test_IntsHead(t *testing.T) {
	for _, test := range intsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Last(); res != test.last {
				t.Errorf("expected %v but got %v", test.last, res)
			}
		})
	}
}

func Test_IntsLast(t *testing.T) {
	for _, test := range intsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Head(); res != test.head {
				t.Errorf("expected %v but got %v", test.head, res)
			}
		})
	}
}