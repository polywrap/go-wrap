package main

import (
	"github.com/polywrap/go-wrap/examples/demo1/wrap/module"
	"github.com/polywrap/go-wrap/polywrap"
)

//export _wrap_invoke
func _wrap_invoke(methodSize, argsSize, envSize uint32) bool {
	args := polywrap.WrapInvokeArgs(methodSize, argsSize)

	if args.Method == "sampleMethod" {
		return polywrap.WrapInvoke(args, envSize, module.SampleMethodWrapped)
	} else {
		return polywrap.WrapInvoke(args, envSize, nil)
	}
}

func main() {
}
