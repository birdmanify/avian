package main

/*
#cgo LDFLAGS: -llog

#include "bridge.h"
*/
import "C"

import (
	"runtime"
	"config"
	"delegate"
	"tunnel"
)

func main() {
	println("This is avian library code, not a finished product.")
	panic("Please Write other main function by yourself!")
}

//export coreInit
func coreInit(home, versionName C.c_string, sdkVersion C.int) {
	h := C.GoString(home)
	v := C.GoString(versionName)
	s := int(sdkVersion)

	delegate.Init(h, v, s)

	reset()
}

//export reset
func reset() {
	config.LoadDefault()
	tunnel.ResetStatistic()
	tunnel.CloseAllConnections()

	runtime.GC()
}

//export forceGc
func forceGc() {
	go func() {
		log.Infoln("[APP] request force GC")

		runtime.GC()
	}()
}
