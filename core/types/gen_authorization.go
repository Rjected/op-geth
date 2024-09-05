// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*authorizationMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (a Authorization) MarshalJSON() ([]byte, error) {
	type Authorization struct {
		ChainID *hexutil.Big
		Address common.Address `json:"address" gencodec:"required"`
		Nonce   hexutil.Uint64 `json:"nonce" gencodec:"required"`
		YParity *hexutil.Big   `json:"yParity" gencodec:"required"`
		R       *hexutil.Big   `json:"r" gencodec:"required"`
		S       *hexutil.Big   `json:"s" gencodec:"required"`
	}
	var enc Authorization
	enc.ChainID = (*hexutil.Big)(a.ChainID)
	enc.Address = a.Address
	enc.Nonce = hexutil.Uint64(a.Nonce)
	enc.YParity = (*hexutil.Big)(a.YParity)
	enc.R = (*hexutil.Big)(a.R)
	enc.S = (*hexutil.Big)(a.S)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (a *Authorization) UnmarshalJSON(input []byte) error {
	type Authorization struct {
		ChainID *hexutil.Big
		Address *common.Address `json:"address" gencodec:"required"`
		Nonce   *hexutil.Uint64 `json:"nonce" gencodec:"required"`
		YParity *hexutil.Big    `json:"yParity" gencodec:"required"`
		R       *hexutil.Big    `json:"r" gencodec:"required"`
		S       *hexutil.Big    `json:"s" gencodec:"required"`
	}
	var dec Authorization
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ChainID != nil {
		a.ChainID = (*big.Int)(dec.ChainID)
	}
	if dec.Address == nil {
		return errors.New("missing required field 'address' for Authorization")
	}
	a.Address = *dec.Address
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' for Authorization")
	}
	a.Nonce = uint64(*dec.Nonce)
	if dec.YParity == nil {
		return errors.New("missing required field 'yParity' for Authorization")
	}
	a.YParity = (*big.Int)(dec.YParity)
	if dec.R == nil {
		return errors.New("missing required field 'r' for Authorization")
	}
	a.R = (*big.Int)(dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' for Authorization")
	}
	a.S = (*big.Int)(dec.S)
	return nil
}