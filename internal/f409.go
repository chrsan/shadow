package internal

import (
	"unsafe"
)

func f409(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
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
	s0i32 = l2
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 128
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl2:
		s0i32 = l0
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = 2
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			l6 = s0i32
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l7 = s0i32
			s1i32 = 1
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			l0 = s0i32
			if s0i32 != 0 {
			lbl5:
				s0i32 = l1
				s1i32 = l0
				l2 = s1i32
				s2i32 = -1
				s1i32 = s1i32 + s2i32
				l0 = s1i32
				s2i32 = 2
				s1i32 = s1i32 << (uint32(s2i32) & 31)
				s0i32 = s0i32 + s1i32
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
				l8 = s0i32
				s0i32 = l0
				l4 = s0i32
				s0i32 = l2
				s1i32 = 1
				s0i32 = s0i32 << (uint32(s1i32) & 31)
				l3 = s0i32
				s1i32 = l7
				if uint32(s0i32) > uint32(s1i32) {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl6
				}
				s0i32 = l8
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
				l10 = s0i32
			lbl7:
				s0i32 = l8
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
				s1i32 = l10
				s2i32 = l3
				s3i32 = l6
				if uint32(s2i32) <= uint32(s3i32) {
					s2i32 = 1
				} else {
					s2i32 = 0
				}
				if s2i32 != 0 {
					s2i32 = l3
					s3i32 = l1
					s4i32 = l3
					s5i32 = 2
					s4i32 = s4i32 << (uint32(s5i32) & 31)
					s3i32 = s3i32 + s4i32
					l4 = s3i32
					s4i32 = -4
					s3i32 = s3i32 + s4i32
					s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
					l9 = s3i32
					s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
					s4i32 = l9
					s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)]))
					l9 = s4i32
					s5i32 = l9
					s6i32 = l4
					s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
					l4 = s6i32
					s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+16)]))
					l11 = s6i32
					if s5i32 == s6i32 {
						s5i32 = 1
					} else {
						s5i32 = 0
					}
					l9 = s5i32
					if s5i32 != 0 {
						// s3i32 = s3i32
					} else {
						s3i32 = s4i32
					}
					s4i32 = l4
					s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
					s5i32 = l11
					s6i32 = l9
					if s6i32 != 0 {
						// s4i32 = s4i32
					} else {
						s4i32 = s5i32
					}
					if s3i32 < s4i32 {
						s3i32 = 1
					} else {
						s3i32 = 0
					}
					s2i32 = s2i32 | s3i32
					l3 = s2i32
				}
				s2i32 = l10
				s3i32 = l1
				s4i32 = l3
				s5i32 = -1
				s4i32 = s4i32 + s5i32
				l4 = s4i32
				s5i32 = 2
				s4i32 = s4i32 << (uint32(s5i32) & 31)
				s3i32 = s3i32 + s4i32
				s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
				l9 = s3i32
				s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
				l11 = s3i32
				if s2i32 == s3i32 {
					s2i32 = 1
				} else {
					s2i32 = 0
				}
				l12 = s2i32
				if s2i32 != 0 {
					// s0i32 = s0i32
				} else {
					s0i32 = s1i32
				}
				s1i32 = l9
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
				s2i32 = l11
				s3i32 = l12
				if s3i32 != 0 {
					// s1i32 = s1i32
				} else {
					s1i32 = s2i32
				}
				if s0i32 >= s1i32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					s0i32 = l2
					s1i32 = -1
					s0i32 = s0i32 + s1i32
					l4 = s0i32
					goto lbl6
				}
				s0i32 = l2
				s1i32 = 2
				s0i32 = s0i32 << (uint32(s1i32) & 31)
				s1i32 = l1
				s0i32 = s0i32 + s1i32
				s1i32 = -4
				s0i32 = s0i32 + s1i32
				s1i32 = l9
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
				s0i32 = l3
				l2 = s0i32
				s1i32 = 1
				s0i32 = s0i32 << (uint32(s1i32) & 31)
				l3 = s0i32
				s1i32 = l7
				if uint32(s0i32) <= uint32(s1i32) {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl7
				}
			lbl6:
				s0i32 = l1
				s1i32 = l4
				s2i32 = 2
				s1i32 = s1i32 << (uint32(s2i32) & 31)
				s0i32 = s0i32 + s1i32
				s1i32 = l8
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
				s0i32 = l0
				if s0i32 != 0 {
					goto lbl5
				}
				s0i32 = l5
				if s0i32 == 0 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl0
				}
			}
		lbl11:
			s0i32 = l1
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l0 = s0i32
			s0i32 = l1
			s1i32 = l1
			s2i32 = l6
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			l2 = s1i32
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l2
			s1i32 = l0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = 2
			if uint32(s0i32) < uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl0
			}
			s0i32 = l1
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l4 = s0i32
			s0i32 = 2
			l0 = s0i32
			s0i32 = 1
			l2 = s0i32
		lbl12:
			s0i32 = l2
			s1i32 = 2
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			s1i32 = l1
			s0i32 = s0i32 + s1i32
			s1i32 = -4
			s0i32 = s0i32 + s1i32
			s1i32 = l1
			s2i32 = l0
			s3i32 = l6
			if uint32(s2i32) >= uint32(s3i32) {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				s2i32 = l0
			} else {
				s2i32 = l0
				s3i32 = l1
				s4i32 = l0
				s5i32 = 2
				s4i32 = s4i32 << (uint32(s5i32) & 31)
				s3i32 = s3i32 + s4i32
				l3 = s3i32
				s4i32 = -4
				s3i32 = s3i32 + s4i32
				s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
				l5 = s3i32
				s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
				s4i32 = l5
				s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)]))
				l5 = s4i32
				s5i32 = l5
				s6i32 = l3
				s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
				l3 = s6i32
				s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+16)]))
				l8 = s6i32
				if s5i32 == s6i32 {
					s5i32 = 1
				} else {
					s5i32 = 0
				}
				l5 = s5i32
				if s5i32 != 0 {
					// s3i32 = s3i32
				} else {
					s3i32 = s4i32
				}
				s4i32 = l3
				s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
				s5i32 = l8
				s6i32 = l5
				if s6i32 != 0 {
					// s4i32 = s4i32
				} else {
					s4i32 = s5i32
				}
				if s3i32 < s4i32 {
					s3i32 = 1
				} else {
					s3i32 = 0
				}
				s2i32 = s2i32 | s3i32
			}
			l3 = s2i32
			s3i32 = -1
			s2i32 = s2i32 + s3i32
			l5 = s2i32
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l3
			l2 = s0i32
			s0i32 = l3
			s1i32 = 1
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			l0 = s0i32
			s1i32 = l6
			if uint32(s0i32) <= uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl12
			}
			s0i32 = l3
			s1i32 = 1
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			l0 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl14
			}
			s0i32 = l4
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			l2 = s0i32
		lbl15:
			s0i32 = l1
			s1i32 = l0
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			l5 = s1i32
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l8 = s0i32
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			s1i32 = l8
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			l7 = s1i32
			s2i32 = l2
			s3i32 = l7
			if s2i32 == s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			l7 = s2i32
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			s1i32 = l4
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			s2i32 = l2
			s3i32 = l7
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
			if s0i32 >= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l3
				s1i32 = -1
				s0i32 = s0i32 + s1i32
				l5 = s0i32
				goto lbl14
			}
			s0i32 = l3
			s1i32 = 2
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			s1i32 = l1
			s0i32 = s0i32 + s1i32
			s1i32 = -4
			s0i32 = s0i32 + s1i32
			s1i32 = l8
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l0
			l3 = s0i32
			s1i32 = 1
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			l0 = s0i32
			if s0i32 != 0 {
				goto lbl15
			}
		lbl14:
			s0i32 = l1
			s1i32 = l5
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			s1i32 = l4
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l6 = s0i32
			if s0i32 != 0 {
				goto lbl11
			}
			goto lbl0
		}
		s0i32 = l1
		s1i32 = l5
		s2i32 = 1
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 2147483644
		s1i32 = s1i32 & s2i32
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		s0i32 = l3
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l1
		s1i32 = l2
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			l6 = s0i32
			goto lbl17
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l5 = s0i32
		s0i32 = l1
		l6 = s0i32
		l3 = s0i32
	lbl19:
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l7 = s1i32
		s2i32 = l5
		s3i32 = l7
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l7 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l5
		s3i32 = l7
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l6
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = l8
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l6 = s0i32
		}
		s0i32 = l3
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = l2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl19
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
	lbl17:
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l3 = s0i32
		s0i32 = l6
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		s2i32 = l6
		s3i32 = -4
		s2i32 = s2i32 + s3i32
		f409(ctx, s0i32, s1i32, s2i32)
		s0i32 = l2
		s1i32 = l6
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s0i32 = s0i32 - s1i32
		l5 = s0i32
		s1i32 = 128
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	l0 = s0i32
lbl21:
	s0i32 = l3
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	l8 = s0i32
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l4 = s1i32
	s2i32 = l4
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	l0 = s3i32
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	l7 = s3i32
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l10 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l7
	s3i32 = l10
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl23:
		s0i32 = l3
		s1i32 = l3
		s2i32 = -4
		s1i32 = s1i32 + s2i32
		l0 = s1i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl24
		}
		s0i32 = l3
		s1i32 = -8
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = l0
		l3 = s0i32
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s1i32 = l4
		s2i32 = l4
		s3i32 = l7
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		l7 = s3i32
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
		l10 = s3i32
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l9 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l10
		s3i32 = l9
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl23
		}
	lbl24:
		s0i32 = l0
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l6
	l0 = s0i32
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l2
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
lbl0:
}
