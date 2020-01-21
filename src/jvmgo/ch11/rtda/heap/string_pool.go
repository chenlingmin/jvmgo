package heap

import "unicode/utf16"

var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	chars := stringToUtf16(goStr)
	jChars := &Object{
		class: loader.LoadClass("[C"),
		data:  chars,
	}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr
}

func stringToUtf16(s string) []uint16 {
	runes := []rune(s) // utf32
	return utf16.Encode(runes)
}

func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s)
	return string(runes)
}

func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	internedStrings[goStr] = jStr
	return jStr
}
