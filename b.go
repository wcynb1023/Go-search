package main

import "runtime"

func functionToDebug() {
	// 在要调试的函数内部使用 runtime.Breakpoint() 设置断点
	runtime.Breakpoint()

	// 函数的其他代码
}
