package impl

import (
	"errors"
	"github.com/redstone-finance/redstone-contracts-wasm/go/src/common_types"
	"github.com/redstone-finance/redstone-contracts-wasm/go/src/types"
)

type ImgContract struct {
	state types.ImgState
}


func (c *ImgContract) Handle(action common_types.Action, actionBytes []byte) (interface{}, common_types.ActionResult, error) {
	fn := action.Function

	// this is how you can access functions imported from js
	/*
		console.Log("Block height", block.Height())
		console.Log("Block indep_hash", block.IndepHash())
		console.Log("Block timestamp", block.Timestamp())

		console.Log("Transaction id", transaction.Id())
		console.Log("Transaction owner", transaction.Owner())
		console.Log("Transaction target", transaction.Target())

		console.Log("Contract id", contract.Id())
		console.Log("Contract owner", contract.Owner())
	*/

	clonedState := c.CloneState().(types.ImgState)

	switch fn {
	case "upLoad":
		var upLoad types.UpLoadAction
		err := upLoad.UnmarshalJSON(actionBytes)
		state, err := UpLoad(clonedState, upLoad)
		return state, nil, err
	case "cropImg":
		var cropImg types.CropImgAction
		err := cropImg.UnmarshalJSON(actionBytes)
		state, err := CropImg(clonedState, cropImg)
		return state, nil, err
	default:
		return nil, nil, errors.New("[RE:WTF] unknown function: " + fn)
	}
}

func (c *ImgContract) InitState(stateJson string) {
	var state types.ImgState
	err := state.UnmarshalJSON([]byte(stateJson))
	if err != nil {
		return // TODO: throw in a similar way as in handle
	}
	c.UpdateState(&state)
}

func (c *ImgContract) UpdateState(newState interface{}) {
	// note: we're first type asserting here to the pointer to types.ImgState
	// - and then retrieving value from the pointer
	c.state = *(newState.(*types.ImgState))
}

func (c *ImgContract) CurrentState() interface{} {
	return c.state
}

func (c *ImgContract) CloneState() interface{} {
	json, _ := c.state.MarshalJSON()
	state := types.ImgState{}
	err := state.UnmarshalJSON(json)
	if err != nil {
		// TODO: return error
		return types.ImgState{}
	}

	return state
}
