package internal

import (
	"unsafe"
)

func f49(ctx *Context, l0 int32, l1 int32, l2 float32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s1i32 = 5
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = 1
	l3 = s0i32
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl5
	case 1:
		goto lbl4
	case 2:
		goto lbl3
	case 3:
		goto lbl2
	case 4:
		goto lbl1
	default:
		goto lbl0
	}
lbl5:
	s0i32 = 1
	l5 = s0i32
	s0i32 = l1
	l3 = s0i32
	goto lbl0
lbl4:
	s0i32 = 2
	l5 = s0i32
	s0i32 = l1
	l3 = s0i32
	goto lbl0
lbl3:
	s0i32 = 2
	l3 = s0i32
	s0i32 = 4
	l5 = s0i32
	goto lbl0
lbl2:
	s0i32 = 3
	l3 = s0i32
	s0i32 = 8
	l5 = s0i32
	goto lbl0
lbl1:
	s0i32 = 0
	l3 = s0i32
lbl0:
	s0i32 = l0
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+86)])) = uint16(s1i32)
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+90)])
	s2i32 = l5
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+90)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s0i32 = f2074(ctx, s0i32)
	s1i32 = l1
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = 3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 44
		s0i32 = s0i32 + s1i32
		s0i32 = f2073(ctx, s0i32)
		s1f32 = l2
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l5 = s0i32
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l5
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		s1i32 = l1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s1i32 = 2
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			s1i32 = l3
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s1i32 = -1
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl7
			}
			s0i32 = l0
			s1i32 = l3
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
			s2i32 = l3
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = f29(ctx, s1i32, s2i32)
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		}
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l0 = s0i32
	s0i32 = l4
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l5
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	return s0i32
lbl8:
	s0i32 = l4
	s1i32 = 345
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 4503
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 4423
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = 4397
	s1i32 = l4
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl7:
	s0i32 = l4
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 4471
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 4423
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 4397
	s1i32 = l4
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
