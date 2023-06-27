package main

import (
	"github.com/StarContract/wasm-warp-avm/src/common"
	"github.com/StarContract/wasm-warp-avm/src/impl"
)

// contract - implementation of the SwContract interface
var contract = impl.PstContract{}

// handles all the WASM-JS related trickery...
func main() {
	common.Run(&contract)
}
