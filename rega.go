//Author: Tianle Shu
//GMIT Software Development Student

package main

import (
	"fmt"
	shunt "./shuntingPackage"
	"bufio"
	"os"
)

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

		//Catenation:
		// `.` means : first followed by next. (e.g. r.s – the concatenation of the languages defined by regular expressions r and s.)
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

		// Alternation:
		// `|` means: or. ( a|b = an 'a' or a 'b'.)
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

		// Zero or more:
		// `*` means:  zero or more times.(a∗ = zero or more a’s.)
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
		
		//Zero or one:
		case '?':
			//pop off 1 frag
			frag := nfastack[len(nfastack)-1]
			//creating new state called initial,setting one edge to frag and a accept state
			initial := state{edge1: frag.initial, edge2: frag.accept}
			//setting a frag accept state arrow back to initial of frag
			frag.accept.edge1 = &initial
			//appending a new concatenate frag to stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})
			
		//Literal characters:
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

	// must handle and debug any issues
	if len(nfastack) != 1{
		fmt.Println("Uh oh: ",len(nfastack),nfastack)
	}
	return nfastack[0]
}

//addState gets the current statewhile checking 
//if there is a arrow and follows 
//and the arrow  to get next states
func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	// s.symbol == 0 means state has an arrow from it
	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)

		// if there is another edge it must be added
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
}

//pomatch compares a pofix expression to a string and returns either true or false
func pomatch(po string, s string) bool {

	// default isMatch to false
	ismatch := false
	// create nfa from regular expression
	ponfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	// add initial state of pofix nfa and all available states
	current = addState(current[:],ponfa.initial, ponfa.accept)

	// loop through the user entered the character string 
	for _, r:= range s {
		// check all current states
		for _, c := range current {
			// if current state is set to the rune  
			// add that state and other state that can be accepted
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		// swap current make the next the new current for next iteraction
		current = next
		next = []*state{}
	}
	// iterate through the current states that means for check 
	// if they are accepted from nfa
	for _, c := range current {

		//if current state == accept state of nfa
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}
	return ismatch
}

// Main function for check the algorithm and output the result
func main() {
	a := 0
	var b string
	var c string
	fmt.Println("Please input choose the number:\n 1.Check Shunting yard algorithm, type '1'.\n 2.Check Postfix regular expression, type '2'.\n 3.if you want to exit just type '3'. ")
	fmt.Scanln(&a)

	//if user type '1' , it will check shunting yard algorithm
	//if user type '2' , it will check nfa algorithm
	//if user type '3' , it will exit the program
	if a == 1 {
		fmt.Println("please enter a Infix expression: ")
		//create a new pointer to bufio.Reader.
		//This line of code will create a reader and bind it with standard input.
		r := bufio.NewReader(os.Stdin)
		expression, _ := r.ReadString('\n') 
		expression  =  shunt.Infix(expression)
		fmt.Println("Infix: ", expression)
		fmt.Println("Postfix: ", shunt.Inpost(expression))
	} else if a == 2 {
		fmt.Println("Please enter a regular string : ")
		fmt.Scanln(&b)
		fmt.Println("Please enter a match the string: ")
		fmt.Scanln(&c)
		//nfa := poregtonfa(b)
		pomatch := pomatch(b,c)
		//fmt.Println(pomatch)
		if pomatch == true {
		fmt.Println("match seccessful.")
	}else {
		fmt.Println("match failure.")
	}
	} else if a == 3 {
		fmt.Println("you exit this program~\n")
	} else {
		fmt.Println("Please enter a correct String~")
	}
	/*nfa := poregtonfa("ab.c*|")
	pomatch := pomatch("ab.c*|","abc")
	fmt.Println(pomatch)
	fmt.Println(nfa)*/
}