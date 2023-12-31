package transaction

import (
	"github.com/StarContract/wasm-warp-avm/src/common/imports"
	"syscall/js"
)

func Id() string {
	return importTransaction().Call("id").String()
}

func Owner() string {
	return importTransaction().Call("owner").String()
}

func Target() string {
	return importTransaction().Call("target").String()
}

func importTransaction() js.Value {
	return imports.RedStone().Get("Transaction")
}
