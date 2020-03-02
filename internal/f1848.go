package internal

import (
	"unsafe"
)

func f1848(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l8 int64
	_ = l8
	var l9 int64
	_ = l9
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 1136
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l4 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l5 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l8 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l6 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l7 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l9 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = l8
	s1i64 = l9
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+40)])
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l6
		s2i32 = 32768
		s1i32 = s1i32 + s2i32
		s2i32 = 16
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l4
		s2i32 = 32768
		s1i32 = s1i32 + s2i32
		s2i32 = 16
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l7
		s2i32 = 32768
		s1i32 = s1i32 + s2i32
		s2i32 = 16
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l5
		s2i32 = 32768
		s1i32 = s1i32 + s2i32
		s2i32 = 16
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l2
		f220(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l2
	s0i32 = f151(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1104)]))
	l2 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1100)]))
	l4 = s0i32
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 32768
	s1i32 = s1i32 + s2i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1120)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = 32768
	s1i32 = s1i32 + s2i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1124)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = 32768
	s1i32 = s1i32 + s2i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1128)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = 32768
	s1i32 = s1i32 + s2i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1132)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 1120
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = l2
	f220(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s0i32 = f74(ctx, s0i32)
	s0i32 = l1
	f40(ctx, s0i32)
lbl0:
	s0i32 = l3
	s1i32 = 1136
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
