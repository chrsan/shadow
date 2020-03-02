package internal

import (
	"unsafe"
)

func f1989(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
	var l7 int64
	_ = l7
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	var s6i64 int64
	_ = s6i64
	s0i32 = l3
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
		l6 = s0i64
		s1i64 = 10
		s0i64 = s0i64 << (uint64(s1i64) & 63)
		s1i64 = 1072693248
		s0i64 = s0i64 & s1i64
		s1i64 = l6
		s2i64 = 1023
		s1i64 = s1i64 & s2i64
		s0i64 = s0i64 | s1i64
		s1i64 = l6
		s2i64 = 20
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i64 = 1124800395214848
		s1i64 = s1i64 & s2i64
		s0i64 = s0i64 | s1i64
		s1i32 = l1
		s1i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		l6 = s1i64
		s2i64 = 10
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i64 = 1072693248
		s1i64 = s1i64 & s2i64
		s2i64 = l6
		s3i64 = 1023
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		s2i64 = l6
		s3i64 = 20
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s3i64 = 1124800395214848
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		s0i64 = s0i64 + s1i64
		s1i32 = l4
		s1i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		l6 = s1i64
		s2i64 = 10
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i64 = 1072693248
		s1i64 = s1i64 & s2i64
		s2i64 = l6
		s3i64 = 1023
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		s2i64 = l6
		s3i64 = 20
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s3i64 = 1124800395214848
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		s2i64 = 1
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s0i64 = s0i64 + s1i64
		l6 = s0i64
	lbl1:
		s0i32 = l0
		s1i32 = l5
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s1i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])))
		l7 = s1i64
		s2i64 = 10
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i64 = 1072693248
		s1i64 = s1i64 & s2i64
		s2i64 = l7
		s3i64 = 1023
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		s2i64 = l7
		s3i64 = 20
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s3i64 = 1124800395214848
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		s2i32 = l1
		s2i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)])))
		l7 = s2i64
		s3i64 = 10
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s3i64 = 1072693248
		s2i64 = s2i64 & s3i64
		s3i64 = l7
		s4i64 = 1023
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		s3i64 = l7
		s4i64 = 20
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 1124800395214848
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		s1i64 = s1i64 + s2i64
		s2i32 = l4
		s2i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)])))
		l7 = s2i64
		s3i64 = 10
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s3i64 = 1072693248
		s2i64 = s2i64 & s3i64
		s3i64 = l7
		s4i64 = 1023
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		s3i64 = l7
		s4i64 = 20
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 1124800395214848
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		s3i64 = 1
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s1i64 = s1i64 + s2i64
		l7 = s1i64
		s2i64 = l6
		s3i32 = l2
		s3i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)])))
		l6 = s3i64
		s4i64 = 10
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 1072693248
		s3i64 = s3i64 & s4i64
		s4i64 = l6
		s5i64 = 1023
		s4i64 = s4i64 & s5i64
		s3i64 = s3i64 | s4i64
		s4i64 = l6
		s5i64 = 20
		s4i64 = s4i64 << (uint64(s5i64) & 63)
		s5i64 = 1124800395214848
		s4i64 = s4i64 & s5i64
		s3i64 = s3i64 | s4i64
		s4i32 = l1
		s4i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)])))
		l6 = s4i64
		s5i64 = 10
		s4i64 = s4i64 << (uint64(s5i64) & 63)
		s5i64 = 1072693248
		s4i64 = s4i64 & s5i64
		s5i64 = l6
		s6i64 = 1023
		s5i64 = s5i64 & s6i64
		s4i64 = s4i64 | s5i64
		s5i64 = l6
		s6i64 = 20
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 1124800395214848
		s5i64 = s5i64 & s6i64
		s4i64 = s4i64 | s5i64
		s3i64 = s3i64 + s4i64
		s4i32 = l4
		s4i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)])))
		l6 = s4i64
		s5i64 = 10
		s4i64 = s4i64 << (uint64(s5i64) & 63)
		s5i64 = 1072693248
		s4i64 = s4i64 & s5i64
		s5i64 = l6
		s6i64 = 1023
		s5i64 = s5i64 & s6i64
		s4i64 = s4i64 | s5i64
		s5i64 = l6
		s6i64 = 20
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 1124800395214848
		s5i64 = s5i64 & s6i64
		s4i64 = s4i64 | s5i64
		s5i64 = 1
		s4i64 = s4i64 << (uint64(s5i64) & 63)
		s3i64 = s3i64 + s4i64
		s4i64 = 1
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s2i64 = s2i64 + s3i64
		s1i64 = s1i64 + s2i64
		l6 = s1i64
		s2i64 = 14
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		s2i64 = 1047552
		s1i64 = s1i64 & s2i64
		s2i64 = l6
		s3i64 = 4
		s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
		s3i64 = 1023
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		s2i64 = l6
		s3i64 = 24
		s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
		s3i64 = 1072693248
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i64)
		s0i32 = l2
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l4
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l1
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i64 = l7
		l6 = s0i64
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
