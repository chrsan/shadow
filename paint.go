package shadow

import "github.com/chrsan/shadow/internal"

type Paint struct {
	c   *internal.Context
	ptr int32
}

func (p *Paint) Dispose() {
	if p != nil {
		p.c.SkiaPaintDestroy(p.ptr)
		p = nil
	}
}

type PaintStyle uint8

const (
	PaintStyleFill PaintStyle = iota
	PaintStyleStroke
	PaintStyleStrokeAndFill
)

func (p *Paint) Style() PaintStyle {
	return PaintStyle(p.c.SkiaPaintGetStyle(p.ptr))
}

func (p *Paint) SetStyle(s PaintStyle) {
	p.c.SkiaPaintSetStyle(p.ptr, int32(s))
}

func (p *Paint) Color() uint32 {
	return uint32(p.c.SkiaPaintGetColor(p.ptr))
}

func (p *Paint) SetColor(r, g, b, a uint8) {
	p.c.SkiaPaintSetColor(p.ptr, int32(r), int32(g), int32(b), int32(a))
}

func (p *Paint) Alpha() uint8 {
	return uint8(p.c.SkiaPaintGetAlpha(p.ptr))
}

func (p *Paint) SetAlpha(a uint8) {
	p.c.SkiaPaintSetAlpha(p.ptr, int32(a))
}

func (p *Paint) IsAntiAlias() bool {
	return p.c.SkiaPaintIsAntiAlias(p.ptr) != 0
}

func (p *Paint) SetAntiAlias(aa bool) {
	var b int32
	if aa {
		b = 1
	}
	p.c.SkiaPaintSetAntiAlias(p.ptr, b)
}

func (p *Paint) StrokeWidth() float32 {
	return p.c.SkiaPaintGetStrokeWidth(p.ptr)
}

func (p *Paint) SetStrokeWidth(w float32) {
	p.c.SkiaPaintSetStrokeWidth(p.ptr, w)
}

type StrokeCap uint8

const (
	StrokeCapButt StrokeCap = iota
	StrokeCapRound
	StrokeCapSquare
)

func (p *Paint) StrokeCap() StrokeCap {
	return StrokeCap(p.c.SkiaPaintGetStrokeCap(p.ptr))
}

func (p *Paint) SetStrokeCap(c StrokeCap) {
	p.c.SkiaPaintSetStrokeCap(p.ptr, int32(c))
}

type StrokeJoin uint8

const (
	StrokeJoinMiter StrokeJoin = iota
	StrokeJoinRound
	StrokeJoinBevel
)

func (p *Paint) StrokeJoin() StrokeJoin {
	return StrokeJoin(p.c.SkiaPaintGetStrokeJoin(p.ptr))
}

func (p *Paint) SetStrokeJoin(j StrokeJoin) {
	p.c.SkiaPaintSetStrokeJoin(p.ptr, int32(j))
}

func (p *Paint) StrokeMiter() float32 {
	return p.c.SkiaPaintGetStrokeMiter(p.ptr)
}

func (p *Paint) SetStrokeMiter(m float32) {
	p.c.SkiaPaintSetStrokeMiter(p.ptr, m)
}

func (p *Paint) Reset() {
	p.c.SkiaPaintReset(p.ptr)
}
