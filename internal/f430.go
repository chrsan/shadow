package internal

import (
	"unsafe"
)

func f430(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s1i64 int64
	_ = s1i64
	var s5i64 int64
	_ = s5i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)]))
	l5 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l2
		s3i32 = l3
		s4i32 = l4
		s5i32 = l0
		s5i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s5i32+104)]))
		s6i32 = l5
		if int(s6i32) < 0 || int(s6i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s6i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s6i32].numParams != 6 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32, int64))(table[s6i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i64)
		goto lbl0
	}
	s0i32 = l0
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)]))
	s0i32 = int32(ctx.Mem[int(s0i32+8)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		l5 = s0i32
		goto lbl2
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l6 = s0i32
	s0i32 = l7
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l0
	s2i32 = 48
	s1i32 = s1i32 + s2i32
	f192(ctx, s0i32, s1i32)
	s0i32 = l5
	s1i32 = l0
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	f172(ctx, s0i32, s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl8
	case 1:
		goto lbl9
	default:
		goto lbl6
	}
lbl9:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l6 = s0i32
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l9
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s1i32 = 3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+196)]))
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l6
	s1i32 = 6
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 15
		s2i32 = 0
		f16(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l5
	s1i32 = 124
	s2i32 = l0
	s3i32 = -64
	s2i32 = s2i32 - s3i32
	f16(ctx, s0i32, s1i32, s2i32)
	goto lbl5
lbl8:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l6 = s0i32
lbl7:
	s0i32 = l5
	s1i32 = l6
	s2i32 = l0
	s3i32 = -64
	s2i32 = s2i32 - s3i32
	f149(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s1i32 = 3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 10
		s2i32 = 0
		f16(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	s1i32 = l5
	f99(ctx, s0i32, s1i32)
lbl6:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s1i32 = 3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 8
		s2i32 = 0
		f16(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+196)]))
	s1f32 = 0
	if s0f32 > s1f32 {
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
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 25
		s2i32 = l0
		s3i32 = 196
		s2i32 = s2i32 + s3i32
		f16(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s2i32 = l0
	s3i32 = -64
	s2i32 = s2i32 - s3i32
	f173(ctx, s0i32, s1i32, s2i32)
lbl5:
	s0i32 = l7
	s1i32 = l5
	f271(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = 832
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)]))
	l5 = s0i32
	s0i32 = l0
	s1i32 = 4728
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint32(s1i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	if s0i32 != 0 {
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l5
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
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint32(s1i32)
lbl2:
	s0i32 = l8
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l4
	s5i32 = l5
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
lbl0:
	s0i32 = l7
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
