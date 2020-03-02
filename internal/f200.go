package internal

import (
	"unsafe"
)

func f200(ctx *Context, l0 int32, l1 float32, l2 float32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 float32
	_ = l5
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
	var s2i64 int64
	_ = s2i64
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l4 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = f24(ctx, s1i32)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1i32 = l4
	s2i32 = 14
	s1i32 = s1i32 & s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l0
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3f32 = l1
		s2f32 = s2f32 + s3f32
		l5 = s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = s2f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2f32 = l2
		s1f32 = s1f32 + s2f32
		goto lbl2
	}
	s1i32 = l4
	s2i32 = 8
	s1i32 = s1i32 & s2i32
	if s1i32 != 0 {
		s1i32 = l3
		s2i32 = 1065353216
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)])) = uint32(s2i32)
		s1i32 = l3
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)])) = uint64(s2i64)
		s1i32 = l3
		s2i64 = 4575657221408423936
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)])) = uint64(s2i64)
		s1i32 = l3
		s2i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
		s1i32 = l3
		s2f32 = l2
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)])) = s2f32
		s1i32 = l3
		s2f32 = l1
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = s2f32
		s1i32 = l3
		s2i32 = 17
		s3i32 = 17
		s4i32 = 16
		s5f32 = l2
		s6f32 = 0
		if s5f32 != s6f32 {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		s4f32 = l1
		s5f32 = 0
		if s4f32 != s5f32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		l4 = s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)])) = uint32(s2i32)
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 & s2i32
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl1
		}
		s1i32 = l0
		s2i32 = l0
		s3i32 = l3
		s4i32 = 8
		s3i32 = s3i32 + s4i32
		f69(ctx, s1i32, s2i32, s3i32)
		goto lbl1
	}
	s1i32 = l0
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4f32 = l1
	s3f32 = s3f32 * s4f32
	s4i32 = l0
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s5f32 = l2
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 + s3f32
	l5 = s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = s2f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s3f32 = l1
	s2f32 = s2f32 * s3f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
lbl2:
	l1 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l0
	s1i32 = l4
	s2i32 = 1
	s1i32 = s1i32 | s2i32
	l0 = s1i32
	s2i32 = l0
	s3i32 = l4
	s4i32 = -2
	s3i32 = s3i32 & s4i32
	s4f32 = l1
	s5f32 = 0
	if s4f32 != s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3f32 = l5
	s4f32 = 0
	if s3f32 != s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
lbl1:
	s0i32 = l3
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
