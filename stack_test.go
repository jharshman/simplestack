package simplestack

import (
	"reflect"
	"testing"
)

const CAP = 10

var (
	tests = []struct {
		name         string
		Data         []StackItem
		length       int
		expectedItem *StackItem
	}{
		{
			name: "small stack same type",
			Data: []StackItem{
				{Data: "hello"},
				{Data: "world"},
			},
			length:       2,
			expectedItem: &StackItem{Data: "world"},
		},
		{
			name: "medium stack same type",
			Data: []StackItem{
				{Data: "hello"},
				{Data: "world"},
				{Data: "gophers"},
				{Data: "are"},
				{Data: "cute"},
			},
			length:       5,
			expectedItem: &StackItem{Data: "cute"},
		},
		{
			name: "medium stack different types",
			Data: []StackItem{
				{Data: "hello"},
				{Data: 1},
				{Data: 1.50},
				{Data: []float64{.1, .2, .3}},
				{Data: []string{"gophers", "are", "cute"}},
			},
			length:       5,
			expectedItem: &StackItem{Data: []string{"gophers", "are", "cute"}},
		},
	}
)

func TestStack_Push(t *testing.T) {
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			s := &Stack{}
			for _, item := range v.Data {
				s.Push(item.Data)
			}

			l := len(*s)
			if v.length != l {
				t.Errorf("expected %d got %d\n", v.length, l)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			s := &Stack{}
			for _, item := range v.Data {
				s.Push(item.Data)
			}

			poppedItem := s.Pop()

			l := len(*s)
			if v.length-1 != l {
				t.Errorf("expected %d got %v\n", v.length-1, l)
			}

			if !reflect.DeepEqual(v.expectedItem, poppedItem) {
				t.Errorf("expected %v got %v\n", v.expectedItem, poppedItem)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			s := &Stack{}
			for _, item := range v.Data {
				s.Push(item.Data)
			}

			peekedItem := s.Peek()
			if !reflect.DeepEqual(v.expectedItem, peekedItem) {
				t.Errorf("expected %+v got %+v\n", v.expectedItem, peekedItem)
			}
		})
	}
}

func TestStack_PopOnEmpty(t *testing.T) {
	s := &Stack{}
	v := s.Pop()
	if v != nil {
		t.Errorf("expected nil, got %v\n", v)
	}
}

func TestStack_PeekOnEmpty(t *testing.T) {
	s := &Stack{}
	v := s.Peek()
	if v != nil {
		t.Errorf("expected nil, got %v\n", v)
	}
}
