package internal

import (
	"math"
	"unsafe"
)

func f1114(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 float32
	_ = l17
	var l18 float32
	_ = l18
	var l19 float32
	_ = l19
	var l20 float32
	_ = l20
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
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
	s0i32 = ctx.g0
	s1i32 = 672
	s0i32 = s0i32 - s1i32
	l8 = s0i32
	ctx.g0 = s0i32
	s0i32 = l8
	s1i32 = 632
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 0
	f392(ctx, s0i32, s1i32, s2i32)
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+632)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l8
	s1i32 = 584
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s0i32 = f46(ctx, s0i32, s1i32)
	l10 = s0i32
	s1i32 = l10
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = -2
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l10
	s1i32 = l10
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = -193
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+576)])) = uint32(s1i32)
	s0i32 = l10
	s1i32 = l8
	s2i32 = 576
	s1i32 = s1i32 + s2i32
	f83(ctx, s0i32, s1i32)
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+576)]))
	l7 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l7
	s1i32 = l7
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
		goto lbl1
	}
	s0i32 = l7
	s1i32 = l7
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
	s0i32 = l8
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+568)])) = uint32(s1i32)
	s0i32 = l10
	s1i32 = l8
	s2i32 = 568
	s1i32 = s1i32 + s2i32
	f235(ctx, s0i32, s1i32)
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+568)]))
	l7 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l7
	s1i32 = l7
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
		goto lbl2
	}
	s0i32 = l7
	s1i32 = l7
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
	s0i32 = l8
	s1i32 = 536
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = l8
	s2i32 = 280
	s1i32 = s1i32 + s2i32
	s2i32 = 256
	s3i32 = 256
	s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
	l14 = s0i32
	l1 = s0i32
	s0i32 = l8
	s1i32 = 264
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l9
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l9
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l9
	l12 = s0i32
	s0i32 = l8
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+204)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l12
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+208)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l1 = s0i32
	s0i32 = l8
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+220)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l10
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+216)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+212)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	l1 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+256)])) = uint64(s1i64)
	s0i32 = l8
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+248)])) = uint64(s1i64)
	s0i32 = l8
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+240)])) = uint64(s1i64)
	s0i32 = l8
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+232)])) = uint64(s1i64)
	s0i32 = l8
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)])) = uint64(s1i64)
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+632)]))
	l1 = s0i32
	s1i32 = l8
	s2i32 = 200
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+64)]))
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
	l16 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = 21300
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l1 = s0i32
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l1
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		}
		s0i32 = l8
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l8
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
		s0i32 = l8
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+632)]))
		l1 = s0i32
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l1
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		}
		s0i32 = l8
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = uint32(s1i32)
		s0i32 = l10
		s1i32 = l8
		s2i32 = 192
		s1i32 = s1i32 + s2i32
		f83(ctx, s0i32, s1i32)
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)]))
		l7 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l7
		s1i32 = l7
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
			goto lbl8
		}
		s0i32 = l7
		s1i32 = l7
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
	lbl8:
		s0i32 = l5
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = 0
		l1 = s0i32
		s0i32 = l4
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl10:
			s0i32 = l8
			s1i64 = 69784829952
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
			s0i32 = l8
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
			s0i32 = l8
			s1i64 = 1065353216
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
			s0i32 = l8
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
			s0i32 = l8
			s1i64 = 1065353216
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
			s0i32 = l8
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = l2
			s2i32 = l1
			s3i32 = 4
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			l4 = s2i32
			s1i32 = s1i32 + s2i32
			s0i32 = f283(ctx, s0i32, s1i32)
			s0i32 = l8
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = l3
			s2i32 = l4
			s1i32 = s1i32 + s2i32
			l4 = s1i32
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s1f32 = -s1f32
			s2i32 = l4
			s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			s2f32 = -s2f32
			f200(ctx, s0i32, s1f32, s2f32)
			s0i32 = l8
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
			f116(ctx, s0i32, s1i32)
			s0i32 = l8
			s1i32 = l8
			s2i32 = 128
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
			s0i32 = l8
			s1i32 = 24
			s0i32 = s0i32 + s1i32
			s1i32 = l4
			s2i32 = l10
			s3i32 = 0
			s4i32 = 0
			f160(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
			s0i32 = l1
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s1i32 = l5
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl10
			}
			goto lbl4
			panic("unreachable executed")
			panic("unreachable executed")
		}
	lbl11:
		s0i32 = l8
		s1i32 = 176
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l1
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		f573(ctx, s0i32, s1i32)
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+632)]))
		l7 = s0i32
		if s0i32 != 0 {
			s0i32 = l7
			s1i32 = l7
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		}
		s0i32 = l8
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = 184
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = l8
		s3i32 = 176
		s2i32 = s2i32 + s3i32
		s3i32 = l8
		s4i32 = 168
		s3i32 = s3i32 + s4i32
		f581(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l10
		s1i32 = l8
		s2i32 = 184
		s1i32 = s1i32 + s2i32
		f83(ctx, s0i32, s1i32)
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)]))
		l9 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl13
		}
		s0i32 = l9
		s1i32 = l9
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l7 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl13
		}
		s0i32 = l9
		s1i32 = l9
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
	lbl13:
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)]))
		l9 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		s0i32 = l9
		s1i32 = l9
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l7 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		s0i32 = l9
		s1i32 = l9
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
	lbl14:
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)]))
		l9 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l9
		s1i32 = l9
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l7 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l9
		s1i32 = l9
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
	lbl15:
		s0i32 = l8
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
		s0i32 = l8
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
		s0i32 = l8
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
		s0i32 = l8
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
		s0i32 = l8
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
		s0i32 = l8
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = l1
		s3i32 = 4
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l7 = s2i32
		s1i32 = s1i32 + s2i32
		s0i32 = f283(ctx, s0i32, s1i32)
		s0i32 = l8
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l7
		s1i32 = s1i32 + s2i32
		l7 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1f32 = -s1f32
		s2i32 = l7
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s2f32 = -s2f32
		f200(ctx, s0i32, s1f32, s2f32)
		s0i32 = l8
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		f116(ctx, s0i32, s1i32)
		s0i32 = l8
		s1i32 = l8
		s2i32 = 128
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = 24
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s2i32 = l10
		s3i32 = 0
		s4i32 = 0
		f160(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l5
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl11
		}
		goto lbl4
	}
	s0i32 = l8
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	s1i32 = f105(ctx)
	s2i32 = 3
	s3i32 = l8
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+212)]))
	s4i32 = 3
	s0i32 = f95(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l9 = s0i32
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+540)]))
		l1 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l11 = s0i32
		s1i32 = 24
		s0i32 = s0i32 | s1i32
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+544)]))
		s2i32 = l1
		s1i32 = s1i32 - s2i32
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l14
			s1i32 = 24
			s2i32 = 4
			f18(ctx, s0i32, s1i32, s2i32)
			s0i32 = 0
			s1i32 = l8
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+540)]))
			l1 = s1i32
			s0i32 = s0i32 - s1i32
			s1i32 = 3
			s0i32 = s0i32 & s1i32
			l11 = s0i32
		}
		s0i32 = l8
		s1i32 = l1
		s2i32 = l11
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s2i32 = 24
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+540)])) = uint32(s1i32)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)]))
		s1i32 = 23
		s2i32 = l1
		f16(ctx, s0i32, s1i32, s2i32)
		s0i32 = l6
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+200)]))
		f99(ctx, s0i32, s1i32)
		goto lbl16
	}
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+632)]))
	l1 = s0i32
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
	s0i32 = (*(*func(*Context, int32) int32)(table[s1i32].f()))(ctx, s0i32)
	l13 = s0i32
	s0i32 = 0
	l1 = s0i32
