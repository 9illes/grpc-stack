package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		desc  string
		input []int32
		want  []int32
	}{
		{
			desc:  "Push an int32 on the stack",
			input: []int32{100},
			want:  []int32{100},
		},
		{
			desc:  "Push multiple int32 on the stack",
			input: []int32{100, 101, 102},
			want:  []int32{100, 101, 102},
		},
	}

	for _, tc := range tests {
		sut := NewStack()
		for _, v := range tc.input {
			sut.Push(v)
		}

		if !reflect.DeepEqual(tc.want, sut.data) {
			t.Fatalf("expected: %v, got: %v", tc.want, sut.data)
		}
	}
}

func TestPopEmptyStack(t *testing.T) {
	sut := NewStack()
	n, empty := sut.Pop()
	if n != 0 {
		t.Fatalf("expected: %v, got: %v", 0, n)
	}
	if !empty {
		t.Fatalf("expected: %v, got: %v", true, empty)
	}
}

func TestStack(t *testing.T) {
	sut := NewStack()
	input := []int32{100, 101, 102}

	for _, v := range input {
		sut.Push(v)
	}

	// pop all numbers
	for i := len(input) - 1; i >= 0; i-- {
		n, empty := sut.Pop()
		if n != input[i] {
			t.Fatalf("expected: %v, got: %v", input[i], n)
		}
		if empty {
			t.Fatalf("expected: %v, got: %v", false, empty)
		}
	}

	// stack is now empty
	n, empty := sut.Pop()
	if n != 0 {
		t.Fatalf("expected: %v, got: %v", 0, n)
	}
	if !empty {
		t.Fatalf("expected: %v, got: %v", true, empty)
	}
}
