package simplestack

type StackItem struct {
	Data interface{}
}

type Stack []*StackItem

// Push adds a value to the stack.
func (s *Stack) Push(v interface{}) {
	item := StackItem{
		Data: v,
	}
	*s = append(*s, &item)
}

// Peek returns the value at the top of the stack
// without removing this value.
func (s *Stack) Peek() *StackItem {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

// Pop returns the value at the top of the stack and removes it.
func (s *Stack) Pop() *StackItem {
	if len(*s) == 0 {
		return nil
	}
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}
