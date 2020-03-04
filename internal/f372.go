package internal

import (
	"math"
	"unsafe"
)

func f372(ctx *Context, l0 float64) float64 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	ctx.g0 = s0i32
	s0f64 = l0
	s0i64 = int64(math.Float64bits(s0f64))
	s1i64 = 32
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	s0i32 = int32(s0i64)
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l2 = s0i32
	s1i32 = 1072243195
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = 1
		s1i32 = l2
		s2i32 = 1044816030
		if uint32(s1i32) < uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl0
		}
		s0f64 = l0
		s1f64 = 0
		s0f64 = f373(ctx, s0f64, s1f64)
		goto lbl0
	}
	s0f64 = l0
	s1f64 = l0
	s0f64 = s0f64 - s1f64
	s1i32 = l2
	s2i32 = 2146435072
	if uint32(s1i32) >= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f64 = l0
	s1i32 = l1
	s0i32 = f1363(ctx, s0f64, s1i32)
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l2 = s0i32
	s1i32 = 2
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl4
		case 1:
			goto lbl3
		default:
			goto lbl5
		}
	lbl5:
		s0i32 = l1
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l1
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s0f64 = f373(ctx, s0f64, s1f64)
		goto lbl0
	lbl4:
		s0i32 = l1
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l1
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s0f64 = f537(ctx, s0f64, s1f64)
		s0f64 = -s0f64
		goto lbl0
	lbl3:
		s0i32 = l1
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l1
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s0f64 = f373(ctx, s0f64, s1f64)
		s0f64 = -s0f64
		goto lbl0
	}
	s0i32 = l1
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l1
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0f64 = f537(ctx, s0f64, s1f64)
lbl0:
	l0 = s0f64
	s0i32 = l1
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0f64 = l0
	return s0f64
}
