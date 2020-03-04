package internal

import (
	"unsafe"
)

func f1896(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l4 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l0
		s2i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)])) = uint32(s2i32)
		s1i32 = l0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		l4 = s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint32(s2i32)
		s1i32 = l4
		s2i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
		s1i32 = l4
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		goto lbl1
	}
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l6 = s1i32
	s2i32 = l2
	if s1i32 >= s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		l2 = s1i32
		goto lbl0
	}
	s1i32 = l4
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	s3i32 = l4
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	l10 = s3i32
	s2i32 = s2i32 - s3i32
	l7 = s2i32
	s3i32 = 2
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	l8 = s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l5 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl6
	}
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	s2i32 = l6
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl6
	}
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l8
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl6
	}
	s1i32 = l7
	s2i32 = 0
	if s1i32 <= s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl5
	}
	s1i32 = l5
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	l7 = s1i32
lbl7:
	s1i32 = l7
	s2i32 = l9
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	l11 = s2i32
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l10
	s3i32 = l11
	s2i32 = s2i32 + s3i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl6
	}
	s1i32 = l8
	s2i32 = l9
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	l9 = s2i32
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl7
	}
	goto lbl5
lbl6:
	s1i32 = l0
	s2i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l10
	s3i32 = l8
	s4i32 = 2
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s3i32 = 4
	s2i32 = s2i32 + s3i32
	l4 = s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint32(s2i32)
	goto lbl4
lbl5:
	s1i32 = l5
	s2i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
lbl4:
	s1i32 = l2
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s2i32 = l6
	if s1i32 > s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l4
		s2i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l4
		s2i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
		s1i32 = l0
		s2i32 = l4
		s3i32 = 12
		s2i32 = s2i32 + s3i32
		l4 = s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint32(s2i32)
	}
	s1i32 = l4
	s2i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
	s1i32 = l4
	s2i32 = 8
	s1i32 = s1i32 + s2i32
lbl1:
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
lbl0:
	s0i32 = l2
	s1i32 = l4
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l2
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l4
	s1i32 = l1
	s2i32 = l3
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	return
lbl9:
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l1
	s2i32 = l3
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l2
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
}
