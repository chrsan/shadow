package internal

import (
	"unsafe"
)

func f1518(ctx *Context, l0 int32, l1 int32) int32 {
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
	var l8 int64
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
	var l16 float32
	_ = l16
	var l17 float32
	_ = l17
	var l18 float32
	_ = l18
	var l19 float32
	_ = l19
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0i64
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l9 = s0i64
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l10 = s0i64
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l11 = s0i64
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l12 = s0i64
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l13 = s0i64
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	l14 = s0i64
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l15 = s0i64
	s0i32 = l2
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s1i64 = l14
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l13
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l12
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l11
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l10
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l9
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l8
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l15
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l3 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s1i32 = f24(ctx, s1i32)
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	}
	s0i32 = l3
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		l4 = s1i32
		if s1i32 != 0 {
			s1i32 = l2
			s2i32 = 72
			s1i32 = s1i32 + s2i32
			l3 = s1i32
			s2i32 = l1
			s3i32 = l4
			f69(ctx, s1i32, s2i32, s3i32)
			s1i32 = l2
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+108)]))
			l1 = s1i32
			s2i32 = 128
			s1i32 = s1i32 & s2i32
			if s1i32 != 0 {
				s1i32 = l2
				s2i32 = l3
				s2i32 = f24(ctx, s2i32)
				l1 = s2i32
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+108)])) = uint32(s2i32)
			}
			s1i32 = l3
			s2i32 = 0
			s3i32 = l1
			s4i32 = 15
			s3i32 = s3i32 & s4i32
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
		} else {
			s1i32 = l1
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	l4 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l5 = s0i32
	s0i32 = 0
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l3 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l6 = s0i32
	s1i32 = 256
	s0i32 = s0i32 | s1i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l1
	s1i32 = s1i32 - s2i32
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 256
		s2i32 = 4
		f18(ctx, s0i32, s1i32, s2i32)
		s0i32 = 0
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l1 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l6 = s0i32
	}
	s0i32 = l3
	s1i32 = l1
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s2i32 = 256
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 0
	s2i32 = 256
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	l6 = s0i32
	s0i32 = l5
	if s0i32 != 0 {
		s0i32 = 0
		l3 = s0i32
		s0i32 = l5
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
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
			goto lbl6
		}
		goto lbl5
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l5 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l7 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l3 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l17 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l18 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l19 = s0f32
	s0i32 = l2
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	l16 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = s1f32
	s0i32 = l2
	s1f32 = l18
	s2f32 = l16
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = s1f32
	s0i32 = l2
	s1f32 = l17
	s2f32 = l16
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = s1f32
	s0i32 = l2
	s1f32 = l19
	s2f32 = l16
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = s1f32
	s0i32 = l7
	s1i32 = l5
	s2i32 = l2
	s3i32 = 112
	s2i32 = s2i32 + s3i32
	f174(ctx, s0i32, s1i32, s2i32)
lbl6:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = 87
	s2i32 = l6
	f16(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = l4
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
		l4 = s0i32
		s0i32 = 0
		l3 = s0i32
		s0i32 = l1
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l4
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		goto lbl8
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l4 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l5 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l3 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l17 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l18 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l19 = s0f32
	s0i32 = l2
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	l16 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = s1f32
	s0i32 = l2
	s1f32 = l18
	s2f32 = l16
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = s1f32
	s0i32 = l2
	s1f32 = l17
	s2f32 = l16
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = s1f32
	s0i32 = l2
	s1f32 = l19
	s2f32 = l16
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = s1f32
	s0i32 = l5
	s1i32 = l4
	s2i32 = l2
	s3i32 = 112
	s2i32 = s2i32 + s3i32
	f174(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = 0
	l3 = s0i32
	goto lbl5
lbl8:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = 88
	s2i32 = l6
	f16(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	f99(ctx, s0i32, s1i32)
	s0i32 = 1
	l3 = s0i32
lbl5:
	s0i32 = l2
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l3
	return s0i32
}
