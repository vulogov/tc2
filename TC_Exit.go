package tc2

import (

)

type TCExitCbFun func(*TCExecListener)

func (l *TCExecListener) ExitRequest() {
  l.TC.Debugf("Exit requesting. Capacity=%v, Active=%v", l.TC.Pool.PoolSize(), l.TC.Pool.ActiveWorkers())
  l.TC.IsExitReq = true
  l.TC.ExReq <- true
}

func (l *TCExecListener) ExitRequested() bool {
  if l.TC.IsExitReq {
    l.TC.IsExitReq = true
    l.TC.ExReq <- true
    return true
  }
  if len(l.TC.ExReq) > 0 {
    e := <- l.TC.ExReq
    if e == true {
      l.TC.Debugf("Exit requested. Capacity=%v, Active=%v", l.TC.Pool.PoolSize(), l.TC.Pool.ActiveWorkers())
      l.TC.IsExitReq = true
      l.TC.ExReq <- true
      return true
    }
  }
  return false
}

func (l *TCExecListener) RunExitCallbacks() {
  for cb := range l.TC.ExitCb.Items() {
    l.TC.Debugf("Running exit callback: %v", cb.Key)
    cb.Value.(TCExitCbFun)(l)
  }
}

func (l *TCExecListener) RegisterExitCallback(name string, fun TCExitCbFun) {
  l.TC.RegisterExitCallback(name, fun)
}

func (tc *TCstate) RegisterExitCallback(name string, fun TCExitCbFun) {
  tc.Debugf("Reistering exit callback: %v", name)
  tc.ExitCb.Del(name)
  tc.ExitCb.Set(name, fun)
}

func InitExitCallbacks(tc *TCstate) {
  tc.Debugf("Initialize exit callbacks")
  tc.RegisterExitCallback("closePool", TCExitCbPoolClose)
}

func (tc *TCstate) ExitRequested() bool {
  if tc.IsExitReq {
    return true
  }
  if len(tc.ExReq) > 0 {
    e := <- tc.ExReq
    if e == true {
      tc.IsExitReq = true
      tc.ExReq <- true
      return true
    }
  }
  return false
}
