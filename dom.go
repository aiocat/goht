// Copyright (c) 2022 aiocat
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package goht

// Wrapper for document element list
func DomList(elements ...DomElement) string {
	return DomElementList(elements).C()
}

// Wrapper for document element
func Dom(key string, value string, props Props) string {
	return DomElement{key, value, props}.C()
}

// Raw wrapper for document element
func RawDom(key string, value string, props Props) DomElement {
	return DomElement{key, value, props}
}

// Raw wrapper for document element list
func RawDomList(elements ...DomElement) DomElementList {
	return DomElementList(elements)
}
