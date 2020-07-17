grammar Jenkins;


L_PAREN: '(';
R_PAREN: ')';
PIPELINE: 'pipeline';
STRING: '\'' ~'\''* '\'';
WHITESPACE: [ \r\n\t]+ -> skip;

start : (
    ref=library_ref
    | pipe=pipelineDecl
     )* EOF;

pipelineDecl:
    PIPELINE '{'
       agent=agentDecl?
       stages=stagesDecl?
    '}'
    ;

agentDecl
    : 'agent' agentName='any'
    ;

stagesDecl
    : 'stages' '{'
        stageDecl*
     '}'
    ;
stageDecl
    :         'stage' L_PAREN name=STRING R_PAREN '{'
                  'options' '{'
                      'timeout' '(' 'time' ':' duration='1' ',' 'unit' ':' unit=STRING ')'
                  '}'
                  'steps' '{'
                      'echo' STRING
                  '}'
              '}'
;

library_ref:
    '@Library' '(' lib=STRING ')' '_'
    ;

