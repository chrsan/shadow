package internal

import (
	"math"
	"unsafe"
)

func f543(ctx *Context, l0 float64, l1 int32) float64 {
	var l2 int32
	_ = l2
	var l3 int64
	_ = l3
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
	var s2f64 float64
	_ = s2f64
	s0f64 = l0
	s0i64 = int64(math.Float64bits(s0f64))
	l3 = s0i64
	s1i64 = 52
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	s0i32 = int32(s0i64)
	s1i32 = 2047
	s0i32 = s0i32 & s1i32
	l2 = s0i32
	s1i32 = 2047
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1f64 = l0
			s2f64 = 0
			if s1f64 == s2f64 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1i32 = 0
			} else {
				s1f64 = l0
				s2f64 = 1.8446744073709552e+19
				s1f64 = s1f64 * s2f64
				s2i32 = l1
				s1f64 = f543(ctx, s1f64, s2i32)
				l0 = s1f64
				s1i32 = l1
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
				s2i32 = -64
				s1i32 = s1i32 + s2i32
			}
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0f64 = l0
			return s0f64
		}
		s0i32 = l1
		s1i32 = l2
		s2i32 = -1022
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i64 = l3
		s1i64 = -9218868437227405313
		s0i64 = s0i64 & s1i64
		s1i64 = 4602678819172646912
		s0i64 = s0i64 | s1i64
		s0f64 = math.Float64frombits(uint64(s0i64))
	} else {
		s0f64 = l0
	}
	return s0f64
}
