// Code generated by compose-gen.
// DO NOT EDIT!

package compose

import (
	"github.com/katydid/katydid/expr/ast"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/types"
)

func composeFloat64(expr *expr.Expr) (funcs.Float64, error) {
	uniq, err := prep(expr, types.SINGLE_DOUBLE)
	if err != nil {
		return nil, err
	}
	if expr.Terminal != nil {
		if expr.GetTerminal().Variable != nil {
			return funcs.NewFloat64Variable(), nil
		} else {
			return funcs.NewConstFloat64(expr.GetTerminal().GetDoubleValue()), nil
		}
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewFloat64Func(uniq, values...)
}

func composeFloat64s(expr *expr.Expr) (funcs.Float64s, error) {
	uniq, err := prep(expr, types.LIST_DOUBLE)
	if err != nil {
		return nil, err
	}
	if expr.List != nil {
		vs, err := newValues(expr.GetList().GetElems())
		if err != nil {
			return nil, err
		}
		bs := make([]funcs.Float64, len(vs))
		var ok bool
		for i := range vs {
			bs[i], ok = vs[i].(funcs.Float64)
			if !ok {
				return nil, &errExpected{types.SINGLE_DOUBLE.String(), expr.String()}
			}
		}
		return funcs.NewListOfFloat64(bs), nil
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewFloat64sFunc(uniq, values...)
}

func composeInt64(expr *expr.Expr) (funcs.Int64, error) {
	uniq, err := prep(expr, types.SINGLE_INT)
	if err != nil {
		return nil, err
	}
	if expr.Terminal != nil {
		if expr.GetTerminal().Variable != nil {
			return funcs.NewInt64Variable(), nil
		} else {
			return funcs.NewConstInt64(expr.GetTerminal().GetIntValue()), nil
		}
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewInt64Func(uniq, values...)
}

func composeInt64s(expr *expr.Expr) (funcs.Int64s, error) {
	uniq, err := prep(expr, types.LIST_INT)
	if err != nil {
		return nil, err
	}
	if expr.List != nil {
		vs, err := newValues(expr.GetList().GetElems())
		if err != nil {
			return nil, err
		}
		bs := make([]funcs.Int64, len(vs))
		var ok bool
		for i := range vs {
			bs[i], ok = vs[i].(funcs.Int64)
			if !ok {
				return nil, &errExpected{types.SINGLE_INT.String(), expr.String()}
			}
		}
		return funcs.NewListOfInt64(bs), nil
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewInt64sFunc(uniq, values...)
}

func composeUint64(expr *expr.Expr) (funcs.Uint64, error) {
	uniq, err := prep(expr, types.SINGLE_UINT)
	if err != nil {
		return nil, err
	}
	if expr.Terminal != nil {
		if expr.GetTerminal().Variable != nil {
			return funcs.NewUint64Variable(), nil
		} else {
			return funcs.NewConstUint64(expr.GetTerminal().GetUintValue()), nil
		}
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewUint64Func(uniq, values...)
}

func composeUint64s(expr *expr.Expr) (funcs.Uint64s, error) {
	uniq, err := prep(expr, types.LIST_UINT)
	if err != nil {
		return nil, err
	}
	if expr.List != nil {
		vs, err := newValues(expr.GetList().GetElems())
		if err != nil {
			return nil, err
		}
		bs := make([]funcs.Uint64, len(vs))
		var ok bool
		for i := range vs {
			bs[i], ok = vs[i].(funcs.Uint64)
			if !ok {
				return nil, &errExpected{types.SINGLE_UINT.String(), expr.String()}
			}
		}
		return funcs.NewListOfUint64(bs), nil
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewUint64sFunc(uniq, values...)
}

func composeBool(expr *expr.Expr) (funcs.Bool, error) {
	uniq, err := prep(expr, types.SINGLE_BOOL)
	if err != nil {
		return nil, err
	}
	if expr.Terminal != nil {
		if expr.GetTerminal().Variable != nil {
			return funcs.NewBoolVariable(), nil
		} else {
			return funcs.NewConstBool(expr.GetTerminal().GetBoolValue()), nil
		}
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewBoolFunc(uniq, values...)
}

func composeBools(expr *expr.Expr) (funcs.Bools, error) {
	uniq, err := prep(expr, types.LIST_BOOL)
	if err != nil {
		return nil, err
	}
	if expr.List != nil {
		vs, err := newValues(expr.GetList().GetElems())
		if err != nil {
			return nil, err
		}
		bs := make([]funcs.Bool, len(vs))
		var ok bool
		for i := range vs {
			bs[i], ok = vs[i].(funcs.Bool)
			if !ok {
				return nil, &errExpected{types.SINGLE_BOOL.String(), expr.String()}
			}
		}
		return funcs.NewListOfBool(bs), nil
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewBoolsFunc(uniq, values...)
}

func composeString(expr *expr.Expr) (funcs.String, error) {
	uniq, err := prep(expr, types.SINGLE_STRING)
	if err != nil {
		return nil, err
	}
	if expr.Terminal != nil {
		if expr.GetTerminal().Variable != nil {
			return funcs.NewStringVariable(), nil
		} else {
			return funcs.NewConstString(expr.GetTerminal().GetStringValue()), nil
		}
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewStringFunc(uniq, values...)
}

func composeStrings(expr *expr.Expr) (funcs.Strings, error) {
	uniq, err := prep(expr, types.LIST_STRING)
	if err != nil {
		return nil, err
	}
	if expr.List != nil {
		vs, err := newValues(expr.GetList().GetElems())
		if err != nil {
			return nil, err
		}
		bs := make([]funcs.String, len(vs))
		var ok bool
		for i := range vs {
			bs[i], ok = vs[i].(funcs.String)
			if !ok {
				return nil, &errExpected{types.SINGLE_STRING.String(), expr.String()}
			}
		}
		return funcs.NewListOfString(bs), nil
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewStringsFunc(uniq, values...)
}

func composeBytes(expr *expr.Expr) (funcs.Bytes, error) {
	uniq, err := prep(expr, types.SINGLE_BYTES)
	if err != nil {
		return nil, err
	}
	if expr.Terminal != nil {
		if expr.GetTerminal().Variable != nil {
			return funcs.NewBytesVariable(), nil
		} else {
			return funcs.NewConstBytes(expr.GetTerminal().GetBytesValue()), nil
		}
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewBytesFunc(uniq, values...)
}

func composeListOfBytes(expr *expr.Expr) (funcs.ListOfBytes, error) {
	uniq, err := prep(expr, types.LIST_BYTES)
	if err != nil {
		return nil, err
	}
	if expr.List != nil {
		vs, err := newValues(expr.GetList().GetElems())
		if err != nil {
			return nil, err
		}
		bs := make([]funcs.Bytes, len(vs))
		var ok bool
		for i := range vs {
			bs[i], ok = vs[i].(funcs.Bytes)
			if !ok {
				return nil, &errExpected{types.SINGLE_BYTES.String(), expr.String()}
			}
		}
		return funcs.NewListOfBytes(bs), nil
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewListOfBytesFunc(uniq, values...)
}
