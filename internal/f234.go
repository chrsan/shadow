package internal

import (
	"math"
	"unsafe"
)

func f234(ctx *Context, l0 int32, l1 int32, l2 float32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
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
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l6 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l8 = s0f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l3 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l7 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l9 = s0f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l4 = s0i32
	s0i32 = l1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l5 = s1i32
	s1i64 = int64(uint32(s1i32))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l0 = s2i32
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l4
	s1i64 = int64(uint32(s1i32))
	s2i32 = l3
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l1
	s1f32 = l9
	s2i32 = l5
	s2f32 = math.Float32frombits(uint32(s2i32))
	l10 = s2f32
	s1f32 = s1f32 - s2f32
	s2f32 = l2
	s1f32 = s1f32 * s2f32
	s2f32 = l10
	s1f32 = s1f32 + s2f32
	l10 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l8
	s3i32 = l0
	s3f32 = math.Float32frombits(uint32(s3i32))
	l11 = s3f32
	s2f32 = s2f32 - s3f32
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	s3f32 = l11
	s2f32 = s2f32 + s3f32
	l11 = s2f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l1
	s1f32 = l7
	s2i32 = l4
	s2f32 = math.Float32frombits(uint32(s2i32))
	s3f32 = l7
	s2f32 = s2f32 - s3f32
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l12 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l6
	s3i32 = l3
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4f32 = l6
	s3f32 = s3f32 - s4f32
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	l13 = s2f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l1
	s1f32 = l10
	s2f32 = l9
	s3f32 = l7
	s4f32 = l9
	s3f32 = s3f32 - s4f32
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	l7 = s2f32
	s3f32 = l10
	s2f32 = s2f32 - s3f32
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l9 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l11
	s3f32 = l8
	s4f32 = l6
	s5f32 = l8
	s4f32 = s4f32 - s5f32
	s5f32 = l2
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	l6 = s3f32
	s4f32 = l11
	s3f32 = s3f32 - s4f32
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	l8 = s2f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l1
	s1f32 = l7
	s2f32 = l12
	s3f32 = l7
	s2f32 = s2f32 - s3f32
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l7 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l6
	s3f32 = l13
	s4f32 = l6
	s3f32 = s3f32 - s4f32
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	l6 = s2f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l1
	s1f32 = l9
	s2f32 = l7
	s3f32 = l9
	s2f32 = s2f32 - s3f32
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l8
	s3f32 = l6
	s4f32 = l8
	s3f32 = s3f32 - s4f32
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
}
