package internal

import (
	"unsafe"
)

func f110(ctx *Context, l0 int32) {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+60)])
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+60)] = uint8(s1i32)
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+24)])
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = 2147483647
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l0
		s2i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
		s1i32 = l0
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint32(s2i32)
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		goto lbl5
	}
	s1i32 = l2
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l3 = s1i32
	s2i32 = 2147483647
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl2
	}
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l2
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		goto lbl7
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l3 = s1i32
lbl7:
	s1i32 = l0
	s2i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint32(s2i32)
	s1i32 = l1
	s2i32 = 16
	s1i32 = s1i32 + s2i32
lbl5:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l0
	s1i32 = 44
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l0
	s1i32 = 28
	s0i32 = s0i32 + s1i32
	l6 = s0i32
lbl9:
	s0i32 = l5
	s1i32 = l6
	s2i32 = l4
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 0
		ctx.Mem[int(s0i32+60)] = uint8(s1i32)
		return
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+24)])
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = 2147483647
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l0
		s2i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
		s1i32 = l0
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint32(s2i32)
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		goto lbl12
	}
	s1i32 = l2
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l3 = s1i32
	s2i32 = 2147483647
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l2
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		goto lbl14
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l3 = s1i32
lbl14:
	s1i32 = l0
	s2i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint32(s2i32)
	s1i32 = l1
	s2i32 = 16
	s1i32 = s1i32 + s2i32
lbl12:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
lbl3:
	return
lbl2:
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	goto lbl0
lbl1:
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+24)] = uint8(s1i32)
	return
lbl0:
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+24)] = uint8(s1i32)
}
