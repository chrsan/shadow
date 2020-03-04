package internal

import (
	"unsafe"
)

func f1395(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
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
	var s3f32 float32
	_ = s3f32
	s0i32 = ctx.g0
	s1i32 = 96
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l4 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = 2
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = l4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l0
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
		s2i32 = l4
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = f29(ctx, s1i32, s2i32)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	}
	s0i32 = l3
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 1
	s0i32 = f329(ctx, s0i32, s1i32, s2i32)
	l6 = s0i32
	s1i32 = l3
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s0i32 = f131(ctx, s0i32, s1i32)
	l1 = s0i32
	s1i32 = 6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 40
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l3
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = l3
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s1i32 = 8
		s0i32 = s0i32 | s1i32
		l5 = s0i32
		s0i32 = 0
		l4 = s0i32
	lbl4:
		s0i32 = l4
		if s0i32 != 0 {
			s0i32 = 0
			l4 = s0i32
			goto lbl2
		}
		s0i32 = 0
		l4 = s0i32
		s0i32 = l1
		s1i32 = 5
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl12
		case 1:
			goto lbl11
		case 2:
			goto lbl9
		case 3:
			goto lbl10
		case 4:
			goto lbl7
		default:
			goto lbl8
		}
	lbl12:
		s0i32 = l2
		s1i32 = l5
		s2i32 = l5
		s3i32 = 1
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l0
		s1i32 = l5
		f103(ctx, s0i32, s1i32)
		goto lbl6
	lbl11:
		s0i32 = l2
		s1i32 = l3
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		s2i32 = l3
		s3i32 = 16
		s2i32 = s2i32 + s3i32
		s3i32 = 3
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l0
		s1i32 = l5
		f103(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = l7
		f103(ctx, s0i32, s1i32)
		goto lbl6
	lbl10:
		s0i32 = l2
		s1i32 = l3
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		s2i32 = l3
		s3i32 = 16
		s2i32 = s2i32 + s3i32
		s3i32 = 4
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l0
		s1i32 = l5
		f103(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = l7
		f103(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = l8
		f103(ctx, s0i32, s1i32)
		goto lbl6
	lbl9:
		s0i32 = l0
		s1i32 = l2
		s2i32 = l3
		s3i32 = 16
		s2i32 = s2i32 + s3i32
		s3i32 = l6
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		f552(ctx, s0i32, s1i32, s2i32, s3f32)
		goto lbl6
	lbl8:
		s0i32 = l9
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		goto lbl2
	lbl7:
		s0i32 = 1
		l4 = s0i32
	lbl6:
		s0i32 = 1
		l9 = s0i32
		s0i32 = l6
		s1i32 = l3
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		s0i32 = f131(ctx, s0i32, s1i32)
		l1 = s0i32
		s1i32 = 6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
	}
	s0i32 = l0
	f559(ctx, s0i32)
	s0i32 = 1
	l4 = s0i32
lbl2:
	s0i32 = l3
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	return s0i32
lbl0:
	s0i32 = l3
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 14791
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 14743
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 14717
	s1i32 = l3
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
