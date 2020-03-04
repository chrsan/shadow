package internal

import (
	"unsafe"
)

func f596(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = 0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	s2i32 = 3
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l2 = s0i32
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl3
	case 1:
		goto lbl2
	case 2:
		goto lbl3
	default:
		goto lbl4
	}
lbl4:
	s0i32 = l1
	s1i32 = l2
	s2i32 = 8
	s3i32 = l2
	s4i32 = 11
	s5i32 = -1
	s6i32 = -1
	s7i32 = -1
	s8i32 = l2
	s9i32 = 0
	s8i32 = f191(ctx, s8i32, s9i32)
	s9i32 = 0
	s3i32 = f21(ctx, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32, s9i32)
	s1i32 = f96(ctx, s1i32, s2i32, s3i32)
	l0 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	goto lbl1
lbl3:
	s0i32 = l1
	s1i32 = l2
	s2i32 = 8
	s3i32 = l2
	s4i32 = l2
	s5i32 = 1
	s4i32 = f191(ctx, s4i32, s5i32)
	s3i32 = f397(ctx, s3i32, s4i32)
	s1i32 = f96(ctx, s1i32, s2i32, s3i32)
	l0 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 1
	goto lbl0
lbl2:
	s0i32 = l3
	s1i32 = l2
	s2i32 = l2
	s3i32 = l2
	s4i32 = 2
	s3i32 = f191(ctx, s3i32, s4i32)
	s2i32 = f607(ctx, s2i32, s3i32)
	f396(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l2
	s2i32 = l2
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s2i32 = f1650(ctx, s2i32, s3i32, s4i32)
	s3i32 = l2
	s4i32 = l1
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s5i32 = l2
	s6i32 = l1
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
	s7i32 = l1
	s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+8)]))
	s5i32 = f53(ctx, s5i32, s6i32, s7i32)
	s3i32 = f53(ctx, s3i32, s4i32, s5i32)
	s4i32 = l2
	s5i32 = l1
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s6i32 = l2
	s7i32 = l1
	s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+4)]))
	s8i32 = l1
	s8i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s8i32+8)]))
	s6i32 = f68(ctx, s6i32, s7i32, s8i32)
	s4i32 = f68(ctx, s4i32, s5i32, s6i32)
	s1i32 = f1648(ctx, s1i32, s2i32, s3i32, s4i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
lbl1:
	s0i32 = 1
lbl0:
	l0 = s0i32
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
