package internal

import (
	"unsafe"
)

func f1087(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
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
	var l12 int32
	_ = l12
	var l13 int64
	_ = l13
	var l14 int64
	_ = l14
	var l15 int64
	_ = l15
	var l16 int64
	_ = l16
	var l17 float32
	_ = l17
	var l18 float32
	_ = l18
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 288
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l11 = s0i32
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		f1711(ctx, s0i32, s1i32)
		s0i32 = l7
		s1i32 = 24
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s2i32 = 120
		s1i32 = s1i32 + s2i32
		s0i32 = f1710(ctx, s0i32, s1i32)
		l8 = s0i32
		s0i32 = l7
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l13 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = l13
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s1i32 = l8
		s2i32 = l7
		f1709(ctx, s0i32, s1i32, s2i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l9 = s0i32
		if s0i32 != 0 {
			s0i32 = l7
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			l10 = s0i32
			s1i32 = l9
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			l12 = s0i32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l17 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l18 = s0f32
			s0i32 = l4
			l3 = s0i32
		lbl2:
			s0i32 = l10
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l9 = s0i32
			s0i32 = l3
			s1f32 = l18
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
			s0i32 = l3
			s1f32 = l17
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
			s0i32 = l3
			s1i32 = 8
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s0f32 = l18
			s1i32 = l9
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
			s0f32 = s0f32 + s1f32
			l18 = s0f32
			s0f32 = l17
			s1i32 = l9
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			s0f32 = s0f32 + s1f32
			l17 = s0f32
			s0i32 = l10
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l10 = s0i32
			s1i32 = l12
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl2
			}
			s0i32 = l2
			s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l13 = s0i64
		}
		s0i32 = l6
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l14 = s0i64
		s0i32 = l5
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l15 = s0i64
		s0i32 = l7
		s1i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+268)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+264)])) = uint32(s1i32)
		s0i32 = l7
		s1i64 = l15
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+272)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = l14
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = l13
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+256)])) = uint64(s1i64)
		s0i64 = l13
		s1i64 = 32
		s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
		s0i32 = int32(s0i64)
		l3 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l2 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l7
			s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+264)]))
			l16 = s0i64
			s0i32 = l2
			s1i64 = l14
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
			s0i32 = l2
			s1i64 = l15
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = l3
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
			s0i32 = l2
			s1i64 = l16
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i64)
			s0i32 = l2
			s1i64 = l13
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i64)
			s0i32 = l1
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l3 = s0i32
			if s0i32 != 0 {
				s0i32 = l3
				s1i32 = l3
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
				s2i32 = 1
				s1i32 = s1i32 + s2i32
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			}
			s0i32 = l2
			s1i32 = l3
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
			s0i32 = l2
			s1i32 = l1
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = l1
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+11)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+39)])) = uint64(s1i64)
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
			s2i32 = 48
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
			goto lbl3
		}
		s0i32 = l0
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l7
		s3i32 = 264
		s2i32 = s2i32 + s3i32
		s3i32 = l7
		s4i32 = 256
		s3i32 = s3i32 + s4i32
		s4i32 = l7
		s5i32 = 272
		s4i32 = s4i32 + s5i32
		s5i32 = l7
		s6i32 = 280
		s5i32 = s5i32 + s6i32
		f354(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	lbl3:
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
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
			goto lbl6
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
	lbl6:
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l0 = s0i32
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l1 = s1i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l3 = s0i32
	lbl7:
		s0i32 = l3
		s1i32 = l0
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l2 = s0i32
		s0i32 = l3
		s1i32 = -4
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l2
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l1
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = 21
			if s0i32 >= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l0
				f13(ctx, s0i32)
			}
			s0i32 = l8
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		}
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+248)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
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
			goto lbl10
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
	lbl10:
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+244)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl11
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
			goto lbl11
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
	lbl11:
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+240)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl12
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
			goto lbl12
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
	lbl12:
		s0i32 = l7
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		f159(ctx, s0i32)
	}
	s0i32 = l7
	s1i32 = 288
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
