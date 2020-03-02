package shadow

import (
	"fmt"
	"unsafe"

	"github.com/chrsan/shadow/internal"
)

type Path struct {
	c   *internal.Context
	ptr int32
}

func (p *Path) Dispose() {
	if p != nil {
		p.c.SkiaPathDestroy(p.ptr)
		p = nil
	}
}

type PathFillType uint8

const (
	PathFillTypeWinding PathFillType = iota
	PathFillTypeEvenOdd
)

func (p *Path) FillType() PathFillType {
	return PathFillType(p.c.SkiaPathGetFillType(p.ptr))
}

func (p *Path) SetFillType(ft PathFillType) {
	p.c.SkiaPathSetFillType(p.ptr, int32(ft))
}

func (p *Path) MoveTo(x, y float32) {
	p.c.SkiaPathMoveTo(p.ptr, x, y)
}

func (p *Path) LineTo(x, y float32) {
	p.c.SkiaPathLineTo(p.ptr, x, y)
}

func (p *Path) QuadTo(x1, y1, x2, y2 float32) {
	p.c.SkiaPathQuadTo(p.ptr, x1, y1, x2, y2)
}

func (p *Path) ConicTo(x1, y1, x2, y2, weight float32) {
	p.c.SkiaPathConicTo(p.ptr, x1, y1, x2, y2, weight)
}

func (p *Path) CubicTo(x1, y1, x2, y2, x3, y3 float32) {
	p.c.SkiaPathCubicTo(p.ptr, x1, y1, x2, y2, x3, y3)
}

func (p *Path) Close() {
	p.c.SkiaPathClose(p.ptr)
}

func (p *Path) Reset() {
	p.c.SkiaPathReset(p.ptr)
}

func (p *Path) Rewind() {
	p.c.SkiaPathRewind(p.ptr)
}

func (p *Path) NumPoints() int {
	return int(p.c.SkiaPathCountPoints(p.ptr))
}

func (p *Path) Point(idx int) Point {
	st := p.c.StackSave()
	defer p.c.StackRestore(st)
	ptr := p.c.StackAlloc(8)
	p.c.SkiaPathGetPoint(ptr, p.ptr, int32(idx))
	return *(*Point)(unsafe.Pointer(&p.c.Mem[ptr]))
}

func (p *Path) Points(peek bool) []Point {
	n := p.c.SkiaPathCountPoints(p.ptr)
	if n == 0 {
		return nil
	}
	st := p.c.StackSave()
	defer p.c.StackRestore(st)
	ptr := p.c.StackAlloc(n * 8)
	m := p.c.SkiaPathGetPoints(p.ptr, ptr, n)
	if m != n {
		panic(fmt.Sprintf("expected %d points, got %d", n, m))
	}
	rs := (*[1 << 29]Point)(unsafe.Pointer(&p.c.Mem[ptr]))[:n:n]
	if peek {
		return rs
	}
	ps := make([]Point, n)
	copy(ps, rs)
	return ps
}

func (p *Path) NumVerbs() int {
	return int(p.c.SkiaPathCountVerbs(p.ptr))
}

type PathVerb uint8

const (
	PathVerbMove PathVerb = iota
	PathVerbLine
	PathVerbQuad
	PathVerbConic
	PathVerbCubic
	PathVerbClose
	PathVerbDone
)

func (p *Path) Verbs(peek bool) []PathVerb {
	n := p.c.SkiaPathCountVerbs(p.ptr)
	if n == 0 {
		return nil
	}
	st := p.c.StackSave()
	defer p.c.StackRestore(st)
	ptr := p.c.StackAlloc(n)
	m := p.c.SkiaPathGetVerbs(p.ptr, ptr, n)
	if m != n {
		panic(fmt.Sprintf("expected %d verbs, got %d", n, m))
	}
	rs := (*[1 << 29]PathVerb)(unsafe.Pointer(&p.c.Mem[ptr]))[:n:n]
	if peek {
		return rs
	}
	vs := make([]PathVerb, n)
	copy(vs, rs)
	return vs
}

func (p *Path) Bounds() Rect {
	st := p.c.StackSave()
	defer p.c.StackRestore(st)
	ptr := p.c.StackAlloc(16)
	p.c.SkiaPathGetBounds(ptr, p.ptr)
	return *(*Rect)(unsafe.Pointer(&p.c.Mem[ptr]))
}

func (p *Path) ComputeTightBounds() Rect {
	st := p.c.StackSave()
	defer p.c.StackRestore(st)
	ptr := p.c.StackAlloc(16)
	p.c.SkiaPathComputeTightBounds(ptr, p.ptr)
	return *(*Rect)(unsafe.Pointer(&p.c.Mem[ptr]))
}
