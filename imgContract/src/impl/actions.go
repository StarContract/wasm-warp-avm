package impl

import (
	"errors"
	"github.com/redstone-finance/redstone-contracts-wasm/go/src/types"
	// "image"
	// "image/png"
	// "encoding/base64"
	// "log"
	// "strings"
	// "bytes"
)

func UpLoad(state types.ImgState, action types.UpLoadAction) (*types.ImgState, error) {
	state.ImageData = action.ImageData
	state.Name = action.Name
	if action.Name == "pig" {
		return nil, errors.New("[CE:ITQ] not pig")
	}
	return &state, nil
}

func CropImg(state types.ImgState, action types.CropImgAction) (*types.ImgState, error) {
	// base64String := state.ImageData
	
	// data, err := base64.StdEncoding.DecodeString(base64String)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// reader := strings.NewReader(string(data))

	// img, _, err := image.Decode(reader)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// width := img.Bounds().Max.X
	// height := img.Bounds().Max.Y

	// newWidth := width / 2
	// newHeight := height / 2

	// croppedImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// for y := 0; y < newHeight; y++ {
	// 	for x := 0; x < newWidth; x++ {
	// 		croppedImg.Set(x, y, img.At(x*2, y*2))
	// 	}
	// }

	// buffer := new(bytes.Buffer)

	// png.Encode(buffer, croppedImg)

	// crop_base64String := base64.StdEncoding.EncodeToString(buffer.Bytes())

	//state.ImageData = crop_base64String
 
	state.ImageData = action.ImageData
	state.Name = action.Name

	if state.Name == "pig" {
		return nil, errors.New("[CE:ITQ] not pig")
	}

	return &state, nil
}
