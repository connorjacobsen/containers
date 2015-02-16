package stack

type Element struct {
	// The stack to which this element belongs.
	stack *Stack

	// The next element in the stack
	// i.e., the one that was pushed before this element.
	next *Element

	// The value stored with this element.
	Value interface{}
}

type Stack struct {
	top  *Element
	size int
}

// Init initializes or clears stack s.
func (s *Stack) Init() *Stack {
	s.top = s.top
	s.size = 0
	return s
}

// Size returns the size of the stack.
func (s *Stack) Size() int {
	return s.size
}

// New returns an initialized Stack.
func New() *Stack { return new(Stack).Init() }

// Peek returns the elem at the top of the stack, without removing
// it from the stack.
func (s *Stack) Peek() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.top.Value
}

// Places e at the top of the stack.
func (s *Stack) add(e *Element) *Element {
	n := s.top
	e.next = n
	e.stack = s
	s.top = e
	s.size++
	return e
}

// Push adds a value to the top of the stack.
// Returns the value added to the stack.
func (s *Stack) Push(value interface{}) interface{} {
	s.add(&Element{Value: value})
	return s.top.Value
}

// Pop removes the value at the top of the stack and returns it.
func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	el := s.top
	s.top = el.next
	s.size--
	return el.Value
}
