package tc2

import (
  "fmt"
  "errors"
)



func (tc *TCstate) SetVariable(name string, data interface{}) {
  tc.Debug("variable",data,"->", name)
  tc.Vars.Delete(name)
  tc.Vars.Store(name, data)
}

func (tc *TCstate) GetVariable(name string) (interface{}, error) {
  if data, ok := tc.Vars.Load(name); ok {
    return data, nil
  }
  return nil, errors.New(fmt.Sprintf("Variable %v not found", name))
}

func (tc *TCstate) HasVariable(name string) bool {
  if _, ok := tc.Vars.Load(name); ok {
    return true
  }
  return false
}
