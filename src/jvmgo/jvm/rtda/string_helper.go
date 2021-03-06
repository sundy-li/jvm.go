package rtda

import (
    "jvmgo/util"
    rtc "jvmgo/jvm/rtda/class"
)

func JStringChars(jStr *rtc.Obj) ([]uint16) {
    charArr := jStr.GetFieldValue("value", "[C").(*rtc.Obj)
    return charArr.Fields().([]uint16)
}

// todo: is there a better way to create String?
func NewJString(goStr string, frame *Frame) (*rtc.Obj) {
    chars := util.StringToUtf16(goStr)
    internedStr := getInternedString(chars)
    if internedStr != nil {
        return internedStr
    }

    classLoader := frame.Method().Class().ClassLoader()
    stringClass := classLoader.StringClass()
    jCharArr := rtc.NewCharArray(chars, classLoader)
    jStr := stringClass.NewObj()
    jStr.SetFieldValue("value", "[C", jCharArr)
    return InternString(chars, jStr)
}

func GoString(jStr *rtc.Obj) string {
    utf16 := JStringChars(jStr)
    return util.Utf16ToString(utf16)
}
