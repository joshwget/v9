%{
package main
import "strconv"

type variable struct {
  Type int
  I int
  B bool
}

func (v *variable) SetNumberValue(i int) {
  v.I = i;
  v.Type = 0;
}

func (v *variable) SetBoolValue(b bool) {
  v.B = b;
  v.Type = 1;
}

var vars map[string]*variable

%}

%union {
  n node
  s string
}

%nonassoc COMP_EQU COMP_NEQU COMP_SEQU COMP_SNEQU
%nonassoc COMP_LESS COMP_LTE COMP_GTR COMP_GTE
%left '+' '-' '*' '/'
%left BOOL_AND
%left BOOL_OR

%token NUM
%token VAR
%token PRINT
%token ID
%token IF
%token WHILE

%token COMP_EQU;
%token COMP_NEQU;
%token COMP_LESS;
%token COMP_LTE;
%token COMP_GTR;
%token COMP_GTE;
%token COMP_SEQU;
%token COMP_SNEQU;

%token BOOL_AND;
%token BOOL_OR;

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
         | var_assign ';' { $$ = $1 }
         | if_statement { $$ = $1 }
         | while_statement { $$ = $1 }
;

var_declare: VAR ID '=' exp {
               vars = make(map[string]*variable)
               vars[$2.s] = new(variable)
               $$.n = &assign{ vars[$2.s], $4.n };
             }
;

var_assign: ID '=' exp {
              $$.n = &assign{ vars[$1.s], $3.n };
            }
;

exp: NUM         { i, _ := strconv.Atoi($1.s); $$.n = IntConstant(i); }
   | exp '+' exp { $$.n = &operation2{ $1.n, $3.n, '+' }; }
   | exp '-' exp { $$.n = &operation2{ $1.n, $3.n, '-' }; }
   | exp '*' exp { $$.n = &operation2{ $1.n, $3.n, '*' }; }
   | exp '/' exp { $$.n = &operation2{ $1.n, $3.n, '/' }; }

   | exp COMP_EQU exp { $$.n = &operation2{ $1.n, $3.n, COMP_EQU }; }
   | exp COMP_NEQU exp { $$.n = &operation2{ $1.n, $3.n, COMP_NEQU }; }
   | exp COMP_LESS exp { $$.n = &operation2{ $1.n, $3.n, COMP_LESS }; }
   | exp COMP_LTE exp { $$.n = &operation2{ $1.n, $3.n, COMP_LTE }; }
   | exp COMP_GTR exp { $$.n = &operation2{ $1.n, $3.n, COMP_GTR }; }
   | exp COMP_GTE exp { $$.n = &operation2{ $1.n, $3.n, COMP_GTE }; }
   | exp COMP_SEQU exp { $$.n = &operation2{ $1.n, $3.n, COMP_SEQU }; }
   | exp COMP_SNEQU exp { $$.n = &operation2{ $1.n, $3.n, COMP_SNEQU }; }

   | exp BOOL_AND exp { $$.n = &operation2{ $1.n, $3.n, BOOL_AND }; }
   | exp BOOL_OR exp { $$.n = &operation2{ $1.n, $3.n, BOOL_OR }; }

   | '(' exp ')' { $$ = $2; }
   | ID { $$.n = &var_usage{ vars[$1.s] }; }
;

command: PRINT '(' exp ')' { $$.n = &print{ $3.n } }
;

if_statement: IF '(' exp ')' '{' statement_list '}' { $$.n = &if_node{ $3.n, $6.n } }
while_statement: WHILE '(' exp ')' '{' statement_list '}' { $$.n = &while_node{ $3.n, $6.n } }

%%
