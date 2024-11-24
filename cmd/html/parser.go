package html

import (
	"io"

	x "golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Response(r io.Reader) ([]byte, error) {
	doc := x.NewTokenizer(r)
	var out []byte
	var add bool
loop:
	for {
		typ := doc.Next()
		tok := doc.Token()

		switch typ {
		case x.ErrorToken:
			return nil, doc.Err()
		case x.StartTagToken:
			if tok.DataAtom == atom.P {
				add = true
			}
		case x.TextToken:
			if add {
				out = append(out, tok.Data...)
			}
		case x.EndTagToken:
			if tok.DataAtom == atom.P {
				break loop
			}
		}
	}

	return out, nil
}
