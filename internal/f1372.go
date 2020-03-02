package internal

import (
	"unsafe"
)

func f1372(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
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
	var l10 int64
	_ = l10
	var l11 int64
	_ = l11
	var l12 int64
	_ = l12
	var l13 int64
	_ = l13
	var l14 int64
	_ = l14
	var l15 int64
	_ = l15
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
	var s4i64 int64
	_ = s4i64
	s0i32 = ctx.g0
	s1i32 = 144
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s0i32 = f1293(ctx, s0i32, s1i32)
	l4 = s0i32
lbl1:
	s0i32 = l4
	s0i32 = f1292(ctx, s0i32)
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l5 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l6 = s0i32
	s0i32 = l2
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])))
	l14 = s0i64
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l7 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l8 = s0i32
	s0i32 = l2
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])))
	l15 = s0i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4388)]))
	l2 = s0i32
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 1
		f57(ctx, s0i32, s1i32)
	}
	s0i32 = l3
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l1
	s3i32 = 0
	s4i32 = 0
	s0i32 = f61(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l2 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l6
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	s0i64 = int64(s0i32)
	s1i64 = l14
	s0i64 = s0i64 + s1i64
	l10 = s0i64
	s1i64 = 2147483647
	s2i64 = l10
	s3i64 = 2147483647
	if s2i64 < s3i64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i64 = s0i64
	} else {
		s0i64 = s1i64
	}
	l10 = s0i64
	s1i64 = -2147483647
	s2i64 = l10
	s3i64 = -2147483647
	if s2i64 > s3i64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i64 = s0i64
	} else {
		s0i64 = s1i64
	}
	l10 = s0i64
	s0i32 = l7
	s1i32 = l8
	s0i32 = s0i32 - s1i32
	s0i64 = int64(s0i32)
	s1i64 = l15
	s0i64 = s0i64 + s1i64
	l11 = s0i64
	s1i64 = 2147483647
	s2i64 = l11
	s3i64 = 2147483647
	if s2i64 < s3i64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i64 = s0i64
	} else {
		s0i64 = s1i64
	}
	l11 = s0i64
	s1i64 = -2147483647
	s2i64 = l11
	s3i64 = -2147483647
	if s2i64 > s3i64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i64 = s0i64
	} else {
		s0i64 = s1i64
	}
	l11 = s0i64
lbl4:
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
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
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	f101(ctx, s0i32, s1i32)
	s0i32 = l3
	s1i64 = l10
	s2i32 = l3
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)])))
	l9 = s2i64
	s1i64 = s1i64 + s2i64
	l12 = s1i64
	s2i64 = 2147483647
	s3i64 = l12
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l12 = s1i64
	s2i64 = -2147483647
	s3i64 = l12
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i64)
	s0i32 = l3
	s1i64 = l11
	s2i32 = l3
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)])))
	l12 = s2i64
	s1i64 = s1i64 + s2i64
	l13 = s1i64
	s2i64 = 2147483647
	s3i64 = l13
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l13 = s1i64
	s2i64 = -2147483647
	s3i64 = l13
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i64)
	s0i32 = l3
	s1i64 = l9
	s2i64 = l14
	s1i64 = s1i64 + s2i64
	l9 = s1i64
	s2i64 = 2147483647
	s3i64 = l9
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l9 = s1i64
	s2i64 = -2147483647
	s3i64 = l9
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i64)
	s0i32 = l3
	s1i64 = l12
	s2i64 = l15
	s1i64 = s1i64 + s2i64
	l9 = s1i64
	s2i64 = 2147483647
	s3i64 = l9
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l9 = s1i64
	s2i64 = -2147483647
	s3i64 = l9
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i64)
	s0i32 = l1
	s1i32 = l3
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s3i32 = 16
	s2i32 = s2i32 + s3i32
	s1i32 = f224(ctx, s1i32, s2i32)
	l5 = s1i32
	s2i32 = 1
	s3i32 = l1
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
	s0i32 = l1
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+72)]))
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l6 = s0i32
	s0i32 = l1
	s1i32 = l1
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
	(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
	s0i32 = l1
	s1i32 = l6
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	f208(ctx, s0i32, s1i32)
	s0i32 = l5
	f40(ctx, s0i32)
	s0i32 = l4
	if s0i32 != 0 {
		goto lbl4
	}
lbl3:
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+64)])
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		f54(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = f23(ctx, s0i32)
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
lbl0:
	s0i32 = l3
	s1i32 = 144
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
