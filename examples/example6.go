package main


import "github.com/vulogov/tc2"

func main() {
  tc := tc2.Init()
  tc.Eval("[main] a b c  ;; [2] a b c  ;; [main] v f g ;; ")
}
