// +build !js

package gl

// Enum is equivalent to GLenum, and is normally used with one of the
// constants defined in this package.
type Enum uint32

// Attrib identifies the location of a specific attribute variable.
type Attrib struct {
	Value uint
}

// Program identifies a compiled shader program.
type Program struct {
	Value uint32
}

// Shader identifies a GLSL shader.
type Shader struct {
	Value uint32
}

// Buffer identifies a GL buffer object.
type Buffer struct {
	Value uint32
}

// Framebuffer identifies a GL framebuffer.
type Framebuffer struct {
	Value uint32
}

// A Renderbuffer is a GL object that holds an image in an internal format.
type Renderbuffer struct {
	Value uint32
}

// A Texture identifies a GL texture unit.
type Texture struct {
	Value uint32
}

// Uniform identifies the location of a specific uniform variable.
type Uniform struct {
	Value int32
}

// VertexArray reprensents a VAO.
type VertexArray struct {
	Value uint32
}

func (v Attrib) Valid() bool       { return v.Value != 0 }
func (v Program) Valid() bool      { return v.Value != 0 }
func (v Shader) Valid() bool       { return v.Value != 0 }
func (v Buffer) Valid() bool       { return v.Value != 0 }
func (v Framebuffer) Valid() bool  { return v.Value != 0 }
func (v Renderbuffer) Valid() bool { return v.Value != 0 }
func (v Texture) Valid() bool      { return v.Value != 0 }
func (v Uniform) Valid() bool      { return v.Value != 0 }
func (v VertexArray) Valid() bool  { return v.Value != 0 }
