package internal

import (
	"math"
	"unsafe"
)

func f270(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	s0i32 = ctx.g0
	s1i32 = 384
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l9 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l7 = s0i32
	s0i32 = l4
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+376)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+368)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s1i32 = f105(ctx)
	s2i32 = 3
	s3i32 = l7
	s4i32 = 3
	s0i32 = f95(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s1i32 = l4
	s2i32 = 368
	s1i32 = s1i32 + s2i32
	f162(ctx, s0i32, s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	s0i32 = l4
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l4
	s2i32 = 336
	s1i32 = s1i32 + s2i32
	l8 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l4
	s2i32 = 80
	s1i32 = s1i32 + s2i32
	s2i32 = 256
	s3i32 = 256
	s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
	l8 = s0i32
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+380)]))
		l12 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l4
		s1f32 = l12
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+368)]))
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l4
		s1f32 = l12
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+372)]))
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l4
		s1f32 = l12
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+376)]))
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l6
		s1i32 = l3
		s2i32 = l4
		f174(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s1i32 = l1
		s2i32 = l3
		s3i32 = l6
		s4i32 = l4
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+380)]))
		s5f32 = 1
		if s4f32 == s5f32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		s5i32 = 1
		s0i32 = f432(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		l1 = s0i32
		goto lbl0
	}
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32) int32)(table[s1i32].f()))(ctx, s0i32)
	l10 = s0i32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+380)]))
	l12 = s0f32
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32) int32)(table[s1i32].f()))(ctx, s0i32)
	l11 = s0i32
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l4
	s2i32 = l5
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+60)]))
	if int(s2i32) < 0 || int(s2i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s2i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s2i32].numParams != 2 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32) int32)(table[s2i32].f()))(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+380)]))
		l13 = s0f32
		s1f32 = 1
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l13
			s0i32 = int32(math.Float32bits(s0f32))
			s1i32 = 0
			s2i32 = l3
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			l2 = s2i32
			s1i32 = s1i32 - s2i32
			s2i32 = 3
			s1i32 = s1i32 & s2i32
			l5 = s1i32
			s2i32 = 4
			s1i32 = s1i32 | s2i32
			s2i32 = l3
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
			s3i32 = l2
			s2i32 = s2i32 - s3i32
			if uint32(s1i32) <= uint32(s2i32) {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				goto lbl4
			}
			s0i32 = l3
			s1i32 = 4
			s2i32 = 4
			f18(ctx, s0i32, s1i32, s2i32)
			s0i32 = 0
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			l2 = s1i32
			s0i32 = s0i32 - s1i32
			s1i32 = 3
			s0i32 = s0i32 & s1i32
			l5 = s0i32
			s0i32 = l4
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+380)]))
		lbl4:
			l7 = s0i32
			s0i32 = l3
			s1i32 = l2
			s2i32 = l5
			s1i32 = s1i32 + s2i32
			l2 = s1i32
			s2i32 = 4
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l2
			s1i32 = l7
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = 92
			s2i32 = l2
			f16(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l0
		s1i32 = l1
		s2i32 = l3
		s3i32 = l6
		s4i32 = l10
		s5f32 = l12
		s6f32 = 1
		if s5f32 == s6f32 {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		s4i32 = s4i32 & s5i32
		s5i32 = l11
		s0i32 = f432(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		l1 = s0i32
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 20
	s2i32 = 4
	s0i32 = f56(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	s0i32 = l3
	s1i32 = l1
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 228
	s2i32 = l1
	s3i32 = l0
	s2i32 = s2i32 - s3i32
	f52(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	s1i32 = 20568
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
lbl0:
	s0i32 = l8
	f43(ctx, s0i32)
	s0i32 = l4
	s1i32 = 384
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l1
	return s0i32
}
