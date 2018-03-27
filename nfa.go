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
func poregtonfa(pofix string) {
	//create a empty array of pointer nfa stack to nfa struct
	nfastack := []*nfa{}
	
	//Loop through the regular expression one rune at a time
	for _,r := range pofix {
		switch r {
		case '.':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			frag1.accept.edge1 = frag2.initial
			
			nfastack = append(nfastack,&nfa{initial: frag1.initial, accept:frag2.accept})

		case '|':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
			
			nfastack =append(nfastack,&nfa{initial: &initial, accept: &accept})

		case '*':
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack,&nfa{initial: &initial, accept: &accept})

		default:
			accept := state{}
			initial := state{symbol: r,edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	return nfastack[0]
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)

}