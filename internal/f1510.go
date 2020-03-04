package internal

import (
	"unsafe"
)

func f1510(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
	s1i32 = 2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 28
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l2
	s2i32 = 56
	s1i32 = s1i32 + s2i32
	s2i32 = 0
	s0i32 = f347(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	s1f32 = 1
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
		s1f32 = 1
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l6 = s0i64
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l3 = s0i32
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l6
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s0i32 = f1955(ctx, s0i32)
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)]))
	l4 = s0i32
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = uint32(s1i32)
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 0
		f91(ctx, s0i32, s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)]))
		l3 = s0i32
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l4 = s0i32
		s0i32 = l1
		s0i32 = f1953(ctx, s0i32)
		l3 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)]))
		l1 = s0i32
		s0i32 = l0
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = uint32(s1i32)
		s0i32 = l1
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = 0
			f91(ctx, s0i32, s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)]))
			l3 = s0i32
		}
		s0i32 = l3
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1f32 = 1
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+60)]))
	s1f32 = s1f32 / s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = s1f32
	s0i32 = l2
	s1f32 = 1
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
	s1f32 = s1f32 / s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = s1f32
	s0i32 = 0
	l1 = s0i32
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l2
	s2i32 = 48
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s0i32 = f1961(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+40)]))
		f475(ctx, s0i32, s1f32, s2f32)
		s0i32 = l0
		s1i32 = 72
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		l0 = s1i32
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
		s4i32 = 0
		s5i32 = 0
		s0i32 = f157(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		l4 = s0i32
		goto lbl8
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)]))
	l4 = s0i32
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = uint32(s1i32)
	s0i32 = 1
	l1 = s0i32
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 0
		f91(ctx, s0i32, s1i32)
	}
lbl8:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l3 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l0
	f12(ctx, s0i32)
lbl11:
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl2:
	s0i32 = 0
	l4 = s0i32
lbl1:
	s0i32 = l2
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	return s0i32
lbl0:
	f1509(ctx)
	panic("unreachable executed")
}
