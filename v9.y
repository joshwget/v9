%{
package main
import "strconv"

type variable struct {
  I int
  B bool
}

var vars map[string]*variable

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

command: PRINT '(' exp ')' { $$.n = &print{ $3.n } }
;
%%
