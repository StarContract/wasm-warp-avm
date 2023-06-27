package contract

import (
	"github.com/StarContract/wasm-warp-avm/src/common/imports"
	"syscall/js"
)

func Id() string {
	return importContract().Call("id").String()
}

func Owner() string {
	return importContract().Call("owner").String()
}

func importContract() js.Value {
	return imports.RedStone().Get("Contract")
}
