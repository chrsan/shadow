package internal

import (
	"math"
	"unsafe"
)

func f395(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l6 = s0i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l7 = s1i32
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	s0i32 = l6
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l5 = s1i32
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 14
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = 14
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l1
		s1i32 = 14
		s2i32 = -1
		s3i32 = -1
		s4i32 = -1
		s5i32 = l6
		s6i32 = l5
		s7i32 = 24
		s6i32 = s6i32 * s7i32
		s5i32 = s5i32 + s6i32
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
		s6i32 = l6
		s7i32 = l7
		s8i32 = 24
		s7i32 = s7i32 * s8i32
		s6i32 = s6i32 + s7i32
		s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+16)]))
		s5f32 = s5f32 - s6f32
		s5i32 = int32(math.Float32bits(s5f32))
		s6i32 = 0
		s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		l5 = s0i32
		goto lbl0
	}
	s0i32 = l8
	s1i32 = 14
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l6
	s1i32 = l7
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = l1
	s1i32 = 18
	s2i32 = l5
	s3i32 = l7
	s4i32 = -1
	s5i32 = 0
	s6i32 = 0
	s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	l5 = s0i32
lbl0:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l5
	s3i32 = l4
	s4i32 = l7
	s1i32 = f26(ctx, s1i32, s2i32, s3i32, s4i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l6 = s0i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l7 = s1i32
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	s0i32 = l6
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l5 = s1i32
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 14
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = 14
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l1
		s1i32 = 14
		s2i32 = -1
		s3i32 = -1
		s4i32 = -1
		s5i32 = l6
		s6i32 = l5
		s7i32 = 24
		s6i32 = s6i32 * s7i32
		s5i32 = s5i32 + s6i32
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
		s6i32 = l6
		s7i32 = l7
		s8i32 = 24
		s7i32 = s7i32 * s8i32
		s6i32 = s6i32 + s7i32
		s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+16)]))
		s5f32 = s5f32 - s6f32
		s5i32 = int32(math.Float32bits(s5f32))
		s6i32 = 0
		s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		l5 = s0i32
		goto lbl3
	}
	s0i32 = l8
	s1i32 = 14
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l6
	s1i32 = l7
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
lbl4:
	s0i32 = l1
	s1i32 = 18
	s2i32 = l5
	s3i32 = l7
	s4i32 = -1
	s5i32 = 0
	s6i32 = 0
	s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	l5 = s0i32
lbl3:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l5
	s3i32 = l4
	s4i32 = l7
	s1i32 = f26(ctx, s1i32, s2i32, s3i32, s4i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l6 = s0i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l7 = s1i32
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	s0i32 = l6
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l5 = s1i32
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 14
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = 14
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l1
		s1i32 = 14
		s2i32 = -1
		s3i32 = -1
		s4i32 = -1
		s5i32 = l6
		s6i32 = l5
		s7i32 = 24
		s6i32 = s6i32 * s7i32
		s5i32 = s5i32 + s6i32
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
		s6i32 = l6
		s7i32 = l7
		s8i32 = 24
		s7i32 = s7i32 * s8i32
		s6i32 = s6i32 + s7i32
		s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+16)]))
		s5f32 = s5f32 - s6f32
		s5i32 = int32(math.Float32bits(s5f32))
		s6i32 = 0
		s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		l5 = s0i32
		goto lbl6
	}
	s0i32 = l8
	s1i32 = 14
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l6
	s1i32 = l7
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
lbl7:
	s0i32 = l1
	s1i32 = 18
	s2i32 = l5
	s3i32 = l7
	s4i32 = -1
	s5i32 = 0
	s6i32 = 0
	s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	l5 = s0i32
lbl6:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l5
	s3i32 = l4
	s4i32 = l7
	s1i32 = f26(ctx, s1i32, s2i32, s3i32, s4i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l5 = s0i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l6 = s1i32
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s0i32 = l5
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l2 = s1i32
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 14
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = 14
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l1
		s1i32 = 14
		s2i32 = -1
		s3i32 = -1
		s4i32 = -1
		s5i32 = l5
		s6i32 = l2
		s7i32 = 24
		s6i32 = s6i32 * s7i32
		s5i32 = s5i32 + s6i32
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
		s6i32 = l5
		s7i32 = l6
		s8i32 = 24
		s7i32 = s7i32 * s8i32
		s6i32 = s6i32 + s7i32
		s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+16)]))
		s5f32 = s5f32 - s6f32
		s5i32 = int32(math.Float32bits(s5f32))
		s6i32 = 0
		s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		l2 = s0i32
		goto lbl9
	}
	s0i32 = l7
	s1i32 = 14
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l5
	s1i32 = l6
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
lbl10:
	s0i32 = l1
	s1i32 = 18
	s2i32 = l2
	s3i32 = l6
	s4i32 = -1
	s5i32 = 0
	s6i32 = 0
	s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	l2 = s0i32
lbl9:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s3i32 = l4
	s4i32 = l6
	s1i32 = f26(ctx, s1i32, s2i32, s3i32, s4i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
}
