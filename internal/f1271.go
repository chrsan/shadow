package internal

import (
	"unsafe"
)

func f1271(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
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
	var s2i64 int64
	_ = s2i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 320
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+288)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+132)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+296)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+304)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+148)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+312)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)])) = uint64(s1i64)
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)]))
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+284)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+288)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+292)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+296)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+300)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+304)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+308)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+312)]))
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l7
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+212)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+204)])) = uint32(s1i32)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+272)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s2i32 = 272
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+208)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+208)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s2i32 = 270
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+200)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 216
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l7
	s3i32 = 32
	s2i32 = s2i32 + s3i32
	s3i32 = l7
	s4i32 = 24
	s3i32 = s3i32 + s4i32
	s4i32 = l7
	s5i32 = 16
	s4i32 = s4i32 + s5i32
	s5i32 = l7
	s6i32 = 8
	s5i32 = s5i32 + s6i32
	s0i32 = f1088(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	l11 = s0i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l1 = s0i32
	lbl2:
		s0i32 = l7
		s1i32 = l2
		s2i32 = l1
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+270)])) = uint16(s1i32)
		s0i32 = l7
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = 144
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l1
		s3i32 = 4
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s0i32 = f283(ctx, s0i32, s1i32)
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l5
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		f142(ctx, s0i32, s1f32, s2f32)
		s0i32 = l7
		s1i32 = 96
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s0i32 = f46(ctx, s0i32, s1i32)
		l10 = s0i32
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)]))
		l9 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l7
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+180)]))
		l8 = s0i32
		s1i32 = 128
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l7
			s1i32 = l7
			s2i32 = 144
			s1i32 = s1i32 + s2i32
			s1i32 = f24(ctx, s1i32)
			l8 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+180)])) = uint32(s1i32)
		}
		s0i32 = l8
		s1i32 = 15
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l7
			s1i64 = 69784829952
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
			s0i32 = l7
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
			s0i32 = l7
			s1i64 = 1065353216
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
			s0i32 = l7
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
			s0i32 = l7
			s1i64 = 1065353216
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
			goto lbl6
		}
		s0i32 = l7
		s1i32 = 144
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s2i32 = 56
		s1i32 = s1i32 + s2i32
		s0i32 = f84(ctx, s0i32, s1i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
	lbl6:
		s0i32 = l7
		s1i32 = 48
		s0i32 = s0i32 + s1i32
		s1i32 = l9
		s2i32 = l7
		s3i32 = 56
		s2i32 = s2i32 + s3i32
		f1451(ctx, s0i32, s1i32, s2i32)
		s0i32 = l10
		s1i32 = l7
		s2i32 = 48
		s1i32 = s1i32 + s2i32
		f83(ctx, s0i32, s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l8 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l8
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l9 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l9
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l8
		s1i32 = l8
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
		goto lbl3
	lbl5:
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = l10
		s1i32 = l7
		s2i32 = 40
		s1i32 = s1i32 + s2i32
		f83(ctx, s0i32, s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l8 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l8
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l9 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l9
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l8
		s1i32 = l8
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
	lbl3:
		s0i32 = l7
		s1i32 = 144
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s2i32 = 280
		s1i32 = s1i32 + s2i32
		f116(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+144)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+176)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+168)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+160)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+152)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s2i32 = 56
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
		s1i32 = l8
		s2i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])) = uint32(s2i32)
		s1i32 = l8
		s2i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
		s1i32 = l8
		s2i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l8
		s2i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
		s1i32 = l8
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+144)]))
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
		s0i32 = l10
		s0i32 = f23(ctx, s0i32)
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l4
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+280)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+312)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+304)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+296)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+288)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
	s0i32 = l11
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
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
		goto lbl0
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
lbl0:
	s0i32 = l7
	s1i32 = 320
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
