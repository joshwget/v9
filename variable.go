package main
import "strconv"

type variable struct {
  Type uint8
  N float32
  B bool
  F node
  O map[string]*variable
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

func (v *variable) SetStringValue(s string) {
  v.S = s;
  v.Type = 4;
}

func (v *variable) SetObjectValue(o map[string]*variable) {
  v.O = o;
  v.Type = 3;
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
  switch v.Type {
    case 3:
      return &v.O
    default:
      return nil
  }
}

func (obj *variable) CreateProp(key string) *variable {
  v := new(variable)
  switch obj.Type {
    case 2, 3:
      obj.O[key] = v
  }
  return v
}

func (obj *variable) GetProp(key string) *variable {
  switch obj.Type {
    case 2, 3:
      val, err := obj.O[key]
      if err {
        return val
      } else {
        val, err := obj.O["prototype"]
        if err {
          return val.GetProp(key)
        } else {
          return nil
        }
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
  if true_constant == nil {
    true_value := new(variable)
    true_value.SetBoolValue(true)
    true_constant = &constant{ true_value }
  }

  return true_constant
}

func FalseConstant() *constant {
  if false_constant == nil {
    false_value := new(variable)
    false_value.SetBoolValue(false)
    false_constant = &constant{ false_value }
  }

  return false_constant
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
    case 4:
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
    case 4:
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
      to.O = from.O
    case 3:
      to.SetObjectValue(from.O)
    case 4:
      to.SetStringValue(from.S)
  }
  return to
}

var vars map[string]*variable
var context *variable
var current_function_in_new *variable

var true_constant *constant
var false_constant *constant
