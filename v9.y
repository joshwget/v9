%{
package main
import "strconv"
%}

%union {
  n node
  s string
}

%right '='
%left '.'
%nonassoc COMP_EQU COMP_NEQU COMP_SEQU COMP_SNEQU
%nonassoc COMP_LESS COMP_LTE COMP_GTR COMP_GTE
%left '+' '-'
%left '*' '/'
%left BOOL_AND
%left BOOL_OR

%token NUM
%token VAR
%token PRINT
%token ID
%token IF
%token WHILE
%token FOR
%token IN
%token TRUE
%token FALSE
%token FUNCTION
%token STRING
%token THIS
%token NEW

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
         | if_statement { $$ = $1 }
         | while_statement { $$ = $1 }
         | for_in_statement { $$ = $1 }
;

var_usage: ID { $$.n = &var_usage{ name: $1.s }; }
         | THIS { $$.n = new(this_node) }
;

var_declare: VAR ID '=' exp {
               if vars == nil {
                 vars = make(map[string]*variable)
               }
               vars[$2.s] = new(variable)
               $$.n = &assign{ left_var: vars[$2.s], right: $4.n };
             }
;

indexable: var_usage { $$ = $1 }
         | indexable '[' exp ']' {
             $$.n = &get_prop{ obj_node: $1.n, node_prop: $3.n }
           }
         | indexable '.' ID {
             $$.n = &get_prop{ obj_node: $1.n, string_prop: $3.s }
           }
;

lhs_okay: var_usage { $$ = $1 }
        | indexable '[' exp ']' {
            $$.n = &create_prop{ obj: $1.n, node_prop: $3.n }
          }
        | indexable '.' ID {
            $$.n = &create_prop{ obj: $1.n, string_prop: $3.s }
          }
;

function_dec: FUNCTION '(' ')' '{' statement_list '}' { $$.n = &function_declare{ $5.n } }
function_call_exp: indexable '(' ')' { $$.n = &function_call{ $1.n } }

exp: NUM         { i, _ := strconv.ParseFloat($1.s, 32); $$.n = NumberConstant(float32(i)); }
   | TRUE  { $$.n = TrueConstant() }
   | FALSE { $$.n = FalseConstant() }
   | STRING { $$.n = StringConstant($1.s[1:len($1.s) - 1]) }
   | '{' '}' { $$.n = MakeEmptyObjectNode() }
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
   | function_dec { $$ = $1 }
   | function_call_exp { $$ = $1 }
   | NEW function_call_exp { $$.n = &new_node{ $2.n } }

   | lhs_okay '=' exp {
       $$.n = &assign{ left_node: $1.n, right: $3.n };
     }
   | indexable { $$ = $1 }
;

command: PRINT '(' exp ')' { $$.n = &print{ $3.n } }
;

if_statement: IF '(' exp ')' '{' statement_list '}' { $$.n = &if_node{ $3.n, $6.n } }
while_statement: WHILE '(' exp ')' '{' statement_list '}' { $$.n = &while_node{ $3.n, $6.n } }
for_in_statement: FOR '(' VAR ID IN exp ')' '{' statement_list '}' {
                    v := new(variable)
                    vars[$4.s] = v
                    $$.n = &for_in_node{ v, $6.n, $9.n }
                  }

%%
