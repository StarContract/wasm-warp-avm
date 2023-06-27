package imports

import (
	"github.com/StarContract/wasm-warp-avm/src/common"
	"syscall/js"
)

func RedStone() js.Value {
	return js.Global().
		Get("redstone").
		Get("go").
		Get(common.GetWasmInstance().ModuleId).
		Get("imports")
}
