package main
import "strconv"
import "fmt"

type variable struct {
  Type int
  N float32
  B bool
  F node
  O map[string]*variable
  R *variable
  S string
}

func (v *variable) SetNumberValue(n float32) {
  v.N = n;
  v.Type = 0;
}

func (v *variable) SetBoolValue(b bool) {
  v.B = b;
  v.Type = 1;
}

func (v *variable) SetFunctionValue(function_node node) {
  v.F = function_node;
  v.Type = 2;
}

func (v *variable) SetReferenceValue(other *variable) {
  v.R = other;
  v.Type = 4;
}

func (v *variable) SetStringValue(s string) {
  v.S = s;
  v.Type = 5;
}

func MakeEmptyObject() *variable {
  v := new(variable)
  v.O = make(map[string]*variable)
  v.Type = 3
  return v
}

func MakeEmptyObjectNode() *constant {
  v := new(variable)
  v.O = make(map[string]*variable)
  v.Type = 3
  return &constant { v }
}

func (v *variable) GetProperties() *map[string]*variable {
  fmt.Println(v.Type)
  switch v.Type {
    case 3:
      return &v.O
    case 4:
      return &v.R.O
    default:
      return nil
  }
}

func (obj *variable) SetProp(key string, val *variable) {
  switch obj.Type {
    case 3:
      obj.O[key] = val
    case 4:
      obj.R.O[key] = val
  }
}

func (obj *variable) GetProp(key string) *variable {
  switch obj.Type {
    case 3:
      val, err := obj.O[key]
      if err {
        return val
      } else {
        return nil
      }
    case 4:
      val, err := obj.R.O[key]
      if err {
        return val
      } else {
        return nil
      }
    default:
      return nil
  }
}

func NumberConstant(n float32) *constant {
  v := new(variable)
  v.SetNumberValue(n)
  return &constant { v }
}

func TrueConstant() *constant {
  v := new(variable)
  v.SetBoolValue(true)
  return &constant { v }
}

func FalseConstant() *constant {
  v := new(variable)
  v.SetBoolValue(false)
  return &constant { v }
}

func StringConstant(s string) *constant {
  v := new(variable)
  v.SetStringValue(s)
  return &constant { v }
}

func BoolCast(v *variable) bool {
  switch v.Type {
    case 0:
      return v.N != 0;
    case 1:
      return v.B;
    case 5:
      return v.S != ""
    default:
      return false;
  }
}

func StringCast(v *variable) string {
  if v == nil {
    return "undefined";
  }

  switch v.Type {
    case 0:
      return strconv.FormatFloat(float64(v.N), 'f', -1, 32);
    case 1:
      if v.B {
        return "true"
      } else {
        return "false"
      }
    case 5:
      return v.S
    default:
      return "bad type";
  }
}

func Assign(from *variable, to *variable) *variable {
  switch from.Type {
    case 0:
      to.SetNumberValue(from.N)
    case 1:
      to.SetBoolValue(from.B)
    case 2:
      to.SetFunctionValue(from.F)
    case 3:
      to.SetReferenceValue(from)
    case 4:
      to.SetReferenceValue(from.R)
    case 5:
      to.SetStringValue(from.S)
  }
  return to
}

var vars map[string]*variable
var context *variable
