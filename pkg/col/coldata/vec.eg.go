// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package coldata

import (
	"fmt"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	// HACK: crlfmt removes the "*/}}" comment if it's the last line in the import
	// block. This was picked because it sorts after "pkg/sql/exec/execgen" and
	// has no deps.
	_ "github.com/cockroachdb/cockroach/pkg/util/bufalloc"
)

func (m *memColumn) Append(args AppendArgs) {
	switch args.ColType {
	case coltypes.Bool:
		fromCol := args.Src.Bool()
		toCol := m.Bool()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			// TODO(asubiotto): We could be more efficient for fixed width types by
			// preallocating a destination slice (not so for variable length types).
			// Improve this.
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case coltypes.Bytes:
		fromCol := args.Src.Bytes()
		toCol := m.Bytes()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol.AppendSlice(fromCol, int(args.DestIdx), int(args.SrcStartIdx), int(args.SrcEndIdx))
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			// TODO(asubiotto): We could be more efficient for fixed width types by
			// preallocating a destination slice (not so for variable length types).
			// Improve this.
			toCol = toCol.Slice(0, int(args.DestIdx))
			for _, selIdx := range sel {
				val := fromCol.Get(int(selIdx))
				toCol.AppendVal(val)
			}
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case coltypes.Decimal:
		fromCol := args.Src.Decimal()
		toCol := m.Decimal()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			// TODO(asubiotto): We could be more efficient for fixed width types by
			// preallocating a destination slice (not so for variable length types).
			// Improve this.
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case coltypes.Int8:
		fromCol := args.Src.Int8()
		toCol := m.Int8()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			// TODO(asubiotto): We could be more efficient for fixed width types by
			// preallocating a destination slice (not so for variable length types).
			// Improve this.
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case coltypes.Int16:
		fromCol := args.Src.Int16()
		toCol := m.Int16()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			// TODO(asubiotto): We could be more efficient for fixed width types by
			// preallocating a destination slice (not so for variable length types).
			// Improve this.
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case coltypes.Int32:
		fromCol := args.Src.Int32()
		toCol := m.Int32()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			// TODO(asubiotto): We could be more efficient for fixed width types by
			// preallocating a destination slice (not so for variable length types).
			// Improve this.
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case coltypes.Int64:
		fromCol := args.Src.Int64()
		toCol := m.Int64()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			// TODO(asubiotto): We could be more efficient for fixed width types by
			// preallocating a destination slice (not so for variable length types).
			// Improve this.
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case coltypes.Float32:
		fromCol := args.Src.Float32()
		toCol := m.Float32()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			// TODO(asubiotto): We could be more efficient for fixed width types by
			// preallocating a destination slice (not so for variable length types).
			// Improve this.
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case coltypes.Float64:
		fromCol := args.Src.Float64()
		toCol := m.Float64()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:int(args.DestIdx)], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			// TODO(asubiotto): We could be more efficient for fixed width types by
			// preallocating a destination slice (not so for variable length types).
			// Improve this.
			toCol = toCol[0:int(args.DestIdx)]
			for _, selIdx := range sel {
				val := fromCol[int(selIdx)]
				toCol = append(toCol, val)
			}
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	default:
		panic(fmt.Sprintf("unhandled type %s", args.ColType))
	}
}

