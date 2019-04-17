package cli

import "strings"

type Writer interface {
	WriteString(s string) (int, error)
}

type Buffer interface {
	Writer
	String() string
}

type Builder struct {
	buffer Buffer
}

func (c *Builder) WriteString(s string) *Builder {
	_, err := c.buffer.WriteString(s)
	if err != nil {
		panic(err)
	}
	return c
}

func (c *Builder) String() string {
	return c.buffer.String()
}

func New(buffer Buffer) *Builder {
	return &Builder{buffer: buffer}
}

func NewStringBuilder() *Builder {
	return New(&strings.Builder{})
}
