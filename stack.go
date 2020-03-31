package simplestack

type stackItem struct {
	data interface{}
}

type stack []*stackItem

// New creates a new stack.
func New(capacity int) *stack {
	s := make(stack, 0, capacity)
	return &s
}

// Push adds a value to the stack.
func (s *stack) Push(v interface{}) {
	item := stackItem{
		data: v,
	}
	*s = append(*s, &item)
}

// Peek returns the value at the top of the stack
// without removing this value.
func (s *stack) Peek() *stackItem {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

// Pop returns the value at the top of the stack and removes it.
func (s *stack) Pop() *stackItem {
	if len(*s) == 0 {
		return nil
	}
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}
