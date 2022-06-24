package main


import "github.com/vulogov/tc2"

func main() {
  tc := tc2.Init()
  tc.Eval(`"42"[3.14]`)
}
