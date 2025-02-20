// Code generated by funcs-gen. DO NOT EDIT.
package funcs

import (
	"github.com/katydid/validator-go/validator/ast"
)

type lenDoubles struct {
	E           Doubles
	hash        uint64
	hasVariable bool
}

func (this *lenDoubles) Eval() (int64, error) {
	e, err := this.E.Eval()
	if err != nil {
		return 0, err
	}
	return int64(len(e)), nil
}

func (this *lenDoubles) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*lenDoubles); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return this.ToExpr().Compare(that.ToExpr())
}

func (this *lenDoubles) ToExpr() *ast.Expr {
	return ast.NewFunction("length", this.E.ToExpr())
}

func (this *lenDoubles) HasVariable() bool {
	return this.hasVariable
}

func (this *lenDoubles) Hash() uint64 {
	return this.hash
}

func init() {
	Register("length", LenDoubles)
}

// LenDoubles returns a function that returns the length of a list of type Doubles
func LenDoubles(e Doubles) Int {
	return TrimInt(&lenDoubles{
		E:           e,
		hash:        Hash("length", e),
		hasVariable: e.HasVariable(),
	})
}

type lenInts struct {
	E           Ints
	hash        uint64
	hasVariable bool
}

func (this *lenInts) Eval() (int64, error) {
	e, err := this.E.Eval()
	if err != nil {
		return 0, err
	}
	return int64(len(e)), nil
}

func (this *lenInts) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*lenInts); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return this.ToExpr().Compare(that.ToExpr())
}

func (this *lenInts) ToExpr() *ast.Expr {
	return ast.NewFunction("length", this.E.ToExpr())
}

func (this *lenInts) HasVariable() bool {
	return this.hasVariable
}

func (this *lenInts) Hash() uint64 {
	return this.hash
}

func init() {
	Register("length", LenInts)
}

// LenInts returns a function that returns the length of a list of type Ints
func LenInts(e Ints) Int {
	return TrimInt(&lenInts{
		E:           e,
		hash:        Hash("length", e),
		hasVariable: e.HasVariable(),
	})
}

type lenUints struct {
	E           Uints
	hash        uint64
	hasVariable bool
}

func (this *lenUints) Eval() (int64, error) {
	e, err := this.E.Eval()
	if err != nil {
		return 0, err
	}
	return int64(len(e)), nil
}

func (this *lenUints) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*lenUints); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return this.ToExpr().Compare(that.ToExpr())
}

func (this *lenUints) ToExpr() *ast.Expr {
	return ast.NewFunction("length", this.E.ToExpr())
}

func (this *lenUints) HasVariable() bool {
	return this.hasVariable
}

func (this *lenUints) Hash() uint64 {
	return this.hash
}

func init() {
	Register("length", LenUints)
}

// LenUints returns a function that returns the length of a list of type Uints
func LenUints(e Uints) Int {
	return TrimInt(&lenUints{
		E:           e,
		hash:        Hash("length", e),
		hasVariable: e.HasVariable(),
	})
}

type lenBools struct {
	E           Bools
	hash        uint64
	hasVariable bool
}

func (this *lenBools) Eval() (int64, error) {
	e, err := this.E.Eval()
	if err != nil {
		return 0, err
	}
	return int64(len(e)), nil
}

func (this *lenBools) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*lenBools); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return this.ToExpr().Compare(that.ToExpr())
}

func (this *lenBools) ToExpr() *ast.Expr {
	return ast.NewFunction("length", this.E.ToExpr())
}

func (this *lenBools) HasVariable() bool {
	return this.hasVariable
}

func (this *lenBools) Hash() uint64 {
	return this.hash
}

func init() {
	Register("length", LenBools)
}

// LenBools returns a function that returns the length of a list of type Bools
func LenBools(e Bools) Int {
	return TrimInt(&lenBools{
		E:           e,
		hash:        Hash("length", e),
		hasVariable: e.HasVariable(),
	})
}

