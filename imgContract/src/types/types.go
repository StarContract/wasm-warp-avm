package types

import (
	"github.com/redstone-finance/redstone-contracts-wasm/go/src/common_types"
)

type ImgState struct {
	ImageData    string            `json:"imageData"`
	Name         string            `json:"name"`
}

type UpLoadAction struct {
	common_types.Action
	ImageData    string            `json:"imageData"`
	Name         string            `json:"name"`
}


type CropImgAction struct {
	common_types.Action
	ImageData    string            `json:"imageData"`
	Name         string            `json:"name"`
}