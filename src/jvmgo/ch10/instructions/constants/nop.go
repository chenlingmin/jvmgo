package constants

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
}



