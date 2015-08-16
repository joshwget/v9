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
  left_var *variable
  left_node node
  right node
}

type var_usage struct {
  in *variable
  name string
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
  obj_node node
  prop string
  in node
}

type create_prop struct {
  obj node
  string_prop string
  node_prop node
}

type get_prop struct {
  obj *variable
  obj_node node
  string_prop string
  node_prop node
}

type for_in_node struct {
  iterator *variable
  source node
  body node
}

type this_node struct { }

type new_node struct {
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
  if n.left_node == nil {
    Assign(right, n.left_var)
    return n.left_var
  } else {
    left_var := n.left_node.Interpret()
    Assign(right, left_var)
    return left_var
  }
}

func (n assign) AddChild(in node) { }

func (n *var_usage) Interpret() *variable {
  if n.name == "" {
    return n.in
  } else {
    return vars[n.name]
  }
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
  out.O = make(map[string]*variable)
  out.O["prototype"] = MakeEmptyObject()
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
  if n.obj_node == nil {
    n.obj.SetProp(n.prop, val)
  } else {
    obj := n.obj_node.Interpret()
    obj.SetProp(n.prop, val)
  }
  return nil
}

func (n set_prop) AddChild(in node) { }

func (n *create_prop) Interpret() *variable {
  obj := n.obj.Interpret()
  if n.node_prop == nil {
    return obj.CreateProp(n.string_prop)
  } else {
    return obj.CreateProp(StringCast(n.node_prop.Interpret()))
  }
}

func (n create_prop) AddChild(in node) { }

func (n *get_prop) Interpret() *variable {
  obj := n.obj_node.Interpret()
  if n.node_prop == nil {
    return obj.GetProp(n.string_prop)
  } else {
    return obj.GetProp(StringCast(n.node_prop.Interpret()))
  }
}

func (n get_prop) AddChild(in node) { }

func (n *for_in_node) Interpret() *variable {
  source := n.source.Interpret()
  for k, _ := range *source.GetProperties() {
    key_variable := new(variable)
    key_variable.SetStringValue(k)
    Assign(key_variable, n.iterator)
    n.body.Interpret()
  }
  return nil
}

func (n for_in_node) AddChild(in node) { }

func (n *this_node) Interpret() *variable {
  return context
}

func (n this_node) AddChild(in node) { }

func (n *new_node) Interpret() *variable {
  context = MakeEmptyObject()
  n.body.Interpret()
  return context
}

func (n new_node) AddChild(in node) { }
