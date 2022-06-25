package tc2

import (

)

func TCExitCbPoolClose(l *TCExecListener) {
  l.TC.Debugf("Closing workers pool. Capacity=%v, Active=%v", l.TC.Pool.PoolSize(), l.TC.Pool.ActiveWorkers())
  l.TC.Pool.Close()
}

func TCExitCbSetExitError(l *TCExecListener) {
  l.TC.Debugf("Set an error state for TC")
  l.TC.MakeError("[ BUND ] instance has been terminated")
}
