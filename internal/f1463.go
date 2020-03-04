package internal

import (
	"unsafe"
)

func f1463(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
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
	var l9 float32
	_ = l9
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l4 = s0i32
	s0i32 = 1
	l2 = s0i32
	s0i32 = 1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l5 = s1i32
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 1941503
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i32
		s0i32 = 0
		l2 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		goto lbl0
	}
	s0i32 = l5
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l2 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	f662(ctx, s0i32, s1i32, s2i32)
	s0i32 = 0
	l2 = s0i32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l9
	s1f32 = 1
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l9 = s0f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l9
	s1f32 = 1
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l9 = s0f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l9
	s1f32 = 1
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l2 = s0i32
lbl3:
	s0i32 = f105(ctx)
	l6 = s0i32
	s0i32 = 3
lbl0:
	l5 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 4
		s2i32 = 0
		f16(ctx, s0i32, s1i32, s2i32)
		s0i32 = 5
		l2 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l5
		s2i32 = 3
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = 5
			s2i32 = 6
			s3i32 = l4
			s3i32 = int32(ctx.Mem[int(s3i32+60)])
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
		} else {
			s1i32 = 5
		}
		s2i32 = 0
		f16(ctx, s0i32, s1i32, s2i32)
		s0i32 = 1
		l2 = s0i32
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l8 = s0i32
	s0i32 = 0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l4 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l7 = s0i32
	s1i32 = 100
	s0i32 = s0i32 | s1i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l4
	s1i32 = s1i32 - s2i32
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 100
		s2i32 = 4
		f18(ctx, s0i32, s1i32, s2i32)
		s0i32 = 0
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l4 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l7 = s0i32
	}
	s0i32 = l1
	s1i32 = l4
	s2i32 = l7
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s2i32 = 100
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l6
	s2i32 = l5
	s3i32 = l8
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = 2
	s0i32 = f95(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l2
	f296(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
