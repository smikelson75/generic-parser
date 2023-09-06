package readers

import (
	"bufio"
	"regexp"
	"strings"
)

type Reader struct {
	buffer uint
	reader *bufio.Reader
	tokens []rune
}

func FromString(input string, buffer uint) *Reader {
	return &Reader{
		buffer: buffer,
		reader: bufio.NewReader(strings.NewReader(input)),
	}
}

func (r *Reader) Read() error {
	for i := 0; i < int(r.buffer); i++ {
		character, _, err := r.reader.ReadRune()
		if err != nil {
			return err
		}
		r.tokens = append(r.tokens, character)
	}
	return nil
}

func (r *Reader) ReadTo(exitRune rune) error {
	for {
		character, _, err := r.reader.ReadRune()
		if err != nil {
			return err
		} else if character == exitRune {
			r.tokens = append(r.tokens, character)
			break
		}

		r.tokens = append(r.tokens, character)
	}

	return nil
}

func (r *Reader) Move(amount int) error {
	if len(r.tokens) > 0 && len(r.tokens) > amount {
		r.tokens = r.tokens[amount:]
		return nil
	}

	return r.ReadTo('\n')
}

func (r *Reader) Get(pattern TokenPattern) (*string, error) {
	line := string(r.tokens)
	tester := regexp.MustCompile(pattern.Pattern())

	if matches := tester.FindStringSubmatch(line); matches != nil {
		return &matches[0], r.Move(len(matches[0]))
	}

	return nil, nil
}
