package tc2

import (

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
