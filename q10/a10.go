package q10

import (
	"bytes"
	"io"
)

const (
	plainOriginalSetType = iota + 1
	starOriginalSetType
	dotStarOriginalSetType
	starSetType
	dotStarSetType
)

type token interface {
	Type() int
	matchAndNext(b byte) bool
	matchedAll() bool
}

type tokenSet interface {
	Type() int
	matchAndNext(b byte) bool
	matchedAll() bool
}
type tokenSets []tokenSet

func (ts *tokenSets) WalkAndMatch(buffer *bytes.Buffer) (matched bool) {
	for {
		b, err := buffer.ReadByte()
		if err == io.EOF {
			return true
		}
		failed := false
		func() {
			for _, tokens := range *ts {
				if next := tokens.matchAndNext(b); next {
					if tokens.matchedAll() {
						continue
					}
					return
				}
				failed = true
				return
			}
		}()
		if failed {
			return false
		}
	}
}

func isMatch(s string, p string) bool {
	tokens := parseToken(p)
	buffer := bytes.NewBufferString(s)
	return tokens.WalkAndMatch(buffer)
}

func parseToken(p string) (tokenSets tokenSets) {
	// put every (star)*n + original in the same outer slice
	// possible sets: original; (star)*n + original; dotStar + original; (star)*n; dotStar
	prefixDotsContinual := true
	var prefixDots int
	pb := []byte(p)
	length := len(pb)
	if length == 1 {
		return append(tokenSets, NewOriginal(0, []byte(p)))
	}
	var left, right int
	var plainTokens []token

	for i := 0; i < length; i++ {
		if i < length-1 && pb[i+1] == '*' {
			right = i + 2
			prefixDotsContinual = true
			plainTokens = append(plainTokens, parsePlainTokens(prefixDots, pb[left:right]))
			prefixDots = 0
			i++
			left = i + 1
		} else {
			right = i + 1
			if prefixDotsContinual && pb[i] == '.' {
				prefixDots++
			} else if prefixDotsContinual && pb[i] != '.' {
				prefixDotsContinual = false
			}
			if i == length-1 {
				plainTokens = append(plainTokens, parsePlainTokens(prefixDots, pb[left:right]))
			}
		}
	}

	return combinePlainTokens(plainTokens)
}

func parsePlainTokens(prefixDots int, p []byte) token {
	if len(p) == 2 {
		if p[1] == '*' {
			if p[0] == '.' {
				return NewDotStar()
			}
			return NewStar(p[0])
		}
	}
	return NewOriginal(prefixDots, p)
}

func combinePlainTokens(tokens []token) (tokenSets tokenSets) {
	var metStar, metDotStar bool
	var stars []*star
	for i, token := range tokens {
		if i == 0 && token.Type() == plainOriginalSetType {
			tokenSets = append(tokenSets, token)
		}
		if i > 0 && token.Type() == plainOriginalSetType {
			if metDotStar {
				tokenSets = append(tokenSets, NewDotStarOriginal(token.(*original)))
			} else if metStar {
				tokenSets = append(tokenSets, NewStarOriginal(token.(*original), stars...))
			}
			stars = nil
			metStar, metDotStar = false, false
		}
		if token.Type() == starSetType {
			metStar = true
			stars = append(stars, token.(*star))
		}
		if token.Type() == dotStarSetType {
			metDotStar = true
		}
		if i == len(tokens)-1 {
			if metDotStar {
				tokenSets = append(tokenSets, NewDotStarOriginal(nil))
			} else if metStar {
				tokenSets = append(tokenSets, NewStarOriginal(nil, stars...))
			}
		}
	}
	return
}

func NewOriginal(prefixDots int, p []byte) *original {
	return &original{
		prefixDots: prefixDots,
		buffer:     bytes.NewBuffer(p),
	}
}

type original struct { // certain string
	prefixDots int // prefixed dots count
	buffer     *bytes.Buffer
	current    byte
	matched    bool
}

func (o *original) matchAndNext(b byte) bool {
	var err error
	if o.prefixDots > 0 {
		o.prefixDots--
		return true
	}
	o.current, err = o.buffer.ReadByte()
	if err == io.EOF {
		o.matched = true
		return true
	}
	return o.current == b || o.current == '.'
}

func (o *original) matchedAll() bool {
	return o.matched
}

func (o *original) Type() int {
	return plainOriginalSetType
}

func NewStar(l byte) *star {
	return &star{letter: l}
}

type star struct { // x*
	letter  byte
	matched bool
}

type stars struct { // (x*)*n
	buffer  *bytes.Buffer
	matched bool
}

func (s *star) Type() int {
	return starSetType
}

func (s *star) matchAndNext(b byte) bool {
	panic("implement me")
}

func (s *star) matchedAll() bool {
	panic("implement me")
}

func NewDotStar() *dotStar {
	return &dotStar{}
}

type dotStar struct{} // .*

func (d *dotStar) Type() int {
	return dotStarSetType
}

func (d *dotStar) matchAndNext(b byte) bool {
	panic("implement me")
}

func (d *dotStar) matchedAll() bool {
	panic("implement me")
}

func NewStarOriginal(original *original, star ...*star) *starOriginal {
	return &starOriginal{stars: star, original: original}
}

type starOriginal struct {
	stars                []*star
	index                int // -1 表示完结
	original             *original
	matchedOriginalIndex []int // 匹配中的original Index
}

func (s *starOriginal) Type() int {
	return starOriginalSetType
}

func (s *starOriginal) matchAndNext(b byte) bool {
	if s.index >= len(s.stars) {
		s.index = -1
		return s.original == nil || s.original.matchAndNext(b)
	}
	if s.index >= 0 && s.index < len(s.stars) {
		if s.index == len(s.stars)-1 {

		}
		if s.index < len(s.stars) && s.stars[s.index].letter == b {
			return true
		}
		if s.stars[s.index].letter != b {
			s.stars[s.index].matched = true
			s.index++
			return s.matchAndNext(b)
		}
		return false
	}
	return s.index == -1 && (s.original == nil || s.original.matchAndNext(b))
}

func (s *starOriginal) matchedAll() bool {
	if s.original != nil {
		return s.original.matchedAll()
	}
	return s.index == -1
}

func NewDotStarOriginal(original *original) *dotStarOriginal {
	return &dotStarOriginal{original: original}
}

type dotStarOriginal struct {
	original *original
}

func (d *dotStarOriginal) Type() int {
	return dotStarOriginalSetType
}

// todo
func (d *dotStarOriginal) matchAndNext(b byte) bool {
	return d.original.matchAndNext(b)
}

func (d *dotStarOriginal) matchedAll() bool {
	return d.original.matchedAll()
}
