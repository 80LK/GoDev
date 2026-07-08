package context

import (
	"strconv"
	"strings"
)

type Context struct {
	level       int
	levelPrefix string
	counter     int
}

func (ctx *Context) NextLevel() *Context {
	return &Context{
		level:       ctx.level + 1,
		levelPrefix: ctx.levelPrefix,
		counter:     1,
	}
}

func (ctx *Context) GetPrefix() string {
	prefix := strings.Repeat(ctx.levelPrefix, ctx.level)

	return prefix
}

func (ctx *Context) GetCounter() string {
	n := strconv.Itoa(ctx.counter) + ". "
	ctx.counter++
	return n
}

func Default() *Context {
	return &Context{
		level:       0,
		levelPrefix: "\t",
		counter:     1,
	}
}

func New(LevelPrefix string) *Context {
	return &Context{
		level:       0,
		levelPrefix: LevelPrefix,
		counter:     1,
	}
}

func Get(ctx *Context) *Context {
	if ctx == nil {
		return Default()
	} else {
		return ctx
	}
}
