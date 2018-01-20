// Code generated by funcs-gen. DO NOT EDIT.
package funcs

import (
	"strings"
)

type rangeDoubles struct {
	List  Doubles
	First Int
	Last  Int
	hash uint64
}

func (this *rangeDoubles) Eval() ([]float64, error) {
	list, err := this.List.Eval()
	if err != nil {
		return nil, err
	}
	first64, err := this.First.Eval()
	if err != nil {
		return nil, err
	}
	first := int(first64)
	if len(list) == 0 {
		return nil, NewRangeCheckErr(first, len(list))
	}
	if first < 0 {
		first = first % len(list)
	}
	if first > len(list) {
		return nil, NewRangeCheckErr(first, len(list))
	}
	last64, err := this.Last.Eval()
	if err != nil {
		return nil, err
	}
	last := int(last64)
	if last < 0 {
		last = last % len(list)
	}
	if last > len(list) {
		return nil, NewRangeCheckErr(last, len(list))
	}
	if first > last {
		return nil, NewRangeErr(first, last)
	}
	return list[first:last], nil
}

func (this *rangeDoubles) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*rangeDoubles); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.First.Compare(other.First); c != 0 {
			return c
		}
		if c := this.Last.Compare(other.Last); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *rangeDoubles) String() string {
	return "range(" + sjoin(this.List, this.First, this.Last) +")"
}

func (this *rangeDoubles) Hash() uint64 {
	return this.hash
}

func init() {
	Register("range", RangeDoubles)
}

//RangeDoubles returns a function that returns a range of elements from a list.
func RangeDoubles(list Doubles, from, to Int) Doubles {
	h := uint64(17)
	h = 31*h + 63639164578
	h = 31*h + from.Hash()
	h = 31*h + to.Hash()
	h = 31*h + list.Hash()
	return &rangeDoubles{
		List:  list,
		First: from,
		Last:  to,
		hash: h,
	}
}

type rangeInts struct {
	List  Ints
	First Int
	Last  Int
	hash uint64
}

func (this *rangeInts) Eval() ([]int64, error) {
	list, err := this.List.Eval()
	if err != nil {
		return nil, err
	}
	first64, err := this.First.Eval()
	if err != nil {
		return nil, err
	}
	first := int(first64)
	if len(list) == 0 {
		return nil, NewRangeCheckErr(first, len(list))
	}
	if first < 0 {
		first = first % len(list)
	}
	if first > len(list) {
		return nil, NewRangeCheckErr(first, len(list))
	}
	last64, err := this.Last.Eval()
	if err != nil {
		return nil, err
	}
	last := int(last64)
	if last < 0 {
		last = last % len(list)
	}
	if last > len(list) {
		return nil, NewRangeCheckErr(last, len(list))
	}
	if first > last {
		return nil, NewRangeErr(first, last)
	}
	return list[first:last], nil
}

func (this *rangeInts) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*rangeInts); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.First.Compare(other.First); c != 0 {
			return c
		}
		if c := this.Last.Compare(other.Last); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *rangeInts) String() string {
	return "range(" + sjoin(this.List, this.First, this.Last) +")"
}

func (this *rangeInts) Hash() uint64 {
	return this.hash
}

func init() {
	Register("range", RangeInts)
}

//RangeInts returns a function that returns a range of elements from a list.
func RangeInts(list Ints, from, to Int) Ints {
	h := uint64(17)
	h = 31*h + 2284164
	h = 31*h + from.Hash()
	h = 31*h + to.Hash()
	h = 31*h + list.Hash()
	return &rangeInts{
		List:  list,
		First: from,
		Last:  to,
		hash: h,
	}
}

type rangeUints struct {
	List  Uints
	First Int
	Last  Int
	hash uint64
}

func (this *rangeUints) Eval() ([]uint64, error) {
	list, err := this.List.Eval()
	if err != nil {
		return nil, err
	}
	first64, err := this.First.Eval()
	if err != nil {
		return nil, err
	}
	first := int(first64)
	if len(list) == 0 {
		return nil, NewRangeCheckErr(first, len(list))
	}
	if first < 0 {
		first = first % len(list)
	}
	if first > len(list) {
		return nil, NewRangeCheckErr(first, len(list))
	}
	last64, err := this.Last.Eval()
	if err != nil {
		return nil, err
	}
	last := int(last64)
	if last < 0 {
		last = last % len(list)
	}
	if last > len(list) {
		return nil, NewRangeCheckErr(last, len(list))
	}
	if first > last {
		return nil, NewRangeErr(first, last)
	}
	return list[first:last], nil
}

func (this *rangeUints) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*rangeUints); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.First.Compare(other.First); c != 0 {
			return c
		}
		if c := this.Last.Compare(other.Last); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *rangeUints) String() string {
	return "range(" + sjoin(this.List, this.First, this.Last) +")"
}

func (this *rangeUints) Hash() uint64 {
	return this.hash
}

func init() {
	Register("range", RangeUints)
}

//RangeUints returns a function that returns a range of elements from a list.
func RangeUints(list Uints, from, to Int) Uints {
	h := uint64(17)
	h = 31*h + 81736761
	h = 31*h + from.Hash()
	h = 31*h + to.Hash()
	h = 31*h + list.Hash()
	return &rangeUints{
		List:  list,
		First: from,
		Last:  to,
		hash: h,
	}
}

