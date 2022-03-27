// Copyright (c) 2022 Bi Anlatsana
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

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
