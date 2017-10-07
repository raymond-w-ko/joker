// This file is generated by generate-std.joke script. Do not edit manually!

package string

import (
  "strings"
  "unicode"
  . "github.com/candid82/joker/core"
)

var stringNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.string"))

var index_of_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 2:
    
    s := ExtractString(args, 0)
    value := ExtractObject(args, 1)
    res := indexOf(s, value, 0)
    return res

  case c == 3:
    
    s := ExtractString(args, 0)
    value := ExtractObject(args, 1)
    from := ExtractInt(args, 2)
    res := indexOf(s, value, from)
    return res

  default:
    PanicArity(c)
  }
  return NIL
}

var lower_case_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractString(args, 0)
    res := strings.ToLower(s)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var reverse_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractString(args, 0)
    res := reverse(s)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var split_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 2:
    
    s := ExtractString(args, 0)
    re := ExtractRegex(args, 1)
    res := split(s, re, 0)
    return res

  case c == 3:
    
    s := ExtractString(args, 0)
    re := ExtractRegex(args, 1)
    n := ExtractInt(args, 2)
    res := split(s, re, n)
    return res

  default:
    PanicArity(c)
  }
  return NIL
}

var isstarts_with_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 2:
    
    s := ExtractString(args, 0)
    substr := ExtractString(args, 1)
    res := strings.HasPrefix(s, substr)
    return MakeBool(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var trim_left_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractString(args, 0)
    res := strings.TrimLeftFunc(s, unicode.IsSpace)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var isends_with_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 2:
    
    s := ExtractString(args, 0)
    substr := ExtractString(args, 1)
    res := strings.HasSuffix(s, substr)
    return MakeBool(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var trim_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractString(args, 0)
    res := strings.TrimSpace(s)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var escape_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 2:
    
    s := ExtractString(args, 0)
    cmap := ExtractCallable(args, 1)
    res := escape(s, cmap)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var join_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    coll := ExtractSeqable(args, 0)
    res := join("", coll)
    return MakeString(res)

  case c == 2:
    
    separator := ExtractString(args, 0)
    coll := ExtractSeqable(args, 1)
    res := join(separator, coll)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var pad_right_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 3:
    
    s := ExtractString(args, 0)
    pad := ExtractString(args, 1)
    n := ExtractInt(args, 2)
    res := padRight(s, pad, n)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var capitalize_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractString(args, 0)
    res := capitalize(s)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var replace_first_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 3:
    
    s := ExtractString(args, 0)
    match := ExtractObject(args, 1)
    repl := ExtractString(args, 2)
    res := replaceFirst(s, match, repl)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var pad_left_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 3:
    
    s := ExtractString(args, 0)
    pad := ExtractString(args, 1)
    n := ExtractInt(args, 2)
    res := padLeft(s, pad, n)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var isblank_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractObject(args, 0)
    res := isBlank(s)
    return MakeBool(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var upper_case_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractString(args, 0)
    res := strings.ToUpper(s)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var isincludes_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 2:
    
    s := ExtractString(args, 0)
    substr := ExtractString(args, 1)
    res := strings.Contains(s, substr)
    return MakeBool(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var last_index_of_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 2:
    
    s := ExtractString(args, 0)
    value := ExtractObject(args, 1)
    res := lastIndexOf(s, value, 0)
    return res

  case c == 3:
    
    s := ExtractString(args, 0)
    value := ExtractObject(args, 1)
    from := ExtractInt(args, 2)
    res := lastIndexOf(s, value, from)
    return res

  default:
    PanicArity(c)
  }
  return NIL
}

var trim_newline_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractString(args, 0)
    res := strings.TrimRight(s, "\n\r")
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var split_lines_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractString(args, 0)
    res := split(s, newLine, 0)
    return res

  default:
    PanicArity(c)
  }
  return NIL
}

var trim_right_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractString(args, 0)
    res := strings.TrimRightFunc(s, unicode.IsSpace)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var replace_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 3:
    
    s := ExtractString(args, 0)
    match := ExtractObject(args, 1)
    repl := ExtractString(args, 2)
    res := replace(s, match, repl)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}


