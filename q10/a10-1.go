package q10

import (
	"fmt"
	"runtime"
	"sync"
)

type starsToken struct {
	letters string
}

func (t *starsToken) tokenName() string {
	return "star"
}

func (t *starsToken) token() string {
	return t.letters
}
func (t *starsToken) matchAll(s *string, left, right int) (matched bool, matchedLen int) {
	fmt.Println("stars matchAll", *s, left, right)
	if left > right || len(*s) == 0 {
		return true, 0
	}
	si := left
	for _, x := range []byte(t.letters) {
		for si < right+1 && (*s)[si] == x {
			si++
		}
		if si == right+1 || (*s)[si] != x {
			continue
		}
		si++
	}
	if si < right+1 {
		return false, 0
	}
	return true, right - left + 1
}
func (t *starsToken) findFirst(s *string, left int) (found bool, index int) {
	panic("implement me")
}

type dotStarToken struct{}

func (t *dotStarToken) tokenName() string {
	return "dotStar"
}

func (t *dotStarToken) token() string {
	return "."
}
func (t *dotStarToken) matchAll(s *string, left, right int) (bool, int) {
	return true, right - left + 1
}
func (t *dotStarToken) findFirst(s *string, left int) (found bool, index int) {
	panic("implement me")
}

type plainToken struct { // certain string
	pattern string
}

func (t *plainToken) tokenName() string {
	return "plain"
}
func (t *plainToken) token() string {
	return t.pattern
}
func (t *plainToken) matchAll(s *string, left, right int) (bool, int) {
	fmt.Println("matchAll plain, ", *s, left, right)
	if left > right || right-left+1 != len(t.pattern) || right >= len(*s) || left >= len(*s) || left < 0 || right < 0 {
		return false, 0
	}
	ti := 0
	for i := left; i < right+1; i++ {
		if ti > len(t.pattern)-1 {
			return false, 0
		}
		fmt.Println("matchAll plain, ", *s, len(*s), t.pattern, ti, i, left, right)
		if t.pattern[ti] != (*s)[i] && t.pattern[ti] != '.' {
			return false, 0
		}
		ti++
	}
	return true, right - left + 1
}
func (t *plainToken) findFirst(s *string, left int) (bool, int) {
	fmt.Println("findFirst, left", left)
	tokenIndex := 0
	for i := left; i < len(*s); i++ {
		if t.pattern[tokenIndex] == (*s)[i] || t.pattern[tokenIndex] == '.' {
			if tokenIndex == len(t.pattern)-1 {
				return true, i
			}
			tokenIndex++
		} else {
			tokenIndex = 0
			if t.pattern[0] == (*s)[i] || t.pattern[0] == '.' {
				tokenIndex = 1
			}
		}
	}
	return false, 0
}

func isMatch(s string, p string) bool {
	runtime.GOMAXPROCS(1)
	reg := NewReg(p)
	return reg.Match(s)
}

type regToken interface {
	tokenName() string
	token() string
	matchAll(s *string, left, right int) (matched bool, matchedLen int)
	findFirst(s *string, left int) (found bool, index int)
}

type Reg struct {
	pattern    string
	tokens     []regToken
	flatTokens []regToken
	wg         sync.WaitGroup
	matched    chan struct{}
}

func (r *Reg) printFlatTokens() {
	fmt.Print("\n", len(r.flatTokens))
	for _, x := range r.flatTokens {
		fmt.Print(" ", x.token())
	}
	fmt.Print("\n")
	for _, x := range r.flatTokens {
		fmt.Print(" ", x.tokenName())
	}
	fmt.Print("\n")
}

func (r *Reg) printCombinedTokens() {
	fmt.Print("\n", len(r.tokens))
	for _, x := range r.tokens {
		fmt.Print(" ", x.token())
	}
	fmt.Print("\n")
	for _, x := range r.tokens {
		fmt.Print(" ", x.tokenName())
	}
	fmt.Print("\n")
}

func NewReg(pattern string) *Reg {
	return (&Reg{pattern: pattern}).parseTokens()
}

func (r *Reg) parseTokens() *Reg {

	length := len(r.pattern)
	if length == 1 {
		r.tokens = []regToken{&plainToken{pattern: r.pattern}}
		return r
	}

	plain := ""
	for i := 0; i < length; i++ {
		if i < length-1 && r.pattern[i+1] == '*' {
			if plain != "" {
				r.flatTokens = append(r.flatTokens, &plainToken{pattern: plain})
			}
			plain = ""
			r.flatTokens = append(r.flatTokens, r.parseToken(i, i+2))
			i++
		} else {
			plain += string(r.pattern[i])
			if i == length-1 {
				if plain != "" {
					r.flatTokens = append(r.flatTokens, &plainToken{pattern: plain})
				}
			}
		}
	}
	r.printFlatTokens()
	combined := r.combineTokens()
	r.printCombinedTokens()
	return combined
}

