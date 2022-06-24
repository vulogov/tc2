package tc2

import (
	"fmt"
	"github.com/gammazero/deque"
	"github.com/google/uuid"
	"github.com/lrita/cmap"
)

type TwoStack struct {
	tc    *TCstate
	R      deque.Deque[interface{}]
	C      cmap.Cmap
	tmp    cmap.Cmap
	ID     string
	Status bool
	Mode   bool
}

const minCap = 1024

//
// Initialize new 2-dimentional stack structure
//
func InitTS(tc *TCstate) *TwoStack {
	ts := &TwoStack{
		tc:     tc,
		Status: true,
		Mode:   true,
		ID:     uuid.New().String(),
	}
	return ts
}

func (ts *TwoStack) Left() {
	ts.tc.console.Debug("Rotating Global Stack left")
	ts.R.Rotate(-1)
}

func (ts *TwoStack) CLeft() {
	q := ts.Q()
	q.Rotate(-1)
}

func (ts *TwoStack) Right() {
	ts.tc.console.Debug("Rotating Global Stack right")
	ts.R.Rotate(1)
}

func (ts *TwoStack) CRight() {
	q := ts.Q()
	q.Rotate(1)
}

func (ts *TwoStack) Normal() {
	ts.Mode = true
}

func (ts *TwoStack) Reverse() {
	ts.Mode = false
}


func (ts *TwoStack) PopFront() interface{} {
	if ts.R.Len() == 0 {
		ts.R.PushBack(deque.New[interface{}](0, minCap))
		return nil
	}
	return ts.R.Front().(*deque.Deque[interface{}]).PopFront()
}

func (ts *TwoStack) Front() interface{} {
	if ts.R.Len() == 0 {
		ts.R.PushBack(deque.New[interface{}](0, minCap))
	}
	if ts.R.Front().(*deque.Deque[interface{}]).Len() > 0 {
		return ts.R.Front().(*deque.Deque[interface{}]).Front()
	} else {
		return nil
	}
}

func (ts *TwoStack) Set(data interface{}) {
	if ts.R.Len() == 0 {
		ts.R.PushBack(deque.New[interface{}](0, minCap))
	}
	if ts.Mode == true {
		ts.R.Back().(*deque.Deque[interface{}]).PushBack(data)
	} else {
		ts.R.Front().(*deque.Deque[interface{}]).PushFront(data)
	}
}

func (ts *TwoStack) Get() (interface{}, error) {
	ts.tc.console.Debug(fmt.Sprintf("Getting from twostack len=%v, mode=%v, len=%v glen=%v", ts.R.Len(), ts.Mode, ts.Len(), ts.GLen()))
	if ts.R.Len() == 0 {
		return nil, fmt.Errorf("Stack is to shallow for .Get() operation")
	}
	if ts.Mode == true {
		if ts.R.Back().(*deque.Deque[interface{}]).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Back().(*deque.Deque[interface{}]).Back(), nil
	} else {
		if ts.R.Front().(*deque.Deque[interface{}]).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Front().(*deque.Deque[interface{}]).Front(), nil
	}
}

func (ts *TwoStack) Take() (interface{}, error) {
	ts.tc.console.Debug(fmt.Sprintf("Taking from twostack len=%v, mode=%v", ts.R.Len(), ts.Mode))
	if ts.R.Len() == 0 {
		return nil, fmt.Errorf("Stack is to shallow for .Take() operation")
	}
	if ts.Mode == true {
		if ts.R.Back().(*deque.Deque[interface{}]).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Back().(*deque.Deque[interface{}]).PopBack(), nil
	} else {
		if ts.R.Front().(*deque.Deque[interface{}]).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Front().(*deque.Deque[interface{}]).PopFront(), nil
	}
}

func (ts *TwoStack) Add() {
	if ts.Mode == true {
		ts.R.PushBack(deque.New[interface{}](0, minCap))
	} else {
		ts.R.PushFront(deque.New[interface{}](0, minCap))
	}
	ts.tc.console.Debug(fmt.Sprintf("Adding new twostack mode=%v glen=%v", ts.Mode, ts.GLen()))
}

func (ts *TwoStack) addQ(q *deque.Deque[interface{}]) {
	if ts.Mode == true {
		ts.R.PushBack(q)
	} else {
		ts.R.PushFront(q)
	}
	ts.tc.console.Debug("Adding custom twostack mode=", ts.Mode)
}

func (ts *TwoStack) Del() {
	if ts.R.Len() > 0 {
		if ts.Mode == true {
			ts.R.PopBack()
		} else {
			ts.R.PopFront()
		}
	}
	ts.tc.console.Debug("Delete twostack mode=",ts.Mode, "glen=", ts.GLen())
}

func (ts *TwoStack) GLen() int {
	return ts.R.Len()
}

func (ts *TwoStack) Len() int {
	if ts.R.Len() > 0 {
		if ts.Mode == true {
			return ts.R.Back().(*deque.Deque[interface{}]).Len()
		} else {
			return ts.R.Front().(*deque.Deque[interface{}]).Len()
		}
	}
	return 0
}

func (ts *TwoStack) Q() *deque.Deque[interface{}] {
	if ts.R.Len() > 0 {
		if ts.Mode == true {
			return ts.R.Back().(*deque.Deque[interface{}])
		} else {
			return ts.R.Front().(*deque.Deque[interface{}])
		}
	}
	return nil
}
