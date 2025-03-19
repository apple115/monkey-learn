package vm

import (
	"monkey/code"
	"monkey/object"
)

// 函数调用是嵌套的，与执行相关的数据——指令和指令指针——以后进先出 (LIFO) 方式被访问。我们很熟悉栈，这是很重要的一点。但是处理两个独立的数据片段从来都不是一件轻松的事情。解决方案是将它们捆绑在一起，这就是“帧”​。

type Frame struct {
	fn *object.CompiledFunction
	//该桢的指令指针
	ip int

	basePointer int
}

func NewFrame(fn *object.CompiledFunction,basePointer int) *Frame {
	return &Frame{fn: fn, ip: -1,basePointer: basePointer}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}

// 定义Frame后，我们就有了两种选择：一种是将虚拟机更改为仅在调用和执行函数时使用帧；
//另一种选择更高效流畅，即更改虚拟机，不仅将帧用于函数，而且将主程序bytecode.Instructions视为函数。