func (m *memColumn) Copy(args CopyArgs) {
	if args.Nils != nil && args.Sel64 == nil {
		panic("Nils set without Sel64")
	}

	m.Nulls().UnsetNullRange(args.DestIdx, args.DestIdx+(args.SrcEndIdx-args.SrcStartIdx))

	switch args.ColType {
	case coltypes.Bool:
		fromCol := args.Src.Bool()
		toCol := m.Bool()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					n := len(toCol)
					toColSliced := toCol[int(args.DestIdx):n]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toColSliced[i] = v
						}
					}
					return
				}
				// Nils but no Nulls.
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils or Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case coltypes.Bytes:
		fromCol := args.Src.Bytes()
		toCol := m.Bytes()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					n := toCol.Len()
					toColSliced := toCol.Slice(int(args.DestIdx), n)
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol.Get(int(selIdx))
							toColSliced.Set(i, v)
						}
					}
					return
				}
				// Nils but no Nulls.
				n := toCol.Len()
				toColSliced := toCol.Slice(int(args.DestIdx), n)
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol.Get(int(selIdx))
						toColSliced.Set(i, v)
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := toCol.Len()
				toColSliced := toCol.Slice(int(args.DestIdx), n)
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol.Get(int(selIdx))
						toColSliced.Set(i, v)
					}
				}
				return
			}
			// No Nils or Nulls.
			n := toCol.Len()
			toColSliced := toCol.Slice(int(args.DestIdx), n)
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol.Get(int(selIdx))
				toColSliced.Set(i, v)
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := toCol.Len()
				toColSliced := toCol.Slice(int(args.DestIdx), n)
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol.Get(int(selIdx))
						toColSliced.Set(i, v)
					}
				}
				return
			}
			// No Nulls.
			n := toCol.Len()
			toColSliced := toCol.Slice(int(args.DestIdx), n)
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol.Get(int(selIdx))
				toColSliced.Set(i, v)
			}
			return
		}
		// No Sel or Sel64.
		toCol.CopySlice(fromCol, int(args.DestIdx), int(args.SrcStartIdx), int(args.SrcEndIdx))
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case coltypes.Decimal:
		fromCol := args.Src.Decimal()
		toCol := m.Decimal()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					n := len(toCol)
					toColSliced := toCol[int(args.DestIdx):n]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toColSliced[i] = v
						}
					}
					return
				}
				// Nils but no Nulls.
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils or Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case coltypes.Int8:
		fromCol := args.Src.Int8()
		toCol := m.Int8()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					n := len(toCol)
					toColSliced := toCol[int(args.DestIdx):n]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toColSliced[i] = v
						}
					}
					return
				}
				// Nils but no Nulls.
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils or Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case coltypes.Int16:
		fromCol := args.Src.Int16()
		toCol := m.Int16()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					n := len(toCol)
					toColSliced := toCol[int(args.DestIdx):n]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toColSliced[i] = v
						}
					}
					return
				}
				// Nils but no Nulls.
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils or Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case coltypes.Int32:
		fromCol := args.Src.Int32()
		toCol := m.Int32()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					n := len(toCol)
					toColSliced := toCol[int(args.DestIdx):n]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toColSliced[i] = v
						}
					}
					return
				}
				// Nils but no Nulls.
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils or Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case coltypes.Int64:
		fromCol := args.Src.Int64()
		toCol := m.Int64()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					n := len(toCol)
					toColSliced := toCol[int(args.DestIdx):n]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toColSliced[i] = v
						}
					}
					return
				}
				// Nils but no Nulls.
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils or Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case coltypes.Float32:
		fromCol := args.Src.Float32()
		toCol := m.Float32()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					n := len(toCol)
					toColSliced := toCol[int(args.DestIdx):n]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toColSliced[i] = v
						}
					}
					return
				}
				// Nils but no Nulls.
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils or Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case coltypes.Float64:
		fromCol := args.Src.Float64()
		toCol := m.Float64()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					n := len(toCol)
					toColSliced := toCol[int(args.DestIdx):n]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							v := fromCol[int(selIdx)]
							toColSliced[i] = v
						}
					}
					return
				}
				// Nils but no Nulls.
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nils or Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				n := len(toCol)
				toColSliced := toCol[int(args.DestIdx):n]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						v := fromCol[int(selIdx)]
						toColSliced[i] = v
					}
				}
				return
			}
			// No Nulls.
			n := len(toCol)
			toColSliced := toCol[int(args.DestIdx):n]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				v := fromCol[int(selIdx)]
				toColSliced[i] = v
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[int(args.DestIdx):], fromCol[int(args.SrcStartIdx):int(args.SrcEndIdx)])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	default:
		panic(fmt.Sprintf("unhandled type %s", args.ColType))
	}
}

func (m *memColumn) Slice(colType coltypes.T, start uint64, end uint64) Vec {
	switch colType {
	case coltypes.Bool:
		col := m.Bool()
		return &memColumn{
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Bytes:
		col := m.Bytes()
		return &memColumn{
			col:   col.Slice(int(start), int(end)),
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Decimal:
		col := m.Decimal()
		return &memColumn{
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Int8:
		col := m.Int8()
		return &memColumn{
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Int16:
		col := m.Int16()
		return &memColumn{
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Int32:
		col := m.Int32()
		return &memColumn{
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Int64:
		col := m.Int64()
		return &memColumn{
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Float32:
		col := m.Float32()
		return &memColumn{
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	case coltypes.Float64:
		col := m.Float64()
		return &memColumn{
			col:   col[int(start):int(end)],
			nulls: m.nulls.Slice(start, end),
		}
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

func (m *memColumn) PrettyValueAt(colIdx uint16, colType coltypes.T) string {
	if m.nulls.NullAt(colIdx) {
		return "NULL"
	}
	switch colType {
	case coltypes.Bool:
		col := m.Bool()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Bytes:
		col := m.Bytes()
		v := col.Get(int(colIdx))
		return fmt.Sprintf("%v", v)
	case coltypes.Decimal:
		col := m.Decimal()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Int8:
		col := m.Int8()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Int16:
		col := m.Int16()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Int32:
		col := m.Int32()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Int64:
		col := m.Int64()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Float32:
		col := m.Float32()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	case coltypes.Float64:
		col := m.Float64()
		v := col[int(colIdx)]
		return fmt.Sprintf("%v", v)
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

// Helper to set the value in a Vec when the type is unknown.
func SetValueAt(v Vec, elem interface{}, rowIdx uint16, colType coltypes.T) {
	switch colType {
	case coltypes.Bool:
		target := v.Bool()
		newVal := elem.(bool)
		target[int(rowIdx)] = newVal
	case coltypes.Bytes:
		target := v.Bytes()
		newVal := elem.([]byte)
		target.Set(int(rowIdx), newVal)
	case coltypes.Decimal:
		target := v.Decimal()
		newVal := elem.(apd.Decimal)
		target[int(rowIdx)] = newVal
	case coltypes.Int8:
		target := v.Int8()
		newVal := elem.(int8)
		target[int(rowIdx)] = newVal
	case coltypes.Int16:
		target := v.Int16()
		newVal := elem.(int16)
		target[int(rowIdx)] = newVal
	case coltypes.Int32:
		target := v.Int32()
		newVal := elem.(int32)
		target[int(rowIdx)] = newVal
	case coltypes.Int64:
		target := v.Int64()
		newVal := elem.(int64)
		target[int(rowIdx)] = newVal
	case coltypes.Float32:
		target := v.Float32()
		newVal := elem.(float32)
		target[int(rowIdx)] = newVal
	case coltypes.Float64:
		target := v.Float64()
		newVal := elem.(float64)
		target[int(rowIdx)] = newVal
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}