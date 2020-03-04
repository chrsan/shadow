package internal

import (
	"unsafe"
)

func f1341(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 80
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s0i32 = int32(ctx.Mem[int(s0i32+85)])
		l5 = s0i32
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	l5 = s0i32
	s0i32 = l3
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l3
	s1i32 = l5
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
lbl0:
	s0i32 = l5
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l3 = s0i32
		s0i32 = l4
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l4
		s1i32 = l3
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0i32 = l4
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+10)])
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l2
	s0i32 = f65(ctx, s0i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l0
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = 6
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l5 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l3
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
		s3i32 = l2
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s4i32 = l2
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
		s3i32 = s3i32 | s4i32
		s2i32 = s2i32 | s3i32
		if s2i32 == 0 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			goto lbl5
		}
	}
	s1i32 = l2
	s2i32 = l3
	s3i32 = l6
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s1i32 = f64(ctx, s1i32, s2i32, s3i32, s4i32)
lbl5:
	s0i32 = f70(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		goto lbl2
	}
lbl4:
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 - s1f32
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0f32 = s0f32 - s1f32
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+10)])
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s1i32 = l2
	f550(ctx, s0i32, s1i32)
	goto lbl2
lbl7:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4388)]))
	l5 = s0i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 1
		f57(ctx, s0i32, s1i32)
	}
	s0i32 = l6
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l2
	s3i32 = 0
	s4i32 = l3
	s0i32 = f61(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l2 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
lbl10:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l3
	s1i32 = l1
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
	s3i32 = 0
	s4i32 = l3
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+104)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 4 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl10
	}
lbl9:
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+64)])
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		f54(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s0i32 = f23(ctx, s0i32)
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
lbl2:
	s0i32 = l6
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
