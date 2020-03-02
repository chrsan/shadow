package internal

import (
	"unsafe"
)

func f939(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s1i32 = 14
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
	s0i32 = 22928
	l1 = s0i32
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl15
	case 1:
		goto lbl14
	case 2:
		goto lbl13
	case 3:
		goto lbl12
	case 4:
		goto lbl11
	case 5:
		goto lbl10
	case 6:
		goto lbl9
	case 7:
		goto lbl8
	case 8:
		goto lbl7
	case 9:
		goto lbl6
	case 10:
		goto lbl5
	case 11:
		goto lbl4
	case 12:
		goto lbl3
	case 13:
		goto lbl2
	default:
		goto lbl1
	}
lbl15:
	s0i32 = 22952
	l1 = s0i32
	goto lbl1
lbl14:
	s0i32 = 22976
	l1 = s0i32
	goto lbl1
lbl13:
	s0i32 = 23000
	l1 = s0i32
	goto lbl1
lbl12:
	s0i32 = 23024
	l1 = s0i32
	goto lbl1
lbl11:
	s0i32 = 23048
	l1 = s0i32
	goto lbl1
lbl10:
	s0i32 = 23072
	l1 = s0i32
	goto lbl1
lbl9:
	s0i32 = 23096
	l1 = s0i32
	goto lbl1
lbl8:
	s0i32 = 23120
	l1 = s0i32
	goto lbl1
lbl7:
	s0i32 = 23144
	l1 = s0i32
	goto lbl1
lbl6:
	s0i32 = 23168
	l1 = s0i32
	goto lbl1
lbl5:
	s0i32 = 23192
	l1 = s0i32
	goto lbl1
lbl4:
	s0i32 = 23216
	l1 = s0i32
	goto lbl1
lbl3:
	s0i32 = 23240
	l1 = s0i32
	goto lbl1
lbl2:
	s0i32 = 23264
	l1 = s0i32
lbl1:
	s0i32 = 8
	s0i32 = f17(ctx, s0i32)
	l0 = s0i32
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	return s0i32
}
