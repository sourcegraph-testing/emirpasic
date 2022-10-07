// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"testing"
	"time"
)

func TestIntComparator(t *testing.T) {

	// i1,i2,expected
	tests := [][]any{
		{1, 1, 0},
		{1, 2, -1},
		{2, 1, 1},
		{11, 22, -1},
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
	}

	for _, test := range tests {
		actual := IntComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestStringComparator(t *testing.T) {

	// s1,s2,expected
	tests := [][]any{
		{"a", "a", 0},
		{"a", "b", -1},
		{"b", "a", 1},
		{"aa", "aab", -1},
		{"", "", 0},
		{"a", "", 1},
		{"", "a", -1},
		{"", "aaaaaaa", -1},
	}

	for _, test := range tests {
		actual := StringComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestTimeComparator(t *testing.T) {

	now := time.Now()

	// i1,i2,expected
	tests := [][]any{
		{now, now, 0},
		{now.Add(24 * 7 * 2 * time.Hour), now, 1},
		{now, now.Add(24 * 7 * 2 * time.Hour), -1},
	}

	for _, test := range tests {
		actual := TimeComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestCustomComparator(t *testing.T) {

	type Custom struct {
		id   int
		name string
	}

	byID := func(a, b any) int {
		c1 := a.(Custom)
		c2 := b.(Custom)
		switch {
		case c1.id > c2.id:
			return 1
		case c1.id < c2.id:
			return -1
		default:
			return 0
		}
	}

	// o1,o2,expected
	tests := [][]any{
		{Custom{1, "a"}, Custom{1, "a"}, 0},
		{Custom{1, "a"}, Custom{2, "b"}, -1},
		{Custom{2, "b"}, Custom{1, "a"}, 1},
		{Custom{1, "a"}, Custom{1, "b"}, 0},
	}

	for _, test := range tests {
		actual := byID(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestInt8ComparatorComparator(t *testing.T) {
	tests := [][]any{
		{int8(1), int8(1), 0},
		{int8(0), int8(1), -1},
		{int8(1), int8(0), 1},
	}
	for _, test := range tests {
		actual := Int8Comparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestInt16Comparator(t *testing.T) {
	tests := [][]any{
		{int16(1), int16(1), 0},
		{int16(0), int16(1), -1},
		{int16(1), int16(0), 1},
	}
	for _, test := range tests {
		actual := Int16Comparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestInt32Comparator(t *testing.T) {
	tests := [][]any{
		{int32(1), int32(1), 0},
		{int32(0), int32(1), -1},
		{int32(1), int32(0), 1},
	}
	for _, test := range tests {
		actual := Int32Comparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestInt64Comparator(t *testing.T) {
	tests := [][]any{
		{int64(1), int64(1), 0},
		{int64(0), int64(1), -1},
		{int64(1), int64(0), 1},
	}
	for _, test := range tests {
		actual := Int64Comparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestUIntComparator(t *testing.T) {
	tests := [][]any{
		{uint(1), uint(1), 0},
		{uint(0), uint(1), -1},
		{uint(1), uint(0), 1},
	}
	for _, test := range tests {
		actual := UIntComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestUInt8Comparator(t *testing.T) {
	tests := [][]any{
		{uint8(1), uint8(1), 0},
		{uint8(0), uint8(1), -1},
		{uint8(1), uint8(0), 1},
	}
	for _, test := range tests {
		actual := UInt8Comparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestUInt16Comparator(t *testing.T) {
	tests := [][]any{
		{uint16(1), uint16(1), 0},
		{uint16(0), uint16(1), -1},
		{uint16(1), uint16(0), 1},
	}
	for _, test := range tests {
		actual := UInt16Comparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestUInt32Comparator(t *testing.T) {
	tests := [][]any{
		{uint32(1), uint32(1), 0},
		{uint32(0), uint32(1), -1},
		{uint32(1), uint32(0), 1},
	}
	for _, test := range tests {
		actual := UInt32Comparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestUInt64Comparator(t *testing.T) {
	tests := [][]any{
		{uint64(1), uint64(1), 0},
		{uint64(0), uint64(1), -1},
		{uint64(1), uint64(0), 1},
	}
	for _, test := range tests {
		actual := UInt64Comparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestFloat32Comparator(t *testing.T) {
	tests := [][]any{
		{float32(1.1), float32(1.1), 0},
		{float32(0.1), float32(1.1), -1},
		{float32(1.1), float32(0.1), 1},
	}
	for _, test := range tests {
		actual := Float32Comparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestFloat64Comparator(t *testing.T) {
	tests := [][]any{
		{float64(1.1), float64(1.1), 0},
		{float64(0.1), float64(1.1), -1},
		{float64(1.1), float64(0.1), 1},
	}
	for _, test := range tests {
		actual := Float64Comparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestByteComparator(t *testing.T) {
	tests := [][]any{
		{byte(1), byte(1), 0},
		{byte(0), byte(1), -1},
		{byte(1), byte(0), 1},
	}
	for _, test := range tests {
		actual := ByteComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestRuneComparator(t *testing.T) {
	tests := [][]any{
		{rune(1), rune(1), 0},
		{rune(0), rune(1), -1},
		{rune(1), rune(0), 1},
	}
	for _, test := range tests {
		actual := RuneComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}
