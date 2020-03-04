package internal

import (
	"math"
	"unsafe"
)

func f447(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int64
	_ = l9
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float32
	_ = l16
	var l17 float32
	_ = l17
	var l18 float32
	_ = l18
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
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	var s7f32 float32
	_ = s7f32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		return
	}
	s0i32 = l3
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0f32
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l8 = s0i32
lbl2:
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l11 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l13 = s0f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l6 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l12 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l14 = s0f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l7 = s0i32
	s0i32 = l1
	l3 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s1i64 = int64(uint32(s1i32))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l0 = s2i32
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l7
	s1i64 = int64(uint32(s1i32))
	s2i32 = l6
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l3
	s1f32 = l10
	s2f32 = l14
	s3i32 = l1
	s3f32 = math.Float32frombits(uint32(s3i32))
	l15 = s3f32
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l15
	s1f32 = s1f32 + s2f32
	l15 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l10
	s3f32 = l13
	s4i32 = l0
	s4f32 = math.Float32frombits(uint32(s4i32))
	l16 = s4f32
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s3f32 = l16
	s2f32 = s2f32 + s3f32
	l16 = s2f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1f32 = l12
	s2f32 = l10
	s3i32 = l7
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4f32 = l12
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l17 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l11
	s3f32 = l10
	s4i32 = l6
	s4f32 = math.Float32frombits(uint32(s4i32))
	s5f32 = l11
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	l18 = s2f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l3
	s1f32 = l15
	s2f32 = l10
	s3f32 = l14
	s4f32 = l10
	s5f32 = l12
	s6f32 = l14
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	l12 = s3f32
	s4f32 = l15
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l14 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l16
	s3f32 = l10
	s4f32 = l13
	s5f32 = l10
	s6f32 = l11
	s7f32 = l13
	s6f32 = s6f32 - s7f32
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	l11 = s4f32
	s5f32 = l16
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	l13 = s2f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s1f32 = l12
	s2f32 = l10
	s3f32 = l17
	s4f32 = l12
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l12 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l11
	s3f32 = l10
	s4f32 = l18
	s5f32 = l11
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	l11 = s2f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l3
	s1f32 = l14
	s2f32 = l10
	s3f32 = l12
	s4f32 = l14
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l13
	s3f32 = l10
	s4f32 = l11
	s5f32 = l13
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l8
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = l3
	s2i32 = 24
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l0 = s0i32
	s0i32 = l2
	s1i32 = l5
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l0
	s2i32 = l2
	s1i32 = s1i32 + s2i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l11 = s1f32
	s0f32 = s0f32 - s1f32
	l10 = s0f32
	s0f32 = -s0f32
	s1f32 = l10
	s2f32 = l10
	s3f32 = 0
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l0 = s2i32
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l10 = s0f32
	s1f32 = 1
	s2f32 = l11
	s1f32 = s1f32 - s2f32
	l11 = s1f32
	s1f32 = -s1f32
	s2f32 = l11
	s3i32 = l0
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l11 = s1f32
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l11
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l10
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l4
	l0 = s0i32
	s0f32 = l10
	s1f32 = l11
	s0f32 = s0f32 / s1f32
	l10 = s0f32
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l10
	s2f32 = l10
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl2
	}
lbl3:
	s0i32 = l3
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	l9 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = l9
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = l9
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
lbl0:
}