lbl16:
	s0i32 = l10
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l17 = s0f32
	s1f32 = 1
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l13 = s0i32
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)]))
		l7 = s0i32
		s0i32 = 0
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+540)]))
		l6 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l11 = s0i32
		s1i32 = 4
		s0i32 = s0i32 | s1i32
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+544)]))
		s2i32 = l6
		s1i32 = s1i32 - s2i32
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l14
			s1i32 = 4
			s2i32 = 4
			f18(ctx, s0i32, s1i32, s2i32)
			s0i32 = 0
			s1i32 = l8
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+540)]))
			l6 = s1i32
			s0i32 = s0i32 - s1i32
			s1i32 = 3
			s0i32 = s0i32 & s1i32
			l11 = s0i32
		}
		s0i32 = l8
		s1i32 = l6
		s2i32 = l11
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+540)])) = uint32(s1i32)
		s0i32 = l6
		s1f32 = l17
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l7
		s1i32 = 92
		s2i32 = l6
		f16(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l0
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l10
	s2i32 = l12
	s3i32 = l13
	s4i32 = l14
	s0i32 = f432(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l12 = s0i32
	s0i32 = l8
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s0i32 = f38(ctx, s0i32)
	l15 = s0i32
	s0i32 = l5
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l7 = s0i32
	lbl22:
		s0i32 = l4
		if s0i32 != 0 {
			s0i32 = l8
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = l4
			s2i32 = l7
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			f185(ctx, s0i32, s1i32)
			s0i32 = l9
			s1i32 = l8
			s2i32 = 128
			s1i32 = s1i32 + s2i32
			f162(ctx, s0i32, s1i32)
			s0i32 = l8
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
			l20 = s0f32
			s0i32 = l8
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)]))
			l19 = s0f32
			s0i32 = l8
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)]))
			l17 = s0f32
			s0i32 = l1
			s1i32 = l8
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
			l18 = s1f32
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
			s0i32 = l1
			s1f32 = l17
			s2f32 = l18
			s1f32 = s1f32 * s2f32
			l21 = s1f32
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
			s0i32 = l1
			s1f32 = l18
			s2f32 = l19
			s1f32 = s1f32 * s2f32
			l22 = s1f32
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
			s0i32 = l1
			s1f32 = l18
			s2f32 = l20
			s1f32 = s1f32 * s2f32
			l19 = s1f32
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
			s0i32 = l1
			s1f32 = l18
			s2f32 = 255
			s1f32 = s1f32 * s2f32
			s2f32 = 0.5
			s1f32 = s1f32 + s2f32
			s1f32 = float32(math.Floor(float64(s1f32)))
			l17 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l17
			s4f32 = 2.1474835e+09
			if s3f32 < s4f32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1f32 = s1f32
			} else {
				s1f32 = s2f32
			}
			l17 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l17
			s4f32 = -2.1474835e+09
			if s3f32 > s4f32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1f32 = s1f32
			} else {
				s1f32 = s2f32
			}
			l17 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l17
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl24
			}
			s1i32 = -2147483648
		lbl24:
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+22)])) = uint16(s1i32)
			s0i32 = l1
			s1f32 = l22
			s2f32 = 255
			s1f32 = s1f32 * s2f32
			s2f32 = 0.5
			s1f32 = s1f32 + s2f32
			s1f32 = float32(math.Floor(float64(s1f32)))
			l17 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l17
			s4f32 = 2.1474835e+09
			if s3f32 < s4f32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1f32 = s1f32
			} else {
				s1f32 = s2f32
			}
			l17 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l17
			s4f32 = -2.1474835e+09
			if s3f32 > s4f32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1f32 = s1f32
			} else {
				s1f32 = s2f32
			}
			l17 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l17
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl26
			}
			s1i32 = -2147483648
		lbl26:
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+18)])) = uint16(s1i32)
			s0i32 = l1
			s1f32 = l21
			s2f32 = 255
			s1f32 = s1f32 * s2f32
			s2f32 = 0.5
			s1f32 = s1f32 + s2f32
			s1f32 = float32(math.Floor(float64(s1f32)))
			l17 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l17
			s4f32 = 2.1474835e+09
			if s3f32 < s4f32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1f32 = s1f32
			} else {
				s1f32 = s2f32
			}
			l17 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l17
			s4f32 = -2.1474835e+09
			if s3f32 > s4f32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1f32 = s1f32
			} else {
				s1f32 = s2f32
			}
			l17 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l17
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl28
			}
			s1i32 = -2147483648
		lbl28:
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint16(s1i32)
			s0i32 = l1
			s1f32 = l19
			s2f32 = 255
			s1f32 = s1f32 * s2f32
			s2f32 = 0.5
			s1f32 = s1f32 + s2f32
			s1f32 = float32(math.Floor(float64(s1f32)))
			l17 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l17
			s4f32 = 2.1474835e+09
			if s3f32 < s4f32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1f32 = s1f32
			} else {
				s1f32 = s2f32
			}
			l17 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l17
			s4f32 = -2.1474835e+09
			if s3f32 > s4f32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1f32 = s1f32
			} else {
				s1f32 = s2f32
			}
			l17 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l17
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl30
			}
			s1i32 = -2147483648
		lbl30:
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint16(s1i32)
		}
		s0i32 = l8
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
		s0i32 = l8
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
		s0i32 = l8
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
		s0i32 = l8
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
		s0i32 = l8
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
		s0i32 = l8
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = l7
		s3i32 = 4
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l6 = s2i32
		s1i32 = s1i32 + s2i32
		s0i32 = f283(ctx, s0i32, s1i32)
		s0i32 = l8
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		l11 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1f32 = -s1f32
		s2i32 = l11
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s2f32 = -s2f32
		f200(ctx, s0i32, s1f32, s2f32)
		s0i32 = l8
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		f116(ctx, s0i32, s1i32)
		s0i32 = l16
		s1i32 = l8
		s2i32 = 128
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		s3i32 = l16
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
		if int(s3i32) < 0 || int(s3i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s3i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s3i32].numParams != 3 {
			panic("argument count mismatch")
		}
		s0i32 = (*(*func(*Context, int32, int32, int32) int32)(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l13 = s0i32
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+164)]))
		l6 = s0i32
		s1i32 = 128
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l8
			s1i32 = l8
			s2i32 = 128
			s1i32 = s1i32 + s2i32
			s1i32 = f24(ctx, s1i32)
			l6 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+164)])) = uint32(s1i32)
		}
		s0i32 = l6
		s1i32 = 16
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l8
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = l8
			s2i32 = 640
			s1i32 = s1i32 + s2i32
			s2i32 = l11
			s0i32 = f42(ctx, s0i32, s1i32, s2i32)
			s0i32 = l8
			s1i32 = 640
			s0i32 = s0i32 + s1i32
			s1i32 = l13
			s2i32 = l12
			f219(ctx, s0i32, s1i32, s2i32)
			goto lbl33
		}
		s0i32 = l11
		s1i32 = l8
		s2i32 = 640
		s1i32 = s1i32 + s2i32
		f269(ctx, s0i32, s1i32)
		s0i32 = l8
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1i32 = l8
		s2i32 = 640
		s1i32 = s1i32 + s2i32
		s2i32 = l8
		s3i32 = 640
		s2i32 = s2i32 + s3i32
		s3i32 = 4
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l15
		f278(ctx, s0i32)
		s0i32 = l15
		s1i32 = l8
		s2i32 = 640
		s1i32 = s1i32 + s2i32
		s2i32 = 4
		s3i32 = 1
		f277(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l15
		s1i32 = l13
		s2i32 = l12
		f417(ctx, s0i32, s1i32, s2i32)
	lbl33:
		s0i32 = l7
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l5
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl22
		}
	}
	s0i32 = l15
	f34(ctx, s0i32)
	goto lbl3
lbl4:
	s0i32 = l8
	s1i32 = 21300
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l1
	f12(ctx, s0i32)
lbl3:
	s0i32 = l14
	f43(ctx, s0i32)
	s0i32 = l10
	s0i32 = f23(ctx, s0i32)
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+632)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = l1
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
lbl0:
	s0i32 = l8
	s1i32 = 672
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
