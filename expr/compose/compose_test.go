//  Copyright 2013 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package compose

import (
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/expr/ast"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/serialize"
	"github.com/katydid/katydid/types"
	"strings"
	"testing"
)

func TestComposeNot(t *testing.T) {
	expr := &expr.Expr{
		Function: &expr.Function{
			Name: "not",
			Params: []*expr.Expr{
				{
					Terminal: &expr.Terminal{
						BoolValue: proto.Bool(false),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	r, err := b.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	str := funcs.Sprint(b.Func)
	if str != "true" {
		t.Fatalf("trimming did not work: %s", str)
	}
}

//contains(toLower(decString(test.Address.Street.value)), toLower("TheStreet"))

func TestComposeContains(t *testing.T) {
	expr := &expr.Expr{
		Function: &expr.Function{
			Name: "contains",
			Params: []*expr.Expr{
				{
					Function: &expr.Function{
						Name: "toLower",
						Params: []*expr.Expr{
							{
								Terminal: &expr.Terminal{
									Variable: &expr.Variable{
										Type: types.SINGLE_STRING,
									},
								},
							},
						},
					},
				},
				{
					Function: &expr.Function{
						Name: "toLower",
						Params: []*expr.Expr{
							{
								Terminal: &expr.Terminal{
									StringValue: proto.String("TheStreet"),
								},
							},
						},
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	r, err := b.Eval(serialize.NewStringValue("TheStreet"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	r, err = b.Eval(serialize.NewStringValue("ThatStreet"))
	if err != nil {
		panic(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
	if strings.Contains(funcs.Sprint(b.Func), "toLower(`TheStreet`)") {
		t.Fatalf("trimming did not work")
	}
}

func TestComposeStringEq(t *testing.T) {
	expr := &expr.Expr{
		Function: &expr.Function{
			Name: "eq",
			Params: []*expr.Expr{
				{
					Function: &expr.Function{
						Name: "toLower",
						Params: []*expr.Expr{
							{
								Terminal: &expr.Terminal{
									Variable: &expr.Variable{
										Type: types.SINGLE_STRING,
									},
								},
							},
						},
					},
				},
				{
					Function: &expr.Function{
						Name: "toLower",
						Params: []*expr.Expr{
							{
								Terminal: &expr.Terminal{
									StringValue: proto.String("TheStreet"),
								},
							},
						},
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	r, err := b.Eval(serialize.NewStringValue("TheStreet"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestComposeListBool(t *testing.T) {
	expr := &expr.Expr{
		Function: &expr.Function{
			Name: "eq",
			Params: []*expr.Expr{
				{
					Function: &expr.Function{
						Name: "length",
						Params: []*expr.Expr{
							{
								List: &expr.List{
									Type: types.LIST_BOOL,
									Elems: []*expr.Expr{
										{
											Terminal: &expr.Terminal{
												BoolValue: proto.Bool(true),
											},
										},
										{
											Terminal: &expr.Terminal{
												BoolValue: proto.Bool(false),
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Terminal: &expr.Terminal{
						IntValue: proto.Int64(2),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	r, err := b.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	str := funcs.Sprint(b.Func)
	if str != "true" {
		t.Fatalf("trimming did not work: %s", str)
	}
}

func TestComposeListInt64(t *testing.T) {
	expr := &expr.Expr{
		Function: &expr.Function{
			Name: "eq",
			Params: []*expr.Expr{
				{
					Function: &expr.Function{
						Name: "elem",
						Params: []*expr.Expr{
							{
								Function: &expr.Function{
									Name: "print",
									Params: []*expr.Expr{
										{
											List: &expr.List{
												Type: types.LIST_INT,
												Elems: []*expr.Expr{
													{
														Terminal: &expr.Terminal{
															IntValue: proto.Int64(1),
														},
													},
													{
														Terminal: &expr.Terminal{
															IntValue: proto.Int64(2),
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Terminal: &expr.Terminal{
									IntValue: proto.Int64(1),
								},
							},
						},
					},
				},
				{
					Terminal: &expr.Terminal{
						IntValue: proto.Int64(2),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	r, err := b.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	t.Logf("%s", funcs.Sprint(b.Func))
}

func TestComposeRegex(t *testing.T) {
	expr := &expr.Expr{
		Function: &expr.Function{
			Name: "regex",
			Params: []*expr.Expr{
				{
					Terminal: &expr.Terminal{
						StringValue: proto.String("ab"),
					},
				},
				{
					Terminal: &expr.Terminal{
						Variable: &expr.Variable{
							Type: types.SINGLE_BYTES,
						},
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	r, err := b.Eval(serialize.NewBytesValue([]byte("a")))
	if err != nil {
		panic(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
}

func TestConst(t *testing.T) {
	expr := &expr.Expr{
		Function: &expr.Function{
			Name: "regex",
			Params: []*expr.Expr{
				{
					Terminal: &expr.Terminal{
						Variable: &expr.Variable{
							Type: types.SINGLE_STRING,
						},
					},
				},
				{
					Terminal: &expr.Terminal{
						BytesValue: []byte{1, 2},
					},
				},
			},
		},
	}
	_, err := NewBool(expr)
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "regex has constant") || !strings.Contains(err.Error(), "has a variable parameter") {
		t.Fatalf("expected more specific error %s", err.Error())
	}
}

func TestTrimInit(t *testing.T) {
	expr := &expr.Expr{
		Function: &expr.Function{
			Name: "regex",
			Params: []*expr.Expr{
				{
					Terminal: &expr.Terminal{
						StringValue: proto.String("ab"),
					},
				},
				{
					Terminal: &expr.Terminal{
						BytesValue: []byte{'a', 'b'},
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	r, err := b.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestNoTrim(t *testing.T) {
	expr := &expr.Expr{
		Function: &expr.Function{
			Name: "eq",
			Params: []*expr.Expr{
				{
					Function: &expr.Function{
						Name: "elem",
						Params: []*expr.Expr{
							{
								List: &expr.List{
									Type: types.LIST_STRING,
									Elems: []*expr.Expr{
										{
											Function: &expr.Function{
												Name: "print",
												Params: []*expr.Expr{
													{
														Terminal: &expr.Terminal{
															Variable: &expr.Variable{
																Type: types.SINGLE_STRING,
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Terminal: &expr.Terminal{
									IntValue: proto.Int64(0),
								},
							},
						},
					},
				},
				{
					Terminal: &expr.Terminal{
						StringValue: proto.String("abc"),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	str := funcs.Sprint(b.Func)
	if str == "false" {
		t.Fatalf("too much trimming")
	}
	t.Logf("trimmed = %s", str)
	r, err := b.Eval(serialize.NewStringValue("abc"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestList(t *testing.T) {
	expr := &expr.Expr{
		Function: &expr.Function{
			Name: "eq",
			Params: []*expr.Expr{
				{
					Function: &expr.Function{
						Name: "elem",
						Params: []*expr.Expr{
							{
								List: &expr.List{
									Type: types.LIST_STRING,
									Elems: []*expr.Expr{
										{
											Terminal: &expr.Terminal{
												StringValue: proto.String("abc"),
											},
										},
									},
								},
							},
							{
								Terminal: &expr.Terminal{
									IntValue: proto.Int64(0),
								},
							},
						},
					},
				},
				{
					Terminal: &expr.Terminal{
						StringValue: proto.String("abc"),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	str := funcs.Sprint(b.Func)
	if str != "true" {
		t.Fatalf("not enough trimming on %s", str)
	}
	r, err := b.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}
