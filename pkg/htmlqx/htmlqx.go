package htmlqx

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/antchfx/htmlquery"
)

type ParsingFailedError struct {
	err error
}

func (e *ParsingFailedError) Error() string {
	return "parsing failed: " + e.err.Error()
}

func NewParsingFailedError(err error) error {
	return &ParsingFailedError{err}
}

var (
	ErrEmpty     = errors.New("empty result")
	ErrAmbiguous = errors.New("ambiguous result")
)

type content []byte

func Parse(b []byte) *content {
	c := make(content, len(b))
	copy(c, b)
	return &c
}

func (c *content) Text(xpath string) (string, error) {
	doc, err := htmlquery.Parse(bytes.NewReader(*c))
	if err != nil {
		return "", NewParsingFailedError(err)
	}

	list := htmlquery.Find(doc, xpath)
	if len(list) == 0 {
		return "", ErrEmpty
	}
	if len(list) > 1 {
		return "", ErrAmbiguous
	}

	return htmlquery.InnerText(list[0]), nil
}

func (c *content) Int(xpath string) (int, error) {
	v, err := c.Text(xpath)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(v)
}

func (c *content) Int64(xpath string) (int64, error) {
	v, err := c.Text(xpath)
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(v, 10, 64)
}
