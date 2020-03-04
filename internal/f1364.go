package internal

import (
	"math"
	"unsafe"
)

func f1364(ctx *Context, l0 int64, l1 int64) float64 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int64
	_ = l4
	var l5 int64
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s0f64 float64
	_ = s0f64
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i64 = l1
	s1i64 = 9223372036854775807
	s0i64 = s0i64 & s1i64
	l5 = s0i64
	s1i64 = -4323737117252386816
	s0i64 = s0i64 + s1i64
	s1i64 = l5
	s2i64 = -4899634919602388992
	s1i64 = s1i64 + s2i64
	if uint64(s0i64) < uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i64 = l1
		s1i64 = 4
		s0i64 = s0i64 << (uint64(s1i64) & 63)
		s1i64 = l0
		s2i64 = 60
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		s0i64 = s0i64 | s1i64
		l4 = s0i64
		s0i64 = l0
		s1i64 = 1152921504606846975
		s0i64 = s0i64 & s1i64
		l0 = s0i64
		s1i64 = 576460752303423489
		if uint64(s0i64) >= uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i64 = l4
			s1i64 = 4611686018427387905
			s0i64 = s0i64 + s1i64
			l4 = s0i64
			goto lbl0
		}
		s0i64 = l4
		s1i64 = -4611686018427387904
		s0i64 = s0i64 - s1i64
		l4 = s0i64
		s0i64 = l0
		s1i64 = 576460752303423488
		s0i64 = s0i64 ^ s1i64
		s1i64 = 0
		if s0i64 != s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i64 = l4
		s1i64 = 1
		s0i64 = s0i64 & s1i64
		s1i64 = l4
		s0i64 = s0i64 + s1i64
		l4 = s0i64
		goto lbl0
	}
	s0i64 = l0
	if s0i64 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i64 = l5
	s2i64 = 9223090561878065152
	if uint64(s1i64) < uint64(s2i64) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i64 = l5
	s3i64 = 9223090561878065152
	if s2i64 == s3i64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i64 = l1
		s1i64 = 4
		s0i64 = s0i64 << (uint64(s1i64) & 63)
		s1i64 = l0
		s2i64 = 60
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		s0i64 = s0i64 | s1i64
		s1i64 = 2251799813685247
		s0i64 = s0i64 & s1i64
		s1i64 = 9221120237041090560
		s0i64 = s0i64 | s1i64
		l4 = s0i64
		goto lbl0
	}
	s0i64 = 9218868437227405312
	l4 = s0i64
	s0i64 = l5
	s1i64 = 4899634919602388991
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = 0
	l4 = s0i64
	s0i64 = l5
	s1i64 = 48
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	s0i32 = int32(s0i64)
	l3 = s0i32
	s1i32 = 15249
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i64 = l0
	s2i64 = l1
	s3i64 = 281474976710655
	s2i64 = s2i64 & s3i64
	s3i64 = 281474976710656
	s2i64 = s2i64 | s3i64
	l4 = s2i64
	s3i32 = 15361
	s4i32 = l3
	s3i32 = s3i32 - s4i32
	f1366(ctx, s0i32, s1i64, s2i64, s3i32)
	s0i32 = l2
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i64 = l0
	s2i64 = l4
	s3i32 = l3
	s4i32 = -15233
	s3i32 = s3i32 + s4i32
	f1365(ctx, s0i32, s1i64, s2i64, s3i32)
	s0i32 = l2
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i64 = 4
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i64
	s2i64 = 60
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s0i64 = s0i64 | s1i64
	l4 = s0i64
	s0i32 = l2
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s0i64 = s0i64 | s1i64
	s1i64 = 0
	if s0i64 != s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s0i64 = int64(uint32(s0i32))
	s1i64 = l0
	s2i64 = 1152921504606846975
	s1i64 = s1i64 & s2i64
	s0i64 = s0i64 | s1i64
	l0 = s0i64
	s1i64 = 576460752303423489
	if uint64(s0i64) >= uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i64 = l4
		s1i64 = 1
		s0i64 = s0i64 + s1i64
		l4 = s0i64
		goto lbl0
	}
	s0i64 = l0
	s1i64 = 576460752303423488
	s0i64 = s0i64 ^ s1i64
	s1i64 = 0
	if s0i64 != s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = l4
	s1i64 = 1
	s0i64 = s0i64 & s1i64
	s1i64 = l4
	s0i64 = s0i64 + s1i64
	l4 = s0i64
lbl0:
	s0i32 = l2
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i64 = l4
	s1i64 = l1
	s2i64 = -9223372036854775808
	s1i64 = s1i64 & s2i64
	s0i64 = s0i64 | s1i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	return s0f64
}
