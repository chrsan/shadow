package internal

import (
	"unsafe"
)

func f1666(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
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
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 288
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = 264
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = 2
	ctx.Mem[int(s0i32+18)] = uint8(s1i32)
	s0i32 = l4
	s1i32 = 288
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint16(s1i32)
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l4
	s1i64 = 4575657222503137280
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = 2
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s0i32 = l4
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l5 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
lbl1:
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l3 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
lbl2:
	s0i32 = l4
	s1f32 = 2048
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l4
	s1i32 = l4
	s1i32 = int32(ctx.Mem[int(s1i32+16)])
	s2i32 = 8
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+16)] = uint8(s1i32)
	s0i32 = l2
	s1i32 = 0
	ctx.Mem[int(s0i32+255)] = uint8(s1i32)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l2
	s2i32 = 80
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
	s1i32 = l3
	s2i64 = 13195221663744
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)])) = uint64(s2i64)
	s1i32 = l3
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])) = uint64(s2i64)
	s1i32 = l3
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
	s1i32 = l3
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
	s1i32 = l3
	s2i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)])) = uint64(s2i64)
	s1i32 = l3
	s2i32 = l2
	s2i32 = f212(ctx, s2i32)
	s3i32 = 0
	s4i32 = 21916
	s5i32 = l2
	s6i32 = 208
	s5i32 = s5i32 + s6i32
	s6i32 = l2
	s7i32 = 200
	s6i32 = s6i32 + s7i32
	f645(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	s0i32 = l3
	s0i32 = f23(ctx, s0i32)
	s0i32 = l2
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = 208
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 72
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	f644(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s1i32 = l0
	s2i32 = l2
	s3i32 = 72
	s2i32 = s2i32 + s3i32
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s1i32 = (*(*func(*Context, int32, int32, int32) int32)(table[s4i32].f()))(ctx, s1i32, s2i32, s3i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	l0 = s0i32
	s1i32 = l2
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	if int(s2i32) < 0 || int(s2i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s2i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s2i32].numParams != 2 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32))(table[s2i32].f()))(ctx, s0i32, s1i32)
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l6 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l8 = s0f32
	s0i32 = l1
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2f32 = 0.00048828125
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l1
	s1f32 = l8
	s2f32 = 0.00048828125
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l1
	s1f32 = l7
	s2f32 = 0.00048828125
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l1
	s1f32 = l6
	s2f32 = 0.00048828125
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	l0 = s0i32
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	s0i32 = l0
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		if int(s1i32) < 0 || int(s1i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s1i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s1i32].numParams != 1 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
	}
	s0i32 = l3
	f159(ctx, s0i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
lbl5:
	s0i32 = l2
	s1i32 = 288
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = 1
	return s0i32
}
