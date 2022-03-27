// Copyright (c) 2022 Bi Anlatsana
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package goht

import "testing"

func TestDomElement(t *testing.T) {
	dom := Html(
		Props{"lang": "en"},
		RawDom(
			"body",
			DomList(
				RawDom("h1", "Hello World", Props{"id": "title"}),
				RawDom("p", "An Example", Props{"id": "paragraph"}),
			),
			Props{},
		),
	)

	println(dom)
}
