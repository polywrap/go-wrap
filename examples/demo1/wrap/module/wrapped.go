package module

import "github.com/polywrap/go-wrap/examples/demo1"

func SampleMethodWrapped(argsBuf []byte, envSize uint32) []byte {
	args := deserializeSampleMethodArgs(argsBuf)

	result := demo1.SampleMethod(args)

	return serializeSampleMethodResult(result)
}
