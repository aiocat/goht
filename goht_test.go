package goht

import "testing"

func TestDomElement(t *testing.T) {
	dom := DomElement{
		"body",
		DomList{
			DomElement{
				"h1", "Hello World", Props{"id": "title"},
			},
			DomElement{
				"p", "An Example?", Props{"id": "example"},
			},
		}.C(),
		Props{},
	}.C()

	println(dom)
}
