package console

import (
	"github.com/StarContract/wasm-warp-avm/src/common/imports"
	"syscall/js"
)

func Log(args ...interface{}) {
	importConsole().Call("log", args[0], args[1:])
}

func importConsole() js.Value {
	return imports.RedStone().Get("console")
}
