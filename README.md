# GraphTheory-Project2018
3rd year Graph Theory Project-2018
### Lecturer: *Ian McLoughlin*
### Student Name: *Tianle Shu*
### Student ID: *G00353418*
![Golang](https://scriptcrunch.com/wp-content/uploads/2017/12/golang.jpg)

## *Introduction*
Write a program in the Go programming language [2] that can build a non-deterministic finite automaton (NFA) from a regular expression, and can use the NFA to check if the regular expression matches any given string of text. You must write the program from scratch and cannot use the regexp package from the Go standard library nor any other external library. <br/>
A regular expression is a string containing a series of characters, some of which may have a special meaning. For example, the three characters “.”, “|”, and “∗” have the special meanings “concatenate”, “or”, and “Kleene star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1, and 1∗ means any number of 1’s. These special characters must be used in your submission.

## *Shunting Yard Algorithm*
####  Explain <br/>
The shunting yard algorithm is a simple technique for parsing infix expressions containing binary operators of varying precedence. In general, the algorithm assigns to each operator its correct operands, taking into account the order of precedence. It can, therefore, be used to evaluate the expression immediately, to convert it into postfix, or to construct the corresponding syntax tree. <br/>
The shunting yard algorithm is not your basic algorithm like mergesort, string search, etc. It is quite advanced as stacks, queues, and arrays are all contained in the same algorithm. Although the algorithm itself is very simple, a solid flexible implementation might be thousands of lines of code.
<br/>

## *Nondeterministic Finite Automaton(NFA)*
####  Explain <br/>

Thompson's construction is an algorithm for transforming a regular expression into an equivalent nondeterministic finite automaton (NFA). This NFA can be used to match strings against the regular expression. <br/>
Regular expressions and nondeterministic finite automata are two representations of formal languages. For instance, text processing utilities use regular expressions to describe advanced search patterns, but NFAs are better suited for execution on a computer. Hence, this algorithm is of practical interest, since it can compile regular expressions into NFAs. From a theoretical point of view, this algorithm is a part of the proof that they both accept exactly the same languages, that is, the regular languages.
<br/>
## *Reference*
[Golang beginner](https://tour.go-zh.org/welcome/1)  <br/>
Better to understand NFA: https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton <br/>

Better to understand  Shunting Yard Algorithm: https://en.wikipedia.org/wiki/Shunting-yard_algorithm <br/>
Explain rune in Golang(Chinese Language): https://golangtc.com/t/528cc004320b52227200000f <br/>
