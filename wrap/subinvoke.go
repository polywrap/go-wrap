package wrap

import (
	"errors"
	"unsafe"
)

//go:wasm-module wrap
//export __wrap_subinvoke
func __wrap_subinvoke(uriPtr, uriLen, methodPtr, methodLen, argsPtr, argsLen uint32) bool

// Subinvoke Result

//go:wasm-module wrap
//export __wrap_subinvoke_result_len
func __wrap_subinvoke_result_len() uint32

//go:wasm-module wrap
//export __wrap_subinvoke_result
func __wrap_subinvoke_result(ptr uint32)

// Subinvoke Error

//go:wasm-module wrap
//export __wrap_subinvoke_error_len
func __wrap_subinvoke_error_len() uint32

//go:wasm-module wrap
//export __wrap_subinvoke_error
func __wrap_subinvoke_error(ptr uint32)

func WrapSubinvoke(uri, method string, args []byte) ([]byte, error) {
	WrapDebugLog("up here")
	uriPtr := unsafe.Pointer(&uri)
	methodPtr := unsafe.Pointer(&method)
	argsPtr := unsafe.Pointer(&args)
	WrapDebugLog("up here")

	result := __wrap_subinvoke(*(*uint32)(uriPtr), uint32(len(uri)), *(*uint32)(methodPtr), uint32(len(method)),
		*(*uint32)(argsPtr), uint32(len(args)))
	WrapDebugLog("down here")

	if !result {
		WrapDebugLog("bad result")
		errorLen := __wrap_subinvoke_error_len()
		WrapDebugLog("bad result")
		errorBuf := make([]byte, errorLen)
		WrapDebugLog("bad result")
		errorPtr := unsafe.Pointer(&errorBuf)
		WrapDebugLog("bad result")

		__wrap_subinvoke_error(*(*uint32)(errorPtr))
		WrapDebugLog("bad result")
		return nil, errors.New(string(errorBuf))
	}

	WrapDebugLog("result")
	resultLen := __wrap_subinvoke_result_len()
	WrapDebugLog("result")
	resultBuf := make([]byte, resultLen)
	WrapDebugLog("result")
	resultPtr := unsafe.Pointer(&resultBuf)
	WrapDebugLog("result")

	__wrap_subinvoke_result(*(*uint32)(resultPtr))
	WrapDebugLog("result")
	return resultBuf, nil
}
