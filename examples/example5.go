package main


import "github.com/vulogov/tc2"

func main() {
  tc := tc2.Init()
  tc.Eval("$abc #1 #cde.cde")
}
