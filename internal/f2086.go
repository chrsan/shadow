package internal

import (
	"unsafe"
)

func f2086(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	s0i32 = ctx.g0
	s1i32 = 192
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = l0
	s2i32 = 0
	s0i32 = f229(ctx, s0i32, s1i32, s2i32)
	l0 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+188)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	l7 = s0i32
lbl1:
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+189)])
	if s0i32 != 0 {
	lbl4:
		s0i32 = l0
		f228(ctx, s0i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+173)])
		l6 = s0i32
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+188)])
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l6
		if s0i32 != 0 {
			goto lbl4
		}
		goto lbl2
	lbl5:
		s0i32 = l6
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
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+188)] = uint8(s1i32)
lbl2:
	s0i32 = l7
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l4
	f506(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l0
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
	s0i32 = l0
	s1i32 = 132
	s0i32 = s0i32 + s1i32
	f112(ctx, s0i32)
	s0i32 = l0
	s1i32 = 21300
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l2 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
	f12(ctx, s0i32)
lbl6:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
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
	s0i32 = l5
	s1i32 = 192
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
