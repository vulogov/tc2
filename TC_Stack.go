package tc2

import (
  "fmt"
  "errors"
  "github.com/google/uuid"
)

func (tc *TCstate) AddNewStack(name string) {
  tc.console.Debug("Creating new stack:", name)
  if tc.ResN.Contains(name) {
    tc.console.Warning("Can not create new stack as it is already exists:", name)
    return
  }
  chcap, err := tc.GetVariable("tc.Chancapacity")
  if err != nil {
    chcap = 4096
  }
  tc.ResN.Add(name)
  tc.ResNames.PushFront(name)
  tc.Res.Add()
  tc.StackList.Delete(name)
  tc.StackChan.Delete(name)
  tc.StackChan.Store(name, make(chan interface{}, chcap.(int)))
  tc.StackList.Store(name, tc.Res.Q())
}

func (tc *TCstate) DelStack(name string) {
  tc.Debugf("Dropping stack: %v", name)
  if tc.ResN.Contains(name) {
    tc.ResNames.PopFront()
    tc.ResN.Remove(name)
  }
  tc.StackList.Delete(name)
  tc.StackChan.Delete(name)
  tc.Res.Del()
  if tc.Res.GLen() == 0 {
    name = uuid.NewString()
    tc.Debugf("Stack is empty adter last delete, creating new empty stack: %v", name)
    tc.AddNewStack(name)
  }
}

func (tc *TCstate) DropLastStack() error {
  if tc.ResNames.Len() == 0 {
    return errors.New("TwoStack is empty")
  }
  name := tc.ResNames.PopFront().(string)
  tc.DelStack(name)
  return nil
}

func (tc *TCstate) RotateStackNames() {
  n := tc.ResNames.PopFront()
  tc.Debugf("Positioning stack through: %v", n)
  tc.ResNames.PushBack(n)
}

func (tc *TCstate) PositionStack(name string) error {
  if tc.ResNames.Len() == 0 {
    return errors.New("TwoStack is empty")
  }
  if tc.ResN.Contains(name) == false {
    return errors.New(fmt.Sprintf("Stack not found for postioning: %v", name))
  }
  for {
    if tc.ResNames.Front().(string) == name {
      tc.Debugf("Stack positioned: %v", name)
      return nil
    }
    tc.Res.Left()
    tc.RotateStackNames()
  }
}

func (tc *TCstate) StacksLeft(n int) {
  for x := 0; x < n; x++ {
    tc.RotateStackNames()
    tc.Res.Left()
  }
  tc.Debugf("Current stack: %v", tc.ResNames.Front().(string))
}

func (tc *TCstate) StacksRight(n int) {
  for x := 0; x < n; x++ {
    tc.RotateStackNames()
    tc.Res.Right()
  }
  tc.Debugf("Current stack: %v", tc.ResNames.Front().(string))
}

func (tc *TCstate) StackLeft(n int) {
  for x := 0; x < n; x++ {
    tc.Res.CLeft()
  }
  tc.Debugf("Current stack: %v", tc.ResNames.Front().(string))
}

func (tc *TCstate) StackRight(n int) {
  for x := 0; x < n; x++ {
    tc.Res.CRight()
  }
  tc.Debugf("Current stack: %v", tc.ResNames.Front().(string))
}
