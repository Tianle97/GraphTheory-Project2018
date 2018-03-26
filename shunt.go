package main

import "fmt"

func inpost(infix string) string {
	specials := map[rune]int{'*':10,'.':9,'|':8}

	postfix := []rune{}
	s := []rune{}

	for _,r := range infix{
		switch {
		case r == '(' :
			s = append(s,r)

		case r == ')' :
			for s[len(s)-1] != '(' {
				postfix = append(postfix , s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = s[:len(s)-1]

		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				postfix = append(postfix , s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = append(s,r)

		default:
			postfix = append(postfix,r)
		}
	}

	for len(s) > 0 {
		postfix = append(postfix , s[len(s)-1])
		s = s[:len(s)-1]
	}

	return string(postfix)
}

func main(){
	//Answer ab.c*
	fmt.Println("Infix:  ","a.b.c*")
	fmt.Println("Postfix: ",inpost("a.b.c*")) 

	//Answer  abd.|
	fmt.Println("Infix: ","(a.(b|d))")
	fmt.Println("Postfix: ",inpost("(a.(b|d))"))

	//Answer  abd|.c*
	fmt.Println("Infix: ","a.(b|d).c*")
	fmt.Println("Postfix: ",inpost("a.(b|d).c*"))

	//Answer  abb.+.c.
	fmt.Println("Infix: ","a.(b.b)+.c")
	fmt.Println("Postfix: ",inpost("a.(b.b)+.c"))
}
