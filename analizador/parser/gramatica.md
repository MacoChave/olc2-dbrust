# Gramática

- [Gramática](#gramática)
  - [Léxico](#léxico)
    - [Reservadas](#reservadas)
    - [Simbolos](#simbolos)
    - [RegEx](#regex)
  - [IGNORADOS](#ignorados)
  - [Expresiones](#expresiones)
  - [Funciones nativas](#funciones-nativas)
    - [Imprimir](#imprimir)
  - [Variables](#variables)
    - [Declaración](#declaración)
    - [Asignación](#asignación)
  - [Sentencias de control](#sentencias-de-control)
    - [If](#if)
    - [Switch](#switch)
    - [For](#for)
    - [While](#while)
    - [Do While](#do-while)
  - [Métodos y Funciones](#métodos-y-funciones)
    - [Método](#método)
    - [Función](#función)
  - [Llamadas](#llamadas)
  - [Modificadores de flujo](#modificadores-de-flujo)
    - [Break](#break)
    - [Continue](#continue)
  - [Arreglos](#arreglos)
    - [Definición](#definición)
    - [Asignación](#asignación-1)
  - [Vectores](#vectores)
    - [Asignación](#asignación-2)
    - [Métodos](#métodos)
  - [Structs](#structs)
  - [Módulos](#módulos)

## Léxico

### Reservadas

- R_TRUE: `true`
- R_FALSE: `false`
- R_LET: `let`
- R_MUT: `mut`
- R_INT: `i64`
- R_FLOAT: `f64`
- R_BOOL: `bool`
- R_CHAR: `char`
- R_STRING: `string`
- R_STR: `&str`
- R_FN: `fn`
- R_AS: `as`
- R_TO_OWNED: `to_owned`
- R_TO_STRING: `to_string`
- R_POW: `pow`
- R_POWF: `powF`
- R_PRINTLN: `println!`
- R_ABS: `abs`
- R_SQRT: `sqrt`
- R_TO_STRING: `to_string`
- R_CLONE: `clone`
- R_NEW: `new`
- R_LEN: `len`
- R_PUSH: `push`
- R_REMOVE: `remove`
- R_CONTAINS: `contains`
- R_INSERT: `insert`
- R_CAPACITY: `capacity`
- R_WITH_CAPACITY: `with_capacity`
- R_IF: `if`
- R_ELSE: `else`
- R_MATCH: `match`
- R_LOOP: `loop`
- R_BREAK: `break`
- R_WHILE: `while`
- R_FOR: `for`
- R_IN: `in`
- R_CONTINUE: `continue`
- R_RETURN: `return`
- R_STRUCT: `struct`
- R_VECTOR: `vector`
- R_VEC: `vec`
- R_MOD: `mod`
- R_PUB: `pub`
- R_: ``

### Simbolos

- S_PUNTO: `.`
- S_COMA: `,`
- S_PTCOMA: `;`
- S_ASIGNAR: `=`
- S_AMP: `&`
- S_MATCH_OR: `|`
- S_MARCH_RET: `=>`
- S_MATCH_DEF: `_`
- S_SUMA: `+`
- S_RESTA: `-`
- S_POR: `*`
- S_DIVISION: `/`
- S_MODULO: `%`
- S_MAYOR: `>`
- S_MENOR: `<`
- S_MAYORI: `>=`
- S_MENORI: `<=`
- S_IGUAL: `==`
- S_DIFERENTE: `!=`
- S_OR: `||`
- S_AND: `&&`
- S_NOT: `!`
- S_MOD_RET: `->`
- S_APAR: `(`
- S_CPAR: `)`
- S_DOSPT: `:`
- S_ALIAS: `::`
- S_ACOR: `[`
- S_CCOR: `]`
- S_ALLAV: `{`
- S_CLLAV: `}`
- S_RANGO: `..`
- S_: ``
- S_: ``
- S_: ``

### RegEx

- ID: `[a-zA-Z_][a-zA-Z0-9_]*`
- INT: `[0-9]+`
- FLOAT: `[0-9]+[.][0-9]+`
- CHAR: `''' ~['] '''`
- STRING: `'"' ~["]* '"'`

## IGNORADOS

- COMENTARIO: `('//' [^.]* [\n])`
- BLANCOS: `[ \\\r\t]+`

## Expresiones

```antlr
exp
    : logica
    | relacional
    | aritmetica

logica
    : logica S_AND logica
    : logica S_OR logica
    | S_NOT logica
    | RELACIONAL

relacional
    : relacional (S_MENOR | S_MAYOR | S_MENORI | S_MAYORI | IGUAL | S_DIFERENTE) relacional
    | aritmetica

aritmetica
    : S_RESTA exp
    | aritmetica (S_POR | S_DIVISION | S_MODULO) aritmetica
    | aritmetica (S_SUMA | S_RESTA) aritmetica
    | S_APAR exp S_CPAR
    | exp_valor

exp_valor
    : primitivo
    | llamada

primitivo
    : INT
    | FLOAT
    | STRING
    | R_TRUE
    | R_FALSE
    | ID

llamada
    : ID S_APAR S_CPAR
    | ID S_APAR argumentos S_CPAR

argumentos
    : argumentos S_COMA exp
    | exp
```

## Funciones nativas

### Imprimir

## Variables

### Declaración

### Asignación

## Sentencias de control

### If

### Switch

### For

### While

### Do While

## Métodos y Funciones

### Método

### Función

## Llamadas

## Modificadores de flujo

### Break

### Continue

## Arreglos

### Definición

### Asignación

## Vectores

### Asignación

### Métodos

## Structs

## Módulos