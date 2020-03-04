package internal

import (
	"unsafe"
)

func f228(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
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
	var s1i64 int64
	_ = s1i64
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+180)]))
	l1 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	s2i32 = -8191
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+180)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+184)]))
		s2i32 = 8191
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint32(s1i32)
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8191
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+180)])) = uint32(s1i32)
lbl0:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)]))
	l3 = s0i32
	s0i32 = l0
	s1i32 = l1
	s2i32 = l4
	if s1i32 >= s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l3
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
		s3i32 = -8191
		s2i32 = s2i32 + s3i32
		if s1i32 >= s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
	} else {
		s1i32 = 0
	}
	ctx.Mem[int(s0i32+188)] = uint8(s1i32)
	s0i32 = l2
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l3
	s2i32 = 2147475456
	s3i32 = l3
	s4i32 = 2147475456
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = 8191
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l1
	s2i32 = 2147475456
	s3i32 = l1
	s4i32 = 2147475456
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = 8191
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 52
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s3i32 = 32
	s2i32 = s2i32 + s3i32
	s0i32 = f1938(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 164
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 1651
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 1603
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = 1577
		s1i32 = l2
		f19(ctx, s0i32, s1i32)
		panic("unreachable executed")
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+132)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+148)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 92
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+180)]))
	s1i32 = s1i32 - s2i32
	s1f32 = float32(s1i32)
	s2i32 = 0
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+184)]))
	s2i32 = s2i32 - s3i32
	s2f32 = float32(s2i32)
	f142(ctx, s0i32, s1f32, s2f32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 1064
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+180)]))
	s1i32 = s1i32 - s2i32
	s2i32 = 0
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+184)]))
	s2i32 = s2i32 - s3i32
	s3i32 = l0
	s4i32 = 132
	s3i32 = s3i32 + s4i32
	l1 = s3i32
	f1931(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l2
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l2
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	f273(ctx, s0i32, s1i32)
	s0i32 = l2
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
