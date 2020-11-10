package main

import(
  "math/rand"
	"time"
  "github.com/p0tpan/ass2ascii"
)

func main() {
 rand.Seed(time.Now().UnixNano())
 ass2Ascii.AssPrinter()
}