func (r *Reg) parseToken(left, right int) regToken {
	//fmt.Println("parseToken", left, right)
	if right-left == 2 {
		if r.pattern[left+1] == '*' {
			if r.pattern[left] == '.' {
				return &dotStarToken{}
			}
			return &starsToken{letters: r.pattern[left : left+1]}
		}
	}
	plain := r.pattern[left:right]
	return &plainToken{pattern: plain}
}

func (r *Reg) combineTokens() *Reg {

	for i, token := range r.flatTokens {
		if i == 0 {
			r.tokens = append(r.tokens, token)
			continue
		}
		if previous := r.tokens[len(r.tokens)-1]; (previous.tokenName() == "star" || previous.tokenName() == "dotStar") &&
			(token.tokenName() == "star" || token.tokenName() == "dotStar") {
			if previous.tokenName() == "star" && token.tokenName() == "star" {
				r.tokens[len(r.tokens)-1].(*starsToken).letters += token.token()
				continue
			}
			if previous.tokenName() == "dotStar" {
				continue
			}
			if token.tokenName() == "dotStar" {
				r.tokens[len(r.tokens)-1] = &dotStarToken{}
				continue
			}
		}
		r.tokens = append(r.tokens, token)
	}
	return r
}

func (r *Reg) Match(s string) bool {
	var tokenNum, sLen = len(r.tokens), len(s)
	var tokenLeft, tokenRight = 0, tokenNum
	var sLeft, sRight = 0, sLen - 1
	if len(r.tokens) == 0 {
		return s == ""
	}
	if token := r.tokens[0]; token.tokenName() == "plain" {
		if ok, _ := token.matchAll(&s, 0, len(token.token())-1); !ok {
			return false
		}
		s = s[len(token.token()):]
		sLen = len(s)
		//sLeft = len(token.token())
		tokenLeft++
	}
	if tokenNum > 1 {
		if token := r.tokens[tokenNum-1]; token.tokenName() == "plain" {
			if ok, _ := token.matchAll(&s, sLen-len(token.token()), sLen-1); !ok {
				return false
			}
			s = s[:sLen-len(token.token())]
			sLen = len(s)
			sRight = sLen - 1
			tokenRight--
		}
	}
	fmt.Println("after trimming", s)
	for i := tokenLeft; i < tokenRight; i++ {
		token := r.tokens[i]
		if i == tokenRight-1 { // last one
			ok, _ := token.matchAll(&s, sLeft, sRight)
			fmt.Println("Last token MatchAll", ok)
			if i == 0 || ok {
				return ok
			}
			return r.WaitForResult()
		}
		r.matched = make(chan struct{})

		plain := r.tokens[i+1]
		fmt.Println("findFirst, sLeft", sLeft, "sRight", sRight)
		ok, index := plain.findFirst(&s, sLeft)
		if !ok {
			fmt.Println("???? is this it then ??????? sLeft", sLeft, "sRight", sRight)
			return false
		}
		if ok, _ := token.matchAll(&s, sLeft, index-len(plain.token())); !ok {
			fmt.Println("stars matchAll failed???", sLeft, index-1, s[sLeft: index-len(plain.token())])
			return r.WaitForResult()
		}
		fmt.Println("iterate", tokenLeft, tokenRight, sLeft+1, index-1)
		sLeft = index+1
		if sLeft <= len(s)-1  {
			r.wrap(&s, tokenLeft, tokenRight, sLeft, sRight)
		}
		i++
	}
	return sLeft >= len(s)
}

func (r *Reg) iterate(s *string, tokenLeft, tokenRight, sLeft, sRight int) {
	fmt.Println("iterate inside", tokenLeft, tokenRight, sLeft, sRight, "s='", (*s)[sLeft:], "'")
	for i := tokenLeft; i < tokenRight; i++ {
		token := r.tokens[i]
		if i == tokenRight-1 { // last one
			ok, _ := token.matchAll(s, sLeft, sRight)
			if ok {
				r.matched <- struct{}{}
			}
			return
		}

		plain := r.tokens[i+1]
		ok, index := plain.findFirst(s, sLeft)
		fmt.Println("!!!!!!!!!!!!!!!iterate findFirst!!!!!!!!!!!!!!!!!", index)
		if index == 16 {
		}
		if !ok {
			return
		}
		if ok, _ := token.matchAll(s, sLeft, index); !ok {
			return
		}
		sLeft = index+1
		if sLeft <= len(*s)-1 {
			r.wrap(s, tokenLeft, tokenRight, sLeft, sRight)
		}
		i++
	}
}

func (r *Reg) wrap(s *string, tokenLeft, tokenRight, sLeft, sRight int) {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		r.iterate(s, tokenLeft, tokenRight, sLeft, sRight)
	}()
}

func (r *Reg) WaitForResult() bool {
	var noMatch = make(chan struct{})
	go func() {
		r.wg.Wait()
		noMatch <- struct{}{}
	}()
	select {
	case <-r.matched:
		return true
	case <-noMatch:
		return false
	}
}
