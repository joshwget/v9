%{
package main
import "fmt"
import "strconv"

type variable struct {
  I int
}

var vars map[string]*variable

type node interface {
  Interpret() *variable
  AddChild(in node)
}

type constant struct {
  val *variable
}

func IntConstant(i int) *constant {
  v := &variable{ i }
  return &constant { v }
}

type math2 struct {
  left, right node
  operator rune
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

func (n constant) Interpret() *variable {
  return &variable{ n.val.I }
}

func (n constant) AddChild(in node) { }

func (n math2) Interpret() *variable {
  left := n.left.Interpret()
  right := n.right.Interpret()
  switch(n.operator) {
    case '+':
      return &variable{ left.I + right.I }
    case '-':
      return &variable{ left.I - right.I }
    case '*':
      return &variable{ left.I * right.I }
    case '/':
      return &variable{ left.I / right.I }
    default:
      return &variable{ left.I + right.I }
  }
}

func (n math2) AddChild(in node) { }

func (n print) Interpret() *variable {
  fmt.Println(n.in.Interpret().I)
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

%}

%union {
  n node
  s string
}

%left '+' '-' '*' '/'

%token NUM
%token VAR
%token PRINT
%token ID

%%

program: statement_list {
           $1.n.Interpret();
         }
;

statement_list: { $$.n = &block{ make([]node, 0) }; }
                | statement_list statement {
                    $1.n.AddChild($2.n);
                    $$ = $1;
                  }
;

statement: exp ';'     { $$ = $1; }
         | command ';' { $$ = $1; }
         | var_declare ';' { $$ = $1; }
;

var_declare: VAR ID '=' exp {
               fmt.Println($2.s);
               vars = make(map[string]*variable)
               vars[$2.s] = new(variable)
               $$.n = &assign{ vars[$2.s], $4.n };
             }
;

exp: NUM         { i, _ := strconv.Atoi($1.s); $$.n = IntConstant(i); }
   | exp '+' exp { $$.n = &math2{ $1.n, $3.n, '+' }; }
   | exp '-' exp { $$.n = &math2{ $1.n, $3.n, '-' }; }
   | exp '*' exp { $$.n = &math2{ $1.n, $3.n, '*' }; }
   | exp '/' exp { $$.n = &math2{ $1.n, $3.n, '/' }; }
   | '(' exp ')' { $$ = $2; }
   | ID { $$.n = &var_usage{ vars[$1.s] }; }

command: PRINT exp { $$.n = &print{ $2.n } }
;
%%
