package control

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

// Return void from method
type RETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}
