package base

import "jvmgo/ch09/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPc := pc + offset
	frame.SetNextPC(nextPc)
}
