package main

import (
	"fmt"
	"hash/fnv"
)

func hashIt(in string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(in))
	out := h.Sum64()
	return out
}

func main() {
	s := "hello"
	fmt.Printf("The FNV64a hash of '%v' is '%v'\n", s, hashIt(s))
}

// go build -gcflags '-m' 03_interface_method.go
/*
# command-line-arguments
./03_interface_method.go:9:17: inlining call to fnv.New64a
./03_interface_method.go:17:12: inlining call to fmt.Printf
./03_interface_method.go:10:9: devirtualizing h.Write to *fnv.sum64a
./03_interface_method.go:11:16: devirtualizing h.Sum64 to *fnv.sum64a
./03_interface_method.go:8:13: in does not escape
./03_interface_method.go:10:16: ([]byte)(in) does not escape
./03_interface_method.go:17:13: s escapes to heap
./03_interface_method.go:17:59: hashIt(s) escapes to heap
./03_interface_method.go:17:12: []interface {}{...} does not escape
<autogenerated>:1: .this does not escape

*/
