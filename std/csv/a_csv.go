// This file is generated by generate-std.joke script. Do not edit manually!

package csv

import (
	. "github.com/candid82/joker/core"
)

var csvNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.csv"))



var csv_seq_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		rdr := ExtractObject(_args, 0)
		_res := csvSeqOpts(rdr, EmptyArrayMap())
		return _res

	case _c == 2:
		rdr := ExtractObject(_args, 0)
		opts := ExtractMap(_args, 1)
		_res := csvSeqOpts(rdr, opts)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var write_string_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		data := ExtractSeqable(_args, 0)
		_res := writeString(data, EmptyArrayMap())
		return MakeString(_res)

	case _c == 2:
		data := ExtractSeqable(_args, 0)
		opts := ExtractMap(_args, 1)
		_res := writeString(data, opts)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

func init() {

	csvNamespace.ResetMeta(MakeMeta(nil, "Reads and writes comma-separated values (CSV) files as defined in RFC 4180.", "1.0"))

	
	csvNamespace.InternVar("csv-seq", csv_seq_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("rdr")), NewVectorFrom(MakeSymbol("rdr"), MakeSymbol("opts"))),
			`Returns the csv records from rdr as a lazy sequence.
  rdr must implement io.Reader.
  opts may have the following keys:

  :comma - field delimiter (defaults to ',').
  Must be a valid char and must not be \r, \n,
  or the Unicode replacement character (0xFFFD).

  :comment - comment character (defaults to 0 meaning no comments).
  Lines beginning with the comment character without preceding whitespace are ignored.
  With leading whitespace the comment character becomes part of the
  field, even if trim-leading-space is true.
  comment must be a valid chat and must not be \r, \n,
  or the Unicode replacement character (0xFFFD).
  It must also not be equal to comma.

  :fields-per-record - number of expected fields per record.
  If fields-per-record is positive, csv-seq requires each record to
  have the given number of fields. If fields-per-record is 0 (default), csv-seq sets it to
  the number of fields in the first record, so that future records must
  have the same field count. If fields-per-record is negative, no check is
  made and records may have a variable number of fields.

  :lazy-quotes - if true, a quote may appear in an unquoted field and a
  non-doubled quote may appear in a quoted field. Default value is false.

  :trim-leading-space - if true, leading white space in a field is ignored.
  This is done even if the field delimiter, comma, is white space.
  Default value is false.`, "1.0"))

	csvNamespace.InternVar("write-string", write_string_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("data")), NewVectorFrom(MakeSymbol("data"), MakeSymbol("opts"))),
			`Writes records to a string in CSV format and returns that string.
  data must be Seqable, each element of which must be Seqable as well.
  opts may have the following keys:

  :comma - field delimiter (defaults to ',')

  :use-crlf - if true, uses \r\n as the line terminator. Default value is false.`, "1.0"))

}
