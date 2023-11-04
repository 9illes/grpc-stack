package stack

type Stack struct {
	data []int32
}

// NewStack creates a new instance of the Stack struct.
func NewStack() *Stack {
	return &Stack{
		data: make([]int32, 0),
	}
}

// Push adds a new value to the stack.
func (s *Stack) Push(value int32) {
	s.data = append(s.data, value)
}

// Pop removes and returns the top element from the stack.
func (s *Stack) Pop() (int32, bool) {
	if len(s.data) == 0 {
		return 0, true
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return value, false
}

// Size returns the size of the Stack.
func (s *Stack) Size() int32 {
	return int32(len(s.data))
}
