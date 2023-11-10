package main

import "runtime"
import "fmt"

type S struct {
  bs [1024]uint8
}

func wasteSpace() *S {
  ss := make([]S, 1024 * 1024 * 128)
  return &ss[2]
}

func main() {
  const LOOP_ITERS = 1024 * 1024
  hold := make([]*S, 0, LOOP_ITERS)
  for i := 0; i < LOOP_ITERS; i++ {
    hold = append(hold, wasteSpace())
    runtime.GC()
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Iter %8d: Alloc: %12d MiB, HeapAlloc: %12d MiB\n", i, m.Alloc / 1024 / 1024, m.HeapAlloc / 1024 / 1024)
  }
}
