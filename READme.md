This is our languages Grammar:

program    -> statement* EOF
statement  -> "let" IDENT "=" expr ";"
           | IDENT "=" expr ";"
           | "print" expr ";"
expr       -> prefix ( binop expr )*          // precedence decides the nesting
prefix     -> NUMBER | IDENT | STRING | "-" prefix | "(" expr ")"