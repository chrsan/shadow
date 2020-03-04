package internal

import (
	"unsafe"
)

func f852(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
	var l6 int32
	_ = l6
	var l7 int64
	_ = l7
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
	var s6i32 int32
	_ = s6i32
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	var s6i64 int64
	_ = s6i64
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
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l6 = s2i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3i32 = l6
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s3i32 = s3i32 * s4i32
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l7 = s2i64
	s3i64 = 65535
	s2i64 = s2i64 & s3i64
	s2f32 = float32(uint64(s2i64))
	s3f32 = 1.5259022e-05
	s2f32 = s2f32 * s3f32
	s3i64 = l7
	s4i64 = 16
	s3i64 = int64(uint64(s3i64) >> (uint64(s4i64) & 63))
	s4i64 = 65535
	s3i64 = s3i64 & s4i64
	s3f32 = float32(uint64(s3i64))
	s4f32 = 1.5259022e-05
	s3f32 = s3f32 * s4f32
	s4i64 = l7
	s5i64 = 32
	s4i64 = int64(uint64(s4i64) >> (uint64(s5i64) & 63))
	s5i64 = 65535
	s4i64 = s4i64 & s5i64
	s4f32 = float32(uint64(s4i64))
	s5f32 = 1.5259022e-05
	s4f32 = s4f32 * s5f32
	s5i64 = l7
	s6i64 = 48
	s5i64 = int64(uint64(s5i64) >> (uint64(s6i64) & 63))
	s5f32 = float32(uint64(s5i64))
	s6f32 = 1.5259022e-05
	s5f32 = s5f32 * s6f32
	s6i32 = l1
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
	if int(s6i32) < 0 || int(s6i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s6i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s6i32].numParams != 6 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, float32, float32, float32, float32))(table[s6i32].f()))(ctx, s0i32, s1i32, s2f32, s3f32, s4f32, s5f32)
}
