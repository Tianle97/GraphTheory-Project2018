//Author: Tianle Shu
//GMIT Software Development Student

package shuntingPackage

//import "fmt"

//Inpost function to convert infix regular expression to postfix 
//(e.g. infix regex "a.b.c*" become to postfix trgrx "ab.c*." )
func Inpost(infix string) string {

	//create a map for procefence( * => . => | )
	specials := map[rune]int{'*':10,'.':9,'|':8}

	//two empty rune arrays postfix and s. 
	//Explain in Chinese about the rune array means and example :
	//http://www.cnblogs.com/howDo/archive/2013/04/20/GoLang-String.html
	postfix := []rune{}
	s := []rune{}

	//range loop converts string to char array (UTF-8)
	// Using range in our loop on our string will cause it to be converted to a rune array
	for _,r := range infix{

		// Switch statement for this Shunt algorithm
		switch {

			//character cases
			case r == '(' :
				s = append(s,r)


			case r == ')' :
				// While the top element of s is not '('
				for s[len(s)-1] != '(' {
					//Append the top element of the s to the postfix of rune array
					postfix = append(postfix , s[len(s)-1])
					//Remove the top element from the s
					s = s[:len(s)-1]
				}
				//pops off next character after brackets
				s = s[:len(s)-1]


			// When the number greater than 0 is returned 
			//(e.g. 'r' is not contained in our special map, append rune array with 'r')
			case specials[r] > 0:
				// While the s still has elements in and while the precedence of the current 
				// character being read is less than the precedence of the character at the top of the s
				for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
					//Append the top element of the s to the postfix of rune array
					postfix = append(postfix , s[len(s)-1])
					//Remove the top element from the s
					s = s[:len(s)-1]
				}
				// Append the characters being read in to the s
				s = append(s,r)

			default:
				postfix = append(postfix,r)
			}
	}

	for len(s) > 0 {

		// Append the top element of the s to postfix rune array
		postfix = append(postfix , s[len(s)-1])
		// Remove the top element from the s
		s = s[:len(s)-1]
	}
	// Return the postfix of rune array and converted into a string
	return string(postfix)
}

func Infix (s string) string {
	if len(s) > 0 {
		s = s[:len(s)-2]
	}
	return s
}
//Main function output result.
/*func main() {
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
}*/