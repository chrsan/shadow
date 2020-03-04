package internal

import (
	"unsafe"
)

func f2102(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
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
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int64
	_ = l10
	var l11 int64
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 176
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l0
	s2i32 = 444
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 4
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+172)]))
	l12 = s0f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+164)]))
	l13 = s0f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l13
	s1f32 = -4.194304e+06
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)]))
	s1f32 = -4.194304e+06
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0f32 = l12
	s1f32 = 4.194304e+06
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)]))
	s1f32 = 4.194304e+06
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl1
	}
lbl2:
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0i64
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l11 = s0i64
	s0i32 = l0
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = l11
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = l10
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l2
	s2i32 = l3
	s3i32 = 80
	s2i32 = s2i32 + s3i32
	s3i32 = l0
	s3i32 = int32(ctx.Mem[int(s3i32+8)])
	s0i32 = f442(ctx, s0i32, s1i32, s2i32, s3i32)
	l5 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		s0i32 = 0
		l1 = s0i32
	lbl4:
		s0i32 = l3
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l10 = s0i64
		s0i32 = l3
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l11 = s0i64
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l1 = s1i32
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1i64 = l11
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i64 = l10
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		l1 = s0i32
		s1i32 = l5
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
	}
	goto lbl0
lbl1:
	s0i32 = l1
	s1i32 = l3
	s2i32 = 80
	s1i32 = s1i32 + s2i32
	s0i32 = f691(ctx, s0i32, s1i32)
	l5 = s0i32
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl5:
	s0i32 = 0
	l1 = s0i32
	s0i32 = l3
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s0i32 = f2114(ctx, s0i32, s1i32)
	l8 = s0i32
	s1i32 = -1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl7:
		s0i32 = l0
		s1i32 = l3
		s2i32 = l1
		s3i32 = 24
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		f2104(ctx, s0i32, s1i32, s2i32)
		s0i32 = l1
		s1i32 = l8
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l9 = s0i32
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l9
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
	}
	s0i32 = l4
	s1i32 = l5
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l1 = s0i32
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl0:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = 6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+444)]))
	l0 = s0i32
	s0i32 = l3
	s1i32 = 176
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = 6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	return s0i32
}