func init() {

stringNamespace.ResetMeta(MakeMeta(nil, "Implements simple functions to manipulate strings.", "1.0"))

stringNamespace.InternVar("index-of", index_of_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("value")), NewVectorFrom(MakeSymbol("s"), MakeSymbol("value"), MakeSymbol("from"))),
    `Return index of value (string or char) in s, optionally searching
  forward from from or nil if not found.`, "1.0"))

stringNamespace.InternVar("lower-case", lower_case_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `Converts string to all lower-case.`, "1.0"))

stringNamespace.InternVar("reverse", reverse_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `Returns s with its characters reversed.`, "1.0"))

stringNamespace.InternVar("split", split_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("re")), NewVectorFrom(MakeSymbol("s"), MakeSymbol("re"), MakeSymbol("n"))),
    `Splits string on a regular expression. Returns vector of the splits.`, "1.0"))

stringNamespace.InternVar("starts-with?", isstarts_with_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("substr"))),
    `True if s starts with substr.`, "1.0"))

stringNamespace.InternVar("trim-left", trim_left_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `Removes whitespace from the left side of string.`, "1.0"))

stringNamespace.InternVar("ends-with?", isends_with_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("substr"))),
    `True if s ends with substr.`, "1.0"))

stringNamespace.InternVar("trim", trim_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `Removes whitespace from both ends of string.`, "1.0"))

stringNamespace.InternVar("escape", escape_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("cmap"))),
    `Return a new string, using cmap to escape each character ch
  from s as follows:

  If (cmap ch) is nil, append ch to the new string.
  If (cmap ch) is non-nil, append (str (cmap ch)) instead.`, "1.0"))

stringNamespace.InternVar("join", join_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("coll")), NewVectorFrom(MakeSymbol("separator"), MakeSymbol("coll"))),
    `Returns a string of all elements in coll, as returned by (seq coll), separated by an optional separator.`, "1.0"))

stringNamespace.InternVar("pad-right", pad_right_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("pad"), MakeSymbol("n"))),
    `Returns s padded with pad at the end to length n.`, "1.0"))

stringNamespace.InternVar("capitalize", capitalize_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `Converts first character of the string to upper-case, all other
  characters to lower-case.`, "1.0"))

stringNamespace.InternVar("replace-first", replace_first_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("match"), MakeSymbol("repl"))),
    `Replaces the first instance of match (String of Regex) with string repl in string s.

  If match is Regex, $1, $2, etc. in the replacement string repl are
  substituted with the string that matched the corresponding
  parenthesized group in the pattern.
  `, "1.0"))

stringNamespace.InternVar("pad-left", pad_left_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("pad"), MakeSymbol("n"))),
    `Returns s padded with pad at the beginning to length n.`, "1.0"))

stringNamespace.InternVar("blank?", isblank_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `True if s is nil, empty, or contains only whitespace.`, "1.0"))

stringNamespace.InternVar("upper-case", upper_case_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `Converts string to all upper-case.`, "1.0"))

stringNamespace.InternVar("includes?", isincludes_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("substr"))),
    `True if s includes substr.`, "1.0"))

stringNamespace.InternVar("last-index-of", last_index_of_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("value")), NewVectorFrom(MakeSymbol("s"), MakeSymbol("value"), MakeSymbol("from"))),
    `Return last index of value (string or char) in s, optionally
  searching backward from from or nil if not found.`, "1.0"))

stringNamespace.InternVar("trim-newline", trim_newline_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `Removes all trailing newline \n or return \r characters from string.`, "1.0"))

stringNamespace.InternVar("split-lines", split_lines_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `Splits string on \n or \r\n. Returns vector of the splits.`, "1.0"))

stringNamespace.InternVar("trim-right", trim_right_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `Removes whitespace from the right side of string.`, "1.0"))

stringNamespace.InternVar("replace", replace_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("match"), MakeSymbol("repl"))),
    `Replaces all instances of match (String of Regex) with string repl in string s.

  If match is Regex, $1, $2, etc. in the replacement string repl are
  substituted with the string that matched the corresponding
  parenthesized group in the pattern.
  `, "1.0"))

}
