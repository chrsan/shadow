package internal

import (
	"unsafe"
)

func f2009(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
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
	s0i32 = ctx.g0
	s1i32 = 192
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i32 = l0
	s2i32 = 0
	s0i32 = f229(ctx, s0i32, s1i32, s2i32)
	l2 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+188)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = 1088
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l2
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	l3 = s0i32
lbl1:
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+189)])
	if s0i32 != 0 {
	lbl4:
		s0i32 = l2
		f228(ctx, s0i32)
		s0i32 = l2
		s0i32 = int32(ctx.Mem[int(s0i32+173)])
		l5 = s0i32
		s0i32 = l2
		s0i32 = int32(ctx.Mem[int(s0i32+188)])
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l5
		if s0i32 != 0 {
			goto lbl4
		}
		goto lbl2
	lbl5:
		s0i32 = l5
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 1
	ctx.Mem[int(s0i32+188)] = uint8(s1i32)
lbl2:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
		s3i32 = l3
		f1077(ctx, s0i32, s1i32, s2i32, s3i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+188)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl0:
	s0i32 = l2
	s1i32 = 132
	s0i32 = s0i32 + s1i32
	f112(ctx, s0i32)
	s0i32 = l2
	s1i32 = 21300
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	f12(ctx, s0i32)
lbl7:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
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
	s0i32 = l0
	f12(ctx, s0i32)
lbl8:
	s0i32 = l4
	s1i32 = 192
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
