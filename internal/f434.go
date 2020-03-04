package internal

import (
	"unsafe"
)

func f434(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int64
	_ = l9
	var l10 int64
	_ = l10
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 112
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l6
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])))
	s1i32 = l3
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
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
	s0i32 = l3
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])))
	s1i32 = l3
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])))
	s0i64 = s0i64 - s1i64
	l10 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = l9
	s1i64 = l10
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
	s0i32 = l6
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	s2i32 = 96
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
lbl0:
	s0i32 = l6
	s1i32 = 72
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	l7 = s0i32
	s0i32 = l6
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	l3 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l8 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l2
		s1i32 = f24(ctx, s1i32)
		l8 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l8
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l1
		f334(ctx, s0i32, s1i32)
		goto lbl2
	}
	s0i32 = l1
	s1i32 = l2
	s2i32 = l3
	f138(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s1i32 = l3
	s1i32 = int32(ctx.Mem[int(s1i32+10)])
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+10)] = uint8(s1i32)
lbl2:
	s0i32 = l4
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+42)])
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l3
			s2i32 = l0
			s3i32 = l5
			f326(ctx, s0i32, s1i32, s2i32, s3i32)
			goto lbl4
		}
		s0i32 = l7
		s1i32 = l0
		s2i32 = l0
		s3i32 = 20
		s2i32 = s2i32 + s3i32
		s3i32 = l0
		s3i32 = int32(ctx.Mem[int(s3i32+40)])
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s0i32 = f81(ctx, s0i32, s1i32)
		s0i32 = l6
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = -1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l1
		l2 = s0i32
		s0i32 = l6
		s1i32 = 28
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 0
		ctx.Mem[int(s0i32+50)] = uint8(s1i32)
		s0i32 = l6
		s1i32 = 257
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint16(s1i32)
		s0i32 = l6
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l7
		s3i32 = l5
		f326(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l0
		s1i32 = l6
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = 1
		f664(ctx, s0i32, s1i32, s2i32)
		s0i32 = l1
		f246(ctx, s0i32)
		s0i32 = l2
		f40(ctx, s0i32)
		goto lbl4
	}
	s0i32 = l7
	s1i32 = l6
	s2i32 = 96
	s1i32 = s1i32 + s2i32
	s0i32 = f81(ctx, s0i32, s1i32)
	s0i32 = l4
	s1i32 = 5
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l3
		s2i32 = l7
		s3i32 = l5
		f326(ctx, s0i32, s1i32, s2i32, s3i32)
		goto lbl4
	}
	s0i32 = l6
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l1
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l1
	l2 = s0i32
	s0i32 = l6
	s1i32 = 28
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l1
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = 0
	ctx.Mem[int(s0i32+50)] = uint8(s1i32)
	s0i32 = l6
	s1i32 = 257
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint16(s1i32)
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = l7
	s3i32 = l5
	f326(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l0
	s1i32 = l6
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	f664(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	f246(ctx, s0i32)
	s0i32 = l2
	f40(ctx, s0i32)
lbl4:
	s0i32 = l3
	f34(ctx, s0i32)
	s0i32 = l7
	f40(ctx, s0i32)
	s0i32 = l6
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
