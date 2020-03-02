package internal

import (
	"unsafe"
)

func f1099(ctx *Context, l0 int32, l1 int32) {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l5 = s0i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 65535
	s0i32 = s0i32 & s1i32
	s1i32 = 8191
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+24)])
		l4 = s0i32
		s0i32 = 0
		goto lbl0
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+24)])
	l4 = s0i32
	s1i32 = 5
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l2 = s0i32
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl3
		case 1:
			goto lbl3
		case 2:
			goto lbl4
		case 3:
			goto lbl6
		case 4:
			goto lbl3
		default:
			goto lbl7
		}
	lbl7:
		s0i32 = l5
		s1i32 = 7
		s0i32 = s0i32 + s1i32
		s1i32 = 3
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		goto lbl2
	lbl6:
		s0i32 = 2
		l2 = s0i32
		goto lbl3
	}
	s0i32 = l3
	s1i32 = 60
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 3323
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 3282
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = 3256
	s1i32 = l3
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl4:
	s0i32 = 4
	l2 = s0i32
lbl3:
	s0i32 = l2
	s1i32 = l5
	s0i32 = s0i32 * s1i32
lbl2:
	l2 = s0i32
	s0i32 = l2
	s1i32 = l0
	s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+2)])))
	s0i32 = s0i32 * s1i32
	l2 = s0i32
	s1i32 = 3
	s0i32 = s0i32 * s1i32
	s1i32 = l2
	s2i32 = l4
	s3i32 = 2
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
lbl0:
	l5 = s0i32
	s0i32 = l4
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s1i32 = 6
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 24
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 24
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 3344
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i32
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = 0
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l4 = s2i32
		s1i32 = s1i32 - s2i32
		s0i32 = s0i32 & s1i32
		l2 = s0i32
		s1i32 = l5
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s1i32 = l2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l1
		s1i32 = l8
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = l4
		s2i32 = s2i32 - s3i32
		if uint32(s1i32) > uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l1
			s2i32 = l5
			s3i32 = l6
			f18(ctx, s1i32, s2i32, s3i32)
			s1i32 = l7
			s2i32 = 0
			s3i32 = l1
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
			l4 = s3i32
			s2i32 = s2i32 - s3i32
			s1i32 = s1i32 & s2i32
		} else {
			s1i32 = l2
		}
		s2i32 = l4
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s2i32 = l5
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		ctx.g0 = s0i32
		return
	}
	s0i32 = l3
	s1i32 = 60
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 3323
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 3282
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 3256
	s1i32 = l3
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl8:
	f63(ctx)
	panic("unreachable executed")
}
