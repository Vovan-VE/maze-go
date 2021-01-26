The Maze Project in Go
======================

[![Ready: 57%](https://img.shields.io/badge/ready-57%25-important)
![CLI: ready](https://img.shields.io/badge/-cli-success)
![generate: ready](https://img.shields.io/badge/-generate-success)
![solve: not ready](https://img.shields.io/badge/-solve-critical)
![export: ready](https://img.shields.io/badge/-export-success)
![import: not ready](https://img.shields.io/badge/-import-critical)
![validate: not ready](https://img.shields.io/badge/-validate-critical)
![test: ready](https://img.shields.io/badge/-test-success)][the-maze-project]

This is an implementation of [The Maze Project][the-maze-project] in Go.

```
$ git clone <repo-url>
$ cd <dir>
$ make
$ bin/maze
$ bin/maze gen -s19x7 -ftext -cwall=▓▓ -cin=in -cout=ex
```

The last command would output something like following:

```
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓              ▓▓  ▓▓                          ▓▓      ▓▓          ▓▓      ▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓  ▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓
in              ▓▓  ▓▓          ▓▓      ▓▓          ▓▓  ▓▓  ▓▓  ▓▓      ▓▓  ▓▓
▓▓  ▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓  ▓▓  ▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓  ▓▓  ▓▓▓▓▓▓  ▓▓
▓▓      ▓▓      ▓▓  ▓▓          ▓▓  ▓▓  ▓▓  ▓▓      ▓▓          ▓▓      ▓▓  ▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓  ▓▓  ▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓▓▓▓▓  ▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓  ▓▓▓▓▓▓
▓▓  ▓▓          ▓▓      ▓▓          ▓▓      ▓▓      ▓▓      ▓▓      ▓▓      ▓▓
▓▓  ▓▓  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ▓▓▓▓▓▓  ▓▓  ▓▓  ▓▓  ▓▓  ▓▓▓▓▓▓  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ▓▓
▓▓      ▓▓      ▓▓              ▓▓      ▓▓  ▓▓  ▓▓  ▓▓      ▓▓  ▓▓          ▓▓
▓▓  ▓▓▓▓▓▓  ▓▓  ▓▓  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓  ▓▓▓▓▓▓  ▓▓  ▓▓▓▓▓▓▓▓▓▓
▓▓  ▓▓      ▓▓      ▓▓  ▓▓      ▓▓              ▓▓      ▓▓      ▓▓          ex
▓▓  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓  ▓▓  ▓▓  ▓▓▓▓▓▓  ▓▓▓▓▓▓▓▓▓▓  ▓▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓
▓▓                          ▓▓      ▓▓      ▓▓              ▓▓          ▓▓  ▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
```

[the-maze-project]: https://github.com/Vovan-VE/maze--main
