# Project 2: Shell Builtins

#Jaylen Neduelan
## Description

his project involves extending a simple shell by adding five shell builtins or shell-adjacent commands, rewritten in Go and integrated into the existing Go shell. The shell supports a variety of builtin commands from shells such as `sh`, `bash`, `csh`, `tcsh`, `ksh`, and `zsh`.

### Builtins Already Implemented
- `cd`
- `env`

### Additional Resources
- [Bourne Shell Builtins](https://www.gnu.org/software/bash/manual/html_node/Bourne-Shell-Builtins.html)
- [Bash Builtins](https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html)
- [Built-in csh and tcsh Commands](https://docstore.mik.ua/orelly/linux/lnut/ch08_09.htm)

## Getting Started

### Prerequisites
Ensure you have Go installed on your system (visit [Go's official site](https://golang.org/dl/) for download instructions). This project uses Go modules, so no additional GOPATH configuration is needed.


For this project we'll be adding commands to a simple shell. 

The shell is already written, but you will choose five (5) shell builtins (or shell-adjacent) commands to rewrite into Go, and integrate into the Go shell.

There are many builtins or shell-adjacent commands to pick from: 
[Bourne Shell Builtins](https://www.gnu.org/software/bash/manual/html_node/Bourne-Shell-Builtins.html), 
[Bash Builtins](https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html,), and 
[Built-in csh and tcsh Commands](https://docstore.mik.ua/orelly/linux/lnut/ch08_09.htm).

Feel free to pick from `sh`, `bash`, `csh`, `tcsh`, `ksh` or `zsh` builtins... or if you have something else in mind, ping me and we'll work it out.

As an example, two shell builtins have already been added to the package builtins:

- `cd`
- `env`


