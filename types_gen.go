// Generated by gen_types. Don't modify manually!

package main

import (
  "fmt"
)

func assertComparable(obj Object, msg string) Comparable {
  switch c := obj.(type) {
  case Comparable:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Comparable", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureComparable(args []Object, index int) Comparable {
  switch c := args[index].(type) {
  case Comparable:
    return c
  default:
    panic(RT.newArgTypeError(index, "Comparable"))
  }
}

func assertVector(obj Object, msg string) *Vector {
  switch c := obj.(type) {
  case *Vector:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Vector", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureVector(args []Object, index int) *Vector {
  switch c := args[index].(type) {
  case *Vector:
    return c
  default:
    panic(RT.newArgTypeError(index, "Vector"))
  }
}

func assertChar(obj Object, msg string) Char {
  switch c := obj.(type) {
  case Char:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Char", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureChar(args []Object, index int) Char {
  switch c := args[index].(type) {
  case Char:
    return c
  default:
    panic(RT.newArgTypeError(index, "Char"))
  }
}

func assertString(obj Object, msg string) String {
  switch c := obj.(type) {
  case String:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "String", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureString(args []Object, index int) String {
  switch c := args[index].(type) {
  case String:
    return c
  default:
    panic(RT.newArgTypeError(index, "String"))
  }
}

func assertSymbol(obj Object, msg string) Symbol {
  switch c := obj.(type) {
  case Symbol:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Symbol", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureSymbol(args []Object, index int) Symbol {
  switch c := args[index].(type) {
  case Symbol:
    return c
  default:
    panic(RT.newArgTypeError(index, "Symbol"))
  }
}

func assertKeyword(obj Object, msg string) Keyword {
  switch c := obj.(type) {
  case Keyword:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Keyword", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureKeyword(args []Object, index int) Keyword {
  switch c := args[index].(type) {
  case Keyword:
    return c
  default:
    panic(RT.newArgTypeError(index, "Keyword"))
  }
}

func assertBool(obj Object, msg string) Bool {
  switch c := obj.(type) {
  case Bool:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Bool", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureBool(args []Object, index int) Bool {
  switch c := args[index].(type) {
  case Bool:
    return c
  default:
    panic(RT.newArgTypeError(index, "Bool"))
  }
}

func assertNumber(obj Object, msg string) Number {
  switch c := obj.(type) {
  case Number:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Number", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureNumber(args []Object, index int) Number {
  switch c := args[index].(type) {
  case Number:
    return c
  default:
    panic(RT.newArgTypeError(index, "Number"))
  }
}

func assertSeqable(obj Object, msg string) Seqable {
  switch c := obj.(type) {
  case Seqable:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Seqable", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureSeqable(args []Object, index int) Seqable {
  switch c := args[index].(type) {
  case Seqable:
    return c
  default:
    panic(RT.newArgTypeError(index, "Seqable"))
  }
}

func assertCallable(obj Object, msg string) Callable {
  switch c := obj.(type) {
  case Callable:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Callable", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureCallable(args []Object, index int) Callable {
  switch c := args[index].(type) {
  case Callable:
    return c
  default:
    panic(RT.newArgTypeError(index, "Callable"))
  }
}

func assertType(obj Object, msg string) *Type {
  switch c := obj.(type) {
  case *Type:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Type", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureType(args []Object, index int) *Type {
  switch c := args[index].(type) {
  case *Type:
    return c
  default:
    panic(RT.newArgTypeError(index, "Type"))
  }
}

func assertMeta(obj Object, msg string) Meta {
  switch c := obj.(type) {
  case Meta:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Meta", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureMeta(args []Object, index int) Meta {
  switch c := args[index].(type) {
  case Meta:
    return c
  default:
    panic(RT.newArgTypeError(index, "Meta"))
  }
}

func assertInt(obj Object, msg string) Int {
  switch c := obj.(type) {
  case Int:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Int", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureInt(args []Object, index int) Int {
  switch c := args[index].(type) {
  case Int:
    return c
  default:
    panic(RT.newArgTypeError(index, "Int"))
  }
}

func assertStack(obj Object, msg string) Stack {
  switch c := obj.(type) {
  case Stack:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Stack", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureStack(args []Object, index int) Stack {
  switch c := args[index].(type) {
  case Stack:
    return c
  default:
    panic(RT.newArgTypeError(index, "Stack"))
  }
}

func assertMap(obj Object, msg string) Map {
  switch c := obj.(type) {
  case Map:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Map", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureMap(args []Object, index int) Map {
  switch c := args[index].(type) {
  case Map:
    return c
  default:
    panic(RT.newArgTypeError(index, "Map"))
  }
}

func assertSet(obj Object, msg string) Set {
  switch c := obj.(type) {
  case Set:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Set", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func ensureSet(args []Object, index int) Set {
  switch c := args[index].(type) {
  case Set:
    return c
  default:
    panic(RT.newArgTypeError(index, "Set"))
  }
}
