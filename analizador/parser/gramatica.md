# Gramática

- [Gramática](#gramática)
  - [Léxico](#léxico)
    - [Simbolos](#simbolos)
    - [Reservadas](#reservadas)
    - [RegEx](#regex)
    - [IGNORADOS](#ignorados)
  - [Sintáctico](#sintáctico)
    - [Expresiones](#expresiones)
    - [Imprimir](#imprimir)
    - [Método main](#método-main)
    - [Funciones nativas](#funciones-nativas)
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
  - [Código fuente](#código-fuente)

## Léxico

### Simbolos

- S_PUNTO: `.`
- S_COMA: `,`
- S_PTCOMA: `;`
- S_ASIGNAR: `=`
- S_AMP: `&`
- S_MATCH_OR: `|`
- S_MATCH_RET: `=>`
- S_MATCH_DEFAULT: `_`
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

### RegEx

- INT: `[0-9]+`
- FLOAT: `[0-9]+[.][0-9]+`
- CHAR: `''' ~['] '''`
- STRING: `'"' ~["]* '"'`
- ID: `[a-zA-Z_][a-zA-Z0-9_]*`

### IGNORADOS

- COMENTARIO: `('//' [^.]* [\n])`
- BLANCOS: `[ \\\r\t]+`

## Sintáctico

### Expresiones

```
exp -> Expresion
    : logica -> Expresion
    | relacional -> Expresion
    | aritmetica -> Expresion

logica -> Expresion
    : logica S_AND logica -> Expresion
    : logica S_OR logica -> Expresion
    | S_NOT logica -> Expresion
    | relacional -> Expresion

relacional -> Expresion
    : relacional (S_MENOR | S_MAYOR | S_MENORI | S_MAYORI | IGUAL | S_DIFERENTE) relacional -> Expresion
    | aritmetica -> Expresion

aritmetica -> Expresion
    : S_RESTA exp -> Expresion
    | aritmetica (S_POR | S_DIVISION | S_MODULO) aritmetica -> Expresion
    | aritmetica (S_SUMA | S_RESTA) aritmetica -> Expresion
    | S_APAR exp S_CPAR -> Expresion
    | exp_valor -> Expresion

exp_valor -> Expresion
    : primitivo -> Primitivo
    | llamada -> Expresion

primitivo -> Expresion
    : INT -> Primitivo
    | FLOAT -> Primitivo
    | STRING -> Primitivo
    | R_TRUE -> Primitivo
    | R_FALSE -> Primitivo
    | ID -> Identificador

llamada -> Expresion
    : ID S_APAR S_CPAR -> Llamada
    | ID S_APAR argumentos S_CPAR -> Llamada

argumentos -> List<Expresion>
    : argumentos S_COMA exp -> List<Expresion>
    | exp -> List<Expresion>
```

### Imprimir

```
imprimir -> R_PRINTLN S_APAR lista_exp S_CPAR ';'

lista_exp -> lista_exp ',' exp
           | exp
```

### Método main

```
main -> FN MAIN ( ) { sentencias }
```

### Funciones nativas

### Variables

#### Declaración

```
declaracion -> LET MUT ID : TIPO = exp
             | LET MUT ID = exp
             | LET ID : TIPO = exp
             | LET ID = exp
```

#### Asignación

```
asignacion -> ID = exp
```

### Sentencias de control

#### If

#### Switch

#### For

#### While

#### Do While

### Métodos y Funciones

#### Método

#### Función

### Llamadas

### Modificadores de flujo

#### Break

#### Continue

### Arreglos

#### Definición

#### Asignación

### Vectores

#### Asignación

#### Métodos

### Structs

### Módulos

## Código fuente

Entrada básica

```rust
fn main() {
  println!("+ -");
}
```

Entrada con structs

```rust
// STRUCT
struct Personaje {
  nombre: String,
  edad: i64,
  descripcion: String
}

// STRUCT
struct Carro {
  placa: String,
  color: String,
  tipo: String
}

fn main() {
  // CONSTRUCCIÓN STRUCT
  let mut p1 = Personaje { nombre: "Fer".to_string(), edad: 18 , descripcion: "No hace nada".to_string() };
  let mut p2 = Personaje { nombre: "Fer".to_string(), edad: 18 , descripcion: "Maneja un carro".to_string() };
  let mut c1 = Carro { placa: "090PLO".to_string(), color: "Gris".to_string(), tipo: "Mecanico".to_string() };
  let mut c2 = Carro { placa: "P0S921".to_string(), color: "Verde".to_string(), tipo: "Automatico".to_string() };

  // ASIGNACIÓN ATRIBUTOS
  p1.edad = 10;
  p2.edad = 20;
  c1.color = "Cafe".to_string();
  c2.color = "Rojo".to_string();

  // ACCESO ATRIBUTO
  println!("{}", p1.edad);
  println!("{}", c1.color);
}
```