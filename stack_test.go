package simplestack

import (
	"reflect"
	"testing"
)

const CAP = 10

var (
	tests = []struct {
		name         string
		data         []stackItem
		length       int
		expectedItem *stackItem
	}{
		{
			name: "small stack same type",
			data: []stackItem{
				{data: "hello"},
				{data: "world"},
			},
			length:       2,
			expectedItem: &stackItem{data: "world"},
		},
		{
			name: "medium stack same type",
			data: []stackItem{
				{data: "hello"},
				{data: "world"},
				{data: "gophers"},
				{data: "are"},
				{data: "cute"},
			},
			length:       5,
			expectedItem: &stackItem{data: "cute"},
		},
		{
			name: "medium stack different types",
			data: []stackItem{
				{data: "hello"},
				{data: 1},
				{data: 1.50},
				{data: []float64{.1, .2, .3}},
				{data: []string{"gophers", "are", "cute"}},
			},
			length: 5,
			expectedItem: &stackItem{data: []string{"gophers", "are", "cute"}},
		},
	}
)

func TestStack_Push(t *testing.T) {
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			s := New(CAP)
			for _, item := range v.data {
				s.Push(item.data)
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
			s := New(CAP)
			for _, item := range v.data {
				s.Push(item.data)
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
			s := New(CAP)
			for _, item := range v.data {
				s.Push(item.data)
			}

			peekedItem := s.Peek()
			if !reflect.DeepEqual(v.expectedItem, peekedItem) {
				t.Errorf("expected %+v got %+v\n", v.expectedItem, peekedItem)
			}
		})
	}
}

func TestStack_PopOnEmpty(t *testing.T) {
	s := New(CAP)
	v := s.Pop()
	if v != nil {
		t.Errorf("expected nil, got %v\n", v)
	}
}

func TestStack_PeekOnEmpty(t *testing.T) {
	s := New(CAP)
	v := s.Peek()
	if v != nil {
		t.Errorf("expected nil, got %v\n", v)
	}
}