package calculator

import (
	"errors"
	"go/constant"
	"go/token"
	"go/types"
)


func Calculate(expression string) (float64, error) {
	tv, err := types.Eval(token.NewFileSet(), nil, token.NoPos, expression)
	if err != nil {
		return 0, errors.New("invalid expression")
	}

	if tv.Value == nil {
		return 0, errors.New("invalid expression")
	}

	value, _ := constant.Float64Val(tv.Value)
	return value, nil
}