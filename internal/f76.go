package internal

import (
	"unsafe"
)

func f76(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
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
	var s10i32 int32
	_ = s10i32
	var s11i32 int32
	_ = s11i32
	var s12i32 int32
	_ = s12i32
	s0i32 = l1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	l15 = s1i32
	s2i32 = 16711680
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 - s1i32
	s1i32 = 255
	s0i32 = s0i32 * s1i32
	s1i32 = 32768
	s0i32 = s0i32 + s1i32
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l15
	s2i32 = 65535
	s1i32 = s1i32 & s2i32
	s2i32 = 255
	s1i32 = s1i32 * s2i32
	s2i32 = 32768
	s1i32 = s1i32 + s2i32
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s0i32 = s0i32 - s1i32
	l13 = s0i32
	s0i32 = l8
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	l12 = s1i32
	s2i32 = l12
	s3i32 = l8
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l10 = s0i32
	s0i32 = l9
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l14 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	l12 = s1i32
	s2i32 = l12
	s3i32 = l9
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l11 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l16 = s0i32
	s0i32 = l14
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l17 = s0i32
	s0i32 = 1
	l12 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l13
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s1i32 = 255
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l12 = s0i32
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s1i32 = l1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = 65536
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	l7 = s2i32
	s3i32 = l7
	s4i32 = 31
	s3i32 = s3i32 >> (uint32(s4i32) & 31)
	l7 = s3i32
	s2i32 = s2i32 + s3i32
	s3i32 = l7
	s2i32 = s2i32 ^ s3i32
	s1i32 = s1i32 - s2i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl2:
	s0i32 = l14
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l12 = s0i32
		goto lbl0
	}
	s0i32 = l14
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l12 = s0i32
		goto lbl0
	}
	s0i32 = 0
	l12 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s1i32 = l1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l14
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = 65536
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l7
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	l1 = s2i32
	s3i32 = l1
	s4i32 = 31
	s3i32 = s3i32 >> (uint32(s4i32) & 31)
	l1 = s3i32
	s2i32 = s2i32 + s3i32
	s3i32 = l1
	s2i32 = s2i32 ^ s3i32
	s1i32 = s1i32 - s2i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l12 = s0i32
