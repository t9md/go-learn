package main

import (
	"fmt"
	"strings"
)

type Indent struct {
	unit int
	char string
	size int
}

func (i *Indent) String() string {
	return strings.Repeat(i.char, i.size)
}
func (i *Indent) Indent() {
	i.size = i.size + i.unit
}
func (i *Indent) UnIndent() {
	i.size = i.size - i.unit
}

func fact(n int, i *Indent) int {
	i.Indent()
	fmt.Printf("%s %d: called by value\n", i, n)
	if n == 0 {
		return 1
	}
	answer := fact(n-1, i)
	fmt.Printf("%s %d: returned from\n", i, n-1)
	i.UnIndent()
	return n * answer
}

func main() {
	fmt.Println(fact(7, &Indent{char: " ", unit: 2, size: 2}))
}

/*
=>
     7: called by value
       6: called by value
         5: called by value
           4: called by value
             3: called by value
               2: called by value
                 1: called by value
                   0: called by value
                   0: returned from
                 1: returned from
               2: returned from
             3: returned from
           4: returned from
         5: returned from
       6: returned from
5040
*/
