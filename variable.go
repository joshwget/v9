package main
import "strconv"

type variable struct {
  Type int
  N float32
  B bool
  F node
  O map[string]*variable
  R *variable
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

func MakeEmptyObject() *constant {
  v := new(variable)
  v.O = make(map[string]*variable)
  v.Type = 3
  return &constant { v }
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

func BoolCast(v *variable) bool {
  switch v.Type {
    case 0:
      return v.N != 0;
    case 1:
      return v.B;
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
    default:
      return "bad type";
  }
}

var vars map[string]*variable
