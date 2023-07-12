package wrap

import (
	"errors"
	"unsafe"

	"github.com/polywrap/go-wrap/msgpack"
)

//go:wasm-module wrap
//export __wrap_getImplementations
func __wrap_getImplementations(uriPtr, uriLen uint32) bool

//go:wasm-module wrap
//export __wrap_getImplementations_result_len
func __wrap_getImplementations_result_len() uint32

//go:wasm-module wrap
//export __wrap_getImplementations_result
func __wrap_getImplementations_result(ptr uint32)

func WrapGetImplementations(uri string) []string {
	uriPtr := unsafe.Pointer(&uri)

	success := __wrap_getImplementations(*(*uint32)(uriPtr), uint32(len(uri)))
	if !success {
		return []string{}
	}

	resultLen := __wrap_getImplementations_result_len()
	resultBuf := make([]byte, resultLen)
	resultPtr := unsafe.Pointer(&resultBuf)

	__wrap_getImplementations_result(*(*uint32)(resultPtr))

	// Deserialize the msgpack buffer,
	// which contains an array of strings
	ctx := msgpack.NewContext("Deserializing __wrap_getImplementations_result buffer")
	reader := msgpack.NewReadDecoder(ctx resultBuf)

	ln := reader.ReadArrayLength()
	result := make([]string, ln)
	for i := uint32(0); i < ln; i++ {
		result[i] = reader.ReadString()
	}

	return result
}
