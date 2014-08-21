package foo

// #include "ick.h"
// #cgo CFLAGS: -DFIXED_SUB=12
import "C"

func baz(i int) int {
     return int(C.ick(C.int(i)))
     //return i-12
}
