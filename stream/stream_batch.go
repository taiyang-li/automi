package stream

import (
	"github.com/vladimirvivien/automi/api"
	"github.com/vladimirvivien/automi/operators/batch"
	"github.com/vladimirvivien/automi/operators/unary"
)

func (s *Stream) Batch() *Stream {
	operator := batch.New(s.ctx)
	return s.appendOp(operator)
}

// GroupByKey groups incoming items that are batched as
// type []map[K]V where parameter key is used to group
// the items when K=key.  Items with same key values are
// grouped in a new map and returned as []map[G]V.
//
// See Also
//
// See batch operator function GroupByKey in
//   "github.com/vladimirvivien/automi/operators/batch/"#GroupByKeyFunc
func (s *Stream) GroupByKey(key interface{}) *Stream {
	operator := unary.New(s.ctx)
	operator.SetOperation(batch.GroupByKeyFunc(key))
	return s.appendOp(operator)
}

// GroupByName groups incoming items that are batched as
// type []T where T is a struct. Parameter name is used to select
// T.name as key to group items with the same value into a map map[key][]T
// that is sent downstream.
//
// See Also
//
// See batch operator function GroupByName in
//    "github.com/vladimirvivien/automi/operators/batch/"
func (s *Stream) GroupByName(name string) *Stream {
	operator := unary.New(s.ctx)
	operator.SetOperation(batch.GroupByNameFunc(name))
	return s.appendOp(operator)
}

// GroupByPos groups incoming items that are batched as
// [][]T. For each i in dimension 1, [i][pos] is selected as key
// and grouped in a map, map[key][]T, that is returned downstream.
//
// See Also
//
// See the batch operator function GroupByPosFunc in
//   "github.com/vladimirvivien/automi/operators/batch/"
func (s *Stream) GroupByPos(pos int) *Stream {
	operator := unary.New(s.ctx)
	operator.SetOperation(batch.GroupByPosFunc(pos))
	return s.appendOp(operator)
}

// SortByKey sorts incoming items that are batched as type []map[K]V
// where K is a comparable type specified by param key and used to
// sort the slice.  The opertor returns a sorted []map[K]V.
//
// See Also
//
// See also the operator function SortByKeyFunc in
//   "github.com/vladimirvivien/operators/batch"
func (s *Stream) SortByKey(key interface{}) *Stream {
	operator := unary.New(s.ctx)
	operator.SetOperation(batch.SortByKeyFunc(key))
	return s.appendOp(operator)
}

// GroupByKey
func (s *Stream) appendOp(operator api.Operator) *Stream {
	s.ops = append(s.ops, operator)
	return s
}