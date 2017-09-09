[![Build Status](https://travis-ci.org/kaizen63/go-roman.svg?branch=master)](https://travis-ci.org/kaizen63/go-roman)
[![Coverage Status](https://coveralls.io/repos/github/kaizen63/go-roman/badge.svg?branch=master)](https://coveralls.io/github/kaizen63/go-roman?branch=master)

# go-roman

A small package to convert integers to roman numbers. The intersting bit is that the conversion into roman is not a mathematical solution. Its pure text processing.
Kudos to Kevlin Hennley for this solution (although it was not in go)
## Example
```
$> go run main.go 10
I
II
III
IV
V
VI
VII
VIII
IX
X
```