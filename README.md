# GraphTheory-Project2018
3rd year Graph Theory Project-2018
#### Lecturer: *Ian McLoughlin*
#### Student Name: *Tianle Shu*
#### Student ID: *G00353418*
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
#### Code Explain <br/>
` "." `: <br/>
Catenation:
```Golang
case '.':
			//pop off 2 frags
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// set  frag1 accept = frag2 initial State
			frag1.accept.edge1 = frag2.initial
			
			//appending a new concatenate frag to stack
			nfastack = append(nfastack,&nfa{initial: frag1.initial, accept:frag2.accept})
```
<br/>

` "*" `:  <br/>
Zero or More:
``` Golang
case '*':
			//pop off 1 frag
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			//create a new state called accept
			accept := state{}
			//creating new state called initial,setting one edge to frag and a accept state
			initial := state{edge1: frag.initial, edge2: &accept}
			//setting 2 frag accept state arrows back to initial of frag
			//and to the new accept state
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			//appending a new concatenate frag to stack
			nfastack = append(nfastack,&nfa{initial: &initial, accept: &accept})
```
<br/>

` "|" `: <br/>
Alternation:
```Golang
case '|':
			//pop off 2 frags
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			//creating new state called initial , both edges pinter to initial state
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			//creating new empty accept state
			accept := state{}
			//using accept states of frag1 and frag2 to new accept state
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
			//appending a new concatenate frag to stack
			nfastack =append(nfastack,&nfa{initial: &initial, accept: &accept})
```
<br/>

` "+" `: <br/>
One or More:
```Golang
//One or more:
		case '+':
			//pop off 1 frag
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			//create a new state called accept
			accept := state{}
			//creating new state called initial,setting one edge to frag and a accept state
			initial := state{edge1: frag.initial, edge2: &accept}
			//setting a frag accept state arrow back to initial of frag
			frag.accept.edge1 = &initial
			//appending a new concatenate frag to stack
			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})
```
<br/>

` "?" `: <br/>
Zero or More:
```Golang
//One or more:
		case '?':
			//pop off 1 frag
			frag := nfastack[len(nfastack)-1]
			//creating new state called initial,setting one edge to frag and a accept state
			initial := state{edge1: frag.initial, edge2: frag.accept}
			//setting a frag accept state arrow back to initial of frag
			frag.accept.edge1 = &initial
			//appending a new concatenate frag to stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})
```

## *Reference*
[Golang beginner](https://tour.go-zh.org/welcome/1)  <br/>
Better to understand NFA: https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton <br/>
Better to understand Shunting Yard Algorithm: https://en.wikipedia.org/wiki/Shunting-yard_algorithm <br/>
Better to understand the some regexp code and the explain: https://swtch.com/~rsc/regexp/regexp1.html <br/>
Explain rune in Golang(Chinese Language): https://golangtc.com/t/528cc004320b52227200000f <br/>
Regex Tester:http://rextester.com/tester <br/>
