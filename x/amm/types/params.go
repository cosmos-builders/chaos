package types

import (
	"fmt"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramstypes.ParamSet = (*Params)(nil)

var (
	KeyFeeRate             = []byte("FeeRate")
	KeyMinInitialLiquidity = []byte("MinInitialLiquidity")
)

var (
	DefaultFeeRate             = sdk.NewDecWithPrec(3, 3) // 0.3%
	DefaultMinInitialLiquidity = sdk.NewInt(1000)
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(feeRate sdk.Dec, minInitialLiquidity sdk.Int) Params {
	return Params{
		FeeRate:             feeRate,
		MinInitialLiquidity: minInitialLiquidity,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultFeeRate, DefaultMinInitialLiquidity)
}

// ParamSetPairs get the params.ParamSet
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyFeeRate, &params.FeeRate, validateFeeRate),
		paramstypes.NewParamSetPair(KeyMinInitialLiquidity, &params.MinInitialLiquidity, validateMinInitialLiquidity),
	}
}

// Validate validates the set of params
func (params Params) Validate() error {
	if err := validateFeeRate(params.FeeRate); err != nil {
		return err
	}
	if err := validateMinInitialLiquidity(params.MinInitialLiquidity); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (params Params) String() string {
	out, _ := yaml.Marshal(params)
	return string(out)
}

func validateFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("fee rate must not be negative: %s", v)
	}
	return nil
}

func validateMinInitialLiquidity(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("min initial liquidity must not be negative: %s", v)
	}
	return nil
}
