//Author: Tianle Shu
//GMIT Software Development Student

package main

import "fmt"

// state struct
//it has 2 edges in which are pointers to others 
type state struct {
	symbol rune
	edge1 *state
	edge2 *state
}

//nfa struct
//it has 2 pointers. One called initial and the second called accept
//2 pointers to state
type nfa struct {
	//these 2 pointers is that would be the initial state 
	//of some sort of linked list, linked data struct here
	initial *state
	accept  *state
}

//poregtonfa function
// that means change the postfix regular expression to an NFA(Non-finite automaton)
// And return a pointer to an NFA struct
func poregtonfa(pofix string) *nfa{
	//create a empty array of pointer nfa stack to nfa struct
	nfastack := []*nfa{}
	
	// through the regular expression one rune at a time
	for _,r := range pofix {
		switch r {
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

		default:
			//create a new state called accept
			accept := state{}
			//creating new initial state 
			//and setting the symbol and use edge1
			initial := state{symbol: r,edge1: &accept}
			//appending a new concatenate frag to stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}
	//returning fully concatenated nfa
	return nfastack[0]
}

//Main function, output result
func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}