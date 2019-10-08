# Wireworld

This repo is for simulating wireworld in different languages

## The game

The [wiki page](https://en.wikipedia.org/wiki/Wireworld) has good examples and a list of all rules.

### Rules

For every tick each cell transfers from the left to the right.

* empty -> empty
* electron head -> electron tail
* electron tail -> conductor
* conductor -> electron head if exactly one or two of the neighbouring cells are electron heads, otherwise it remains a conductor

"Neighbors" use Moor neighborhood so diagnols count.

### Input

Each program will take a .wir input file. These files show the starting "grid" for a wireworld circuit.

* Newlines "\n" indicate a new row (we aren't windows friendly here)
* 0 is an empty cell
* 1 is an electron head
* 2 is an electron tail
* 3 is a conductor

For the sake of clarity, if/when a program prints images of the simulation the following color scheme should be used:
* 0 - black
* 1 - blue
* 2 - red
* 3 - yellow

## Running

### Go

1. Follow the [install guide](https://golang.org/doc/install)
2. ```cd Wireworld/go/```
3. ```export GOPATH=/path/to/Wireworld/go```
5. ```go run src/wireworld.go --input_path=../circuits/one-charge.wir```
