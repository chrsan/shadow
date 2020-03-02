package internal

import (
	"unsafe"
)

func f456(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l1 = s0i32
	s1i32 = 204
	s0i32 = i32DivU(s0i32, s1i32)
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l4 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l2 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 44
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l0
	s2i32 = 44
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l1
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s2i32 = 204
	s1i32 = i32DivU(s1i32, s2i32)
	l5 = s1i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l3
	s2i32 = l5
	s3i32 = 204
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 - s2i32
	s2i32 = 20
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l2
	s2i32 = l6
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l1
	s3i32 = l6
	s4i32 = 204
	s3i32 = s3i32 * s4i32
	s2i32 = s2i32 - s3i32
	s3i32 = 20
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl2:
	s0i32 = l1
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0i32 = s0i32 - s1i32
	s1i32 = 4080
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l1 = s0i32
		s0i32 = l3
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l3 = s0i32
	}
	s0i32 = l1
	s1i32 = l5
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl0:
	s0i32 = l7
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	s0i32 = s0i32 - s1i32
	s1i32 = 2
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l1 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl5:
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		f12(ctx, s0i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l4 = s0i32
		s1i32 = l2
		s0i32 = s0i32 - s1i32
		s1i32 = 2
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l1 = s0i32
		s1i32 = 2
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
	}
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 102
		s2i32 = 204
		s3i32 = l1
		s4i32 = 1
		s3i32 = s3i32 - s4i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	}
	s0i32 = l2
	s1i32 = l4
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
lbl8:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	f12(ctx, s0i32)
	s0i32 = l2
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l4
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l2 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l1 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s1i32 = l2
	s2i32 = l2
	s3i32 = l1
	s2i32 = s2i32 - s3i32
	s3i32 = -4
	s2i32 = s2i32 + s3i32
	s3i32 = 2
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = -1
	s2i32 = s2i32 ^ s3i32
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
lbl7:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		f12(ctx, s0i32)
	}
}
