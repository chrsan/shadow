package internal

import (
	"unsafe"
)

func f1552(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int64
	_ = l5
	var l6 int64
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l3 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l4 = s0i32
	if s0i32 != 0 {
	lbl1:
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l3 = s0i32
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l1
			s2i32 = l2
			s3i32 = l3
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+44)]))
			if int(s3i32) < 0 || int(s3i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s3i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s3i32].numParams != 3 {
				panic("argument count mismatch")
			}
			(*(*func(*Context, int32, int32, int32))(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	} else {
		s0i32 = l3
	}
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l2
	f437(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l2 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l4 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l5 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l3 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l1 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l6 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0f32 = 0
	s1i64 = l5
	s2i64 = l6
	s1i64 = s1i64 | s2i64
	s2i64 = 2147483648
	s1i64 = s1i64 + s2i64
	s2i64 = 4294967295
	if uint64(s1i64) > uint64(s2i64) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s0f32 = float32(s0i32)
	l7 = s0f32
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s0f32 = float32(s0i32)
	l8 = s0f32
	s0i32 = l4
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s0f32 = float32(s0i32)
	l9 = s0f32
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s0f32 = float32(s0i32)
	goto lbl3
lbl4:
	s0f32 = 0
lbl3:
	l10 = s0f32
	s0i32 = l0
	s1f32 = l9
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4412)])) = s1f32
	s0i32 = l0
	s1i32 = 4424
	s0i32 = s0i32 + s1i32
	s1f32 = l7
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1i32 = 4420
	s0i32 = s0i32 + s1i32
	s1f32 = l8
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1i32 = 4416
	s0i32 = s0i32 + s1i32
	s1f32 = l10
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
}
