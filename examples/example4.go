package main


import "github.com/vulogov/tc2"

func main() {
  tc := tc2.Init()
  tc.Eval(": func : a c b ;; [] : funcB : d d ;; a b c d ;; [Abc] :: v v  ;; [Cde] c 1 2 3 ;;  c d f ;;")
}
