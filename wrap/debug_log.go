package wrap

import "unsafe"

//go:wasm-module wrap
//export __wrap_debug_log
func __wrap_debug_log(ptr, length uint32)

func WrapDebugLog(msg string) {
	msgBuf := []byte(msg)
	__wrap_debug_log(*(*uint32)(unsafe.Pointer(&msgBuf)), uint32(len(msgBuf)))
}