type lenStrings struct {
	E           Strings
	hash        uint64
	hasVariable bool
}

func (this *lenStrings) Eval() (int64, error) {
	e, err := this.E.Eval()
	if err != nil {
		return 0, err
	}
	return int64(len(e)), nil
}

func (this *lenStrings) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*lenStrings); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return this.ToExpr().Compare(that.ToExpr())
}

func (this *lenStrings) ToExpr() *ast.Expr {
	return ast.NewFunction("length", this.E.ToExpr())
}

func (this *lenStrings) HasVariable() bool {
	return this.hasVariable
}

func (this *lenStrings) Hash() uint64 {
	return this.hash
}

func init() {
	Register("length", LenStrings)
}

// LenStrings returns a function that returns the length of a list of type Strings
func LenStrings(e Strings) Int {
	return TrimInt(&lenStrings{
		E:           e,
		hash:        Hash("length", e),
		hasVariable: e.HasVariable(),
	})
}

type lenListOfBytes struct {
	E           ListOfBytes
	hash        uint64
	hasVariable bool
}

func (this *lenListOfBytes) Eval() (int64, error) {
	e, err := this.E.Eval()
	if err != nil {
		return 0, err
	}
	return int64(len(e)), nil
}

func (this *lenListOfBytes) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*lenListOfBytes); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return this.ToExpr().Compare(that.ToExpr())
}

func (this *lenListOfBytes) ToExpr() *ast.Expr {
	return ast.NewFunction("length", this.E.ToExpr())
}

func (this *lenListOfBytes) HasVariable() bool {
	return this.hasVariable
}

func (this *lenListOfBytes) Hash() uint64 {
	return this.hash
}

func init() {
	Register("length", LenListOfBytes)
}

// LenListOfBytes returns a function that returns the length of a list of type ListOfBytes
func LenListOfBytes(e ListOfBytes) Int {
	return TrimInt(&lenListOfBytes{
		E:           e,
		hash:        Hash("length", e),
		hasVariable: e.HasVariable(),
	})
}

type lenString struct {
	E           String
	hash        uint64
	hasVariable bool
}

func (this *lenString) Eval() (int64, error) {
	e, err := this.E.Eval()
	if err != nil {
		return 0, err
	}
	return int64(len(e)), nil
}

func (this *lenString) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*lenString); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return this.ToExpr().Compare(that.ToExpr())
}

func (this *lenString) ToExpr() *ast.Expr {
	return ast.NewFunction("length", this.E.ToExpr())
}

func (this *lenString) HasVariable() bool {
	return this.hasVariable
}

func (this *lenString) Hash() uint64 {
	return this.hash
}

func init() {
	Register("length", LenString)
}

// LenString returns a function that returns the length of a list of type String
func LenString(e String) Int {
	return TrimInt(&lenString{
		E:           e,
		hash:        Hash("length", e),
		hasVariable: e.HasVariable(),
	})
}

type lenBytes struct {
	E           Bytes
	hash        uint64
	hasVariable bool
}

func (this *lenBytes) Eval() (int64, error) {
	e, err := this.E.Eval()
	if err != nil {
		return 0, err
	}
	return int64(len(e)), nil
}

func (this *lenBytes) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*lenBytes); ok {
		if c := this.E.Compare(other.E); c != 0 {
			return c
		}
		return 0
	}
	return this.ToExpr().Compare(that.ToExpr())
}

func (this *lenBytes) ToExpr() *ast.Expr {
	return ast.NewFunction("length", this.E.ToExpr())
}

func (this *lenBytes) HasVariable() bool {
	return this.hasVariable
}

func (this *lenBytes) Hash() uint64 {
	return this.hash
}

func init() {
	Register("length", LenBytes)
}

// LenBytes returns a function that returns the length of a list of type Bytes
func LenBytes(e Bytes) Int {
	return TrimInt(&lenBytes{
		E:           e,
		hash:        Hash("length", e),
		hasVariable: e.HasVariable(),
	})
}
