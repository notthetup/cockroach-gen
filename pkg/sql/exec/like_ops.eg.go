// Code generated by execgen; DO NOT EDIT.

package exec

import (
	"bytes"
	"context"
	"regexp"

	"github.com/cockroachdb/cockroach/pkg/sql/exec/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/types"
)

type selPrefixBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg []byte
}

func (p *selPrefixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx uint16
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = bytes.HasPrefix(arg, p.constArg)
					if cmp && !nulls.NullAt(i) {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasPrefix(arg, p.constArg)
					if cmp && !nulls.NullAt(uint16(i)) {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		} else {

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p selPrefixBytesBytesConstOp) Init() {
	p.input.Init()
}

type projPrefixBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg []byte

	outputIdx int
}

func (p projPrefixBytesBytesConstOp) EstimateStaticMemoryUsage() int {
	return EstimateBatchSizeBytes([]types.T{types.Bool}, coldata.BatchSize)
}

func (p projPrefixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return batch
	}
	if p.outputIdx == batch.Width() {
		batch.AppendCol(types.Bool)
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	projCol := projVec.Bool()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel {
			arg := col.Get(int(i))
			projCol[i] = bytes.HasPrefix(arg, p.constArg)
		}
	} else {
		col = col.Slice(0, int(n))
		colLen := col.Len()
		_ = projCol[colLen-1]
		for i := 0; i < col.Len(); i++ {
			arg := col.Get(i)
			projCol[i] = bytes.HasPrefix(arg, p.constArg)
		}
	}
	if vec.Nulls().MaybeHasNulls() {
		nulls := vec.Nulls().Copy()
		projVec.SetNulls(&nulls)
	}
	return batch
}

func (p projPrefixBytesBytesConstOp) Init() {
	p.input.Init()
}

type selSuffixBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg []byte
}

func (p *selSuffixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx uint16
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = bytes.HasSuffix(arg, p.constArg)
					if cmp && !nulls.NullAt(i) {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasSuffix(arg, p.constArg)
					if cmp && !nulls.NullAt(uint16(i)) {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		} else {

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p selSuffixBytesBytesConstOp) Init() {
	p.input.Init()
}

type projSuffixBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg []byte

	outputIdx int
}

func (p projSuffixBytesBytesConstOp) EstimateStaticMemoryUsage() int {
	return EstimateBatchSizeBytes([]types.T{types.Bool}, coldata.BatchSize)
}

func (p projSuffixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return batch
	}
	if p.outputIdx == batch.Width() {
		batch.AppendCol(types.Bool)
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	projCol := projVec.Bool()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel {
			arg := col.Get(int(i))
			projCol[i] = bytes.HasSuffix(arg, p.constArg)
		}
	} else {
		col = col.Slice(0, int(n))
		colLen := col.Len()
		_ = projCol[colLen-1]
		for i := 0; i < col.Len(); i++ {
			arg := col.Get(i)
			projCol[i] = bytes.HasSuffix(arg, p.constArg)
		}
	}
	if vec.Nulls().MaybeHasNulls() {
		nulls := vec.Nulls().Copy()
		projVec.SetNulls(&nulls)
	}
	return batch
}

func (p projSuffixBytesBytesConstOp) Init() {
	p.input.Init()
}

type selRegexpBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg *regexp.Regexp
}

func (p *selRegexpBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx uint16
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = p.constArg.Match(arg)
					if cmp && !nulls.NullAt(i) {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = p.constArg.Match(arg)
					if cmp && !nulls.NullAt(uint16(i)) {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		} else {

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = p.constArg.Match(arg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = p.constArg.Match(arg)
					if cmp {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p selRegexpBytesBytesConstOp) Init() {
	p.input.Init()
}

type projRegexpBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg *regexp.Regexp

	outputIdx int
}

func (p projRegexpBytesBytesConstOp) EstimateStaticMemoryUsage() int {
	return EstimateBatchSizeBytes([]types.T{types.Bool}, coldata.BatchSize)
}

func (p projRegexpBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return batch
	}
	if p.outputIdx == batch.Width() {
		batch.AppendCol(types.Bool)
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	projCol := projVec.Bool()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel {
			arg := col.Get(int(i))
			projCol[i] = p.constArg.Match(arg)
		}
	} else {
		col = col.Slice(0, int(n))
		colLen := col.Len()
		_ = projCol[colLen-1]
		for i := 0; i < col.Len(); i++ {
			arg := col.Get(i)
			projCol[i] = p.constArg.Match(arg)
		}
	}
	if vec.Nulls().MaybeHasNulls() {
		nulls := vec.Nulls().Copy()
		projVec.SetNulls(&nulls)
	}
	return batch
}

func (p projRegexpBytesBytesConstOp) Init() {
	p.input.Init()
}

type selNotPrefixBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg []byte
}

func (p *selNotPrefixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx uint16
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = !bytes.HasPrefix(arg, p.constArg)
					if cmp && !nulls.NullAt(i) {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasPrefix(arg, p.constArg)
					if cmp && !nulls.NullAt(uint16(i)) {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		} else {

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = !bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p selNotPrefixBytesBytesConstOp) Init() {
	p.input.Init()
}

type projNotPrefixBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg []byte

	outputIdx int
}

func (p projNotPrefixBytesBytesConstOp) EstimateStaticMemoryUsage() int {
	return EstimateBatchSizeBytes([]types.T{types.Bool}, coldata.BatchSize)
}

func (p projNotPrefixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return batch
	}
	if p.outputIdx == batch.Width() {
		batch.AppendCol(types.Bool)
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	projCol := projVec.Bool()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel {
			arg := col.Get(int(i))
			projCol[i] = !bytes.HasPrefix(arg, p.constArg)
		}
	} else {
		col = col.Slice(0, int(n))
		colLen := col.Len()
		_ = projCol[colLen-1]
		for i := 0; i < col.Len(); i++ {
			arg := col.Get(i)
			projCol[i] = !bytes.HasPrefix(arg, p.constArg)
		}
	}
	if vec.Nulls().MaybeHasNulls() {
		nulls := vec.Nulls().Copy()
		projVec.SetNulls(&nulls)
	}
	return batch
}

func (p projNotPrefixBytesBytesConstOp) Init() {
	p.input.Init()
}

type selNotSuffixBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg []byte
}

