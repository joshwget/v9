package main
import "fmt"
import "strconv"

type node interface {
  Interpret() *variable
  AddChild(in node)
}

type constant struct {
  val *variable
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

type operation2 struct {
  left, right node
  operator int
}

type print struct {
  in node
}

type block struct {
  ins []node
}

type assign struct {
  left *variable
  right node
}

type var_usage struct {
  in *variable
}

type if_node struct {
  cond node
  body node
}

type while_node struct {
  cond node
  body node
}

func (n constant) Interpret() *variable {
  return n.val
}

func (n constant) AddChild(in node) { }

func (n operation2) Interpret() *variable {
  left := n.left.Interpret()
  right := n.right.Interpret()
  out := new(variable)
  switch(n.operator) {
    case '+':
      out.SetNumberValue(left.N + right.N)
    case '-':
      out.SetNumberValue(left.N - right.N)
    case '*':
      out.SetNumberValue(left.N * right.N)
    case '/':
      out.SetNumberValue(left.N / right.N)
    case COMP_EQU:
      out.SetBoolValue(left.N == right.N)
    case COMP_NEQU:
      out.SetBoolValue(left.N != right.N)
    case COMP_SEQU:
      out.SetBoolValue(left.N == right.N)
    case COMP_SNEQU:
      out.SetBoolValue(left.N != right.N)
    case COMP_LESS:
      out.SetBoolValue(left.N < right.N)
    case COMP_LTE:
      out.SetBoolValue(left.N <= right.N)
    case COMP_GTR:
      out.SetBoolValue(left.N > right.N)
    case COMP_GTE:
      out.SetBoolValue(left.N >= right.N)
    case BOOL_AND:
      out.SetBoolValue(left.B && right.B)
    case BOOL_OR:
      out.SetBoolValue(left.B || right.B)
    default:
      out.SetNumberValue(left.N + right.N)
  }

  return out
}

func (n operation2) AddChild(in node) { }

func (n print) Interpret() *variable {
  val := n.in.Interpret()
  fmt.Println(StringCast(val))
  return nil
}

func (n print) AddChild(in node) { }

func (n block) Interpret() *variable {
  for _, in := range n.ins {
    in.Interpret()
  }
  return nil
}

func (n *block) AddChild(in node) {
  n.ins = append(n.ins, in)
}

func (n *assign) Interpret() *variable {
  right := n.right.Interpret()
  n.left.N = right.N
  return n.left
}

func (n assign) AddChild(in node) { }

func (n *var_usage) Interpret() *variable {
  return n.in
}

func (n var_usage) AddChild(in node) { }

func (n *if_node) Interpret() *variable {
  cond := n.cond.Interpret()
  if BoolCast(cond) {
    n.body.Interpret()
  }
  return nil
}

func (n if_node) AddChild(in node) { }

func (n *while_node) Interpret() *variable {
  cond := n.cond.Interpret()
  for BoolCast(cond) {
    n.body.Interpret()
  }
  return nil
}

func (n while_node) AddChild(in node) { }
