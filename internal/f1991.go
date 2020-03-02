package internal

import (
	"unsafe"
)

func f1991(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int64
	_ = l4
	var l5 int64
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
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
	s0i32 = l3
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
		l4 = s0i64
		s1i64 = 10
		s0i64 = s0i64 << (uint64(s1i64) & 63)
		s1i64 = 1072693248
		s0i64 = s0i64 & s1i64
		s1i64 = l4
		s2i64 = 1023
		s1i64 = s1i64 & s2i64
		s0i64 = s0i64 | s1i64
		s1i64 = l4
		s2i64 = 20
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i64 = 1124800395214848
		s1i64 = s1i64 & s2i64
		s0i64 = s0i64 | s1i64
		s1i64 = l4
		s2i64 = 30
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i64 = 3458764513820540928
		s1i64 = s1i64 & s2i64
		s0i64 = s0i64 | s1i64
		l4 = s0i64
		s0i32 = 0
		l2 = s0i32
	lbl1:
		s0i32 = l0
		s1i32 = l2
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i64 = l4
		s2i32 = l1
		s2i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)])))
		l4 = s2i64
		s3i64 = 10
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s3i64 = 1072693248
		s2i64 = s2i64 & s3i64
		s3i64 = l4
		s4i64 = 1023
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		s3i64 = l4
		s4i64 = 20
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 1124800395214848
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		s3i64 = l4
		s4i64 = 30
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 3458764513820540928
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		l4 = s2i64
		s1i64 = s1i64 + s2i64
		s2i32 = l1
		s2i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)])))
		l5 = s2i64
		s3i64 = 10
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s3i64 = 1072693248
		s2i64 = s2i64 & s3i64
		s3i64 = l5
		s4i64 = 1023
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		s3i64 = l5
		s4i64 = 20
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 1124800395214848
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		s3i64 = l5
		s4i64 = 30
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 3458764513820540928
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		s3i64 = 1
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s1i64 = s1i64 + s2i64
		l5 = s1i64
		s2i64 = 12
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		s2i64 = 1047552
		s1i64 = s1i64 & s2i64
		s2i64 = l5
		s3i64 = 2
		s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
		s3i64 = 1023
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		s2i64 = l5
		s3i64 = 22
		s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
		s3i64 = 1072693248
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		s2i64 = l5
		s3i64 = 32
		s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
		s3i64 = 3221225472
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i64)
		s0i32 = l1
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
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
