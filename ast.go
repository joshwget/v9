package main
import "fmt"

type node interface {
  Interpret() *variable
  AddChild(in node)
}

type constant struct {
  val *variable
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

type function_declare struct {
  body node
}

type function_call struct {
  function_node *variable
}

type set_prop struct {
  obj *variable
  prop string
  in node
}

type get_prop struct {
  obj *variable
  prop string
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
  switch right.Type {
    case 0:
      n.left.SetNumberValue(right.N)
    case 1:
      n.left.SetBoolValue(right.B)
    case 2:
      n.left.SetFunctionValue(right.F)
    case 3:
      n.left.SetReferenceValue(right)
    case 4:
      n.left.SetReferenceValue(right.R)
    case 5:
      n.left.SetStringValue(right.S)
  }
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

func (n *function_declare) Interpret() *variable {
  out := new(variable)
  out.SetFunctionValue(n.body)
  return out
}

func (n function_declare) AddChild(in node) { }

func (n *function_call) Interpret() *variable {
  n.function_node.F.Interpret()
  return nil
}

func (n function_call) AddChild(in node) { }

func (n *set_prop) Interpret() *variable {
  val := n.in.Interpret()
  n.obj.SetProp(n.prop, val)
  return nil
}

func (n set_prop) AddChild(in node) { }

func (n *get_prop) Interpret() *variable {
  return n.obj.GetProp(n.prop)
}

func (n get_prop) AddChild(in node) { }
