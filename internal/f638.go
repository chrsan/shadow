package internal

import (
	"unsafe"
)

func f638(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = 19200
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1212
	s0i32 = s0i32 + s1i32
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 1200
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		f13(ctx, s0i32)
	}
	s0i32 = l0
	s1i32 = 1188
	s0i32 = s0i32 + s1i32
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 1176
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		f13(ctx, s0i32)
	}
	s0i32 = l0
	s1i32 = 1132
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l1
		f13(ctx, s0i32)
	}
	s0i32 = l0
	s1i32 = 1128
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l1
		f13(ctx, s0i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1084)]))
	l1 = s0i32
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1084)])) = uint32(s1i32)
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = f41(ctx, s0i32)
		f12(ctx, s0i32)
	}
	s0i32 = l0
	s1i32 = 1032
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l0
	s1i32 = 1056
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 != 0 {
	lbl6:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1036)]))
		f112(ctx, s0i32)
		s0i32 = l1
		f367(ctx, s0i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1056)]))
		if s0i32 != 0 {
			goto lbl6
		}
	}
	s0i32 = l1
	f522(ctx, s0i32)
	s0i32 = l0
	s1i32 = 156
	s0i32 = s0i32 + s1i32
	s0i32 = f41(ctx, s0i32)
	s0i32 = l0
	s1i32 = 19780
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
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
		goto lbl7
	}
	s0i32 = l1
	f12(ctx, s0i32)
lbl7:
	s0i32 = l0
	return s0i32
}
