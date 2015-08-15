package main
import "fmt"

type node interface {
  Interpret() *variable
  AddChild(in node)
}

type constant struct {
  val *variable
}

func IntConstant(i int) *constant {
  v := &variable{ I: i }
  return &constant { v }
}

func BoolCast(v *variable) bool {
  switch v.Type {
    case 0:
      return v.I != 0;
    case 1:
      return v.B;
    default:
      return false;
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
  return &variable{ I: n.val.I }
}

func (n constant) AddChild(in node) { }

func (n operation2) Interpret() *variable {
  left := n.left.Interpret()
  right := n.right.Interpret()
  out := new(variable)
  switch(n.operator) {
    case '+':
      out.SetNumberValue(left.I + right.I)
    case '-':
      out.SetNumberValue(left.I - right.I)
    case '*':
      out.SetNumberValue(left.I * right.I)
    case '/':
      out.SetNumberValue(left.I / right.I)
    case COMP_EQU:
      out.SetBoolValue(left.I == right.I)
    case COMP_NEQU:
      out.SetBoolValue(left.I != right.I)
    case COMP_SEQU:
      out.SetBoolValue(left.I == right.I)
    case COMP_SNEQU:
      out.SetBoolValue(left.I != right.I)
    case COMP_LESS:
      out.SetBoolValue(left.I < right.I)
    case COMP_LTE:
      out.SetBoolValue(left.I <= right.I)
    case COMP_GTR:
      out.SetBoolValue(left.I > right.I)
    case COMP_GTE:
      out.SetBoolValue(left.I >= right.I)
    case BOOL_AND:
      out.SetBoolValue(left.B && right.B)
    case BOOL_OR:
      out.SetBoolValue(left.B || right.B)
    default:
      out.SetNumberValue(left.I + right.I)
  }

  return out
}

func (n operation2) AddChild(in node) { }

func (n print) Interpret() *variable {
  val := n.in.Interpret()
  switch val.Type {
    case 0:
      fmt.Println(val.I);
    case 1:
      if val.B {
        fmt.Println("true");
      } else {
        fmt.Println("false");
      }
    default:
      fmt.Println("Invalid type")
  }
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
  n.left.I = right.I
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
