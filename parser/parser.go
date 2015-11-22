package parser
import (
	"strings"
	"errors"
)

// Parser represents parser interface for other
type Parser interface {
	GetTitle(string) string
	Parse([]byte) string
}

var NotFound = errors.New("Parser not found")

// New created new parser
func New(name string) (Parser, error) {
	switch {
	case strings.HasSuffix(name, ".md"):
		return NewMdParser(), nil
	}

	return nil, NotFound
}