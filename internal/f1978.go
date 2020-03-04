package internal

import (
	"unsafe"
)

func f1978(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s9i32 int32
	_ = s9i32
	var s10i32 int32
	_ = s10i32
	var s11i32 int32
	_ = s11i32
	var s12i32 int32
	_ = s12i32
	var s13i32 int32
	_ = s13i32
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = 21300
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
	s0i32 = l6
	l7 = s0i32
	s0i32 = l6
	s1i32 = 19456
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l6
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	l8 = s1i32
	s0i32 = f206(ctx, s0i32, s1i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = l0
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		s3i32 = 0
		f136(ctx, s0i32, s1i32, s2i32, s3i32)
	}
	s0i32 = l6
	s1i32 = l0
	s2i32 = 116
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l0
	s2i32 = 1064
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1084)]))
	l0 = s1i32
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s2i32 = 0
	s3i32 = l0
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+48)]))
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s4i32 = l1
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s5i32 = l1
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
	s6i32 = l1
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+20)]))
	s7i32 = l1
	s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+24)]))
	s8i32 = l4
	s9i32 = l1
	s9i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s9i32+28)]))
	s10i32 = l1
	s10i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s10i32+52)]))
	s11i32 = l5
	s12i32 = l2
	s13i32 = l3
	f1111(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32, s9i32, s10i32, s11i32, s12i32, s13i32)
	s0i32 = l6
	s1i32 = 21300
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
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
		goto lbl1
	}
	s0i32 = l0
	f12(ctx, s0i32)
lbl1:
	s0i32 = l6
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
