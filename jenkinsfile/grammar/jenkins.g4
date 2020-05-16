grammar jenkins;


STRING: '\'' ~'\''* '\'';
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : ref=library_ref EOF;

library_ref:
    '@Library' '(' lib=STRING ')' '_'
    ;

