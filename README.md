# chris

Pratt parser implementation in Go

The core implementation details follows the advice by
Bob Nystrom detailed in
his [article on Pratt parsing](http://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/)

My notes on Pratt parsing can be
found [here.](https://woojiahao.notion.site/Pratt-Parsing-Notes-a3ccdbc32a424be6bcf67f52769ebd94)

## Architecture

```mermaid
flowchart LR
    Lexer-->Parser-->Compiler
```

Lexer acts as an iterator over a given expression and converts each character/word into a given token. It ignores
whitespaces and will parse numbers and words as a whole chunk.

Parser reads the token stream from a given Lexer and applies grammar to the tokens to generate an AST tree. It is not
responsible for checking if the keywords are valid. It just needs to know that the expression can generate a valid
AST tree.

Compiler receives the generated AST tree from the Parser and performs operations on the given AST tree and the
respective nodes. `chris`, however is not a compiler, but a parser, so it will not compile the given AST.

### Parser

Parser logic is performed by something known as "Parselets". Effectively, they are the components that handles behavior
of each token. This is slightly different to having functions per non-terminal character in our grammar.

## Sample

```text
1 + 2 * 3               := 1 + (2 * 3)
sin(pi/4)               := sin((pi/4))
2^x + cos(pi/4 + 15)    := (2^x) + cos(((pi/4) + 15))
```

### Operators/Symbols

| Symbol     | Purpose                                                     | Position     | Precedence |
|------------|-------------------------------------------------------------|--------------|------------|
| +          | Addition                                                    | Infix        | 2          |
| -          | Subtraction                                                 | Prefix/Infix | 2          |
| *          | Multiplication                                              | Infix        | 3          |
| /          | Division                                                    | Infix        | 3          |
| ^          | Exponent                                                    | Infix        | 4          |
| (          | Create sub-expression or encapsulate a function's arguments | Prefix       | 5          |
| )          | End sub-expression                                          | -            | -1         |
| =          | Assignment                                                  | Infix        | 1          |
| <keyword>  | Keyword that corresponds to a function                      | Infix        | -1         |
| <number>   | Number                                                      | Prefix       | -1         |
| <Variable> | Single character to represent a variable                    | Prefix       | -1         |

### Keywords

```text
sin, cos, tan, sec, csc, cot, pi
```