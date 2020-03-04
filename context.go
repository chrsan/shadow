package shadow

import (
	"unsafe"

	"github.com/chrsan/shadow/internal"
)

type Context struct {
	c *internal.Context
}

func NewContext() *Context {
	c := internal.NewContext(importedFuncs{})
	c.Start()
	return &Context{c}
}

func (c *Context) Dispose() error {
	return c.c.Dispose()
}

func (c *Context) InverseMatrix(m Matrix) (Matrix, bool) {
	st := c.c.StackSave()
	defer c.c.StackRestore(st)
	sz := c.c.SkiaMatrixSize()
	// Stack alloc two matrixes.
	ptr := c.c.StackAlloc(sz * 2)
	c.c.SkiaMatrixInit(ptr)
	c.c.SkiaMatrixInit(ptr + sz)
	// Set source matrix.
	c.c.SkiaMatrixSet(ptr, m.ScaleX, m.SkewX, m.TransX, m.SkewY, m.ScaleY, m.TransY)
	if c.c.SkiaMatrixCreateInverse(ptr, ptr+sz) == 0 {
		return Matrix{}, false
	}
	inv := *(*Matrix)(unsafe.Pointer(&c.c.Mem[ptr+sz]))
	return inv, true
}

func (c *Context) SurfaceIsBGRA() bool {
	return c.c.SkiaIsSurfaceBgra() != 0
}

func (c *Context) NewRGBASurface(width, height int) *Surface {
	ptr := c.c.SkiaSurfaceCreateRgba(int32(width), int32(height))
	if ptr == 0 {
		panic("could not create surface")
	}
	cptr := c.c.SkiaSurfaceGetCanvas(ptr)
	if cptr == 0 {
		panic("could not get canvas")
	}
	return &Surface{c: c.c, ptr: ptr, Canvas: &Canvas{c: c.c, ptr: cptr}}
}

func (c *Context) NewRGBAPremultipliedSurface(width, height int) *Surface {
	ptr := c.c.SkiaSurfaceCreateRgbaPremultiplied(int32(width), int32(height))
	if ptr == 0 {
		panic("could not create surface")
	}
	cptr := c.c.SkiaSurfaceGetCanvas(ptr)
	if cptr == 0 {
		panic("could not get canvas")
	}
	return &Surface{c: c.c, ptr: ptr, Canvas: &Canvas{c: c.c, ptr: cptr}}
}

func (c *Context) NewPaint() *Paint {
	ptr := c.c.SkiaPaintCreate()
	if ptr == 0 {
		panic("could not create paint")
	}
	return &Paint{c: c.c, ptr: ptr}
}

func (c *Context) StackAllocPaint() (paint *Paint, dispose func()) {
	st := c.c.StackSave()
	ptr := c.c.StackAlloc(c.c.SkiaPaintSize())
	c.c.SkiaPaintInit(ptr)
	paint = &Paint{c: c.c, ptr: ptr}
	dispose = func() {
		c.c.StackRestore(st)
	}
	return
}

func (c *Context) NewPath() *Path {
	ptr := c.c.SkiaPathCreate()
	if ptr == 0 {
		panic("could not create path")
	}
	return &Path{c: c.c, ptr: ptr}
}

func (c *Context) StackAllocPath() (path *Path, dispose func()) {
	st := c.c.StackSave()
	ptr := c.c.StackAlloc(c.c.SkiaPathSize())
	c.c.SkiaPathInit(ptr)
	path = &Path{c: c.c, ptr: ptr}
	dispose = func() {
		c.c.StackRestore(st)
	}
	return
}
