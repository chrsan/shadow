package shadow

import (
	"unsafe"

	"github.com/chrsan/shadow/internal"
)

type Canvas struct {
	c   *internal.Context
	ptr int32
}

func (c *Canvas) Clear(r, g, b, a uint8) {
	c.c.SkiaCanvasClear(c.ptr, int32(r), int32(g), int32(b), int32(a))
}

func (c *Canvas) Flush() {
	c.c.SkiaCanvasFlush(c.ptr)
}

func (c *Canvas) Save() {
	c.c.SkiaCanvasSave(c.ptr)
}

func (c *Canvas) Restore() {
	c.c.SkiaCanvasRestore(c.ptr)
}

func (c *Canvas) Scale(sx, sy float32) {
	c.c.SkiaCanvasScale(c.ptr, sx, sy)
}

func (c *Canvas) Translate(dx, dy float32) {
	c.c.SkiaCanvasTranslate(c.ptr, dx, dy)
}

func (c *Canvas) SetMatrix(m Matrix) {
	st, ptr := c.allocMatrix()
	defer c.c.StackRestore(st)
	c.c.SkiaMatrixSet(ptr, m.ScaleX, m.SkewX, m.TransX, m.SkewY, m.ScaleY, m.TransY)
	c.c.SkiaCanvasSetMatrix(c.ptr, ptr)
}

func (c *Canvas) Concat(m Matrix) {
	st, ptr := c.allocMatrix()
	defer c.c.StackRestore(st)
	c.c.SkiaMatrixSet(ptr, m.ScaleX, m.SkewX, m.TransX, m.SkewY, m.ScaleY, m.TransY)
	c.c.SkiaCanvasConcat(c.ptr, ptr)
}

func (c *Canvas) TotalMatrix() Matrix {
	st, ptr := c.allocMatrix()
	defer c.c.StackRestore(st)
	c.c.SkiaMatrixReset(ptr)
	c.c.SkiaCanvasGetTotalMatrix(c.ptr, ptr)
	return *(*Matrix)(unsafe.Pointer(&c.c.Mem[ptr]))
}

func (c *Canvas) allocMatrix() (st, ptr int32) {
	st = c.c.StackSave()
	sz := c.c.SkiaMatrixSize()
	ptr = c.c.StackAlloc(sz)
	return
}

func (c *Canvas) DrawPath(path *Path, paint *Paint) {
	c.c.SkiaCanvasDrawPath(c.ptr, path.ptr, paint.ptr)
}

func (c *Canvas) DrawRect(x, y, w, h float32, paint *Paint) {
	c.c.SkiaCanvasDrawRect(c.ptr, x, y, w, h, paint.ptr)
}
