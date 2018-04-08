package main

import (
	"fmt"
)

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct {
	initial *state
	accept  *state
}

func poregtonfa(pofix string) *nfa {
	nfastack := []*nfa{}

	for _, r := range pofix {
		switch r {

		case '|':
			// Take from top of stack and store in frag2
			// Take the last thing on the stack, it's already on frag2
			// Repeat for frag1
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// New initial state that points to the initial of the frag at edge1
			// and points to the new initial state at edge2
			// Join the frag at edge1 to the new accept state
			// Join the frag at edge2 to the new accept state
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			// Append the new nfa accept state and initial state we created above to the nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '.':
			// Take the last thing off the stack and store in frag2
			// Take the last thing on the stack, it's already on frag2
			// Repeat for frag1
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// Join the frag at edge1 to the inital state at frag2
			frag1.accept.edge1 = frag2.initial

			// Append the new frag2 accept state
			// and frag1 initial state we created above to the nfastack
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

		case '?':
			// Take the last thing off the stack and store in frag1
			// Take the last thing on the stack, it's already on frag1
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// New initial state that points to the initial of the frag at edge1
			// and points to the new accept state at edge2
			// Join the frag at edge1 to the new accept state
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = &accept

			// Append the new nfa accept state and initial state we created above to the nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '*':
			// Take the last thing off the stack and store in frag1
			// Take the last thing on the stack, it's already on frag1
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// New initial state that points to the initial of the frag at edge1 and points to the new accept state at edge2
			// Join the frag at edge2 to the new accept state
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			// Append the new nfa accept state and initial state we created above to the nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		}

	}
	return nfastack[0]
}

// Converts infix to postfix
func intopost(infix string) string {
	specials := map[rune]int{'*': 10, '?': 10, '.': 9, '|': 8}

	pofix, s := []rune{}, []rune{}

	for _, r := range infix {
		switch {
		case r == '(':
			s = append(s, r)

		case r == ')':
			for s[len(s)-1] != '(' {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = s[:len(s)-1]

		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		default:
			pofix = append(pofix, r)
		}
	}

	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix)
}

// Gets current array
// Adds state s to it
func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	// Checks if s has e arrows from it and makes sure its not an accept state
	if s != a && s.symbol == 0 {

		// Follow first edge and call same function for state pointed to by edge 1
		l = addState(l, s.edge1, a)
		// Check if nil,if not do same thing
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	return l
}

// Figure out if po matches s
func pomatch(po string, s string) bool {
	ismatch := false

	// create nfa from regex
	ponfa := poregtonfa(po)

	// Keep track of current state and the next state
	current := []*state{}
	next := []*state{}

	// addState is passed the initial and accept state
	current = addState(current[:], ponfa.initial, ponfa.accept)

	// Generate next state from current repeatedly
	// Read s char at a time
	// Take all current states,check if labelled by char from s
	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		// Swap current for next and reset next to empty array
		current, next = next, []*state{}
	}

	// Loop through to see if any are the accept state
	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}

func main() {

	var userInput string
	var regexInput string

	fmt.Println("Enter String You Want to Test: ")
	fmt.Scan(&userInput)

	fmt.Println("Enter the Regex You want to Test the String Against: ")
	fmt.Scan(&regexInput)

	fmt.Println(pomatch(regexInput, userInput))
}
