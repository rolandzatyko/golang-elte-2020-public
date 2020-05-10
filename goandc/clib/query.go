package main

//go:generate gcc -g -Wall -c query.c

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: query.o
#include <stdlib.h>
#include "query.h"
*/
import "C"

import (
	"flag"
	"fmt"
	"unsafe"
)

var forSize [1]*C.char

func Query(q string) ([]string, error) {
	cq := C.CString(q)
	defer C.free(unsafe.Pointer(cq))
	var respout **C.char
	l := C.query(cq, &respout)
	if l <= 0 {
		return nil, fmt.Errorf("%q bad query", q)
	}
	var resp []string
	for i := 0; i < int(l); i++ {
		respi := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(respout)) + uintptr(i)*unsafe.Sizeof(forSize[0])))
		resp = append(resp, C.GoString(respi))
		C.free(unsafe.Pointer(respi))
	}
	C.free(unsafe.Pointer(respout))
	return resp, nil
}

func main() {
	flag.Parse()
	for _, q := range flag.Args() {
		resp, err := Query(q)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			continue
		}
		fmt.Printf("%q -> %q\n", q, resp)
	}
}
