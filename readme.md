# Cairn

**Cairn** is a stack-based personal programming language, written by [Stephen Malone][sm] in [Go 1.22][go]. It's designed to be a fun, retro programming environment for my own hobbyist coding interests.

- See [`changes.md`][ch] for the complete changelog.
- See [`license.md`][li] for the open-source license (BSD-3).

## Syntax

Cairn has an extremely simple syntax with only three forms: **comments**, **numbers** and **symbols**.

- **Comments** start with `//` and exclude the remaining line.
- **Numbers** are unsigned eight-bit integers from 0 to 255.
- **Symbols** are references to built-in or user-defined functions.

By convention, all Cairn code is upper-case. Built-in functions are all three letters, but user-defined functions can be any length.

## Machine

### Memory

Cairn operates inside a fantasy virtual machine with two memory types: **registers** and the **stack**.

- **Registers** are fixed variables that can each store one integer.
- The **stack** is a last-in-first-out stack of stored integers.

There are 8 registers (named `R0` to `R7`) and the stack can hold up to 65,536 integers.

### Input / Output

Input and output are handled [Brainfuck][bf]-style with a single stream each for input and output. By default these are `STDIN` and `STDOUT` but they can be overridden with specified files.

### Logic

Zero (`0`) integers are considered **false**, all other integers are **true**. If a command returns a boolean value, it will always return zero (`0`) for false and one (`1`) for true.

## Commands

**Commands** are built-in functions with callable symbols. Each command operates on the stack, *popping* arguments off and *pushing* their results back on.

In the following tables, the "form" column shows each command's effect. The items before the arrow show the stack *before* the command is executed (with the rightmost item at the top) and the items after show the stack *after* execution.

An underscore (`_`) indicates a command does not take or return arguments.
An ellipsis (`...`) indicates a variable number of arguments.

### Integer Commands

Name  | Form      | Description
----- | --------- | -----------
`ADD` | `a b → c` | Return `a` + `b`.
`SUB` | `a b → c` | Return `a` - `b`.
`MOD` | `a b → c` | Return `a` % `b`.
`GTE` | `a b → c` | Return `a` >= `b`.
`LTE` | `a b → c` | Return `a` <= `b`.

### Memory Commands

Name  | Form        | Description
----- | ----------- | -----------
`CLR` | `... → _`   | Clear the stack.
`DUP` | `a → a a`   | Duplicate the top item.
`DRP` | `a b → a`   | Delete the top item.
`SWP` | `a b → b a` | Swap the top two items.
`GET` | `a → b`     | Return the value of register `a`.
`SET` | `a b → _`   | Set the value `a` to register `b`.

### Logic Commands

Name  | Form      | Description
----- | --------- | -----------
`EQU` | `a b → c` | Return `1` if `a` equals `b`.
`NEQ` | `a b → c` | Return `1` if `a` does not equal `b`.
`AND` | `a b → c` | Return `1` if both `a` and `b` are true.
`ORR` | `a b → c` | Return `1` if either `a` or `b` are true.
`XOR` | `a b → c` | Return `1` if only `a` or only `b` is true.
`NOT` | `a → b`   | Return `1` if `a` is false, `0` if true.

### Input / Output Commands

Name  | Form       | Description
----- | ---------- | -----------
`INN` | `_ → a`    | Return an input ASCII character as an integer.
`OUT` | `a → _`    | Write `a` as an ASCII character to output.
`NUM` | `a → a...` | Return the ASCII character(s) of the integer `a`.

### Flow Control Commands

These commands are special as they wrap smaller pieces of code and execute them according to specific conditions. Each flow command must end with the symbol `END` after the arguments.

#### `IFT [CODE] END` · `a → _`

Execute `[CODE]` if the top item is true.

#### `IFF [CODE] END` · `a → _`

Execute `[CODE]` if the top item is false.

#### `FOR [REG] [CODE] END` · `_ → _`

Execute `[CODE]` in a continuous loop until the value of register `[REG]` (0 to 7) is false.

#### `DEF [NAME] [CODE] END` · `_ → _`

Set the user-defined function `[NAME]` to `[CODE]`. After executing this command, the symbol `[NAME]` will always execute `[CODE]`. Functions can be re-defined at any time.

## Contributing

Please add all bug reports and feature requests to the [issue tracker][is], thank you.

[bf]: https://esolangs.org/wiki/Brainfuck
[ch]: https://github.com/wirehaiku/cairn/blob/main/changes.md
[go]: https://golang.org/doc/go1.22
[is]: https://github.com/wirehaiku/cairn/issues
[li]: https://github.com/wirehaiku/cairn/blob/main/license.md
[sm]: https://mastodon.social/@stvmln
