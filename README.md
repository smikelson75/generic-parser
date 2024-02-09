# Text Parser

Implements the ability to parse text from a file. 

## Tokenization

### Handlers

Handlers to tokenize text for parsing into an AST (abstract syntax tree).

| Handler | Task |
| ------- | ---- |
| Complex Identifier | Provides a specific token for an identifier that is surrounded by double-quotes. These will be marked as Identifiers.|
| Identifiers | Provides a token for recognizing indentifiers. In this version, identifiers are defined as a sequence of characters that start with a letter and are followed by zero or more letters, digits, or underscores. |
| Number  | Provides a token for recognizing a number (e.g. signed or unsigned integers, floats, etc.). |
| String  | Provides a token that recognizes a string (using single-quotes). | 
| Whitespace | Provides a token that recognizes whitespace being discovered in the text. |

### Tokens

Tokens are groups of text with a specific meaning. For example, the text "123" would be recognized as a Number token.

## Parser

### Nodes

Nodes are the building blocks of an AST. Nodes are created by the parser and are used to represent the structure of the text. 