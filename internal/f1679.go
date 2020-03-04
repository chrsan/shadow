package internal

import (
	"math"
	"unsafe"
)

func f1679(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float32, l6 float32, l7 int32, l8 int32) {
	var l9 int64
	_ = l9
	var l10 int64
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
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
	var s5i32 int32
	_ = s5i32
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
	s0i32 = ctx.g0
	s1i32 = 208
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l6 = s0f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2f32 = 1
	s3f32 = l6
	s2f32 = s2f32 - s3f32
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 0.00024414062
	if s2f32 <= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l7
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l9 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l10 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = uint64(s1i64)
	s0i32 = 0
	s1i64 = l9
	s1i32 = int32(s1i64)
	s1f32 = math.Float32frombits(uint32(s1i32))
	l6 = s1f32
	s2i64 = l10
	s3i64 = 32
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	s2i32 = int32(s2i64)
	s2f32 = math.Float32frombits(uint32(s2i32))
	l11 = s2f32
	s1f32 = s1f32 * s2f32
	s2i64 = l9
	s3i64 = 32
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	s2i32 = int32(s2i64)
	s2f32 = math.Float32frombits(uint32(s2i32))
	l12 = s2f32
	s3i64 = l10
	s3i32 = int32(s3i64)
	s3f32 = math.Float32frombits(uint32(s3i32))
	l13 = s3f32
	s2f32 = s2f32 * s3f32
	if s1f32 > s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s0i32 = l7
	s1f32 = l12
	s1f32 = -s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+204)])) = s1f32
	s0i32 = l7
	s1f32 = l6
	s1f32 = -s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = s1f32
	s0i32 = l7
	s1f32 = l11
	s1f32 = -s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+196)])) = s1f32
	s0i32 = l7
	s1f32 = l13
	s1f32 = -s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = s1f32
	s0i32 = l0
	l2 = s0i32
	s0i32 = l1
	l0 = s0i32
	s0i32 = l2
	l1 = s0i32
	s0i32 = 1
lbl1:
	l2 = s0i32
	s0i32 = l7
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 152
	s0i32 = s0i32 + s1i32
	s1f32 = l5
	s2f32 = l5
	f240(ctx, s0i32, s1f32, s2f32)
	s0i32 = l7
	s1i32 = 152
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	f142(ctx, s0i32, s1f32, s2f32)
	s0i32 = l7
	s1i32 = 200
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s2i32 = 192
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s3i32 = l7
	s4i32 = 152
	s3i32 = s3i32 + s4i32
	s4i32 = l7
	s0i32 = f687(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l8 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 0
	l4 = s0i32
lbl2:
	s0i32 = l0
	s1i32 = l7
	s2i32 = l4
	s3i32 = 28
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s3i32 = l2
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	s4i32 = l2
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
	s5i32 = l2
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+24)]))
	s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l8
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)]))
	l6 = s0f32
	s0i32 = l7
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+196)]))
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+196)])) = s1f32
	s0i32 = l7
	s1f32 = l6
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = s1f32
	s0i32 = l1
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s0i32 = l1
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l7
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+192)]))
	s1f32 = s1f32 - s2f32
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = l7
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+196)]))
	s2f32 = s2f32 - s3f32
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
lbl0:
	s0i32 = l7
	s1i32 = 208
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