func (p *selNotSuffixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx uint16
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = !bytes.HasSuffix(arg, p.constArg)
					if cmp && !nulls.NullAt(i) {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasSuffix(arg, p.constArg)
					if cmp && !nulls.NullAt(uint16(i)) {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		} else {

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = !bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p selNotSuffixBytesBytesConstOp) Init() {
	p.input.Init()
}

type projNotSuffixBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg []byte

	outputIdx int
}

func (p projNotSuffixBytesBytesConstOp) EstimateStaticMemoryUsage() int {
	return EstimateBatchSizeBytes([]types.T{types.Bool}, coldata.BatchSize)
}

func (p projNotSuffixBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return batch
	}
	if p.outputIdx == batch.Width() {
		batch.AppendCol(types.Bool)
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	projCol := projVec.Bool()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel {
			arg := col.Get(int(i))
			projCol[i] = !bytes.HasSuffix(arg, p.constArg)
		}
	} else {
		col = col.Slice(0, int(n))
		colLen := col.Len()
		_ = projCol[colLen-1]
		for i := 0; i < col.Len(); i++ {
			arg := col.Get(i)
			projCol[i] = !bytes.HasSuffix(arg, p.constArg)
		}
	}
	if vec.Nulls().MaybeHasNulls() {
		nulls := vec.Nulls().Copy()
		projVec.SetNulls(&nulls)
	}
	return batch
}

func (p projNotSuffixBytesBytesConstOp) Init() {
	p.input.Init()
}

type selNotRegexpBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg *regexp.Regexp
}

func (p *selNotRegexpBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	for {
		batch := p.input.Next(ctx)
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx uint16
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = !p.constArg.Match(arg)
					if cmp && !nulls.NullAt(i) {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !p.constArg.Match(arg)
					if cmp && !nulls.NullAt(uint16(i)) {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		} else {

			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(int(i))
					cmp = !p.constArg.Match(arg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				col = col.Slice(0, int(n))
				for i := 0; i < col.Len(); i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !p.constArg.Match(arg)
					if cmp {
						sel[idx] = uint16(i)
						idx++
					}
				}
			}

		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

func (p selNotRegexpBytesBytesConstOp) Init() {
	p.input.Init()
}

type projNotRegexpBytesBytesConstOp struct {
	OneInputNode

	colIdx   int
	constArg *regexp.Regexp

	outputIdx int
}

func (p projNotRegexpBytesBytesConstOp) EstimateStaticMemoryUsage() int {
	return EstimateBatchSizeBytes([]types.T{types.Bool}, coldata.BatchSize)
}

func (p projNotRegexpBytesBytesConstOp) Next(ctx context.Context) coldata.Batch {
	batch := p.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return batch
	}
	if p.outputIdx == batch.Width() {
		batch.AppendCol(types.Bool)
	}
	vec := batch.ColVec(p.colIdx)
	col := vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	projCol := projVec.Bool()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel {
			arg := col.Get(int(i))
			projCol[i] = !p.constArg.Match(arg)
		}
	} else {
		col = col.Slice(0, int(n))
		colLen := col.Len()
		_ = projCol[colLen-1]
		for i := 0; i < col.Len(); i++ {
			arg := col.Get(i)
			projCol[i] = !p.constArg.Match(arg)
		}
	}
	if vec.Nulls().MaybeHasNulls() {
		nulls := vec.Nulls().Copy()
		projVec.SetNulls(&nulls)
	}
	return batch
}

func (p projNotRegexpBytesBytesConstOp) Init() {
	p.input.Init()
}
