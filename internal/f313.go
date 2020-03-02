package internal

import (
	"unsafe"
)

func f313(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
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
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
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
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0f32
	s0i32 = l3
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l7 = s1f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+60)]))
	l8 = s3f32
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l7
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s4i32 = l0
	s5i32 = -64
	s4i32 = s4i32 - s5i32
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	l7 = s4f32
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s0i32 = f327(ctx, s0i32, s1f32, s2f32)
	if s0i32 != 0 {
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0f32
		s0i32 = l3
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s1f32 = l6
		s1f32 = -s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0f32
		s0i32 = l2
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2f32 = l5
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l2
		s1f32 = l6
		s2f32 = l5
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		goto lbl1
	}
	s0i32 = 0
	s1i32 = 25464
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+88)]))
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
lbl1:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0f32
		s0i32 = l0
		s1f32 = l7
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = s1f32 + s2f32
		l6 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = s1f32
		s0i32 = l0
		s1f32 = l8
		s2f32 = l5
		s1f32 = s1f32 + s2f32
		l5 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = s1f32
		s0i32 = l0
		s1i32 = 108
		s0i32 = s0i32 + s1i32
		s1f32 = l5
		s2f32 = l6
		f176(ctx, s0i32, s1f32, s2f32)
		s0i32 = l0
		s1i32 = 96
		s0i32 = s0i32 + s1i32
		s1f32 = l8
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1f32 = s1f32 - s2f32
		s2f32 = l7
		s3i32 = l2
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
		s2f32 = s2f32 - s3f32
		f176(ctx, s0i32, s1f32, s2f32)
		goto lbl3
	}
	s0i32 = l0
	s1i32 = 108
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 96
	s1i32 = s1i32 + s2i32
	s2i32 = l0
	s3i32 = 44
	s2i32 = s2i32 + s3i32
	s3i32 = l0
	s4i32 = 60
	s3i32 = s3i32 + s4i32
	s4i32 = l3
	s5i32 = l0
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s6i32 = l0
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
	s7i32 = l0
	s7i32 = int32(ctx.Mem[int(s7i32+84)])
	s8i32 = l4
	s9i32 = l0
	s9i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s9i32+92)]))
	if int(s9i32) < 0 || int(s9i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s9i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s9i32].numParams != 9 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32, int32, float32, float32, int32, int32))(table[s9i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5f32, s6f32, s7i32, s8i32)
lbl3:
	s0i32 = l0
	s1i32 = l4
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = 1
lbl0:
	return s0i32
}
