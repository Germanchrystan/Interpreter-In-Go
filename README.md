# Interpreter in GO
## Mate programming language


## 2.Parsers

According to Wikipedia

`
A parser is a software component that takes input data (frequently text) and builds a data structure – often some kind of parse tree, abstract syntax tree or other hierarchical structure – giving a structural representation of the input, checking for correct syntax in the process. [...] The parser is often preceded by a separate lexical analyser, which creates tokens from the sequence of input characters;
`

A parser turns its input unto a data structure that represents the input.

A JSON parser takes text as input and builds a data structure that represents the input. That’s exactly what the parser of a programming language does. The difference is that in the case of a JSON parser you can see the data structure when looking at the input. Whereas if you look at this
if ((5 + 2 * 3) == 91) { return computeStuff(input1, input2); }
it’s not immediately obvious how this could be represented with a data structure.

In most interpreters and compilers the data structure used for the internal representation of
the source code is called a “syntax tree” or an “abstract syntax tree” (AST for short). The “abstract” is based on the fact that certain details visible in the source code are omitted in the AST. Semicolons, newlines, whitespace, comments, braces, bracket and parentheses – depending on the language and the parser these details are not represented in the AST, but merely guide
the parser when constructing it.