lbl0:
	s0i32 = l10
	s1i32 = l11
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l8
	s1i32 = l2
	s2i32 = l2
	s3i32 = l8
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l8 = s0i32
	s1i32 = l9
	s2i32 = l3
	s3i32 = l9
	s4i32 = l3
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l9 = s1i32
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l9
		s1i32 = l11
		s2i32 = l11
		s3i32 = l9
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l1 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l2 = s0i32
		s1i32 = l8
		s2i32 = l10
		s3i32 = l10
		s4i32 = l8
		if s3i32 > s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		l3 = s3i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l7 = s1i32
		s2i32 = l7
		s3i32 = l2
		if s2i32 < s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l11
		s2i32 = l9
		s3i32 = l1
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l1 = s1i32
		s2i32 = l10
		s3i32 = l8
		s4i32 = l3
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		l2 = s2i32
		s3i32 = l1
		s4i32 = l2
		if s3i32 < s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		s0i32 = i32DivS(s0i32, s1i32)
		l8 = s0i32
		l9 = s0i32
	}
	s0i32 = l10
	s1i32 = l11
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2i32 = l8
	s3i32 = l9
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l15
	s1i32 = 16
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l1 = s0i32
	s0i32 = l11
	s1i32 = l9
	s2i32 = l9
	s3i32 = l11
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l2 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l14 = s0i32
	s0i32 = l8
	s1i32 = l10
	s2i32 = l8
	s3i32 = l10
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l15 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l3 = s0i32
	s0i32 = l9
	s1i32 = l11
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l7 = s0i32
	s1i32 = -65536
	s0i32 = s0i32 & s1i32
	l2 = s0i32
	s1i32 = l10
	s2i32 = l8
	s3i32 = l15
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l9 = s1i32
	s2i32 = 65535
	s1i32 = s1i32 + s2i32
	l15 = s1i32
	s2i32 = -65536
	s1i32 = s1i32 & s2i32
	l8 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = l3
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l8
		s1i32 = l3
		s0i32 = s0i32 - s1i32
		l11 = s0i32
		s1i32 = 65535
		s0i32 = s0i32 + s1i32
		s1i32 = 16
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s1i32 = 1
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l10
			s1i32 = 1
			s0i32 = s0i32 - s1i32
			if s0i32 != 0 {
				s0i32 = l3
				s1i32 = 16
				s0i32 = s0i32 >> (uint32(s1i32) & 31)
				l10 = s0i32
				s0i32 = l13
				s1i32 = 255
				s0i32 = s0i32 & s1i32
				s1i32 = 255
				if s0i32 != s1i32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				s1i32 = l12
				s0i32 = s0i32 | s1i32
				l16 = s0i32
				s0i32 = l8
				s1i32 = l9
				s0i32 = s0i32 - s1i32
				s1i32 = l11
				s0i32 = s0i32 + s1i32
				s1i32 = 2
				s0i32 = i32DivS(s0i32, s1i32)
				s1i32 = 8
				s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
				l3 = s0i32
				s0i32 = l6
				if s0i32 != 0 {
					s0i32 = l5
					s1i32 = l10
					s0i32 = s0i32 + s1i32
					l9 = s0i32
					s0i32 = l16
					if s0i32 == 0 {
						s0i32 = 1
					} else {
						s0i32 = 0
					}
					if s0i32 != 0 {
						s0i32 = l9
						s1i32 = l3
						ctx.Mem[int(s0i32+0)] = uint8(s1i32)
						goto lbl8
					}
					s0i32 = l9
					s1i32 = l9
					s1i32 = int32(ctx.Mem[int(s1i32+0)])
					s2i32 = l3
					s3i32 = 255
					s2i32 = s2i32 & s3i32
					s3i32 = l13
					s4i32 = 255
					s3i32 = s3i32 & s4i32
					s2i32 = s2i32 * s3i32
					s3i32 = 8
					s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
					s1i32 = s1i32 + s2i32
					l3 = s1i32
					s2i32 = 255
					s3i32 = l3
					s4i32 = 255
					if uint32(s3i32) < uint32(s4i32) {
						s3i32 = 1
					} else {
						s3i32 = 0
					}
					if s3i32 != 0 {
						// s1i32 = s1i32
					} else {
						s1i32 = s2i32
					}
					ctx.Mem[int(s0i32+0)] = uint8(s1i32)
					goto lbl8
				}
				s0i32 = l16
				if s0i32 == 0 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					s0i32 = l4
					s1i32 = 0
					s2i32 = l4
					s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
					s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
					if int(s2i32) < 0 || int(s2i32) >= len(table) {
						panic("table entry out of bounds")
					}
					if table[s2i32].numParams == -1 {
						panic("table entry is nil")
					}
					if table[s2i32].numParams != 2 {
						panic("argument count mismatch")
					}
					s0i32 = (*(*func(*Context, int32, int32) int32)(table[s2i32].f()))(ctx, s0i32, s1i32)
					l9 = s0i32
					s1i32 = l10
					s2i32 = l1
					s3i32 = 1
					s4i32 = l3
					s5i32 = 255
					s4i32 = s4i32 & s5i32
					s5i32 = l9
					s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
					s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
					if int(s5i32) < 0 || int(s5i32) >= len(table) {
						panic("table entry out of bounds")
					}
					if table[s5i32].numParams == -1 {
						panic("table entry is nil")
					}
					if table[s5i32].numParams != 5 {
						panic("argument count mismatch")
					}
					(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
					goto lbl8
				}
				s0i32 = l4
				s1i32 = l10
				s2i32 = l1
				s3i32 = l3
				s4i32 = 255
				s3i32 = s3i32 & s4i32
				s4i32 = l13
				s5i32 = 255
				s4i32 = s4i32 & s5i32
				s3i32 = s3i32 * s4i32
				s4i32 = 8
				s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
				s4i32 = l4
				s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
				s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+64)]))
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
				goto lbl8
			}
			s0i32 = l13
			s1i32 = l9
			s2i32 = l3
			s1i32 = s1i32 - s2i32
			s2i32 = l11
			s3i32 = -65536
			s2i32 = s2i32 + s3i32
			l10 = s2i32
			s1i32 = s1i32 - s2i32
			s2i32 = 11
			s1i32 = s1i32 >> (uint32(s2i32) & 31)
			l9 = s1i32
			s2i32 = l16
			s3i32 = 11
			s2i32 = s2i32 >> (uint32(s3i32) & 31)
			l11 = s2i32
			s1i32 = s1i32 * s2i32
			s2i32 = l9
			s1i32 = s1i32 * s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s0i32 = s0i32 - s1i32
			l9 = s0i32
			s0i32 = l10
			s1i32 = 11
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l10 = s0i32
			s1i32 = l11
			s0i32 = s0i32 * s1i32
			s1i32 = l10
			s0i32 = s0i32 * s1i32
			s1i32 = 8
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			l10 = s0i32
			s0i32 = l3
			s1i32 = 16
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l3 = s0i32
			s0i32 = l6
			if s0i32 != 0 {
				s0i32 = l3
				s1i32 = l5
				s0i32 = s0i32 + s1i32
				l3 = s0i32
				s1i32 = l3
				s1i32 = int32(ctx.Mem[int(s1i32+0)])
				s2i32 = l10
				s3i32 = 255
				s2i32 = s2i32 & s3i32
				s1i32 = s1i32 + s2i32
				l10 = s1i32
				s2i32 = 255
				s3i32 = l10
				s4i32 = 255
				if uint32(s3i32) < uint32(s4i32) {
					s3i32 = 1
				} else {
					s3i32 = 0
				}
				if s3i32 != 0 {
					// s1i32 = s1i32
				} else {
					s1i32 = s2i32
				}
				ctx.Mem[int(s0i32+0)] = uint8(s1i32)
				s0i32 = l3
				s1i32 = l3
				s1i32 = int32(ctx.Mem[int(s1i32+1)])
				s2i32 = l9
				s3i32 = 255
				s2i32 = s2i32 & s3i32
				s1i32 = s1i32 + s2i32
				l3 = s1i32
				s2i32 = 255
				s3i32 = l3
				s4i32 = 255
				if uint32(s3i32) < uint32(s4i32) {
					s3i32 = 1
				} else {
					s3i32 = 0
				}
				if s3i32 != 0 {
					// s1i32 = s1i32
				} else {
					s1i32 = s2i32
				}
				ctx.Mem[int(s0i32+1)] = uint8(s1i32)
				goto lbl8
			}
			s0i32 = l13
			s1i32 = 255
			s0i32 = s0i32 & s1i32
			s1i32 = 255
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = l12
			s0i32 = s0i32 | s1i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l4
				s1i32 = 0
				s2i32 = l4
				s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
				s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
				if int(s2i32) < 0 || int(s2i32) >= len(table) {
					panic("table entry out of bounds")
				}
				if table[s2i32].numParams == -1 {
					panic("table entry is nil")
				}
				if table[s2i32].numParams != 2 {
					panic("argument count mismatch")
				}
				s0i32 = (*(*func(*Context, int32, int32) int32)(table[s2i32].f()))(ctx, s0i32, s1i32)
				l11 = s0i32
				s1i32 = l3
				s2i32 = l1
				s3i32 = l10
				s4i32 = 255
				s3i32 = s3i32 & s4i32
				s4i32 = l9
				s5i32 = 255
				s4i32 = s4i32 & s5i32
				s5i32 = l11
				s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
				s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+36)]))
				if int(s5i32) < 0 || int(s5i32) >= len(table) {
					panic("table entry out of bounds")
				}
				if table[s5i32].numParams == -1 {
					panic("table entry is nil")
				}
				if table[s5i32].numParams != 5 {
					panic("argument count mismatch")
				}
				(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
				goto lbl8
			}
			s0i32 = l4
			s1i32 = l3
			s2i32 = l1
			s3i32 = l10
			s4i32 = 255
			s3i32 = s3i32 & s4i32
			s4i32 = l4
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+64)]))
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
			s0i32 = l4
			s1i32 = l3
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			s2i32 = l1
			s3i32 = l9
			s4i32 = 255
			s3i32 = s3i32 & s4i32
			s4i32 = l4
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+64)]))
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
			goto lbl8
		}
		s0i32 = l4
		s1i32 = l1
		s2i32 = l3
		s3i32 = l8
		s4i32 = l9
		s5i32 = l8
		s6i32 = l16
		s7i32 = 2147483647
		s8i32 = l13
		s9i32 = 255
		s8i32 = s8i32 & s9i32
		s9i32 = l5
		s10i32 = l6
		s11i32 = l12
		s12i32 = 1
		f32(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32, s9i32, s10i32, s11i32, s12i32)
	lbl8:
		s0i32 = l2
		s1i32 = l8
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l15
		s1i32 = 16
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l3 = s0i32
		s0i32 = l2
		s1i32 = l8
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s1i32 = 16
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l9 = s0i32
		s0i32 = l6
		if s0i32 != 0 {
			s0i32 = l8
			s1i32 = 1
			if s0i32 < s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl16
			}
			s0i32 = l13
			s1i32 = 255
			s0i32 = s0i32 & s1i32
			l10 = s0i32
			s0i32 = 0
			l8 = s0i32
		lbl18:
			s0i32 = l5
			s1i32 = l3
			s2i32 = l8
			s1i32 = s1i32 + s2i32
			s0i32 = s0i32 + s1i32
			l11 = s0i32
			s1i32 = l11
			s1i32 = int32(ctx.Mem[int(s1i32+0)])
			s2i32 = l10
			s1i32 = s1i32 + s2i32
			l11 = s1i32
			s2i32 = 255
			s3i32 = l11
			s4i32 = 255
			if uint32(s3i32) < uint32(s4i32) {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
			ctx.Mem[int(s0i32+0)] = uint8(s1i32)
			s0i32 = l8
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l8 = s0i32
			s1i32 = l9
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl18
			}
			goto lbl16
		}
		s0i32 = l13
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		l8 = s0i32
		s1i32 = 255
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = l12
		s0i32 = s0i32 | s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 0
			s2i32 = l4
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
			if int(s2i32) < 0 || int(s2i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s2i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s2i32].numParams != 2 {
				panic("argument count mismatch")
			}
			s0i32 = (*(*func(*Context, int32, int32) int32)(table[s2i32].f()))(ctx, s0i32, s1i32)
			l8 = s0i32
			s1i32 = l3
			s2i32 = l1
			s3i32 = l9
			s4i32 = l8
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
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
			goto lbl16
		}
		s0i32 = l4
		s1i32 = l3
		s2i32 = l1
		s3i32 = l9
		s4i32 = l8
		s5i32 = l4
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+68)]))
		if int(s5i32) < 0 || int(s5i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s5i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s5i32].numParams != 5 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	lbl16:
		s0i32 = l14
		s1i32 = l2
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l14
		s1i32 = l2
		s0i32 = s0i32 - s1i32
		l9 = s0i32
		s1i32 = 65535
		s0i32 = s0i32 + s1i32
		s1i32 = 16
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = 1
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = 1
			s0i32 = s0i32 - s1i32
			if s0i32 != 0 {
				s0i32 = l7
				s1i32 = 16
				s0i32 = s0i32 >> (uint32(s1i32) & 31)
				l3 = s0i32
				s0i32 = l13
				s1i32 = 255
				s0i32 = s0i32 & s1i32
				s1i32 = 255
				if s0i32 != s1i32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				s1i32 = l12
				s0i32 = s0i32 | s1i32
				l8 = s0i32
				s0i32 = l7
				s1i32 = l2
				s0i32 = s0i32 - s1i32
				s1i32 = l9
				s0i32 = s0i32 + s1i32
				s1i32 = 2
				s0i32 = i32DivS(s0i32, s1i32)
				s1i32 = 8
				s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
				l2 = s0i32
				s0i32 = l6
				if s0i32 != 0 {
					s0i32 = l3
					s1i32 = l5
					s0i32 = s0i32 + s1i32
					l1 = s0i32
					s0i32 = l8
					if s0i32 == 0 {
						s0i32 = 1
					} else {
						s0i32 = 0
					}
					if s0i32 != 0 {
						s0i32 = l1
						s1i32 = l2
						ctx.Mem[int(s0i32+0)] = uint8(s1i32)
						goto lbl5
					}
					s0i32 = l1
					s1i32 = l1
					s1i32 = int32(ctx.Mem[int(s1i32+0)])
					s2i32 = l2
					s3i32 = 255
					s2i32 = s2i32 & s3i32
					s3i32 = l13
					s4i32 = 255
					s3i32 = s3i32 & s4i32
					s2i32 = s2i32 * s3i32
					s3i32 = 8
					s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
					s1i32 = s1i32 + s2i32
					l1 = s1i32
					s2i32 = 255
					s3i32 = l1
					s4i32 = 255
					if uint32(s3i32) < uint32(s4i32) {
						s3i32 = 1
					} else {
						s3i32 = 0
					}
					if s3i32 != 0 {
						// s1i32 = s1i32
					} else {
						s1i32 = s2i32
					}
					ctx.Mem[int(s0i32+0)] = uint8(s1i32)
					goto lbl5
				}
				s0i32 = l8
				if s0i32 == 0 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					s0i32 = l4
					s1i32 = 0
					s2i32 = l4
					s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
					s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
					if int(s2i32) < 0 || int(s2i32) >= len(table) {
						panic("table entry out of bounds")
					}
					if table[s2i32].numParams == -1 {
						panic("table entry is nil")
					}
					if table[s2i32].numParams != 2 {
						panic("argument count mismatch")
					}
					s0i32 = (*(*func(*Context, int32, int32) int32)(table[s2i32].f()))(ctx, s0i32, s1i32)
					l4 = s0i32
					s1i32 = l3
					s2i32 = l1
					s3i32 = 1
					s4i32 = l2
					s5i32 = 255
					s4i32 = s4i32 & s5i32
					s5i32 = l4
					s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
					s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
					if int(s5i32) < 0 || int(s5i32) >= len(table) {
						panic("table entry out of bounds")
					}
					if table[s5i32].numParams == -1 {
						panic("table entry is nil")
					}
					if table[s5i32].numParams != 5 {
						panic("argument count mismatch")
					}
					(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
					goto lbl5
				}
				s0i32 = l4
				s1i32 = l3
				s2i32 = l1
				s3i32 = l2
				s4i32 = 255
				s3i32 = s3i32 & s4i32
				s4i32 = l13
				s5i32 = 255
				s4i32 = s4i32 & s5i32
				s3i32 = s3i32 * s4i32
				s4i32 = 8
				s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
				s4i32 = l4
				s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
				s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+64)]))
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
				goto lbl5
			}
			s0i32 = l14
			s1i32 = l7
			s0i32 = s0i32 - s1i32
			s1i32 = l2
			s2i32 = l7
			s1i32 = s1i32 - s2i32
			s2i32 = 65536
			s1i32 = s1i32 + s2i32
			l2 = s1i32
			s0i32 = s0i32 - s1i32
			s1i32 = 11
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l3 = s0i32
			s1i32 = l17
			s2i32 = 11
			s1i32 = s1i32 >> (uint32(s2i32) & 31)
			l8 = s1i32
			s0i32 = s0i32 * s1i32
			s1i32 = l3
			s0i32 = s0i32 * s1i32
			s1i32 = 8
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			l3 = s0i32
			s0i32 = l13
			s1i32 = l2
			s2i32 = 11
			s1i32 = s1i32 >> (uint32(s2i32) & 31)
			l2 = s1i32
			s2i32 = l8
			s1i32 = s1i32 * s2i32
			s2i32 = l2
			s1i32 = s1i32 * s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s0i32 = s0i32 - s1i32
			l8 = s0i32
			s0i32 = l7
			s1i32 = 16
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l2 = s0i32
			s0i32 = l6
			if s0i32 != 0 {
				s0i32 = l2
				s1i32 = l5
				s0i32 = s0i32 + s1i32
				l1 = s0i32
				s1i32 = l1
				s1i32 = int32(ctx.Mem[int(s1i32+0)])
				s2i32 = l8
				s3i32 = 255
				s2i32 = s2i32 & s3i32
				s1i32 = s1i32 + s2i32
				l2 = s1i32
				s2i32 = 255
				s3i32 = l2
				s4i32 = 255
				if uint32(s3i32) < uint32(s4i32) {
					s3i32 = 1
				} else {
					s3i32 = 0
				}
				if s3i32 != 0 {
					// s1i32 = s1i32
				} else {
					s1i32 = s2i32
				}
				ctx.Mem[int(s0i32+0)] = uint8(s1i32)
				s0i32 = l1
				s1i32 = l1
				s1i32 = int32(ctx.Mem[int(s1i32+1)])
				s2i32 = l3
				s3i32 = 255
				s2i32 = s2i32 & s3i32
				s1i32 = s1i32 + s2i32
				l1 = s1i32
				s2i32 = 255
				s3i32 = l1
				s4i32 = 255
				if uint32(s3i32) < uint32(s4i32) {
					s3i32 = 1
				} else {
					s3i32 = 0
				}
				if s3i32 != 0 {
					// s1i32 = s1i32
				} else {
					s1i32 = s2i32
				}
				ctx.Mem[int(s0i32+1)] = uint8(s1i32)
				goto lbl5
			}
			s0i32 = l13
			s1i32 = 255
			s0i32 = s0i32 & s1i32
			s1i32 = 255
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = l12
			s0i32 = s0i32 | s1i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l4
				s1i32 = 0
				s2i32 = l4
				s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
				s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
				if int(s2i32) < 0 || int(s2i32) >= len(table) {
					panic("table entry out of bounds")
				}
				if table[s2i32].numParams == -1 {
					panic("table entry is nil")
				}
				if table[s2i32].numParams != 2 {
					panic("argument count mismatch")
				}
				s0i32 = (*(*func(*Context, int32, int32) int32)(table[s2i32].f()))(ctx, s0i32, s1i32)
				l4 = s0i32
				s1i32 = l2
				s2i32 = l1
				s3i32 = l8
				s4i32 = 255
				s3i32 = s3i32 & s4i32
				s4i32 = l3
				s5i32 = 255
				s4i32 = s4i32 & s5i32
				s5i32 = l4
				s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
				s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+36)]))
				if int(s5i32) < 0 || int(s5i32) >= len(table) {
					panic("table entry out of bounds")
				}
				if table[s5i32].numParams == -1 {
					panic("table entry is nil")
				}
				if table[s5i32].numParams != 5 {
					panic("argument count mismatch")
				}
				(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
				goto lbl5
			}
			s0i32 = l4
			s1i32 = l2
			s2i32 = l1
			s3i32 = l8
			s4i32 = 255
			s3i32 = s3i32 & s4i32
			s4i32 = l4
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+64)]))
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
			s0i32 = l4
			s1i32 = l2
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			s2i32 = l1
			s3i32 = l3
			s4i32 = 255
			s3i32 = s3i32 & s4i32
			s4i32 = l4
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+64)]))
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
			goto lbl5
		}
		s0i32 = l4
		s1i32 = l1
		s2i32 = l2
		s3i32 = l7
		s4i32 = l2
		s5i32 = l14
		s6i32 = 2147483647
		s7i32 = l17
		s8i32 = l13
		s9i32 = 255
		s8i32 = s8i32 & s9i32
		s9i32 = l5
		s10i32 = l6
		s11i32 = l12
		s12i32 = 1
		f32(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32, s9i32, s10i32, s11i32, s12i32)
		goto lbl5
	}
	s0i32 = l4
	s1i32 = l1
	s2i32 = l3
	s3i32 = l7
	s4i32 = l9
	s5i32 = l14
	s6i32 = l16
	s7i32 = l17
	s8i32 = l13
	s9i32 = 255
	s8i32 = s8i32 & s9i32
	s9i32 = l5
	s10i32 = l6
	s11i32 = l12
	s12i32 = 1
	f32(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32, s9i32, s10i32, s11i32, s12i32)
lbl5:
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
}
