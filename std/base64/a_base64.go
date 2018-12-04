// This file is generated by generate-std.joke script. Do not edit manually!

package base64

import (
	
	. "github.com/candid82/joker/core"
)

var base64Namespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.base64"))

var decode_string_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 1:
		
		s := ExtractString(_args, 0)
		_res := decodeString(s)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var encode_string_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 1:
		
		s := ExtractString(_args, 0)
		_res := encodeString(s)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}


func init() {

base64Namespace.ResetMeta(MakeMeta(nil, "Implements base64 encoding as specified by RFC 4648.", "1.0"))

base64Namespace.InternVar("decode-string", decode_string_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("s"))),
		`Returns the bytes represented by the base64 string s.`, "1.0"))

base64Namespace.InternVar("encode-string", encode_string_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("s"))),
		`Returns the base64 encoding of s.`, "1.0"))

}
