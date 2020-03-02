package shadow

import (
	"unsafe"

	"github.com/chrsan/shadow/internal"
)

type Surface struct {
	c      *internal.Context
	ptr    int32
	Canvas *Canvas
}

func (s *Surface) Dispose() {
	if s != nil {
		s.c.SkiaSurfaceDestroy(s.ptr)
		s.Canvas = nil
		s = nil
	}
}

func (s *Surface) Data(peek bool) []uint8 {
	st := s.c.StackSave()
	defer s.c.StackRestore(st)
	sptr := s.c.StackAlloc(8)
	s.c.SkiaSurfaceReadPixels(s.ptr, sptr)
	dptr := *(*int32)(unsafe.Pointer(&s.c.Mem[sptr]))
	sz := *(*int32)(unsafe.Pointer(&s.c.Mem[sptr+4]))
	if dptr == 0 || sz == 0 {
		return nil
	}
	if peek {
		return (*[1 << 29]uint8)(unsafe.Pointer(&s.c.Mem[dptr]))[:sz:sz]
	}
	data := make([]uint8, sz)
	copy(data, s.c.Mem[dptr:])
	return data
}
