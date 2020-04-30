package q10

import (
	"fmt"
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
func (t *starsToken) matchAll(s *string, left, right int) (bool, int) {
	fmt.Println("left, right", left, right)
	if left > right {
		return true, 0
	}
	si := left
	for _, x := range []byte(t.letters) {
		if (*s)[si] != x {
			si++
		}
	}
	if si < right {
		return false, 0
	}
	return true, right - left + 1
}
func (t *starsToken) findFirst(s *string, left, right int, previousLetter byte) (found bool, index int) {
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
func (t *dotStarToken) findFirst(s *string, left, right int, previousLetter byte) (found bool, index int) {
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
	ti := 0
	for i := left; i < right+1; i++ {
		if t.pattern[ti] != (*s)[i] && t.pattern[ti] != '.' {
			return false, 0
		}
		ti++
	}
	return true, right - left + 1
}
func (t *plainToken) findFirst(s *string, left, right int, previousLetter byte) (bool, int) {
	ti := 0
	for i := left; i < right+1; i++ {
		if t.pattern[ti] == (*s)[i] || t.pattern[ti] == '.' {
			if ti == len(t.pattern)-1 {
				return true, i
			}
			ti++
		} else {
			ti = 0
		}
		//if previousLetter != '.' && (*s)[i] != previousLetter {
		//	return false, 0
		//}
	}
	return false, 0
}

func isMatch2(s string, p string) bool {
	reg := NewReg(p)
	return reg.Match(s)
}

type regToken interface {
	tokenName() string
	token() string
	matchAll(s *string, left, right int) (matched bool, matchedLen int)
	findFirst(s *string, left, right int, previousLetter byte) (found bool, index int)
}

type Reg struct {
	pattern    string
	tokens     []regToken
	flatTokens []regToken
	wg         sync.WaitGroup
	matched    chan struct{}
	strOver    chan struct{}
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
	var left, right int

	plain := ""
	for i := 0; i < length; i++ {
		if i < length-1 && r.pattern[i+1] == '*' {
			if plain != "" {
				r.flatTokens = append(r.flatTokens, &plainToken{pattern: plain})
			}
			plain = ""
			right = i + 2
			r.flatTokens = append(r.flatTokens, r.parseToken(left, right))
			i++
			left = i + 2
		} else {
			plain += string(r.pattern[i])
			if i == length-1 {
				if plain != "" {
					r.flatTokens = append(r.flatTokens, &plainToken{pattern: plain})
				}
			}
		}
	}
	fmt.Println(len(r.flatTokens), r.flatTokens[0].token(), r.flatTokens[1].token(), r.flatTokens[2].token(), r.flatTokens[3].token())
	fmt.Println(len(r.flatTokens), r.flatTokens[0].tokenName(), r.flatTokens[1].tokenName(), r.flatTokens[2].tokenName(), r.flatTokens[3].tokenName())
	return r.combineTokens()
}

func (r *Reg) parseToken(left, right int) regToken {
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
	//fmt.Println(r.tokens[0].token(), r.tokens[1].token())
	return r
}

func (r *Reg) Match(s string) bool {
	var tokenNum, sLen = len(r.tokens), len(s)
	var tokenLeft, tokenRight = 0, tokenNum
	var sLeft, sRight = 0, sLen - 1
	if token := r.tokens[0]; token.tokenName() == "plain" {
		if ok, _ := r.tokens[0].matchAll(&s, 0, len(token.token())-1); !ok {
			return false
		}
		s = s[len(token.token())-1:]
		sLeft = len(token.token())
		tokenLeft++
	}
	if tokenNum > 1 {
		if token := r.tokens[tokenNum-1]; token.tokenName() == "plain" {
			if ok, _ := token.matchAll(&s, sLen-len(token.token()), sLen-1); !ok {
				return false
			}
			s = s[:sLen-len(token.token())]
			sRight -= len(token.token())
			tokenRight--
		}
	}
	for i := tokenLeft; i < tokenRight; i++ {
		token := r.tokens[i]
		if i == tokenRight-1 { // last one
			ok, _ := token.matchAll(&s, sLeft, sRight)
			if r.matched == nil {
				return ok
			}
			if ok {
				r.wrap(func() {
					r.matched <- struct{}{}
				})
			}
			return r.WaitForResult()
		}
		r.matched, r.strOver = make(chan struct{}), make(chan struct{})

		plain := r.tokens[i+1]
		ok, index := plain.findFirst(&s, sLeft, sRight, token.token()[len(token.token())-1])
		if !ok {
			return false
		}
		r.wrap(func() {
			r.iterate(&s, i, tokenRight, index+1, sRight)
		})
		token.matchAll(&s, sLeft, index)
		i++
	}
	return true
}

func (r *Reg) iterate(s *string, tokenLeft, tokenRight, sLeft, sRight int) {
	for i := tokenLeft; i < tokenRight; i++ {
		token := r.tokens[i]
		if i == tokenRight-1 { // last one
			ok, _ := token.matchAll(s, sLeft, sRight)
			if ok {
				r.matched <- struct{}{}
			}
			return
		}
		r.matched = make(chan struct{})

		plain := r.tokens[i+1]
		ok, index := plain.findFirst(s, sLeft, sRight, token.token()[len(token.token())-1])
		if !ok {
			return
		}
		r.wrap(func() {
			r.iterate(s, i, tokenRight, index+1, sRight)
		})
		token.matchAll(s, sLeft, index)
		i++
	}
}

func (r *Reg) wrap(f func()) {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		f()
	}()
}

func (r *Reg) WaitForResult() bool {
	var noMatch = make(chan struct{})
	go func() {
		<-r.strOver // wait for str
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