type rangeBools struct {
	List  Bools
	First Int
	Last  Int
	hash uint64
}

func (this *rangeBools) Eval() ([]bool, error) {
	list, err := this.List.Eval()
	if err != nil {
		return nil, err
	}
	first64, err := this.First.Eval()
	if err != nil {
		return nil, err
	}
	first := int(first64)
	if len(list) == 0 {
		return nil, NewRangeCheckErr(first, len(list))
	}
	if first < 0 {
		first = first % len(list)
	}
	if first > len(list) {
		return nil, NewRangeCheckErr(first, len(list))
	}
	last64, err := this.Last.Eval()
	if err != nil {
		return nil, err
	}
	last := int(last64)
	if last < 0 {
		last = last % len(list)
	}
	if last > len(list) {
		return nil, NewRangeCheckErr(last, len(list))
	}
	if first > last {
		return nil, NewRangeErr(first, last)
	}
	return list[first:last], nil
}

func (this *rangeBools) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*rangeBools); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.First.Compare(other.First); c != 0 {
			return c
		}
		if c := this.Last.Compare(other.Last); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *rangeBools) String() string {
	return "range(" + sjoin(this.List, this.First, this.Last) +")"
}

func (this *rangeBools) Hash() uint64 {
	return this.hash
}

func init() {
	Register("range", RangeBools)
}

//RangeBools returns a function that returns a range of elements from a list.
func RangeBools(list Bools, from, to Int) Bools {
	h := uint64(17)
	h = 31*h + 64369321
	h = 31*h + from.Hash()
	h = 31*h + to.Hash()
	h = 31*h + list.Hash()
	return &rangeBools{
		List:  list,
		First: from,
		Last:  to,
		hash: h,
	}
}

type rangeStrings struct {
	List  Strings
	First Int
	Last  Int
	hash uint64
}

func (this *rangeStrings) Eval() ([]string, error) {
	list, err := this.List.Eval()
	if err != nil {
		return nil, err
	}
	first64, err := this.First.Eval()
	if err != nil {
		return nil, err
	}
	first := int(first64)
	if len(list) == 0 {
		return nil, NewRangeCheckErr(first, len(list))
	}
	if first < 0 {
		first = first % len(list)
	}
	if first > len(list) {
		return nil, NewRangeCheckErr(first, len(list))
	}
	last64, err := this.Last.Eval()
	if err != nil {
		return nil, err
	}
	last := int(last64)
	if last < 0 {
		last = last % len(list)
	}
	if last > len(list) {
		return nil, NewRangeCheckErr(last, len(list))
	}
	if first > last {
		return nil, NewRangeErr(first, last)
	}
	return list[first:last], nil
}

func (this *rangeStrings) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*rangeStrings); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.First.Compare(other.First); c != 0 {
			return c
		}
		if c := this.Last.Compare(other.Last); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *rangeStrings) String() string {
	return "range(" + sjoin(this.List, this.First, this.Last) +")"
}

func (this *rangeStrings) Hash() uint64 {
	return this.hash
}

func init() {
	Register("range", RangeStrings)
}

//RangeStrings returns a function that returns a range of elements from a list.
func RangeStrings(list Strings, from, to Int) Strings {
	h := uint64(17)
	h = 31*h + 77092305506
	h = 31*h + from.Hash()
	h = 31*h + to.Hash()
	h = 31*h + list.Hash()
	return &rangeStrings{
		List:  list,
		First: from,
		Last:  to,
		hash: h,
	}
}

type rangeListOfBytes struct {
	List  ListOfBytes
	First Int
	Last  Int
	hash uint64
}

func (this *rangeListOfBytes) Eval() ([][]byte, error) {
	list, err := this.List.Eval()
	if err != nil {
		return nil, err
	}
	first64, err := this.First.Eval()
	if err != nil {
		return nil, err
	}
	first := int(first64)
	if len(list) == 0 {
		return nil, NewRangeCheckErr(first, len(list))
	}
	if first < 0 {
		first = first % len(list)
	}
	if first > len(list) {
		return nil, NewRangeCheckErr(first, len(list))
	}
	last64, err := this.Last.Eval()
	if err != nil {
		return nil, err
	}
	last := int(last64)
	if last < 0 {
		last = last % len(list)
	}
	if last > len(list) {
		return nil, NewRangeCheckErr(last, len(list))
	}
	if first > last {
		return nil, NewRangeErr(first, last)
	}
	return list[first:last], nil
}

func (this *rangeListOfBytes) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*rangeListOfBytes); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.First.Compare(other.First); c != 0 {
			return c
		}
		if c := this.Last.Compare(other.Last); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *rangeListOfBytes) String() string {
	return "range(" + sjoin(this.List, this.First, this.Last) +")"
}

func (this *rangeListOfBytes) Hash() uint64 {
	return this.hash
}

func init() {
	Register("range", RangeListOfBytes)
}

//RangeListOfBytes returns a function that returns a range of elements from a list.
func RangeListOfBytes(list ListOfBytes, from, to Int) ListOfBytes {
	h := uint64(17)
	h = 31*h + 65169257167589942
	h = 31*h + from.Hash()
	h = 31*h + to.Hash()
	h = 31*h + list.Hash()
	return &rangeListOfBytes{
		List:  list,
		First: from,
		Last:  to,
		hash: h,
	}
}
