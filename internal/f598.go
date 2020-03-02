package internal

import (
	"math"
	"unsafe"
)

func f598(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s6f32 float32
	_ = s6f32
	var s7f32 float32
	_ = s7f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l5 = s0i32
	s1i32 = l4
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s0i32 = l1
	s1i32 = l5
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l6 = s2i32
	s3i32 = 24
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 14
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l8 = s1f32
		s1i32 = l7
		s2i32 = 14
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl2
		}
		s1i32 = l0
		s2i32 = 14
		s3i32 = -1
		s4i32 = -1
		s5i32 = -1
		s6f32 = l8
		s7i32 = l5
		s8i32 = l4
		s9i32 = 24
		s8i32 = s8i32 * s9i32
		s7i32 = s7i32 + s8i32
		s7f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s7i32+16)]))
		s6f32 = s6f32 * s7f32
		s6i32 = int32(math.Float32bits(s6f32))
		s7i32 = 0
		s1i32 = f21(ctx, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32)
		goto lbl0
	}
	s1i32 = l7
	s2i32 = 14
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s1i32 = l5
	s2i32 = l4
	s3i32 = 24
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2f32 = 1
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s1i32 = l6
	goto lbl0
lbl2:
	s1i32 = l4
	s2f32 = l8
	s3f32 = 1
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl0
	}
lbl1:
	s1i32 = l0
	s2i32 = 21
	s3i32 = l6
	s4i32 = l4
	s5i32 = -1
	s6i32 = 0
	s7i32 = 0
	s1i32 = f21(ctx, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32)
lbl0:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l1 = s0i32
	s1i32 = l4
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = l2
	s1i32 = l1
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l5 = s2i32
	s3i32 = 24
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 14
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l7
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l8 = s1f32
		s1i32 = l6
		s2i32 = 14
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl6
		}
		s1i32 = l0
		s2i32 = 14
		s3i32 = -1
		s4i32 = -1
		s5i32 = -1
		s6f32 = l8
		s7i32 = l1
		s8i32 = l4
		s9i32 = 24
		s8i32 = s8i32 * s9i32
		s7i32 = s7i32 + s8i32
		s7f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s7i32+16)]))
		s6f32 = s6f32 * s7f32
		s6i32 = int32(math.Float32bits(s6f32))
		s7i32 = 0
		s1i32 = f21(ctx, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32)
		goto lbl4
	}
	s1i32 = l6
	s2i32 = 14
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl5
	}
	s1i32 = l1
	s2i32 = l4
	s3i32 = 24
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2f32 = 1
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl5
	}
	s1i32 = l5
	goto lbl4
lbl6:
	s1i32 = l4
	s2f32 = l8
	s3f32 = 1
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl4
	}
lbl5:
	s1i32 = l0
	s2i32 = 21
	s3i32 = l5
	s4i32 = l4
	s5i32 = -1
	s6i32 = 0
	s7i32 = 0
	s1i32 = f21(ctx, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32)
lbl4:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l1 = s0i32
	s1i32 = l4
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s0i32 = l1
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l5 = s1i32
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 14
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l8 = s0f32
		s0i32 = l2
		s1i32 = 14
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l3
		s1i32 = l0
		s2i32 = 14
		s3i32 = -1
		s4i32 = -1
		s5i32 = -1
		s6f32 = l8
		s7i32 = l1
		s8i32 = l4
		s9i32 = 24
		s8i32 = s8i32 * s9i32
		s7i32 = s7i32 + s8i32
		s7f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s7i32+16)]))
		s6f32 = s6f32 * s7f32
		s6i32 = int32(math.Float32bits(s6f32))
		s7i32 = 0
		s1i32 = f21(ctx, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		return
	}
	s0i32 = l2
	s1i32 = 14
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l1
	s1i32 = l4
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1f32 = 1
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	goto lbl8
lbl10:
	s0f32 = l8
	s1f32 = 1
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l4
	l5 = s0i32
lbl9:
	s0i32 = l3
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	return
lbl8:
	s0i32 = l3
	s1i32 = l0
	s2i32 = 21
	s3i32 = l5
	s4i32 = l4
	s5i32 = -1
	s6i32 = 0
	s7i32 = 0
	s1i32 = f21(ctx, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
}
