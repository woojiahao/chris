# chris
Pratt parser implementation in Go

The core implementation details follows the advice by
Bob Nystrom detailed in his [article on Pratt parsing](http://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/)

## Architecture

```marmaid
flowchart LR
    Lexer-->Parser-->Compiler
```

Lexer acts as an iterator over a given expression and converts each character/word into a given token. It ignores 
whitespaces and will parse numbers and words as a whole chunk.

Parser reads the token stream from a given Lexer and applies grammar to the tokens to generate an AST tree

Compiler receives the generated AST tree from the Parser and performs operations on the given AST tree and the respective
nodes.