package lang

import (
    //"math"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _runtime(freeMemory, "freeMemory", "()J")
}

func _runtime(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/Runtime", name, desc, method)
}

// public native long freeMemory();
// ()J
func freeMemory(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopRef() // this
    stack.PushLong(int64(1000000)) // todo
}
